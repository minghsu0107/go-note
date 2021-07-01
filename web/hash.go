package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// SHA1 hashes are frequently used to compute short identities for binary or text blobs.
// For example, the git revision control system uses SHA1s extensively
// to identify versioned files and directories.

func main() {
	s := "sha1 this string"

	h1 := sha1.New()
	h2 := md5.New()

	// Write expects bytes
	h1.Write([]byte(s))
	h2.Write([]byte(s))

	//  The argument to Sum can be used to append to an existing byte slice:
	//  it usually isn’t needed.
	bs1 := h1.Sum(nil)
	bs2 := h2.Sum(nil)

	str1 := fmt.Sprintf("%x", bs1)
	fmt.Println(str1)

	str2 := fmt.Sprintf("%x", bs2)
	fmt.Println(str2)
}
