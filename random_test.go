package gohelper

import (
	"fmt"
	"testing"
)

func TestNewRandomAction(t *testing.T) {
	r := NewRandomAction()
	fmt.Println(r.GeneratorInt())
	fmt.Println(r.GeneratorIntn(10))
	fmt.Println(r.GeneratorIntn(10))
}
