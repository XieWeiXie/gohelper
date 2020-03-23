package gohelper

import (
	"fmt"
	"runtime"
)

type RuntimeFunction struct {

}


func (r *RuntimeFunction) Caller(depth int){
	ptr, file, line, ok := runtime.Caller(depth)
	if !ok {
		return
	}
	name := runtime.FuncForPC(ptr).Name()
	fmt.Println(fmt.Sprintf("file: %s line: %d function name: %s", file, line, name))
}