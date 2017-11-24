package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

const service = "http://localhost:8080"

func main() {
	var err error
	var f *os.File
	var fi os.FileInfo
	var bar *pb.ProgressBar

	if f, err = os.Open("out-3.ogv"); err != nil {
		log.Fatal(err)
	}
	if fi, err = f.Stat(); err != nil {
		log.Fatal(err)
	}
	bar = pb.New64(fi.Size()).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.Start()

	r, w := io.Pipe()
	mpw := multipart.NewWriter(w)
	go func() {
		var part io.Writer
		defer w.Close()
		defer f.Close()

		if part, err = mpw.CreateFormFile("file", fi.Name()); err != nil {
			log.Fatal(err)
		}
		part = io.MultiWriter(part, bar)
		if _, err = io.Copy(part, f); err != nil {
			log.Fatal(err)
		}
		if err = mpw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	resp, err := http.Post(service, mpw.FormDataContentType(), r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(ret))
}
