package om

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/rueian/rueidis/internal/proto"
)

var ErrVersionMismatch = errors.New("object version mismatched, please retry")

type ExecFn func(ctx context.Context, keys, args []string) proto.Result
type ScriptExecFn func(script string) ExecFn

type HashObjectAdapter interface {
	Save(ctx context.Context, key string, fields map[string]string) error
	Fetch(ctx context.Context, key string) (map[string]proto.Message, error)
	FetchCache(ctx context.Context, key string, ttl time.Duration) (map[string]proto.Message, error)
	Remove(ctx context.Context, key string) error
}

func NewHashRepository(prefix string, schema interface{}, adapter HashObjectAdapter, scriptExecFn ScriptExecFn) *HashRepository {
	repo := &HashRepository{
		prefix:  prefix,
		typ:     reflect.TypeOf(schema),
		adapter: adapter,
		execFn:  scriptExecFn(saveScript),
	}
	if _, ok := schema.(HashConverter); !ok {
		repo.factory = newHashConvFactory(repo.typ)
	}
	return repo
}

type HashRepository struct {
	prefix  string
	typ     reflect.Type
	factory *hashConvFactory

	adapter HashObjectAdapter
	execFn  ExecFn
}

func (r *HashRepository) key(id string) (key string) {
	sb := strings.Builder{}
	sb.Grow(len(r.prefix) + len(id) + 1)
	sb.WriteString(r.prefix)
	sb.WriteString(":")
	sb.WriteString(id)
	return sb.String()
}

func (r *HashRepository) conv(v reflect.Value) (entity interface{}, conv HashConverter) {
	if r.factory != nil {
		return v.Interface(), r.factory.NewConverter(v)
	}
	entity = v.Interface()
	return entity, entity.(HashConverter)
}

func (r *HashRepository) Make() (entity interface{}) {
	v := reflect.New(r.typ)
	entity, conv := r.conv(v)
	_ = conv.FromHash(id(), nil)
	return entity
}

func (r *HashRepository) fromHash(id string, record map[string]proto.Message) (v interface{}, err error) {
	fields := make(map[string]string, len(record))
	for k, v := range record {
		if !v.IsNil() {
			fields[k] = v.String
		}
	}

	v, conv := r.conv(reflect.New(r.typ))
	if err := conv.FromHash(id, fields); err != nil {
		return nil, err
	}
	return v, nil
}

func (r *HashRepository) Fetch(ctx context.Context, id string) (v interface{}, err error) {
	record, err := r.adapter.Fetch(ctx, r.key(id))
	if err != nil {
		return nil, err
	}
	return r.fromHash(id, record)
}

func (r *HashRepository) FetchCache(ctx context.Context, id string, ttl time.Duration) (v interface{}, err error) {
	record, err := r.adapter.FetchCache(ctx, r.key(id), ttl)
	if err != nil {
		return nil, err
	}
	return r.fromHash(id, record)
}

func (r *HashRepository) Save(ctx context.Context, entity interface{}) (err error) {
	var conv HashConverter

	if r.factory != nil {
		conv = r.factory.NewConverter(reflect.ValueOf(entity))
	} else {
		conv = entity.(HashConverter)
	}

	id, fields := conv.ToHash()
	if ver, ok := fields[VersionField]; ok {
		args := make([]string, 0, len(fields)*2)
		args = append(args, VersionField, ver)
		for f, v := range fields {
			if f == VersionField {
				continue
			}
			args = append(args, f, v)
		}
		fields[VersionField], err = r.execFn(ctx, []string{r.key(id)}, args).ToString()
		if proto.IsRedisNil(err) {
			return ErrVersionMismatch
		}
		if err != nil {
			return err
		}
		return conv.FromHash(id, fields)
	}
	return r.adapter.Save(ctx, r.key(id), fields)
}

func (r *HashRepository) Remove(ctx context.Context, id string) error {
	return r.adapter.Remove(ctx, r.key(id))
}

var saveScript = fmt.Sprintf(`
local v = redis.call('HGET',KEYS[1],'%s')
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  if redis.call('HSET',KEYS[1],unpack(ARGV)) then return ARGV[2] end
end
return nil
`, VersionField)
