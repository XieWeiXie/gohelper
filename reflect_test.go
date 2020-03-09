package gohelper

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestReflectAction(t *testing.T) {
	r := NewReflectAction("XieWei", 18)
	v := unsafe.Pointer(&r)
	first := unsafe.Pointer(uintptr(v) + unsafe.Offsetof(r.Name))
	valueFirst := *(*string)(first)
	fmt.Println(valueFirst)
	second := unsafe.Pointer(uintptr(v) + unsafe.Offsetof(r.Age))
	fmt.Println(*(*int)(second))
	fmt.Println(unsafe.Sizeof(""), unsafe.Sizeof(0))
}
