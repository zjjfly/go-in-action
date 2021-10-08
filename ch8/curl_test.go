package ch8

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestCurl(t *testing.T) {
	resp, err := http.Get("http://www.a9vg.com")
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("resp.html")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, resp.Body)
	if err := resp.Body.Close(); err != nil {
		t.Log(err)
	}
}
