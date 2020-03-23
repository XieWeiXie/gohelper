package gohelper

import "testing"

func TestRuntimeFunction_Caller(t *testing.T) {
	var r RuntimeFunction
	for i:=0;i<2;i++{
		r.Caller(i)
	}
}
