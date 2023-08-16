package redis_filter

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/demdxx/gocast/v2"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed lua/set_bitmap.lua
var luaSetBitMap string

//go:embed lua/get_bitmap.lua
var luaGetBitMap string

//go:embed lua/flush_bitmap.lua
var luaFlushBitMap string

//go:embed lua/expire_bitmap.lua
var luaExpireBitMap string

type RedisBloomFilter struct {
	m           uint32    //预估一共有多少个bit 位
	k           int       //hash 次数(加密次数)
	keyPrefix   string    //redis中 键的前缀
	encrypt     Encryptor //hash 函数
	client      redis.Cmdable
	evalTimeOut time.Duration
	keySLice    []string
}

func NewRedisBloomFilter(count uint32, encryptTimes int, keyPrefix string, encrypt Encryptor, client redis.Cmdable, duration time.Duration) *RedisBloomFilter {

	r := &RedisBloomFilter{
		m:           count,
		k:           encryptTimes,
		keyPrefix:   keyPrefix,
		encrypt:     encrypt,
		client:      client,
		evalTimeOut: duration,
	}
	//redis 键的格式是 前缀:编号   编号 从 0 开始
	for i := 0; i < encryptTimes; i++ {
		r.keySLice = append(r.keySLice, fmt.Sprintf("%s:%d", r.keyPrefix, i))
	}
	return r
}

func (r *RedisBloomFilter) Set(key string) (bool, error) {
	//TODO implement me
	return r.Eval(luaSetBitMap, r.getEncryptedSlicing(key))
}

func (r *RedisBloomFilter) Get(key string) (bool, error) {
	//TODO implement me
	return r.Eval(luaGetBitMap, r.getEncryptedSlicing(key))
}

func (r *RedisBloomFilter) Delete() (int64, error) {
	//TODO implement me

	//lua 操作
	//return r.Eval(luaFlushBitMap, 1, []uint32{})

	//正常操作
	ctx, cancel := context.WithTimeout(context.Background(), r.evalTimeOut)
	defer cancel()
	return r.client.Del(ctx, r.keySLice...).Result()
}

func (r *RedisBloomFilter) Expire(milliseconds uint32) {
	//lua 脚本操作
	//ctx, cancel := context.WithTimeout(context.Background(), r.evalTimeOut)
	//defer cancel()
	//argvStr := ""
	//for i := range r.keySLice {
	//	argvStr += gocast.Str(r.keySLice[i])
	//	if i != len(r.keySLice)-1 {
	//		argvStr += ","
	//	}
	//}
	//return r.client.Eval(ctx, luaExpireBitMap, r.keySLice, milliseconds).Bool()
	//正常操作也可以
	for i := range r.keySLice {
		ctx, cancel := context.WithTimeout(context.Background(), r.evalTimeOut)
		_, _ = r.client.PExpire(ctx, r.keySLice[i], time.Millisecond*time.Duration(milliseconds)).Result()
		cancel()
	}
	return
}

func (r *RedisBloomFilter) getEncryptedSlicing(val string) []uint32 {
	res := make([]uint32, 0, r.k)
	for i := 0; i < r.k; i++ {
		tmPInt := r.encrypt.Encrypt(val)
		val = gocast.Str(tmPInt)
		res = append(res, tmPInt%r.m)
	}
	return res
}

func (r *RedisBloomFilter) Eval(luaScript string, bitSlice []uint32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.evalTimeOut)
	defer cancel()
	argvStr := ""
	for i := range bitSlice {
		argvStr += gocast.Str(bitSlice[i])
		if i != len(bitSlice)-1 {
			argvStr += ","
		}
	}
	return r.client.Eval(ctx, luaScript, r.keySLice, argvStr).Bool()
}
