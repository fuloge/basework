package auth

import (
	"fmt"

	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	jwt := New()
	token, _ := jwt.CreateToken(12345, time.Now().Add(time.Hour*time.Duration(1)).Unix())
	fmt.Println(token)

	time.Sleep(time.Second * 10)
	if res := jwt.TokenIsInvalid(token.Token); res == true {
		fmt.Println("token is invalid")
	} else {
		fmt.Println("token is valid")
	}
}

func TestJWT2(t *testing.T) {
	jwt := New()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU4MjI3NjgsImlhdCI6MTU4NTgxOTE2OCwic3ViIjoiaGFwcHkgYSBsaWZlIiwidWlkIjoxMjM0NX0.t8BafhjwLzXgCwPJ-LQ4wMB5T_C7pPp5T4Khig5moIc"

	if res := jwt.TokenIsInvalid(token); res == true {
		fmt.Println("token is invalid")
	} else {
		fmt.Println("token is valid")
	}
}

func TestGoogleAuth_GetCode(t *testing.T) {
	gauth := GoogleAuth{}
	sect := gauth.GetSecret()
	println("sect", sect)
	code, _ := gauth.GetCode(sect)
	println("code", code)
	time.Sleep(time.Second * 10)
	ok, _ := gauth.VerifyCode(sect, code)
	println("verify code:", ok)

	//er := gauth.CreateQRcode("xcv", sect, 100)
	//println(er)

}
