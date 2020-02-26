package gohelper

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

const (
	// SHA256 校验和算法-sha256
	SHA256 = "SHA256"
	// SHA1 校验和算法-sha1
	SHA1 = "SHA1"
)

type Hash struct {
	file    string
	hashSum string
}

func NewHash(algorithm string, file string) *Hash {
	var h hash.Hash
	// 判断两个 utf-8 编码字符串，大小写不敏感
	if strings.EqualFold(algorithm, SHA256) {
		h = sha256.New()
	}
	if strings.EqualFold(algorithm, SHA1) {
		h = sha1.New()
	}

	f, err := os.Open(file)
	if err != nil {
		return &Hash{}
	}
	io.Copy(h, f)
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return &Hash{
		file:    file,
		hashSum: string(sum),
	}
}

func (h Hash) Equal(given string) bool {
	return h.hashSum == given
}

func (h Hash) HashSum() string {
	return h.hashSum
}
