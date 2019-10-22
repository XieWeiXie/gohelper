package gohelper

import (
	"os"
)

type Path struct {
	p string
}

func NewPath(v string) *Path {
	return &Path{p: v}
}

func (P Path) Stat() bool {
	_, e := os.Stat(P.p)
	if e == nil {
		return true
	}
	return false
}
