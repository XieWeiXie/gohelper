package gohelper

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

type Md5Action struct {
	V   string
	hex string
}

func NewMd5Action(v string) *Md5Action {
	return &Md5Action{
		V:   v,
		hex: "",
	}
}
func (M *Md5Action) Do() {
	m := md5.New()
	m.Write([]byte(M.V))
	M.hex = hex.EncodeToString(m.Sum(nil))
}
func (M *Md5Action) String() string {
	return M.hex
}

func (M *Md5Action) Random(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
