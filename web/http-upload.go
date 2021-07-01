package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func uploadPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("<html><title>Go upload</title><body><form action='http://localhost:8080/receive' method='post' enctype='multipart/form-data'><label for='file'>Filename:</label><input type='file' name='myfile'><input type='submit' value='Click to Upload'></form></body></html>")))
}
func uploadProgress(w http.ResponseWriter, r *http.Request) {

	mr, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ticker := time.Tick(time.Millisecond) // <-- use this in production
	ticker := time.Tick(50 * time.Millisecond) // this is for demo purpose with longer delay

	for {

		var read int64
		part, err := mr.NextPart()

		if err == io.EOF {
			fmt.Printf("\nDone!")
			break
		}

		dst, err := os.OpenFile("./" + part.FileName(), os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			return
		}

		for {
			buffer := make([]byte, 100000)
			cBytes, err := part.Read(buffer)
			if err == io.EOF {
				fmt.Printf("\nLast buffer read!")
				break
			}
			read = read + int64(cBytes)

			if read > 0 {
				<-ticker
				fmt.Printf("\rUploading progress %v", read) // for console
				dst.Write(buffer[0:cBytes])
			} else {
				break
			}

		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadPage)
	mux.HandleFunc("/receive", uploadProgress)

	http.ListenAndServe(":8080", mux)
}
