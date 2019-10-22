package gohelper

import (
	"fmt"
	"log"
	"testing"
)

func TestNewExecutor(t *testing.T) {
	s, e := ExecutorRun("ls", "")
	if e != nil {
		log.Println(e)
		return
	}
	fmt.Println(s)
}
