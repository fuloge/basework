package rsa

import "testing"

func TestRsaEncrypt(t *testing.T) {
	enc, _ := RsaEncrypt([]byte("hello"))
	xx, _ := RsaDecrypt(enc)
	println(xx)
}
