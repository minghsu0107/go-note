package main

import (
    "html/template"
    "io"
    "log"
    "net/http"
    "os"
    "bufio"
)

// compiling/caching the template
var templates = template.Must(template.New("tmpl").Parse(`
<html>
  <head>
    <title>File Upload Demo</title>
    <style>
        body {
            font-family: Sans-serif;
            padding-top: 40px;
            padding-bottom: 40px;
            background-color: #ffffff;
        }
        h1 {text-align: center; margin-bottom: 30px;}
        .message {font-weight:bold}
        fieldset {width:50%}
    </style>
  </head>
  <body>
    <div class="container">
      <h1>File Upload Demo</h1>
      <div class="message">{{.}}</div>
      <form class="form-signin" method="post" action="/upload" enctype="multipart/form-data">
          <fieldset>
            <input type="file" name="myfiles" id="myfiles" multiple="multiple">
            <input type="submit" name="submit" value="Submit">
          </fieldset>
      </form>
    </div>
  </body>
</html>
`))

// limit the POST body size to 32.5 Mb
// throw error if client attempt to send more than that
func maxBytes(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
            r.Body = http.MaxBytesReader(w, r.Body, 32<<20+512)
            if err := r.ParseForm(); err != nil {
                http.Error(w, "Bad Request", http.StatusBadRequest)
                return
            }
            f(w, r)
    }
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    // GET to display the upload form.
    case "GET":
        err := templates.Execute(w, nil)
        if err != nil {
            log.Print(err)
        }
        // POST analyzes each part of the MultiPartReader (ie the uploaded file(s))
        // and saves them to disk.
    case "POST":
        // grab the request.MultipartReader
        reader, err := r.MultipartReader()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // copy each part to destination.
        for {
            part, err := reader.NextPart()
            if err == io.EOF {
                break
            }

            log.Println(part.FormName())

            // if part.FileName() is empty, skip this iteration.
            if part.FileName() == "" {
                continue
            }

            buf := bufio.NewReader(part)
            sniff, _ := buf.Peek(512)
            contentType := http.DetectContentType(sniff)
            log.Print(contentType)
            if contentType != "image/jpeg" {
                http.Error(w, "file type not allowed", http.StatusBadRequest)
                return
            }

            // prepare the dst
            dst, err := os.Create("./" + part.FileName())
            defer dst.Close()
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            // limit file size to 32 Mb
            // with 1 byte offset to see if part reader still has some data left
            var maxSize int64 = 32 << 20
            lmt := io.MultiReader(buf, io.LimitReader(part, maxSize - 511))
            written, err := io.Copy(dst, lmt);
            if err != nil && err != io.EOF {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            if written > maxSize {
                os.Remove(dst.Name())
                http.Error(w, "file size over limit", http.StatusBadRequest)
                return
            }

        }
        // displaying a success message.
        err = templates.Execute(w, "Upload successful.")
        if err != nil {
            log.Print(err)
        }
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
func main() {
    http.HandleFunc("/upload", maxBytes(uploadHandler))
    log.Print("Listening on port:8082...")
    // Listen on port 8080
    http.ListenAndServe(":8082", nil)
}