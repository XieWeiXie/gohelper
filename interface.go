package gohelper

type Action interface {
	Do()
	String() string
}
