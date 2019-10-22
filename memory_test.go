package gohelper

import (
	"fmt"
	"testing"
)

func TestNewMemory(t *testing.T) {
	m := NewMemory()
	fmt.Println(m.Used)
}
