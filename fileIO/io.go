package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	// ouput "some" to stdout
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 32)
	// Seek(offset int64, whence int)
	// whence: 0=start, 1=current, 2=tail
	r.Seek(0, 0)
	lr = io.LimitReader(r, 4)
	// ouput "some" to stdout
	if _, err := io.CopyBuffer(os.Stdout, lr, buf); err != nil {
		log.Fatal(err)
	}

	r.Seek(0, 0)
	mw := io.MultiWriter(os.Stdout, os.Stdout, os.Stdout)
	r.WriteTo(mw)

	r.Seek(0, 0)
	sr := io.NewSectionReader(r, 6, 5)
	n, err := io.Copy(os.Stdout, sr) // World
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
