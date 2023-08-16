package local_filter

import (
	"math"

	"github.com/spaolacci/murmur3"
)

// MyEncryptor 加密对象
type MyEncryptor struct {
}

func NewEncryptor() Encryptor {
	return &MyEncryptor{}
}

func (e *MyEncryptor) Encrypt(origin string) int32 {
	hasher := murmur3.New32()
	_, _ = hasher.Write([]byte(origin))
	return int32(hasher.Sum32() % math.MaxInt32)
}
