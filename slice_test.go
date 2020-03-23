package gohelper

import (
	"fmt"
	"testing"
)

func TestNewSliceCap(t *testing.T) {
	s := NewSlice(512,512)
	s.v = append(s.v, 1)
	fmt.Println(s.Len(), s.Cap())
	s1 := NewSlice(1024,1024)
	s1.v = append(s1.v, 1)
	fmt.Println(s1.Len(), s1.Cap()) // 1024 * 1.25
}
