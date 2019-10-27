package gohelper

import (
	"testing"
)

func TestNewSizeofAction(t *testing.T) {
	type N struct {
		Name string
	}
	var n = N{
		Name: "string",
	}
	s := NewSizeofAction(n)
	s.GetSize()
}
