package templatemethod

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

// inherite base template
type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

// implement download method
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

// override save method
func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

// implement download method
func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
