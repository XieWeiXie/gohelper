package gohelper

import (
	"fmt"
	"testing"
)

func TestNewBase64Action(t *testing.T) {
	b := NewBase64Action("golang")
	b.Do()
	fmt.Println(b.EncodeString)
}
