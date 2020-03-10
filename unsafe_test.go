package gohelper

import (
	"fmt"
	"testing"
	"unsafe"
)

type Model struct {
	Name  string
	Age   int
	Slice []string
}

func TestNewUnsafeAction(t *testing.T) {
	model := Model{
		Name:  "XieWei",
		Age:   20,
		Slice: []string{"1", "2", "3", "4"},
	}
	un := NewUnsafeAction(model)
	mm := un.model.(Model)
	unPointer := unsafe.Pointer(&mm)
	m := uintptr(unPointer) + unsafe.Offsetof(mm.Name)
	fmt.Println(*(*string)(unsafe.Pointer(m)))
	age := uintptr(unPointer) + unsafe.Offsetof(mm.Age)
	fmt.Println(*(*int)(unsafe.Pointer(age)))
	slice := uintptr(unPointer) + unsafe.Offsetof(mm.Slice)
	fmt.Println(*(*[]string)(unsafe.Pointer(slice)))
	length := slice + uintptr(8)
	fmt.Println(*(*int)(unsafe.Pointer(length)))
	capacity := slice + uintptr(16)
	fmt.Println(*(*int)(unsafe.Pointer(capacity)))

	um := unsafe.Pointer(&model)
	name := uintptr(um) + unsafe.Offsetof(model.Name)
	fmt.Println(*(*string)(unsafe.Pointer(name)))
}
