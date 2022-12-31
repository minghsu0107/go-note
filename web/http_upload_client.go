package main

import (
	"net/http"
	"mime/multipart"
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"bytes"
)

const fileName = "/Users/xuhaoming/Downloads/photo.jpg"
const uploadURL = "http://localhost:8080/receive"

func UploadFile(url, fileName string) error {
	var buf = new(bytes.Buffer) // buf is a pointer
	m := multipart.NewWriter(buf)

	part, err := m.CreateFormFile("myFile", "foo.jpg")
	if err != nil {
		return err
	}

	file, err := os.Open(fileName)
	if err != nil {
	    return err
	}
	defer file.Close()

	if _, err = io.Copy(part, file); err != nil {
	    return err
	}

	m.Close() // do remember to close

	_, err = http.Post(url, m.FormDataContentType(), buf)
	if err != nil {
		return err
	}
	return nil
}

// transfer data in chunks unboundedly, without specifying Content-Lengthof request body
// so the header would contain "Transfer-Encoding: chunked"
func UploadFileSaveMem(url, fileName string) {
	// use pipe to avoid extra buffer
	// func Pipe() (*PipeReader, *PipeWriter)
	// Each Write to the PipeWriter blocks until it has satisfied one or more Reads 
	// from the PipeReader that fully consume the written data. 
	// The data is copied directly from the Write to the corresponding Read (or Reads); 
	// there is no internal buffering.
	r, w := io.Pipe() // write to pipe through w and read from pipe through r
	m := multipart.NewWriter(w) // write to w
	go func() {
	    defer w.Close() // close last; close the pipe
	    defer m.Close() // close first; ends writing to w

	    // creates a new multipart section 
	    // with a form-data header (the provided field name and file name)
	    // the body of the part should be written to the returned writer
	    part, err := m.CreateFormFile("myFile", "foo.jpg") // part is a writer that writes to w
	    if err != nil {
	        return
	    }
	    file, err := os.Open(fileName)
	    if err != nil {
	        return
	    }
	    fmt.Println(file.Name()) // /Users/xuhaoming/Downloads/photo.jpg
	    defer file.Close()

	    // copy file content to part
	    if _, err = io.Copy(part, file); err != nil {
	        return
	    }
	}()

	// multipart/form-data; boundary=d32ec6cacfa3c9e86547f76ba289c880b1d953b369dae569397366c54c39
	fmt.Println(m.FormDataContentType())
	// d32ec6cacfa3c9e86547f76ba289c880b1d953b369dae569397366c54c39
	fmt.Println(m.Boundary())

	resp, err := http.Post(url, m.FormDataContentType(), r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body) // resp.Body type: io.ReadCloser
	fmt.Println("response Body:", string(body))
}
func main() {
	UploadFile(uploadURL, fileName)
	UploadFileSaveMem(uploadURL, fileName)
}
