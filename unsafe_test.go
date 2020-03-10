package gohelper

import (
	"fmt"
	"reflect"
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
	typ := reflect.TypeOf(mm)
	for i := 0; i < typ.NumField(); i++ {
		fmt.Println(typ.Field(i).Name, typ.Field(i).Offset)
	}
	unPointer := unsafe.Pointer(&mm)
	m := uintptr(unPointer) + unsafe.Offsetof(mm.Name)
	fmt.Println(*(*string)(unsafe.Pointer(m)))
	age := uintptr(unPointer) + unsafe.Offsetof(mm.Age)
	//age := uintptr(unPointer) + unsafe.Sizeof(8)*2
	fmt.Println(unsafe.Sizeof(8)*2, unsafe.Offsetof(mm.Age), *(*int)(unsafe.Pointer(age)))

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
