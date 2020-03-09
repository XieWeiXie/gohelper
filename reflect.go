package gohelper

import (
	"fmt"
	"reflect"
)

type ReflectAction struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewReflectAction(name string, age int) ReflectAction {
	return ReflectAction{
		Name: name,
		Age:  age,
	}
}

func (r ReflectAction) Kind() {
	v := reflect.TypeOf(r)
	fmt.Println(v.Kind())
}
