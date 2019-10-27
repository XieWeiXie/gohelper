package gohelper

import (
	"fmt"
	"unsafe"
)

type SizeofAction struct {
	Value interface{}
	size  int
}

func NewSizeofAction(v interface{}) *SizeofAction {
	return &SizeofAction{
		Value: v,
	}
}

func (S *SizeofAction) Do() {

}
func (S *SizeofAction) String() interface{} {
	return nil
}

func (S *SizeofAction) GetSize() {
	fmt.Println(unsafe.Sizeof(S.Value))
}
