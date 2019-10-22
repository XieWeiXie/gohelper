package gohelper

import (
	"fmt"
	"runtime"
)

type Memory struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

var DefaultMemory = Memory{}

func NewMemory() *Memory {
	var memStatus runtime.MemStats
	runtime.ReadMemStats(&memStatus)
	fmt.Println(memStatus)
	return &Memory{
		All:  0,
		Used: 0,
		Free: 0,
		Self: memStatus.Alloc,
	}
}
