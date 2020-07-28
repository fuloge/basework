package rsa

import (
	"strconv"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	enc, _ := RsaEncrypt([]byte(""))
	xx, _ := RsaDecrypt(enc)
	println(xx)

	i, _ := strconv.ParseInt("", 10, 64)
	println(i)
}
