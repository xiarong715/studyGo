package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

var gJWTHeader = `{"alg":"ES256","typ":"JWT"}`

func main() {
	payload := `{"u":"xiarong"}`
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(privateKey)

	ecdsa.Sign(rand.Reader, privateKey, []byte(payload))
}
