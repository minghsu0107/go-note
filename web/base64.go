package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	var sEnc string
	sEnc = b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+

	sDec, _ := b64.StdEncoding.DecodeString(sEnc) // sDec is []byte
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // YWJjMTIzIT8kKiYoKSctPUB-
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
