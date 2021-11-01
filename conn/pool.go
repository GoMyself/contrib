package conn

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	statusCmdType = iota + 1
	stringCmdType
	intCmdType
	floatCmdType
	boolCmdType
)

const (
	bitOpAnd = iota + 1
	bitOpOr
	bitOpXor
)

var (
	errWrongArguments       = errors.New("wrong number of arguments")
	errShardPoolUnSupported = errors.New("shard pool didn't support the command")
	errCrossMultiShards     = errors.New("cross multi shards was not allowed")
)

type ConnFactory interface {
	getSlaveConn(key ...string) (*redis.Client, error)
	getMasterConn(key ...string) (*redis.Client, error)
	close()
}

func newErrorStringIntMapCmd(err error) *redis.StringIntMapCmd {
	cmd := &redis.StringIntMapCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorBoolSliceCmd(err error) *redis.BoolSliceCmd {
	cmd := &redis.BoolSliceCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorIntCmd(err error) *redis.IntCmd {
	cmd := &redis.IntCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorFloatCmd(err error) *redis.FloatCmd {
	cmd := &redis.FloatCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorSliceCmd(err error) *redis.SliceCmd {
	cmd := &redis.SliceCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorStringStringMapCmd(err error) *redis.StringStringMapCmd {
	cmd := &redis.StringStringMapCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorIntSliceCmd(err error) *redis.IntSliceCmd {
	cmd := &redis.IntSliceCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorDurationCmd(err error) *redis.DurationCmd {
	cmd := &redis.DurationCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorBoolCmd(err error) *redis.BoolCmd {
	cmd := &redis.BoolCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorStatusCmd(err error) *redis.StatusCmd {
	cmd := &redis.StatusCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorStringCmd(err error) *redis.StringCmd {
	cmd := &redis.StringCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorStringSliceCmd(err error) *redis.StringSliceCmd {
	cmd := &redis.StringSliceCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorStringStructMapCmd(err error) *redis.StringStructMapCmd {
	cmd := &redis.StringStructMapCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorZSliceCmd(err error) *redis.ZSliceCmd {
	cmd := &redis.ZSliceCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorScanCmd(err error) *redis.ScanCmd {
	cmd := &redis.ScanCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorGeoCmd(err error) *redis.GeoPosCmd {
	cmd := &redis.GeoPosCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorGeoLocationCmd(err error) *redis.GeoLocationCmd {
	cmd := &redis.GeoLocationCmd{}
	cmd.SetErr(err)
	return cmd
}

func newErrorCmd(err error) *redis.Cmd {
	cmd := &redis.Cmd{}
	cmd.SetErr(err)
	return cmd
}

type Pool struct {
	connFactory ConnFactory
}

func NewHA(cfg *HAConfig) (*Pool, error) {
	factory, err := NewHAConnFactory(cfg)
	if err != nil {
		return nil, err
	}
	return &Pool{
		connFactory: factory,
	}, nil
}

func NewShard(cfg *ShardConfig) (*Pool, error) {
	factory, err := NewShardConnFactory(cfg)
	if err != nil {
		return nil, err
	}
	return &Pool{
		connFactory: factory,
	}, nil
}

func (p *Pool) Close() {
	p.connFactory.close()
}

func (p *Pool) WithMaster(key ...string) (*redis.Client, error) {
	return p.connFactory.getMasterConn(key...)
}

func (p *Pool) Pipeline() redis.Pipeliner {
	//if _, ok := p.connFactory.(*ShardConnFactory); ok {
	//	return nil, errShardPoolUnSupported
	//}
	conn, _ := p.connFactory.getMasterConn()
	return conn.Pipeline()
}

func (p *Pool) Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return nil, errShardPoolUnSupported
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.Pipelined(ctx, fn)
}

func (p *Pool) TxPipeline() redis.Pipeliner {
	//if _, ok := p.connFactory.(*ShardConnFactory); ok {
	//	return nil, errShardPoolUnSupported
	//}
	conn, _ := p.connFactory.getMasterConn()
	return conn.TxPipeline()
}

func (p *Pool) TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return nil, errShardPoolUnSupported
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.TxPipelined(ctx, fn)
}

func (p *Pool) Ping() *redis.StatusCmd {
	// FIXME: use config to determine whether no key would access the master
	conn, err := p.connFactory.getMasterConn()
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.Ping(ctx)
}

func (p *Pool) Get(key string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.Get(ctx, key)
}

func (p *Pool) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.Set(ctx, key, value, expiration)
}

func (p *Pool) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.SetNX(ctx, key, value, expiration)
}

func (p *Pool) SetXX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.SetXX(ctx, key, value, expiration)
}

func (p *Pool) SetRange(key string, offset int64, value string) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SetRange(ctx, key, offset, value)
}

func (p *Pool) StrLen(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.StrLen(ctx, key)
}

func (p *Pool) Echo(message interface{}) *redis.StringCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStringCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.Echo(ctx, message)
}

func (p *Pool) Del(keys ...string) (int64, error) {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.Del(ctx, keys...).Result()
	}

	fn := func(factory *ShardConnFactory, keyList ...string) redis.Cmder {
		conn, _ := factory.getMasterConn(keyList[0])
		return conn.Del(ctx, keyList...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	return factory.doMultiIntCommand(fn, keys...)
}

func (p *Pool) Unlink(keys ...string) (int64, error) {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.Unlink(ctx, keys...).Result()
	}

	fn := func(factory *ShardConnFactory, keyList ...string) redis.Cmder {
		conn, _ := factory.getMasterConn(keyList[0])
		return conn.Unlink(ctx, keyList...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	return factory.doMultiIntCommand(fn, keys...)
}

func (p *Pool) Touch(keys ...string) (int64, error) {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.Touch(ctx, keys...).Result()
	}

	fn := func(factory *ShardConnFactory, keyList ...string) redis.Cmder {
		conn, _ := factory.getMasterConn(keyList[0])
		return conn.Touch(ctx, keyList...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	return factory.doMultiIntCommand(fn, keys...)
}

func (p *Pool) MGet(keys ...string) ([]interface{}, error) {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.MGet(ctx, keys...).Result()
	}

	fn := func(factory *ShardConnFactory, keyList ...string) redis.Cmder {
		conn, err := factory.getSlaveConn(keyList[0])
		if err != nil {
			return newErrorCmd(err)
		}
		return conn.MGet(ctx, keyList...)
	}

	factory := p.connFactory.(*ShardConnFactory)
	results := factory.doMultiKeys(fn, keys...)
	keyVals := make(map[string]interface{}, 0)
	for _, result := range results {
		vals, err := result.(*redis.SliceCmd).Result()
		if err != nil {
			return nil, err
		}
		for i, val := range vals {
			args := result.Args()
			keyVals[args[i+1].(string)] = val
		}
	}
	vals := make([]interface{}, len(keys))
	for i, key := range keys {
		vals[i] = nil
		if val, ok := keyVals[key]; ok {
			vals[i] = val
		}
	}
	return vals, nil
}

func appendArgs(dst, src []interface{}) []interface{} {
	if len(src) == 1 {
		switch v := src[0].(type) {
		case []string:
			for _, s := range v {
				dst = append(dst, s)
			}
			return dst
		case map[string]interface{}:
			for k, v := range v {
				dst = append(dst, k, v)
			}
			return dst
		}
	}

	dst = append(dst, src...)
	return dst
}

// MSet is like Set but accepts multiple values:
//   - MSet("key1", "value1", "key2", "value2")
//   - MSet([]string{"key1", "value1", "key2", "value2"})
//   - MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (p *Pool) MSet(values ...interface{}) *redis.StatusCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.MSet(ctx, values...)
	}

	args := make([]interface{}, 0, len(values))
	args = appendArgs(args, values)
	if len(args) == 0 || len(args)%2 != 0 {
		return newErrorStatusCmd(errWrongArguments)
	}
	factory := p.connFactory.(*ShardConnFactory)
	index2Values := make(map[uint32][]interface{})
	for i := 0; i < len(args); i += 2 {
		ind := factory.cfg.HashFn([]byte(fmt.Sprint(args[i]))) % uint32(len(factory.shards))
		if _, ok := index2Values[ind]; !ok {
			index2Values[ind] = make([]interface{}, 0)
		}
		index2Values[ind] = append(index2Values[ind], args[i], args[i+1])
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var result *redis.StatusCmd
	for ind, vals := range index2Values {
		wg.Add(1)
		conn, _ := factory.shards[ind].getMasterConn()
		go func(conn *redis.Client, vals ...interface{}) {
			defer wg.Done()
			status := conn.MSet(ctx, vals...)
			mu.Lock()
			if result == nil || status.Err() != nil {
				result = status
			}
			mu.Unlock()
		}(conn, vals...)
	}
	wg.Wait()
	return result
}

// MSetNX is like SetNX but accepts multiple values:
//   - MSetNX("key1", "value1", "key2", "value2")
//   - MSetNX([]string{"key1", "value1", "key2", "value2"})
//   - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (p *Pool) MSetNX(values ...interface{}) *redis.BoolCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.MSetNX(ctx, values...)
	}

	args := make([]interface{}, 0, len(values))
	args = appendArgs(args, values)
	if len(args) == 0 || len(args)%2 != 0 {
		return newErrorBoolCmd(errWrongArguments)
	}

	factory := p.connFactory.(*ShardConnFactory)
	keys := make([]string, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		keys[i/2] = fmt.Sprint(args[i])
	}
	if factory.isCrossMultiShards(keys...) {
		// we can't guarantee the atomic when msetnx across multi shards
		return newErrorBoolCmd(errCrossMultiShards)
	}
	conn, _ := factory.getMasterConn(keys[0])
	return conn.MSetNX(ctx, values...)
}

func (p *Pool) Dump(key string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.Dump(ctx, key)
}

func (p *Pool) Exists(keys ...string) (int64, error) {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.Exists(ctx, keys...).Result()
	}

	fn := func(factory *ShardConnFactory, keyList ...string) redis.Cmder {
		conn, err := factory.getSlaveConn(keyList[0])
		if err != nil {
			return newErrorCmd(err)
		}
		return conn.Exists(ctx, keyList...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	return factory.doMultiIntCommand(fn, keys...)
}

func (p *Pool) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.Expire(ctx, key, expiration)
}

func (p *Pool) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.ExpireAt(ctx, key, tm)
}

func (p *Pool) TTL(key string) *redis.DurationCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorDurationCmd(err)
	}
	return conn.TTL(ctx, key)
}

func (p *Pool) ObjectRefCount(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ObjectRefCount(ctx, key)
}

func (p *Pool) ObjectEncoding(key string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.ObjectEncoding(ctx, key)
}

func (p *Pool) ObjectIdleTime(key string) *redis.DurationCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorDurationCmd(err)
	}
	return conn.ObjectIdleTime(ctx, key)
}

func (p *Pool) Rename(key, newkey string) *redis.StatusCmd {
	if factory, ok := p.connFactory.(*ShardConnFactory); ok {
		if factory.isCrossMultiShards(key, newkey) {
			return newErrorStatusCmd(errCrossMultiShards)
		}
	}
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.Rename(ctx, key, newkey)
}

func (p *Pool) RenameNX(key, newkey string) *redis.BoolCmd {
	if factory, ok := p.connFactory.(*ShardConnFactory); ok {
		ind := factory.cfg.HashFn([]byte(key)) % uint32(len(factory.shards))
		newInd := factory.cfg.HashFn([]byte(newkey)) % uint32(len(factory.shards))
		if ind != newInd {
			return newErrorBoolCmd(errCrossMultiShards)
		}
	}
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.RenameNX(ctx, key, newkey)
}

func (p *Pool) Sort(key string, sort *redis.Sort) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.Sort(ctx, key, sort)
}

func (p *Pool) SortStore(key, store string, sort *redis.Sort) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.SortStore(ctx, key, store, sort)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(key, store) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SortStore(ctx, key, store, sort)
}

func (p *Pool) SortInterfaces(key string, sort *redis.Sort) *redis.SliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorSliceCmd(err)
	}
	return conn.SortInterfaces(ctx, key, sort)
}

func (p *Pool) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.Eval(ctx, script, keys, args...)
}

func (p *Pool) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.EvalSha(ctx, sha1, keys, args...)
}

func (p *Pool) ScriptExists(hashes ...string) *redis.BoolSliceCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorBoolSliceCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.ScriptExists(ctx, hashes...)
}

func (p *Pool) ScriptFlush() *redis.StatusCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStatusCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.ScriptFlush(ctx)
}

func (p *Pool) ScriptKill() *redis.StatusCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStatusCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.ScriptKill(ctx)
}

func (p *Pool) ScriptLoad(script string) *redis.StringCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStringCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.ScriptLoad(ctx, script)
}

func (p *Pool) DebugObject(key string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.DebugObject(ctx, key)
}

func (p *Pool) MemoryUsage(key string, samples ...int) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.MemoryUsage(ctx, key, samples...)
}

func (p *Pool) Publish(channel string, message interface{}) *(redis.IntCmd) {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorIntCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.Publish(ctx, channel, message)
}

func (p *Pool) PubSubChannels(pattern string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStringSliceCmd(errShardPoolUnSupported)
	}
	conn, err := p.connFactory.getSlaveConn()
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.PubSubChannels(ctx, pattern)
}

func (p *Pool) PubSubNumSub(channels ...string) *redis.StringIntMapCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStringIntMapCmd(errShardPoolUnSupported)
	}
	conn, err := p.connFactory.getSlaveConn()
	if err != nil {
		return newErrorStringIntMapCmd(err)
	}
	return conn.PubSubNumSub(ctx, channels...)
}

func (p *Pool) PubSubNumPat() *redis.IntCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorIntCmd(errShardPoolUnSupported)
	}
	conn, err := p.connFactory.getSlaveConn()
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.PubSubNumPat(ctx)
}

func (p *Pool) Type(key string) *redis.StatusCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.Type(ctx, key)
}

func (p *Pool) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorScanCmd(errShardPoolUnSupported)
	}
	conn, err := p.connFactory.getMasterConn()
	if err != nil {
		return newErrorScanCmd(err)
	}
	return conn.Scan(ctx, cursor, match, count)
}

func (p *Pool) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorScanCmd(err)
	}
	return conn.SScan(ctx, key, cursor, match, count)
}

func (p *Pool) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorScanCmd(err)
	}
	return conn.HScan(ctx, key, cursor, match, count)
}

func (p *Pool) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorScanCmd(err)
	}
	return conn.ZScan(ctx, key, cursor, match, count)
}

func (p *Pool) Append(key, value string) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.Append(ctx, key, value)
}

func (p *Pool) GetRange(key string, start, end int64) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.GetRange(ctx, key, start, end)
}

func (p *Pool) GetSet(key string, value interface{}) *redis.StringCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.GetSet(ctx, key, value)
}

func (p *Pool) BitCount(key string, bitCount *redis.BitCount) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.BitCount(ctx, key, bitCount)
}

func (p *Pool) BitPos(key string, bit int64, pos ...int64) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.BitPos(ctx, key, bit, pos...)
}

func (p *Pool) BitField(key string, args ...interface{}) *redis.IntSliceCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntSliceCmd(err)
	}
	return conn.BitField(ctx, key, args...)
}

func (p *Pool) GetBit(key string, offset int64) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.GetBit(ctx, key, offset)
}

func (p *Pool) SetBit(key string, offset int64, value int) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SetBit(ctx, key, offset, value)
}

func (p *Pool) BitOp(op int, destKey string, keys ...string) *redis.IntCmd {
	if factory, ok := p.connFactory.(*ShardConnFactory); ok {
		allKeys := append(keys, destKey)
		if factory.isCrossMultiShards(allKeys...) {
			return newErrorIntCmd(errCrossMultiShards)
		}
	}
	conn, err := p.connFactory.getMasterConn(destKey)
	if err != nil {
		return newErrorIntCmd(err)
	}
	switch op {
	case bitOpAnd:
		return conn.BitOpAnd(ctx, destKey, keys...)
	case bitOpOr:
		return conn.BitOpOr(ctx, destKey, keys...)
	case bitOpXor:
		return conn.BitOpXor(ctx, destKey, keys...)
	default:
		return newErrorIntCmd(errors.New("unknown op type"))
	}
}

func (p *Pool) BitOpAnd(destKey string, keys ...string) *redis.IntCmd {
	return p.BitOp(bitOpAnd, destKey, keys...)
}

func (p *Pool) BitOpOr(destKey string, keys ...string) *redis.IntCmd {
	return p.BitOp(bitOpOr, destKey, keys...)
}

func (p *Pool) BitOpXor(destKey string, keys ...string) *redis.IntCmd {
	return p.BitOp(bitOpXor, destKey, keys...)
}

func (p *Pool) BitOpNot(destKey string, key string) *redis.IntCmd {
	if factory, ok := p.connFactory.(*ShardConnFactory); ok {
		if factory.isCrossMultiShards(destKey, key) {
			return newErrorIntCmd(errCrossMultiShards)
		}
	}
	conn, err := p.connFactory.getMasterConn(destKey)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.BitOpNot(ctx, destKey, key)
}

func (p *Pool) Decr(key string) *redis.IntCmd {
	return p.DecrBy(key, 1)
}

func (p *Pool) Incr(key string) *redis.IntCmd {
	return p.DecrBy(key, -1)
}

func (p *Pool) IncrBy(key string, increment int64) *redis.IntCmd {
	return p.DecrBy(key, -1*increment)
}

func (p *Pool) DecrBy(key string, decrement int64) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.DecrBy(ctx, key, decrement)
}

func (p *Pool) IncrByFloat(key string, value float64) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.IncrByFloat(ctx, key, value)
}

func (p *Pool) HSet(key, field string, value interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.HSet(ctx, key, field, value)
}

func (p *Pool) HDel(key string, fields ...string) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.HDel(ctx, key, fields...)
}

func (p *Pool) HExists(key, field string) *redis.BoolCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.HExists(ctx, key, field)
}

func (p *Pool) HGet(key, field string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.HGet(ctx, key, field)
}

func (p *Pool) HGetAll(key string) *redis.StringStringMapCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringStringMapCmd(err)
	}
	return conn.HGetAll(ctx, key)
}

func (p *Pool) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.HIncrBy(ctx, key, field, incr)
}

func (p *Pool) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.HIncrByFloat(ctx, key, field, incr)
}

func (p *Pool) HKeys(key string) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.HKeys(ctx, key)
}

func (p *Pool) HLen(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.HLen(ctx, key)
}

func (p *Pool) HMGet(key string, fields ...string) *redis.SliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorSliceCmd(err)
	}
	return conn.HMGet(ctx, key, fields...)
}

func (p *Pool) HMSet(key string, values ...interface{}) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.HMSet(ctx, key, values...)
}

func (p *Pool) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.HSetNX(ctx, key, field, value)
}

func (p *Pool) HVals(key string) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.HVals(ctx, key)
}

func (p *Pool) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.BLPop(ctx, timeout, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(keys...) {
		return newErrorStringSliceCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(keys[0])
	return conn.BLPop(ctx, timeout, keys...)
}

func (p *Pool) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.BRPop(ctx, timeout, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(keys...) {
		return newErrorStringSliceCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(keys[0])
	return conn.BRPop(ctx, timeout, keys...)
}

func (p *Pool) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.BRPopLPush(ctx, source, destination, timeout)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(source, destination) {
		return newErrorStringCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(source)
	return conn.BRPopLPush(ctx, source, destination, timeout)
}

func (p *Pool) LIndex(key string, index int64) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.LIndex(ctx, key, index)
}

func (p *Pool) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.LInsert(ctx, key, op, pivot, value)
}

func (p *Pool) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return p.LInsert(key, "BEFORE", pivot, value)
}

func (p *Pool) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return p.LInsert(key, "AFTER", pivot, value)
}

func (p *Pool) LLen(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.LLen(ctx, key)
}

func (p *Pool) LPop(key string) *redis.StringCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.LPop(ctx, key)
}

func (p *Pool) LPush(key string, values ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.LPush(ctx, key, values...)
}

func (p *Pool) LPushX(key string, values ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.LPushX(ctx, key, values...)
}

func (p *Pool) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.LRange(ctx, key, start, stop)
}

func (p *Pool) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.LRem(ctx, key, count, value)
}

func (p *Pool) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.LSet(ctx, key, index, value)
}

func (p *Pool) LTrim(key string, start, stop int64) *redis.StatusCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStatusCmd(err)
	}
	return conn.LTrim(ctx, key, start, stop)
}

func (p *Pool) RPop(key string) *redis.StringCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.RPop(ctx, key)
}

func (p *Pool) RPopLPush(source, destination string) *redis.StringCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.RPopLPush(ctx, source, destination)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(source, destination) {
		return newErrorStringCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(source)
	return conn.RPopLPush(ctx, source, destination)
}

func (p *Pool) RPush(key string, values ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.RPush(ctx, key, values...)
}

func (p *Pool) RPushX(key string, values ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.RPushX(ctx, key, values...)
}

func (p *Pool) SAdd(key string, members ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SAdd(ctx, key, members...)
}

func (p *Pool) SCard(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SCard(ctx, key)
}

func (p *Pool) SDiff(keys ...string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, err := p.connFactory.getSlaveConn()
		if err != nil {
			return newErrorStringSliceCmd(err)
		}
		return conn.SDiff(ctx, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(keys...) {
		return newErrorStringSliceCmd(errCrossMultiShards)
	}
	conn, err := p.connFactory.getSlaveConn(keys[0])
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SDiff(ctx, keys...)
}

func (p *Pool) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.SDiffStore(ctx, destination, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)

	if factory.isCrossMultiShards(append(keys, destination)...) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(destination)
	return conn.SDiffStore(ctx, destination, keys...)
}

func (p *Pool) SInter(keys ...string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, err := p.connFactory.getSlaveConn()
		if err != nil {
			return newErrorStringSliceCmd(err)
		}
		return conn.SInter(ctx, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(keys...) {
		return newErrorStringSliceCmd(errCrossMultiShards)
	}
	conn, err := p.connFactory.getSlaveConn(keys[0])
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SInter(ctx, keys...)
}

func (p *Pool) SInterStore(destination string, keys ...string) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.SInterStore(ctx, destination, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(append(keys, destination)...) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(destination)
	return conn.SInterStore(ctx, destination, keys...)
}

func (p *Pool) SIsMember(key string, member interface{}) *redis.BoolCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorBoolCmd(err)
	}
	return conn.SIsMember(ctx, key, member)
}

func (p *Pool) SMembers(key string) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SMembers(ctx, key)
}

func (p *Pool) SMembersMap(key string) *redis.StringStructMapCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringStructMapCmd(err)
	}
	return conn.SMembersMap(ctx, key)
}

func (p *Pool) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.SMove(ctx, source, destination, member)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(source, destination) {
		return newErrorBoolCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(source)
	return conn.SMove(ctx, source, destination, member)
}

func (p *Pool) SPop(key string) *redis.StringCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.SPop(ctx, key)
}

func (p *Pool) SPopN(key string, count int64) *redis.StringSliceCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SPopN(ctx, key, count)
}

func (p *Pool) SRandMember(key string) *redis.StringCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringCmd(err)
	}
	return conn.SRandMember(ctx, key)
}

func (p *Pool) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SRandMemberN(ctx, key, count)
}

func (p *Pool) SRem(key string, members ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.SRem(ctx, key, members...)
}

func (p *Pool) SUnion(keys ...string) *redis.StringSliceCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, err := p.connFactory.getSlaveConn()
		if err != nil {
			return newErrorStringSliceCmd(err)
		}
		return conn.SUnion(ctx, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(keys...) {
		return newErrorStringSliceCmd(errCrossMultiShards)
	}
	conn, err := p.connFactory.getSlaveConn(keys[0])
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.SUnion(ctx, keys...)
}

func (p *Pool) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.SUnionStore(ctx, destination, keys...)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if factory.isCrossMultiShards(append(keys, destination)...) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(destination)
	return conn.SUnionStore(ctx, destination, keys...)
}

func (p *Pool) ZAdd(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAdd(ctx, key, members...)
}

func (p *Pool) ZAddNX(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAddNX(ctx, key, members...)
}

func (p *Pool) ZAddXX(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAddXX(ctx, key, members...)
}

func (p *Pool) ZAddCh(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAddCh(ctx, key, members...)
}

func (p *Pool) ZAddNXCh(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAddNXCh(ctx, key, members...)
}

func (p *Pool) ZAddXXCh(key string, members ...*redis.Z) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZAddXXCh(ctx, key, members...)
}

func (p *Pool) ZIncr(key string, member *redis.Z) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.ZIncr(ctx, key, member)
}

func (p *Pool) ZIncrNX(key string, member *redis.Z) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.ZIncrNX(ctx, key, member)
}

func (p *Pool) ZIncrXX(key string, member *redis.Z) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.ZIncrXX(ctx, key, member)
}

func (p *Pool) ZCard(key string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZCard(ctx, key)
}

func (p *Pool) ZCount(key, min, max string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZCount(ctx, key, min, max)
}

func (p *Pool) ZLexCount(key, min, max string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZLexCount(ctx, key, min, max)
}

func (p *Pool) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.ZIncrBy(ctx, key, increment, member)
}

func (p *Pool) ZPopMax(key string, count ...int64) *redis.ZSliceCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZPopMax(ctx, key, count...)
}

func (p *Pool) ZPopMin(key string, count ...int64) *redis.ZSliceCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZPopMin(ctx, key, count...)
}

func (p *Pool) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRange(ctx, key, start, stop)
}

func (p *Pool) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZRangeWithScores(ctx, key, start, stop)
}

func (p *Pool) ZRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRangeByScore(ctx, key, opt)
}

func (p *Pool) ZRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRangeByLex(ctx, key, opt)
}

func (p *Pool) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZRangeByScoreWithScores(ctx, key, opt)
}

func (p *Pool) ZRank(key, member string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRank(ctx, key, member)
}

func (p *Pool) ZRem(key string, members ...interface{}) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRem(ctx, key, members...)
}

func (p *Pool) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRemRangeByRank(ctx, key, start, stop)
}

func (p *Pool) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRemRangeByScore(ctx, key, min, max)
}

func (p *Pool) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRemRangeByLex(ctx, key, min, max)
}

func (p *Pool) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRevRange(ctx, key, start, stop)
}

func (p *Pool) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZRevRangeWithScores(ctx, key, start, stop)
}

func (p *Pool) ZRevRangeByScore(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRevRangeByScore(ctx, key, opt)
}

func (p *Pool) ZRevRangeByLex(key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.ZRevRangeByLex(ctx, key, opt)
}

func (p *Pool) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorZSliceCmd(err)
	}
	return conn.ZRevRangeByScoreWithScores(ctx, key, opt)
}

func (p *Pool) ZRevRank(key, member string) *redis.IntCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.ZRevRank(ctx, key, member)
}

func (p *Pool) ZScore(key, member string) *redis.FloatCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.ZScore(ctx, key, member)
}

func (p *Pool) ZUnionStore(dest string, store *redis.ZStore) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.ZUnionStore(ctx, dest, store)
	}
	factory := p.connFactory.(*ShardConnFactory)
	keys := append(store.Keys, dest)
	if factory.isCrossMultiShards(keys...) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(keys[0])
	return conn.ZUnionStore(ctx, dest, store)
}

func (p *Pool) ZInterStore(destination string, store *redis.ZStore) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.ZInterStore(ctx, destination, store)
	}
	factory := p.connFactory.(*ShardConnFactory)
	keys := append(store.Keys, destination)
	if factory.isCrossMultiShards(keys...) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(keys[0])
	return conn.ZInterStore(ctx, destination, store)
}

func (p *Pool) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	conn, err := p.connFactory.getMasterConn(key)
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.GeoAdd(ctx, key, geoLocation...)
}

func (p *Pool) GeoPos(key string, members ...string) *redis.GeoPosCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorGeoCmd(err)
	}
	return conn.GeoPos(ctx, key, members...)
}

func (p *Pool) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorGeoLocationCmd(err)
	}
	return conn.GeoRadius(ctx, key, longitude, latitude, query)
}

func (p *Pool) GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.GeoRadiusStore(ctx, key, longitude, latitude, query)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if query.Store != "" && factory.isCrossMultiShards(key, query.Store) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	if query.StoreDist != "" && factory.isCrossMultiShards(key, query.StoreDist) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(key)
	return conn.GeoRadiusStore(ctx, key, longitude, latitude, query)
}

func (p *Pool) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorGeoLocationCmd(err)
	}
	return conn.GeoRadiusByMember(ctx, key, member, query)
}

func (p *Pool) GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	if _, ok := p.connFactory.(*HAConnFactory); ok {
		conn, _ := p.connFactory.getMasterConn()
		return conn.GeoRadiusByMemberStore(ctx, key, member, query)
	}
	factory := p.connFactory.(*ShardConnFactory)
	if query.Store != "" && factory.isCrossMultiShards(key, query.Store) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	if query.StoreDist != "" && factory.isCrossMultiShards(key, query.StoreDist) {
		return newErrorIntCmd(errCrossMultiShards)
	}
	conn, _ := p.connFactory.getMasterConn(key)
	return conn.GeoRadiusByMemberStore(ctx, key, member, query)
}

func (p *Pool) GeoDist(key string, member1, member2, unit string) *redis.FloatCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorFloatCmd(err)
	}
	return conn.GeoDist(ctx, key, member1, member2, unit)
}

func (p *Pool) GeoHash(key string, members ...string) *redis.StringSliceCmd {
	conn, err := p.connFactory.getSlaveConn(key)
	if err != nil {
		return newErrorStringSliceCmd(err)
	}
	return conn.GeoHash(ctx, key, members...)
}

func (p *Pool) PFAdd(key string, els ...interface{}) *redis.IntCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorIntCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.PFAdd(ctx, key, els...)
}

func (p *Pool) PFCount(keys ...string) *redis.IntCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorIntCmd(errShardPoolUnSupported)
	}
	conn, err := p.connFactory.getSlaveConn()
	if err != nil {
		return newErrorIntCmd(err)
	}
	return conn.PFCount(ctx, keys...)
}

func (p *Pool) PFMerge(dest string, keys ...string) *redis.StatusCmd {
	if _, ok := p.connFactory.(*ShardConnFactory); ok {
		return newErrorStatusCmd(errShardPoolUnSupported)
	}
	conn, _ := p.connFactory.getMasterConn()
	return conn.PFMerge(ctx, dest, keys...)
}
