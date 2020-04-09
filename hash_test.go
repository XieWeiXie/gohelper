package gohelper

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

func TestNewHash(t *testing.T) {
	h := NewHash("sha256", "/Users/xiewei/.g/downloads/go1.14.darwin-amd64.tar.gz")
	fmt.Println(h.Equal("2472dcd681b761501fadb35ec361503efd27de2ba2270b2fe35cb6ece7362243"))
	fmt.Println(strings.EqualFold("SSHA256", SHA256))

}

func TestHash(t *testing.T) {
	fmt.Println(computeAcceptKey("w4v7O6xFTi36lq3RNcgctw=="))
}

var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey))
	h.Write(keyGUID)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
