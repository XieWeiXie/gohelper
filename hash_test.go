package gohelper

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewHash(t *testing.T) {
	h := NewHash("sha256", "/Users/xiewei/.g/downloads/go1.14.darwin-amd64.tar.gz")
	fmt.Println(h.Equal("2472dcd681b761501fadb35ec361503efd27de2ba2270b2fe35cb6ece7362243"))
	fmt.Println(strings.EqualFold("SSHA256", SHA256))

}
