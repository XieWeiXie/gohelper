package gohelper

import (
	"fmt"
	"testing"
)

func TestNewPath(t *testing.T) {
	p := NewPath("os.go")
	fmt.Println(p.Stat())
}
