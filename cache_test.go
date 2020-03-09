package gohelper

import (
	"fmt"
	"testing"
)

func TestCache(test *testing.T) {
	m := NewMapCache()
	m.Set("1", "Hello")
	fmt.Println(m.Get("1"))

	m1 := SyncMapCache{}
	m1.Set("2", "Hello2")
	fmt.Println(m1.Get("2"))

	m2 := StructCache{}
	m2.Set("3", "Hello3")
	m2.Set("4", "Hello4")
	fmt.Println(m2.Get("3"))
}
