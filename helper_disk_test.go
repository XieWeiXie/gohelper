package go_helper

import (
	"fmt"
	"testing"
)

func TestDiskForPathAll(t *testing.T) {
	path := "/"
	d := NewDisk(path)
	fmt.Println(d)
	fmt.Println(d.Free / 1024 / 1024 / 1024)
	fmt.Println(d.Status(90))
}
