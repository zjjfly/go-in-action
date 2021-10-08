package ch8

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func Test_write(t *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprint(&b, "World!")
	b.WriteTo(os.Stdout)
}
