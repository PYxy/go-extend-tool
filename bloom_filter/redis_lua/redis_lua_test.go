package redis_lua

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

//go:embed lua/string_redis.lua
var luaStringRedis string

func Test_Lua(t *testing.T) {
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
	fmt.Println(rdb.Eval(ctx, luaStringRedis, []string{"ljy"}).Result())
}
