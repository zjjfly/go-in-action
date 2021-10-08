package ch8

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	flag := log.Ldate | log.Ltime | log.Lshortfile
	Trace = log.New(ioutil.Discard, "TRACE: ", flag)
	Info = log.New(os.Stdout, "INFO: ", flag)
	Warning = log.New(os.Stdout, "WARNING: ", flag)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flag)
}

func Test_customized_log(t *testing.T) {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
