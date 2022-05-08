package rueidiscompat

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/rueian/rueidis"
)

type StringCmd struct {
	res rueidis.RedisResult
	val string
	err error
}

func newStringCmd(res rueidis.RedisResult) *StringCmd {
	val, err := res.ToString()
	return &StringCmd{res: res, val: val, err: err}
}

func (cmd *StringCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StringCmd) Val() string {
	return cmd.val
}

func (cmd *StringCmd) Err() error {
	return cmd.err
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringCmd) Bytes() ([]byte, error) {
	return []byte(cmd.val), cmd.err
}

func (cmd *StringCmd) Bool() (bool, error) {
	return cmd.res.ToBool()
}

func (cmd *StringCmd) Int() (int, error) {
	i, err := cmd.res.ToInt64()
	return int(i), err
}

func (cmd *StringCmd) Int64() (int64, error) {
	return cmd.res.ToInt64()
}

func (cmd *StringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *StringCmd) Float32() (float32, error) {
	f, err := cmd.res.ToFloat64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func (cmd *StringCmd) Float64() (float64, error) {
	return cmd.res.ToFloat64()
}

func (cmd *StringCmd) Time() (time.Time, error) {
	if cmd.err != nil {
		return time.Time{}, cmd.err
	}
	return time.Parse(time.RFC3339Nano, cmd.Val())
}

func (cmd *StringCmd) String() string {
	return cmd.val
}

type BoolCmd struct {
	res rueidis.RedisResult
	val bool
	err error
}

func newBoolCmd(res rueidis.RedisResult) *BoolCmd {
	val, err := res.AsBool()
	if rueidis.IsRedisNil(err) {
		val = false
		err = nil
	}
	return &BoolCmd{res: res, val: val, err: err}
}

func (cmd *BoolCmd) SetVal(val bool) {
	cmd.val = val
}

func (cmd *BoolCmd) Val() bool {
	return cmd.val
}

func (cmd *BoolCmd) Err() error {
	return cmd.err
}

func (cmd *BoolCmd) Result() (bool, error) {
	return cmd.val, cmd.err
}

func (cmd *BoolCmd) String() (string, error) {
	return cmd.res.ToString()
}

type IntCmd struct {
	res rueidis.RedisResult
	val int64
	err error
}

func newIntCmd(res rueidis.RedisResult) *IntCmd {
	val, err := res.AsInt64()
	return &IntCmd{res: res, val: val, err: err}
}

func (cmd *IntCmd) SetVal(val int64) {
	cmd.val = val
}

func (cmd *IntCmd) Val() int64 {
	return cmd.val
}

func (cmd *IntCmd) Err() error {
	return cmd.err
}

func (cmd *IntCmd) Result() (int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntCmd) Uint64() (uint64, error) {
	return uint64(cmd.val), cmd.err
}

func (cmd *IntCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StatusCmd struct {
	res rueidis.RedisResult
	val string
	err error
}

func newStatusCmd(res rueidis.RedisResult) *StatusCmd {
	val, err := res.ToString()
	return &StatusCmd{res: res, val: val, err: err}
}

func (cmd *StatusCmd) SetVal(val string) {
	cmd.val = val
}

func (cmd *StatusCmd) Val() string {
	return cmd.val
}

func (cmd *StatusCmd) Err() error {
	return cmd.err
}

func (cmd *StatusCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *StatusCmd) String() (string, error) {
	return cmd.res.ToString()
}

type SliceCmd struct {
	res rueidis.RedisResult
	val []interface{}
	err error
}

func newSliceCmd(res rueidis.RedisResult) *SliceCmd {
	val, err := res.ToArray()
	slice := &SliceCmd{res: res, val: make([]interface{}, len(val)), err: err}
	for i, v := range val {
		if s, err := v.ToString(); err == nil {
			slice.val[i] = s
		}
	}
	return slice
}

func (cmd *SliceCmd) SetVal(val []interface{}) {
	cmd.val = val
}

func (cmd *SliceCmd) Val() []interface{} {
	return cmd.val
}

func (cmd *SliceCmd) Err() error {
	return cmd.err
}

func (cmd *SliceCmd) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *SliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StringSliceCmd struct {
	res rueidis.RedisResult
	val []string
	err error
}

func newStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	val, err := res.AsStrSlice()
	return &StringSliceCmd{res: res, val: val, err: err}
}

func (cmd *StringSliceCmd) SetVal(val []string) {
	cmd.val = val
}

func (cmd *StringSliceCmd) Val() []string {
	return cmd.val
}

func (cmd *StringSliceCmd) Err() error {
	return cmd.err
}

func (cmd *StringSliceCmd) Result() ([]string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type IntSliceCmd struct {
	res rueidis.RedisResult
	val []int64
	err error
}

func newIntSliceCmd(res rueidis.RedisResult) *IntSliceCmd {
	val, err := res.AsIntSlice()
	return &IntSliceCmd{res: res, val: val, err: err}
}

func (cmd *IntSliceCmd) SetVal(val []int64) {
	cmd.val = val
}

func (cmd *IntSliceCmd) Val() []int64 {
	return cmd.val
}

func (cmd *IntSliceCmd) Err() error {
	return cmd.err
}

func (cmd *IntSliceCmd) Result() ([]int64, error) {
	return cmd.val, cmd.err
}

func (cmd *IntSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type BoolSliceCmd struct {
	res rueidis.RedisResult
	val []bool
	err error
}

func newBoolSliceCmd(res rueidis.RedisResult) *BoolSliceCmd {
	ints, err := res.AsIntSlice()
	if err != nil {
		return &BoolSliceCmd{res: res, err: err}
	}
	val := make([]bool, 0, len(ints))
	for _, i := range ints {
		val = append(val, i == 1)
	}
	return &BoolSliceCmd{res: res, val: val, err: err}
}

func (cmd *BoolSliceCmd) SetVal(val []bool) {
	cmd.val = val
}

func (cmd *BoolSliceCmd) Val() []bool {
	return cmd.val
}

func (cmd *BoolSliceCmd) Err() error {
	return cmd.err
}

func (cmd *BoolSliceCmd) Result() ([]bool, error) {
	return cmd.val, cmd.err
}

func (cmd *BoolSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type FloatSliceCmd struct {
	res rueidis.RedisResult
	val []float64
	err error
}

func newFloatSliceCmd(res rueidis.RedisResult) *FloatSliceCmd {
	val, err := res.AsFloatSlice()
	return &FloatSliceCmd{res: res, val: val, err: err}
}

func (cmd *FloatSliceCmd) SetVal(val []float64) {
	cmd.val = val
}

func (cmd *FloatSliceCmd) Val() []float64 {
	return cmd.val
}

func (cmd *FloatSliceCmd) Err() error {
	return cmd.err
}

func (cmd *FloatSliceCmd) Result() ([]float64, error) {
	return cmd.val, cmd.err
}

func (cmd *FloatSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type ZSliceCmd struct {
	res rueidis.RedisResult
	val []Z
	err error
}

func newZSliceCmd(res rueidis.RedisResult) *ZSliceCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &ZSliceCmd{res: res, err: err}
	}
	val := make([]Z, 0, len(arr)/2)
	for i, j := 0, 1; i < len(arr); i, j = i+2, j+2 {
		member, err := arr[i].ToString()
		if err != nil {
			return &ZSliceCmd{res: res, err: err}
		}
		score, err := arr[j].AsFloat64()
		if err != nil {
			return &ZSliceCmd{res: res, err: err}
		}
		val = append(val, Z{
			Member: member,
			Score:  score,
		})
	}
	return &ZSliceCmd{res: res, val: val, err: err}
}

func (cmd *ZSliceCmd) SetVal(val []Z) {
	cmd.val = val
}

func (cmd *ZSliceCmd) Val() []Z {
	return cmd.val
}

func (cmd *ZSliceCmd) Err() error {
	return cmd.err
}

func (cmd *ZSliceCmd) Result() ([]Z, error) {
	return cmd.val, cmd.err
}

func (cmd *ZSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type FloatCmd struct {
	res rueidis.RedisResult
	val float64
	err error
}

func newFloatCmd(res rueidis.RedisResult) *FloatCmd {
	val, err := res.ToFloat64()
	return &FloatCmd{res: res, val: val, err: err}
}

func (cmd *FloatCmd) SetVal(val float64) {
	cmd.val = val
}

func (cmd *FloatCmd) Val() float64 {
	return cmd.val
}

func (cmd *FloatCmd) Err() error {
	return cmd.err
}

func (cmd *FloatCmd) Result() (float64, error) {
	return cmd.val, cmd.err
}

func (cmd *FloatCmd) String() (string, error) {
	return cmd.res.ToString()
}

type ScanCmd struct {
	res    rueidis.RedisResult
	cursor uint64
	keys   []string
	err    error
}

func newScanCmd(res rueidis.RedisResult) *ScanCmd {
	ret, err := res.ToArray()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	cursor, err := ret[0].AsInt64()
	if err != nil {
		return &ScanCmd{res: res, err: err}
	}
	keys, err := ret[1].AsStrSlice()
	return &ScanCmd{res: res, cursor: uint64(cursor), keys: keys, err: err}
}

func (cmd *ScanCmd) SetVal(keys []string, cursor uint64) {
	cmd.keys = keys
	cmd.cursor = cursor
}

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) {
	return cmd.keys, cmd.cursor
}

func (cmd *ScanCmd) Err() error {
	return cmd.err
}

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	return cmd.keys, cmd.cursor, cmd.err
}

func (cmd *ScanCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StringStringMapCmd struct {
	res rueidis.RedisResult
	val map[string]string
	err error
}

func newStringStringMapCmd(res rueidis.RedisResult) *StringStringMapCmd {
	val, err := res.AsStrMap()
	return &StringStringMapCmd{res: res, val: val, err: err}
}

func (cmd *StringStringMapCmd) SetVal(val map[string]string) {
	cmd.val = val
}

func (cmd *StringStringMapCmd) Val() map[string]string {
	return cmd.val
}

func (cmd *StringStringMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringStringMapCmd) Result() (map[string]string, error) {
	return cmd.val, cmd.err
}

func (cmd *StringStringMapCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StringIntMapCmd struct {
	res rueidis.RedisResult
	val map[string]int64
	err error
}

func newStringIntMapCmd(res rueidis.RedisResult) *StringIntMapCmd {
	val, err := res.AsIntMap()
	return &StringIntMapCmd{res: res, val: val, err: err}
}

func (cmd *StringIntMapCmd) SetVal(val map[string]int64) {
	cmd.val = val
}

func (cmd *StringIntMapCmd) Val() map[string]int64 {
	return cmd.val
}

func (cmd *StringIntMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringIntMapCmd) Result() (map[string]int64, error) {
	return cmd.val, cmd.err
}

func (cmd *StringIntMapCmd) String() (string, error) {
	return cmd.res.ToString()
}

type StringStructMapCmd struct {
	res rueidis.RedisResult
	val map[string]struct{}
	err error
}

func newStringStructMapCmd(res rueidis.RedisResult) *StringStructMapCmd {
	strSlice, err := res.AsStrSlice()
	if err != nil {
		return &StringStructMapCmd{res: res, err: err}
	}
	val := make(map[string]struct{}, len(strSlice))
	for _, v := range strSlice {
		val[v] = struct{}{}
	}
	return &StringStructMapCmd{res: res, val: val, err: err}
}

func (cmd *StringStructMapCmd) SetVal(val map[string]struct{}) {
	cmd.val = val
}

func (cmd *StringStructMapCmd) Val() map[string]struct{} {
	return cmd.val
}

func (cmd *StringStructMapCmd) Err() error {
	return cmd.err
}

func (cmd *StringStructMapCmd) Result() (map[string]struct{}, error) {
	return cmd.val, cmd.err
}

func (cmd *StringStructMapCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XMessageSliceCmd struct {
	res rueidis.RedisResult
	val []XMessage
	err error
}

func newXMessageSliceCmd(res rueidis.RedisResult) *XMessageSliceCmd {
	val, err := res.AsXRangeSlice()
	slice := &XMessageSliceCmd{res: res, val: make([]XMessage, len(val)), err: err}
	for i, r := range val {
		slice.val[i] = newXMessage(r)
	}
	return slice
}

func newXMessage(r rueidis.XRange) XMessage {
	m := XMessage{ID: r.ID, Values: make(map[string]interface{}, len(r.FieldValues))}
	for k, v := range r.FieldValues {
		m.Values[k] = v
	}
	return m
}

func (cmd *XMessageSliceCmd) SetVal(val []XMessage) {
	cmd.val = val
}

func (cmd *XMessageSliceCmd) Val() []XMessage {
	return cmd.val
}

func (cmd *XMessageSliceCmd) Err() error {
	return cmd.err
}

func (cmd *XMessageSliceCmd) Result() ([]XMessage, error) {
	return cmd.val, cmd.err
}

func (cmd *XMessageSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XStream struct {
	Stream   string
	Messages []XMessage
}

type XStreamSliceCmd struct {
	res rueidis.RedisResult
	val []XStream
	err error
}

func newXStreamSliceCmd(res rueidis.RedisResult) *XStreamSliceCmd {
	arrs, err := res.ToArray()
	if err != nil {
		return &XStreamSliceCmd{res: res, err: err}
	}
	val := make([]XStream, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		if len(arr) != 2 {
			return &XStreamSliceCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
		}
		stream, err := arr[0].ToString()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		ranges, err := arr[1].AsXRangeSlice()
		if err != nil {
			return &XStreamSliceCmd{res: res, err: err}
		}
		msgs := make([]XMessage, 0, len(ranges))
		for _, r := range ranges {
			msgs = append(msgs, newXMessage(r))
		}
		val = append(val, XStream{Stream: stream, Messages: msgs})
	}
	return &XStreamSliceCmd{res: res, val: val, err: err}
}

func (cmd *XStreamSliceCmd) SetVal(val []XStream) {
	cmd.val = val
}

func (cmd *XStreamSliceCmd) Val() []XStream {
	return cmd.val
}

func (cmd *XStreamSliceCmd) Err() error {
	return cmd.err
}

func (cmd *XStreamSliceCmd) Result() ([]XStream, error) {
	return cmd.val, cmd.err
}

func (cmd *XStreamSliceCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XPending struct {
	Count     int64
	Lower     string
	Higher    string
	Consumers map[string]int64
}

type XPendingCmd struct {
	res rueidis.RedisResult
	val XPending
	err error
}

func newXPendingCmd(res rueidis.RedisResult) *XPendingCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	if len(arr) != 4 {
		return &XPendingCmd{res: res, err: fmt.Errorf("got %d, wanted 4", len(arr))}
	}
	count, err := arr[0].ToInt64()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	lower, err := arr[1].ToString()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	higher, err := arr[2].ToString()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	val := XPending{
		Count:  count,
		Lower:  lower,
		Higher: higher,
	}
	consumerArr, err := arr[3].ToArray()
	if err != nil {
		return &XPendingCmd{res: res, err: err}
	}
	for _, v := range consumerArr {
		consumer, err := v.ToArray()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		if len(arr) != 2 {
			return &XPendingCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
		}
		consumerName, err := consumer[0].ToString()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		consumerPending, err := consumer[1].AsInt64()
		if err != nil {
			return &XPendingCmd{res: res, err: err}
		}
		if val.Consumers == nil {
			val.Consumers = make(map[string]int64)
		}
		val.Consumers[consumerName] = consumerPending
	}
	return &XPendingCmd{res: res, val: val, err: err}
}

func (cmd *XPendingCmd) SetVal(val XPending) {
	cmd.val = val
}

func (cmd *XPendingCmd) Val() XPending {
	return cmd.val
}

func (cmd *XPendingCmd) Err() error {
	return cmd.err
}

func (cmd *XPendingCmd) Result() (XPending, error) {
	return cmd.val, cmd.err
}

func (cmd *XPendingCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	res rueidis.RedisResult
	val []XPendingExt
	err error
}

func newXPendingExtCmd(res rueidis.RedisResult) *XPendingExtCmd {
	arrs, err := res.ToArray()
	if err != nil {
		return &XPendingExtCmd{res: res, err: err}
	}
	val := make([]XPendingExt, 0, len(arrs))
	for _, v := range arrs {
		arr, err := v.ToArray()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		if len(arr) != 4 {
			return &XPendingExtCmd{res: res, err: fmt.Errorf("got %d, wanted 4", len(arr))}
		}
		id, err := arr[0].ToString()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		consumer, err := arr[1].ToString()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		idle, err := arr[2].ToInt64()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		retryCount, err := arr[3].ToInt64()
		if err != nil {
			return &XPendingExtCmd{res: res, err: err}
		}
		val = append(val, XPendingExt{
			ID:         id,
			Consumer:   consumer,
			Idle:       time.Duration(idle) * time.Millisecond,
			RetryCount: retryCount,
		})
	}
	return &XPendingExtCmd{res: res, val: val, err: err}
}

func (cmd *XPendingExtCmd) SetVal(val []XPendingExt) {
	cmd.val = val
}

func (cmd *XPendingExtCmd) Val() []XPendingExt {
	return cmd.val
}

func (cmd *XPendingExtCmd) Err() error {
	return cmd.err
}

func (cmd *XPendingExtCmd) Result() ([]XPendingExt, error) {
	return cmd.val, cmd.err
}

func (cmd *XPendingExtCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XAutoClaimCmd struct {
	res   rueidis.RedisResult
	start string
	val   []rueidis.XRange
	err   error
}

func newXAutoClaimCmd(res rueidis.RedisResult) *XAutoClaimCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	if len(arr) != 2 {
		return &XAutoClaimCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	val, err := arr[1].AsXRangeSlice()
	if err != nil {
		return &XAutoClaimCmd{res: res, err: err}
	}
	return &XAutoClaimCmd{res: res, val: val, start: start, err: err}
}

func (cmd *XAutoClaimCmd) SetVal(val []rueidis.XRange, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimCmd) Val() (messages []rueidis.XRange, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimCmd) Err() error {
	return cmd.err
}

func (cmd *XAutoClaimCmd) Result() (messages []rueidis.XRange, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

func (cmd *XAutoClaimCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XAutoClaimJustIDCmd struct {
	res   rueidis.RedisResult
	start string
	val   []string
	err   error
}

func newXAutoClaimJustIDCmd(res rueidis.RedisResult) *XAutoClaimJustIDCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	if len(arr) != 2 {
		return &XAutoClaimJustIDCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	start, err := arr[0].ToString()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	val, err := arr[1].AsStrSlice()
	if err != nil {
		return &XAutoClaimJustIDCmd{res: res, err: err}
	}
	return &XAutoClaimJustIDCmd{res: res, val: val, start: start, err: err}
}

func (cmd *XAutoClaimJustIDCmd) SetVal(val []string, start string) {
	cmd.val = val
	cmd.start = start
}

func (cmd *XAutoClaimJustIDCmd) Val() (ids []string, start string) {
	return cmd.val, cmd.start
}

func (cmd *XAutoClaimJustIDCmd) Err() error {
	return cmd.err
}

func (cmd *XAutoClaimJustIDCmd) Result() (ids []string, start string, err error) {
	return cmd.val, cmd.start, cmd.err
}

func (cmd *XAutoClaimJustIDCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XInfoGroup struct {
	Name            string
	Consumers       int64
	Pending         int64
	LastDeliveredID string
}

type XInfoGroupsCmd struct {
	res rueidis.RedisResult
	val []XInfoGroup
	err error
}

func newXInfoGroupsCmd(res rueidis.RedisResult) *XInfoGroupsCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoGroupsCmd{res: res, err: err}
	}
	groupInfos := make([]XInfoGroup, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &XInfoGroupsCmd{res: res, err: err}
		}
		if len(info) != 8 {
			return &XInfoGroupsCmd{res: res, err: fmt.Errorf("got %d, wanted 8", len(arr))}
		}
		var group XInfoGroup
		for i, j := 0, 1; i < 8; i, j = i+2, j+2 {
			key, err := info[i].ToString()
			if err != nil {
				return &XInfoGroupsCmd{res: res, err: err}
			}
			val, err := info[j].ToString()
			if err != nil {
				return &XInfoGroupsCmd{res: res, err: err}
			}
			switch key {
			case "name":
				group.Name = val
			case "consumers":
				group.Consumers, err = strconv.ParseInt(val, 0, 64)
				if err != nil {
					return &XInfoGroupsCmd{res: res, err: err}
				}
			case "pending":
				group.Pending, err = strconv.ParseInt(val, 0, 64)
				if err != nil {
					return &XInfoGroupsCmd{res: res, err: err}
				}
			case "last-delivered-id":
				group.LastDeliveredID = val
			default:
				return &XInfoGroupsCmd{res: res, err: fmt.Errorf("unexpected content %s", key)}
			}
		}

		groupInfos = append(groupInfos, group)
	}
	return &XInfoGroupsCmd{res: res, val: groupInfos, err: err}
}

func (cmd *XInfoGroupsCmd) SetVal(val []XInfoGroup) {
	cmd.val = val
}

func (cmd *XInfoGroupsCmd) Val() []XInfoGroup {
	return cmd.val
}

func (cmd *XInfoGroupsCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoGroupsCmd) Result() ([]XInfoGroup, error) {
	return cmd.val, cmd.err
}

func (cmd *XInfoGroupsCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XInfoStream struct {
	Length          int64
	RadixTreeKeys   int64
	RadixTreeNodes  int64
	Groups          int64
	LastGeneratedID string
	FirstEntry      XMessage
	LastEntry       XMessage
}
type XInfoStreamCmd struct {
	res rueidis.RedisResult
	val XInfoStream
	err error
}

func newXInfoStreamCmd(res rueidis.RedisResult) *XInfoStreamCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoStreamCmd{res: res, err: err}
	}
	if len(arr) != 14 {
		return &XInfoStreamCmd{res: res, err: fmt.Errorf("got %d, wanted 14", len(arr))}
	}
	var val XInfoStream
	for i, j := 0, 1; i < 14; i, j = i+2, j+2 {
		key, err := arr[i].ToString()
		if err != nil {
			return &XInfoStreamCmd{res: res, err: err}
		}
		switch key {
		case "length":
			val.Length, err = arr[j].ToInt64()
		case "radix-tree-keys":
			val.RadixTreeKeys, err = arr[j].ToInt64()
		case "radix-tree-nodes":
			val.RadixTreeNodes, err = arr[j].ToInt64()
		case "groups":
			val.Groups, err = arr[j].ToInt64()
		case "last-generated-id":
			val.LastGeneratedID, err = arr[j].ToString()
		case "first-entry":
			r, err := arr[j].AsXRange()
			if rueidis.IsRedisNil(err) {
				err = nil
			}
			val.FirstEntry = newXMessage(r)
		case "last-entry":
			r, err := arr[j].AsXRange()
			if rueidis.IsRedisNil(err) {
				err = nil
			}
			val.LastEntry = newXMessage(r)
		default:
			err = fmt.Errorf("unexpected content %s", key)
		}
		if err != nil {
			return &XInfoStreamCmd{res: res, err: err}
		}
	}
	return &XInfoStreamCmd{res: res, val: val, err: err}
}

func (cmd *XInfoStreamCmd) SetVal(val XInfoStream) {
	cmd.val = val
}

func (cmd *XInfoStreamCmd) Val() XInfoStream {
	return cmd.val
}

func (cmd *XInfoStreamCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoStreamCmd) Result() (XInfoStream, error) {
	return cmd.val, cmd.err
}

func (cmd *XInfoStreamCmd) String() (string, error) {
	return cmd.res.ToString()
}

type XInfoStreamConsumerPending struct {
	ID            string
	DeliveryTime  time.Time
	DeliveryCount int64
}

type XInfoStreamGroupPending struct {
	ID            string
	Consumer      string
	DeliveryTime  time.Time
	DeliveryCount int64
}

type XInfoStreamConsumer struct {
	Name     string
	SeenTime time.Time
	PelCount int64
	Pending  []XInfoStreamConsumerPending
}

type XInfoStreamGroup struct {
	Name            string
	LastDeliveredID string
	PelCount        int64
	Pending         []XInfoStreamGroupPending
	Consumers       []XInfoStreamConsumer
}

type XInfoStreamFull struct {
	Length          int64
	RadixTreeKeys   int64
	RadixTreeNodes  int64
	LastGeneratedID string
	Entries         []XMessage
	Groups          []XInfoStreamGroup
}

type XInfoStreamFullCmd struct {
	res rueidis.RedisResult
	val XInfoStreamFull
	err error
}

func newXInfoStreamFullCmd(res rueidis.RedisResult) *XInfoStreamFullCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoStreamFullCmd{res: res, err: err}
	}
	if len(arr) != 12 {
		return &XInfoStreamFullCmd{res: res, err: fmt.Errorf("got %d, wanted 12", len(arr))}
	}
	var val XInfoStreamFull
	for i, j := 0, 1; i < 12; i, j = i+2, j+2 {
		key, err := arr[i].ToString()
		if err != nil {
			return &XInfoStreamFullCmd{res: res, err: err}
		}
		switch key {
		case "length":
			val.Length, err = arr[j].ToInt64()
		case "radix-tree-keys":
			val.RadixTreeKeys, err = arr[j].ToInt64()
		case "radix-tree-nodes":
			val.RadixTreeNodes, err = arr[j].ToInt64()
		case "last-generated-id":
			val.LastGeneratedID, err = arr[j].ToString()
		case "entries":
			ranges, err := arr[j].AsXRangeSlice()
			if err != nil {
				return &XInfoStreamFullCmd{res: res, err: err}
			}
			val.Entries = make([]XMessage, 0, len(ranges))
			for _, r := range ranges {
				val.Entries = append(val.Entries, newXMessage(r))
			}
		case "groups":
			val.Groups, err = readStreamGroups(arr[j])
		default:
			err = fmt.Errorf("unexpected content %s", key)
		}
		if err != nil {
			return &XInfoStreamFullCmd{res: res, err: err}
		}
	}

	return &XInfoStreamFullCmd{res: res, val: val, err: err}
}

func (cmd *XInfoStreamFullCmd) SetVal(val XInfoStreamFull) {
	cmd.val = val
}

func (cmd *XInfoStreamFullCmd) Val() XInfoStreamFull {
	return cmd.val
}

func (cmd *XInfoStreamFullCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoStreamFullCmd) Result() (XInfoStreamFull, error) {
	return cmd.val, cmd.err
}

func (cmd *XInfoStreamFullCmd) String() (string, error) {
	return cmd.res.ToString()
}

func readStreamGroups(res rueidis.RedisMessage) ([]XInfoStreamGroup, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	groups := make([]XInfoStreamGroup, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return nil, err
		}
		if len(info) != 10 {
			return nil, fmt.Errorf("got %d, wanted 10", len(arr))
		}
		var group XInfoStreamGroup
		for i, j := 0, 1; i < 10; i, j = i+2, j+2 {
			key, err := info[i].ToString()
			if err != nil {
				return nil, err
			}
			switch key {
			case "name":
				group.Name, err = info[j].ToString()
			case "last-delivered-id":
				group.LastDeliveredID, err = info[j].ToString()
			case "pel-count":
				group.PelCount, err = info[j].ToInt64()
			case "pending":
				group.Pending, err = readXInfoStreamGroupPending(info[j])
			case "consumers":
				group.Consumers, err = readXInfoStreamConsumers(info[j])
			default:
				err = fmt.Errorf("unexpected content %s", key)
			}
			if err != nil {
				return nil, err
			}
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func readXInfoStreamGroupPending(res rueidis.RedisMessage) ([]XInfoStreamGroupPending, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	pending := make([]XInfoStreamGroupPending, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return nil, err
		}
		if len(info) != 4 {
			return nil, fmt.Errorf("got %d, wanted 4", len(arr))
		}
		var p XInfoStreamGroupPending
		p.ID, err = info[0].ToString()
		if err != nil {
			return nil, err
		}
		p.Consumer, err = info[1].ToString()
		if err != nil {
			return nil, err
		}
		delivery, err := info[2].ToInt64()
		if err != nil {
			return nil, err
		}
		p.DeliveryTime = time.Unix(delivery/1000, delivery%1000*int64(time.Millisecond))
		p.DeliveryCount, err = info[3].ToInt64()
		if err != nil {
			return nil, err
		}
		pending = append(pending, p)
	}
	return pending, nil
}

func readXInfoStreamConsumers(res rueidis.RedisMessage) ([]XInfoStreamConsumer, error) {
	arr, err := res.ToArray()
	if err != nil {
		return nil, err
	}
	consumer := make([]XInfoStreamConsumer, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return nil, err
		}
		if len(info) != 8 {
			return nil, fmt.Errorf("got %d, wanted 8", len(arr))
		}
		var c XInfoStreamConsumer
		for i, j := 0, 1; i < 8; i, j = i+2, j+2 {
			cKey, err := info[i].ToString()
			if err != nil {
				return nil, err
			}
			switch cKey {
			case "name":
				c.Name, err = info[j].ToString()
			case "seen-time":
				seen, err := info[j].ToInt64()
				if err != nil {
					return nil, err
				}
				c.SeenTime = time.Unix(seen/1000, seen%1000*int64(time.Millisecond))
			case "pel-count":
				c.PelCount, err = info[j].ToInt64()
			case "pending":
				pending, err := info[j].ToArray()
				if err != nil {
					return nil, err
				}
				c.Pending = make([]XInfoStreamConsumerPending, 0, len(pending))
				for _, v := range pending {
					pendingInfo, err := v.ToArray()
					if err != nil {
						return nil, err
					}
					if len(pendingInfo) != 3 {
						return nil, fmt.Errorf("got %d, wanted 3", len(arr))
					}
					var p XInfoStreamConsumerPending
					p.ID, err = pendingInfo[0].ToString()
					if err != nil {
						return nil, err
					}
					delivery, err := pendingInfo[1].ToInt64()
					if err != nil {
						return nil, err
					}
					p.DeliveryTime = time.Unix(delivery/1000, delivery%1000*int64(time.Millisecond))
					p.DeliveryCount, err = pendingInfo[2].ToInt64()
					if err != nil {
						return nil, err
					}
					c.Pending = append(c.Pending, p)
				}
			default:
				err = fmt.Errorf("unexpected content %s", cKey)
			}
			if err != nil {
				return nil, err
			}
		}
		consumer = append(consumer, c)
	}
	return consumer, nil
}

type XInfoConsumer struct {
	Name    string
	Pending int64
	Idle    int64
}
type XInfoConsumersCmd struct {
	res rueidis.RedisResult
	val []XInfoConsumer
	err error
}

func newXInfoConsumersCmd(res rueidis.RedisResult) *XInfoConsumersCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &XInfoConsumersCmd{res: res, err: err}
	}
	val := make([]XInfoConsumer, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &XInfoConsumersCmd{res: res, err: err}
		}
		if len(info) != 6 {
			return &XInfoConsumersCmd{res: res, err: fmt.Errorf("got %d, wanted 6", len(arr))}
		}
		var consumer XInfoConsumer
		for i, j := 0, 1; i < 6; i, j = i+2, j+2 {
			key, err := info[i].ToString()
			if err != nil {
				return &XInfoConsumersCmd{res: res, err: err}
			}
			val, err := info[j].ToString()
			if err != nil {
				return &XInfoConsumersCmd{res: res, err: err}
			}
			switch key {
			case "name":
				consumer.Name = val
			case "pending":
				consumer.Pending, err = strconv.ParseInt(val, 0, 64)
				if err != nil {
					return &XInfoConsumersCmd{res: res, err: err}
				}
			case "idle":
				consumer.Idle, err = strconv.ParseInt(val, 0, 64)
				if err != nil {
					return &XInfoConsumersCmd{res: res, err: err}
				}
			default:
				return &XInfoConsumersCmd{res: res, err: fmt.Errorf("unexpected content %s", key)}
			}
		}
		val = append(val, consumer)
	}
	return &XInfoConsumersCmd{res: res, val: val, err: err}
}

func (cmd *XInfoConsumersCmd) SetVal(val []XInfoConsumer) {
	cmd.val = val
}

func (cmd *XInfoConsumersCmd) Val() []XInfoConsumer {
	return cmd.val
}

func (cmd *XInfoConsumersCmd) Err() error {
	return cmd.err
}

func (cmd *XInfoConsumersCmd) Result() ([]XInfoConsumer, error) {
	return cmd.val, cmd.err
}

func (cmd *XInfoConsumersCmd) String() (string, error) {
	return cmd.res.ToString()
}

// Z represents sorted set member.
type Z struct {
	Score  float64
	Member interface{}
}

// ZWithKey represents sorted set member including the name of the key where it was popped.
type ZWithKey struct {
	Z
	Key string
}

// ZStore is used as an arg to ZInter/ZInterStore and ZUnion/ZUnionStore.
type ZStore struct {
	Keys    []string
	Weights []int64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type ZWithKeyCmd struct {
	res rueidis.RedisResult
	val ZWithKey
	err error
}

func newZWithKeyCmd(res rueidis.RedisResult) *ZWithKeyCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &ZWithKeyCmd{res: res, err: err}
	}
	if len(arr) != 3 {
		return &ZWithKeyCmd{res: res, err: fmt.Errorf("got %d, wanted 3", len(arr))}
	}
	val := ZWithKey{}
	val.Key, err = arr[0].ToString()
	if err != nil {
		return &ZWithKeyCmd{res: res, err: err}
	}
	val.Member, err = arr[0].ToString()
	if err != nil {
		return &ZWithKeyCmd{res: res, err: err}
	}
	val.Score, err = arr[0].AsFloat64()
	if err != nil {
		return &ZWithKeyCmd{res: res, err: err}
	}
	return &ZWithKeyCmd{res: res, val: val, err: err}
}

func (cmd *ZWithKeyCmd) SetVal(val ZWithKey) {
	cmd.val = val
}

func (cmd *ZWithKeyCmd) Val() ZWithKey {
	return cmd.val
}

func (cmd *ZWithKeyCmd) Err() error {
	return cmd.err
}

func (cmd *ZWithKeyCmd) Result() (ZWithKey, error) {
	return cmd.val, cmd.err
}

func (cmd *ZWithKeyCmd) String() (string, error) {
	return cmd.res.ToString()
}

type TimeCmd struct {
	res rueidis.RedisResult
	val time.Time
	err error
}

func newTimeCmd(res rueidis.RedisResult) *TimeCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &TimeCmd{res: res, err: err}
	}
	if len(arr) != 2 {
		return &TimeCmd{res: res, err: fmt.Errorf("got %d, wanted 2", len(arr))}
	}
	sec, err := arr[0].AsInt64()
	if err != nil {
		return &TimeCmd{res: res, err: err}
	}
	microSec, err := arr[1].AsInt64()
	if err != nil {
		return &TimeCmd{res: res, err: err}
	}
	return &TimeCmd{res: res, val: time.Unix(sec, microSec*1000), err: err}
}

func (cmd *TimeCmd) SetVal(val time.Time) {
	cmd.val = val
}

func (cmd *TimeCmd) Val() time.Time {
	return cmd.val
}

func (cmd *TimeCmd) Err() error {
	return cmd.err
}

func (cmd *TimeCmd) Result() (time.Time, error) {
	return cmd.val, cmd.err
}

func (cmd *TimeCmd) String() (string, error) {
	return cmd.res.ToString()
}

type ClusterNode struct {
	ID   string
	Addr string
}

type ClusterSlot struct {
	Start int64
	End   int64
	Nodes []ClusterNode
}

type ClusterSlotsCmd struct {
	res rueidis.RedisResult
	val []ClusterSlot
	err error
}

func newClusterSlotsCmd(res rueidis.RedisResult) *ClusterSlotsCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &ClusterSlotsCmd{res: res, err: err}
	}
	val := make([]ClusterSlot, 0, len(arr))
	for _, v := range arr {
		slot, err := v.ToArray()
		if err != nil {
			return &ClusterSlotsCmd{res: res, err: err}
		}
		if len(slot) < 2 {
			return &ClusterSlotsCmd{res: res, err: fmt.Errorf("got %d, excpected atleast 2", len(slot))}
		}
		start, err := slot[0].ToInt64()
		if err != nil {
			return &ClusterSlotsCmd{res: res, err: err}
		}
		end, err := slot[1].ToInt64()
		if err != nil {
			return &ClusterSlotsCmd{res: res, err: err}
		}
		nodes := make([]ClusterNode, len(slot)-2)
		for i, j := 2, 0; i < len(nodes); i, j = i+1, j+1 {
			node, err := slot[i].ToArray()
			if err != nil {
				return &ClusterSlotsCmd{res: res, err: err}
			}
			if len(node) != 2 && len(node) != 3 {
				return &ClusterSlotsCmd{res: res, err: fmt.Errorf("got %d, expected 2 or 3", len(node))}
			}
			ip, err := node[0].ToString()
			if err != nil {
				return &ClusterSlotsCmd{res: res, err: err}
			}
			port, err := node[1].ToInt64()
			if err != nil {
				return &ClusterSlotsCmd{res: res, err: err}
			}
			nodes[j].Addr = net.JoinHostPort(ip, strconv.FormatInt(port, 10))
			if len(node) == 3 {
				id, err := node[2].ToString()
				if err != nil {
					return &ClusterSlotsCmd{res: res, err: err}
				}
				nodes[j].ID = id
			}
		}
		val = append(val, ClusterSlot{
			Start: start,
			End:   end,
			Nodes: nodes,
		})
	}
	return &ClusterSlotsCmd{res: res, val: val, err: err}
}

func (cmd *ClusterSlotsCmd) SetVal(val []ClusterSlot) {
	cmd.val = val
}

func (cmd *ClusterSlotsCmd) Val() []ClusterSlot {
	return cmd.val
}

func (cmd *ClusterSlotsCmd) Err() error {
	return cmd.err
}

func (cmd *ClusterSlotsCmd) Result() ([]ClusterSlot, error) {
	return cmd.val, cmd.err
}

func (cmd *ClusterSlotsCmd) String() (string, error) {
	return cmd.res.ToString()
}

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoPosCmd struct {
	res rueidis.RedisResult
	val []*GeoPos
	err error
}

func newGeoPosCmd(res rueidis.RedisResult) *GeoPosCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &GeoPosCmd{res: res, err: err}
	}
	val := make([]*GeoPos, 0, len(arr))
	for _, v := range arr {
		loc, err := v.ToArray()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				val = append(val, nil)
				continue
			}
			return &GeoPosCmd{res: res, err: err}
		}
		if len(loc) != 2 {
			return &GeoPosCmd{res: res, err: fmt.Errorf("got %d, expected 2", len(loc))}
		}
		long, err := loc[0].AsFloat64()
		if err != nil {
			return &GeoPosCmd{res: res, err: err}
		}
		lat, err := loc[1].AsFloat64()
		if err != nil {
			return &GeoPosCmd{res: res, err: err}
		}
		val = append(val, &GeoPos{
			Longitude: long,
			Latitude:  lat,
		})
	}
	return &GeoPosCmd{res: res, val: val, err: err}
}

func (cmd *GeoPosCmd) SetVal(val []*GeoPos) {
	cmd.val = val
}

func (cmd *GeoPosCmd) Val() []*GeoPos {
	return cmd.val
}

func (cmd *GeoPosCmd) Err() error {
	return cmd.err
}

func (cmd *GeoPosCmd) Result() ([]*GeoPos, error) {
	return cmd.val, cmd.err
}

func (cmd *GeoPosCmd) String() (string, error) {
	return cmd.res.ToString()
}

type GeoLocationCmd struct {
	res rueidis.RedisResult
	val []GeoLocation
	err error
}

func newGeoLocationCmd(res rueidis.RedisResult, withDist, withGeoHash, withCoord bool) *GeoLocationCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &GeoLocationCmd{res: res, err: err}
	}
	val := make([]GeoLocation, 0, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &GeoLocationCmd{res: res, err: err}
		}
		var loc GeoLocation
		var i int
		loc.Name, err = info[i].ToString()
		i++
		if err != nil {
			return &GeoLocationCmd{res: res, err: err}
		}
		if withDist {
			loc.Dist, err = info[i].AsFloat64()
			i++
			if err != nil {
				return &GeoLocationCmd{res: res, err: err}
			}
		}
		if withGeoHash {
			loc.GeoHash, err = info[i].AsInt64()
			i++
			if err != nil {
				return &GeoLocationCmd{res: res, err: err}
			}
		}
		if withCoord {
			cord, err := info[i].ToArray()
			if err != nil {
				return &GeoLocationCmd{res: res, err: err}
			}
			if len(cord) != 2 {
				return &GeoLocationCmd{res: res, err: fmt.Errorf("got %d, expected 2", len(info))}
			}
			loc.Longitude, err = cord[0].AsFloat64()
			if err != nil {
				return &GeoLocationCmd{res: res, err: err}
			}
			loc.Latitude, err = cord[1].AsFloat64()
			if err != nil {
				return &GeoLocationCmd{res: res, err: err}
			}
		}
		val = append(val, loc)
	}
	return &GeoLocationCmd{res: res, val: val, err: err}
}

func (cmd *GeoLocationCmd) SetVal(val []GeoLocation) {
	cmd.val = val
}

func (cmd *GeoLocationCmd) Val() []GeoLocation {
	return cmd.val
}

func (cmd *GeoLocationCmd) Err() error {
	return cmd.err
}

func (cmd *GeoLocationCmd) Result() ([]GeoLocation, error) {
	return cmd.val, cmd.err
}

func (cmd *GeoLocationCmd) String() (string, error) {
	return cmd.res.ToString()
}

type CommandInfo struct {
	Name        string
	Arity       int8
	Flags       []string
	ACLFlags    []string
	FirstKeyPos int8
	LastKeyPos  int8
	StepCount   int8
	ReadOnly    bool
}

type CommandsInfoCmd struct {
	res rueidis.RedisResult
	val map[string]CommandInfo
	err error
}

func newCommandsInfoCmd(res rueidis.RedisResult) *CommandsInfoCmd {
	arr, err := res.ToArray()
	if err != nil {
		return &CommandsInfoCmd{res: res, err: err}
	}
	val := make(map[string]CommandInfo, len(arr))
	for _, v := range arr {
		info, err := v.ToArray()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		switch len(info) {
		case 6, 7:
		default:
			return &CommandsInfoCmd{res: res, err: fmt.Errorf("got %d, wanted 7", len(info))}
		}
		var cmd CommandInfo
		cmd.Name, err = info[0].ToString()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		arity, err := info[1].ToInt64()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		cmd.Arity = int8(arity)
		cmd.Flags, err = info[2].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				cmd.Flags = []string{}
			} else {
				return &CommandsInfoCmd{res: res, err: err}
			}
		}
		firstKeyPos, err := info[3].ToInt64()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		cmd.FirstKeyPos = int8(firstKeyPos)
		lastKeyPos, err := info[4].ToInt64()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		cmd.LastKeyPos = int8(lastKeyPos)
		stepCount, err := info[5].ToInt64()
		if err != nil {
			return &CommandsInfoCmd{res: res, err: err}
		}
		cmd.StepCount = int8(stepCount)
		for _, flag := range cmd.Flags {
			if flag == "readonly" {
				cmd.ReadOnly = true
				break
			}
		}
		if len(arr) == 6 {
			val[cmd.Name] = cmd
			continue
		}
		cmd.ACLFlags, err = info[6].AsStrSlice()
		if err != nil {
			if rueidis.IsRedisNil(err) {
				cmd.ACLFlags = []string{}
			} else {
				return &CommandsInfoCmd{res: res, err: err}
			}
		}
		val[cmd.Name] = cmd
	}
	return &CommandsInfoCmd{res: res, val: val, err: err}
}

func (cmd *CommandsInfoCmd) SetVal(val map[string]CommandInfo) {
	cmd.val = val
}

func (cmd *CommandsInfoCmd) Val() map[string]CommandInfo {
	return cmd.val
}

func (cmd *CommandsInfoCmd) Err() error {
	return cmd.err
}

func (cmd *CommandsInfoCmd) Result() (map[string]CommandInfo, error) {
	return cmd.val, cmd.err
}

func (cmd *CommandsInfoCmd) String() (string, error) {
	return cmd.res.ToString()
}

type Sort struct {
	By            string
	Offset, Count int64
	Get           []string
	Order         string
	Alpha         bool
}

// SetArgs provides arguments for the SetArgs function.
type SetArgs struct {
	// Mode can be `NX` or `XX` or empty.
	Mode string

	// Zero `TTL` or `Expiration` means that the key has no expiration time.
	TTL      time.Duration
	ExpireAt time.Time

	// When Get is true, the command returns the old value stored at key, or nil when key did not exist.
	Get bool

	// KeepTTL is a Redis KEEPTTL option to keep existing TTL, it requires your redis-server version >= 6.0,
	// otherwise you will receive an error: (error) ERR syntax error.
	KeepTTL bool
}

type BitCount struct {
	Start, End int64
}

//type BitPos struct {
//	BitCount
//	Byte bool
//}

type BitFieldArg struct {
	Encoding string
	Offset   int64
}

type BitField struct {
	Get       *BitFieldArg
	Set       *BitFieldArg
	IncrBy    *BitFieldArg
	Increment int64
	Overflow  string
}

type LPosArgs struct {
	Rank, MaxLen int64
}

// Note: MaxLen/MaxLenApprox and MinID are in conflict, only one of them can be used.
type XAddArgs struct {
	Stream     string
	NoMkStream bool
	MaxLen     int64 // MAXLEN N

	MinID string
	// Approx causes MaxLen and MinID to use "~" matcher (instead of "=").
	Approx bool
	Limit  int64
	ID     string
	Values interface{}
}

type XReadArgs struct {
	Streams []string // list of streams
	Count   int64
	Block   time.Duration
}

type XReadGroupArgs struct {
	Group    string
	Consumer string
	Streams  []string // list of streams
	Count    int64
	Block    time.Duration
	NoAck    bool
}

type XPendingExtArgs struct {
	Stream   string
	Group    string
	Idle     time.Duration
	Start    string
	End      string
	Count    int64
	Consumer string
}

type XClaimArgs struct {
	Stream   string
	Group    string
	Consumer string
	MinIdle  time.Duration
	Messages []string
}

type XAutoClaimArgs struct {
	Stream   string
	Group    string
	MinIdle  time.Duration
	Start    string
	Count    int64
	Consumer string
}

type XMessage struct {
	ID     string
	Values map[string]interface{}
}

// Note: The GT, LT and NX options are mutually exclusive.
type ZAddArgs struct {
	NX      bool
	XX      bool
	LT      bool
	GT      bool
	Ch      bool
	Incr    bool
	Members []Z
}

// ZRangeArgs is all the options of the ZRange command.
// In version> 6.2.0, you can replace the(cmd):
//		ZREVRANGE,
//		ZRANGEBYSCORE,
//		ZREVRANGEBYSCORE,
//		ZRANGEBYLEX,
//		ZREVRANGEBYLEX.
// Please pay attention to your redis-server version.
//
// Rev, ByScore, ByLex and Offset+Count options require redis-server 6.2.0 and higher.
type ZRangeArgs struct {
	Key string

	Start interface{}
	Stop  interface{}

	// The ByScore and ByLex options are mutually exclusive.
	ByScore bool
	ByLex   bool

	Rev bool

	// limit offset count.
	Offset int64
	Count  int64
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

// GeoRadiusQuery is used with GeoRadius to query geospatial index.
type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int64
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

// GeoSearchQuery is used for GEOSearch/GEOSearchStore command query.
type GeoSearchQuery struct {
	Member string

	// Latitude and Longitude when using FromLonLat option.
	Longitude float64
	Latitude  float64

	// Distance and unit when using ByRadius option.
	// Can use m, km, ft, or mi. Default is km.
	Radius     float64
	RadiusUnit string

	// Height, width and unit when using ByBox option.
	// Can be m, km, ft, or mi. Default is km.
	BoxWidth  float64
	BoxHeight float64
	BoxUnit   string

	// Can be ASC or DESC. Default is no sort order.
	Sort     string
	Count    int64
	CountAny bool
}

type GeoSearchLocationQuery struct {
	GeoSearchQuery

	WithCoord bool
	WithDist  bool
	WithHash  bool
}

type GeoSearchStoreQuery struct {
	GeoSearchQuery

	// When using the StoreDist option, the command stores the items in a
	// sorted set populated with their distance from the center of the circle or box,
	// as a floating-point number, in the same unit specified for that shape.
	StoreDist bool
}

func (q *GeoRadiusQuery) args() []string {
	args := make([]string, 0, 2)
	args = append(args, strconv.FormatFloat(q.Radius, 'f', -1, 64))
	if q.Unit != "" {
		args = append(args, q.Unit)
	} else {
		args = append(args, "KM")
	}
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithGeoHash {
		args = append(args, "WITHASH")
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Store != "" {
		args = append(args, "STORE")
		args = append(args, q.Store)
	}
	if q.StoreDist != "" {
		args = append(args, "STOREDIST")
		args = append(args, q.StoreDist)
	}
	return args
}

func (q *GeoSearchQuery) args() []string {
	args := make([]string, 0, 2)
	if q.Member != "" {
		args = append(args, "FROMMEMBER", q.Member)
	} else {
		args = append(args, "FROMLONLAT", strconv.FormatFloat(q.Longitude, 'f', -1, 64), strconv.FormatFloat(q.Latitude, 'f', -1, 64))
	}
	if q.Radius > 0 {
		if q.RadiusUnit == "" {
			q.RadiusUnit = "KM"
		}
		args = append(args, "BYRADIUS", strconv.FormatFloat(q.Radius, 'f', -1, 64), q.RadiusUnit)
	} else {
		if q.BoxUnit == "" {
			q.BoxUnit = "KM"
		}
		args = append(args, "BYXBOX", strconv.FormatFloat(q.BoxWidth, 'f', -1, 64), strconv.FormatFloat(q.BoxHeight, 'f', -1, 64), q.BoxUnit)
	}
	if q.Sort != "" {
		args = append(args, q.Sort)
	}
	if q.Count > 0 {
		args = append(args, "COUNT", strconv.FormatInt(q.Count, 10))
		if q.CountAny {
			args = append(args, "ANY")
		}
	}
	return args
}

func (q *GeoSearchLocationQuery) args() []string {
	args := q.GeoSearchQuery.args()
	if q.WithCoord {
		args = append(args, "WITHCOORD")
	}
	if q.WithDist {
		args = append(args, "WITHDIST")
	}
	if q.WithHash {
		args = append(args, "WITHASH")
	}
	return args
}

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		// too small, truncate too 1ms
		return 1
	}
	return int64(dur / time.Millisecond)
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		// too small ,truncate too 1s
		return 1
	}
	return int64(dur / time.Second)
}