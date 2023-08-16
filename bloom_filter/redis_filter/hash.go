package redis_filter

import (
	"github.com/spaolacci/murmur3"
	"math"
)

type RedisEncryptor struct {
}

func NewRedisEncryptor() *RedisEncryptor {
	return &RedisEncryptor{}
}

func (r *RedisEncryptor) Encrypt(src string) uint32 {
	hasher := murmur3.New32()
	_, _ = hasher.Write([]byte(src))
	return hasher.Sum32() % math.MaxUint32
}
