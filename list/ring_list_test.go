package list

import (
	"fmt"
	"testing"
)

func TestNewRingList(t *testing.T) {
	ringList := NewRingList[int](10)
	for i := 0; i < 15; i++ {
		fmt.Println(ringList.In(i))
	}
	for i := 0; i < 15; i++ {
		fmt.Println(ringList.Out())
	}
}

func TestRingQueue_In(t *testing.T) {

}

func TestRingQueue_IsEmpty(t *testing.T) {

}

func TestRingQueue_Out(t *testing.T) {

}
