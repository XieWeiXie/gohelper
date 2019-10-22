package gohelper

import (
	"fmt"
	"testing"
)

func TestC_Add(t *testing.T) {
	c := NewConCurrency(16)
	for i := 0; i < c.number; i++ {
		c.Add(1)
		go func(i int) {
			defer c.Done()
			fmt.Println(i)
		}(i)
	}
	c.Wait()
}
