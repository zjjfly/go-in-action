package ch8

import (
	"log"
	"testing"
)

func init() {
	log.SetPrefix("TRACE:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func Test_log(t *testing.T) {
	log.Println("message")
	log.Fatalln("fatal message")
	log.Panicln("panic message")
}
