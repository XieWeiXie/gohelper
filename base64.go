package gohelper

import (
	"encoding/base64"
	"fmt"
	"log"
)

type Base64Action struct {
	Input        string `json:"input"`
	EncodeString string `json:"encode_string"`
}

func NewBase64Action(input string) *Base64Action {
	return &Base64Action{
		Input:        input,
		EncodeString: "",
	}
}
func (B *Base64Action) Do() {
	B.Encode()
}

func (B *Base64Action) String() string {
	return fmt.Sprintf("Raw: %s, Encode: %s", B.Input, B.EncodeString)
}
func (B *Base64Action) Encode() string {
	v := []byte(B.Input)
	s := base64.StdEncoding.EncodeToString(v)
	B.EncodeString = s
	return s
}

func (B Base64Action) Decode(v string) string {
	b, e := base64.StdEncoding.DecodeString(v)
	if e != nil {
		log.Println(fmt.Sprintf("DecodeString fail: %s", e.Error()))
		return ""
	}
	return string(b)
}
