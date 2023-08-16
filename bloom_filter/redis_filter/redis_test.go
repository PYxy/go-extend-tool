package redis_filter

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRedisBloomFilter(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		_ = rdb.Close()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis 连接失败:%v", err))
	}

	redis1 := NewRedisBloomFilter(100000, 2, "ljy", NewRedisEncryptor(), rdb, time.Second)
	//redis1.getEncryptedSlicing("ljy")
	//fmt.Println(redis1.Eval(luaSetBitMap, redis1.getEncryptedSlicing("ljy")))
	//fmt.Println("获取")
	//fmt.Println(redis1.Eval(luaGetBitMap, redis1.getEncryptedSlicing("ljy")))
	key := "你好"
	flag, err := redis1.Set(key)
	if err != nil {
		fmt.Println("key:", key, "添加失败")
		return
	}
	fmt.Println("添加结果", flag)

	flag, err = redis1.Get(key)
	if err != nil {
		fmt.Println("key:", key, "获取失败")
		return
	}
	if flag {
		fmt.Println(key, " 查找成功,需要二次检查(存在假阳)")
	} else {
		fmt.Println(key, " 不存在")
	}

	//清空
	//fmt.Println(redis1.Eval(luaFlushBitMap,[]uint32{}))

	//fmt.Println(redis1.Expire(36000)) // 36 秒
	//redis1.Expire(36000)
	//
	fmt.Println(redis1.Delete())
}

func Test_Set(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		_ = rdb.Close()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis 连接失败:%v", err))
	}
	bloomfilter := NewRedisBloomFilter(100000, 2, "user:name", NewRedisEncryptor(), rdb, time.Second)
	tests := []struct {
		key        string
		wartResult bool
		wartError  error
		name       string
	}{
		{
			"小白",
			true,
			nil,
			"第一个测试",
		},
		{
			"小白",
			true,
			nil,
			"重复测试",
		},
		{
			"小宏",
			true,
			nil,
			"重复测试",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flag, err := bloomfilter.Set(test.key)
			assert.Equal(t, test.wartError, err)

			assert.Equal(t, test.wartResult, flag)
		})
	}
}

func Test_Get(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		_ = rdb.Close()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis 连接失败:%v", err))
	}
	bloomfilter := NewRedisBloomFilter(100000, 2, "user:name", NewRedisEncryptor(), rdb, time.Second)
	tests := []struct {
		key        string
		wartResult bool
		wartError  error
		name       string
	}{
		{
			"小白",
			true,
			nil,
			"第一个测试",
		},
		{
			"小白",
			true,
			nil,
			"重复测试",
		},
		{
			"小",
			false,
			nil,
			"查找没设置过的key",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flag, err := bloomfilter.Get(test.key)
			assert.Equal(t, test.wartError, err)

			assert.Equal(t, test.wartResult, flag)
		})
	}
}

// 前置测试 先跑一次Test_Set
func Test_Delete(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		_ = rdb.Close()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis 连接失败:%v", err))
	}
	bloomfilter := NewRedisBloomFilter(100000, 2, "user:name", NewRedisEncryptor(), rdb, time.Second)
	tests := []struct {
		key        string
		wartResult bool
		wartError  error
		name       string
	}{
		{
			"小白",
			false,
			nil,
			"第一个测试",
		},
		{
			"小白",
			false,
			nil,
			"重复测试",
		},
		{
			"小宏",
			false,
			nil,
			"重复测试",
		},
		{
			"小",
			false,
			nil,
			"重复测试",
		},
	}
	//直接删除键
	_, _ = bloomfilter.Delete()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			flag, err := bloomfilter.Get(test.key)
			assert.Equal(t, test.wartError, err)

			assert.Equal(t, test.wartResult, flag)
		})
	}
}

// 前置测试 先跑一次Test_Set
func Test_Expire(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer func() {
		_ = rdb.Close()
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("redis 连接失败:%v", err))
	}
	bloomfilter := NewRedisBloomFilter(100000, 2, "user:name", NewRedisEncryptor(), rdb, time.Second)
	tests := []struct {
		key        string
		wartResult bool
		wartError  error
		name       string
	}{
		{
			"小白",
			false,
			nil,
			"第一个测试",
		},
		{
			"小白",
			false,
			nil,
			"重复测试",
		},
		{
			"小宏",
			false,
			nil,
			"重复测试",
		},
		{
			"小",
			false,
			nil,
			"查找没设置过的key",
		},
	}
	//直接设置超短过期时间
	bloomfilter.Expire(0)
	time.Sleep(time.Second)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			flag, err := bloomfilter.Get(test.key)
			assert.Equal(t, test.wartError, err)

			assert.Equal(t, test.wartResult, flag)
		})
	}
}
