package local_filter

import (
	"fmt"
	"github.com/demdxx/gocast/v2"
	"testing"
)

// https://hur.st/bloomfilter/?n=10000000&p=&m=512M&k=6
func TestLocalBloomService_Exist(t *testing.T) {
	bloomService := NewLocalBloomFilter(50000, 3, NewEncryptor())

	fmt.Println(gocast.Int("aa"))
	fmt.Println(1 << 6)
	fmt.Println(64&32 == 0)
	var a int32
	a = 101
	fmt.Printf("int32: %b %d\n", a, a)
	fmt.Println("& 是 两数各对应的二进位相与。（两位均为1才为1）")
	bloomService.Set("12")
	fmt.Println(bloomService.Exist("12"))
	fmt.Println("✓")
}
