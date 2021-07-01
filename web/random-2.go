package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func sessionId() string {
	b := make([]byte, 32)
	//ReadFull reads len(b) bytes from rand.Reader and writes to b
	//rand.Reader is a robust and globally shared random generator
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	fmt.Println(b)
	// [59 93 7 32 208 226 198 2 69 198 151 17 74 163 25 150 38 122 235 227 13 55 202 200 64 37 233 218 44 57 252 30]
	return base64.URLEncoding.EncodeToString(b)
}

func main() {
	fmt.Println(sessionId()) // O10HINDixgJFxpcRSqMZliZ66-MNN8rIQCXp2iw5_B4=
	fmt.Println(sessionId()) // Fl1vA1DAsRo6wwBKM5j-Kh8Yuq--3pfQJeRMVtG--Ls=
}
