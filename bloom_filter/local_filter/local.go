package local_filter

import (
	"fmt"
	"github.com/demdxx/gocast/v2"
)

type LocalBloomService struct {
	m, k, n   int32
	bitmap    []int32
	encryptor Encryptor
}

// NewLocalBloomFilter /*
func NewLocalBloomFilter(m, k int32, encryptor Encryptor) *LocalBloomService {
	/*
		• m：bimap 的长度，由用户输入
		• k：hash 函数的个数，由用户输入
		• n：布隆过滤器中的元素个数，由布隆过滤器统计
		• bitmap：位图，类型为 []int，其中使用到每个 int 元素的 32 个 bit 位，因此有 []int 长度为 m/32. 构造时为避免除不尽的问题，切片长度额外增大 1
		• encryptor：散列函数编码模块
	*/
	return &LocalBloomService{
		m:         m,
		k:         k,
		bitmap:    make([]int32, m/32+1),
		encryptor: encryptor,
	}
}

func (l *LocalBloomService) Exist(val string) bool {
	fmt.Println(l.getKEncrypted(val))
	for _, offset := range l.getKEncrypted(val) {
		index := offset >> 5     // 等价于 / 32  (计算落在那个组) 下标对应切片中的那个元素 ,每个元素里面有32个位置
		bitOffset := offset & 31 // 等价于 % 32  (在组的第几位) 等价于 offset - index * 32

		if l.bitmap[index]&(1<<bitOffset) == 0 { //1<<bitOffset
			return false
		}
	}

	return true
}

func (l *LocalBloomService) getKEncrypted(val string) []int32 {
	encrypteds := make([]int32, 0, l.k)
	origin := val

	for i := 0; int32(i) < l.k; i++ {
		encrypted := l.encryptor.Encrypt(origin)
		encrypteds = append(encrypteds, encrypted%l.m) //这里得出每个值在 bitmap  中的具体下标
		if int32(i) == l.k-1 {
			break
		}
		origin = gocast.Str(encrypted)
	}
	return encrypteds
}

func (l *LocalBloomService) Set(val string) {
	l.n++
	//[]int32  中每一个元素都占4个byte  就是每个元素有32个bit位
	for _, offset := range l.getKEncrypted(val) {
		index := offset >> 5     // 等价于 / 32
		bitOffset := offset & 31 // 等价于 % 32
		//按位或后赋值(参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
		l.bitmap[index] |= 1 << bitOffset

	}

}
