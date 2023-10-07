package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

var header = `{"alg":"ES256","typ":"JWT"}`
var base64priv = "LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUxxRExYdVBpSHA3TU1QTU5PNG5jTnR3Z1FnWUNyRVE4bmJhdm9FWm1qczVvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFMnRzbE91Z2hmQXNXK28rVDQvcDkrVExMdWxMR0JKNkhjUngvTW1IS3QyREhkODdYakFiYgpBdzdXamRTdWZVRjAxMjhlMVZTL01QT3NFRUNNZDhvNlJRPT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo="

func main() {
	payload := `{"u":"gogs-lsy"}`
	token := make_token(payload)
	fmt.Println(token)
}

func make_token(payload string) string {
	base64header := base64.RawURLEncoding.EncodeToString([]byte(header))
	base64payload := base64.RawURLEncoding.EncodeToString([]byte(payload))
	base64headerpayload := base64header + "." + base64payload
	privpemData, err := base64.StdEncoding.DecodeString(base64priv)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	privkeypem, _ := pem.Decode(privpemData) // NOTE
	if privkeypem == nil {
		fmt.Println("pem not found")
		return ""
	}
	privateKey, err := x509.ParseECPrivateKey(privkeypem.Bytes)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	mac := sha256.New()
	mac.Write([]byte(base64headerpayload))
	expectedMAC := mac.Sum(nil)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, expectedMAC)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	rb := make([]byte, 32)
	sb := make([]byte, 32)
	r.FillBytes(rb)
	s.FillBytes(sb)

	return base64headerpayload + "." + base64.RawURLEncoding.EncodeToString(append(rb, sb...))
}
