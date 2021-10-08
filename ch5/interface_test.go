package ch5

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func Test_curl(t *testing.T) {
	resp, err := http.Get("http://www.a9vg.com")
	if err != nil {
		t.Fail()
	}
	//Copy方法声明的第一个参数类型是Writer,os.Stdout的类型是File,它实现了Writer接口的方法,所以可以作为Writer的一个实现
	//第二个参数类型是Reader接口,而http.Response.Body就是这个接口类型的值
	io.Copy(os.Stdout, resp.Body)
	if err := resp.Body.Close(); err != nil {
		t.Fail()
	}

	//Copy可以用于很多标准库的类型
	var buffer bytes.Buffer
	//buffer指针实现了Write接口
	//如果用户定义的类型实现了一个接口声明的一组方法,那么就认为这个类型实现了接口,可以把这个类型的值赋值给这个接口类型的变量
	buffer.Write([]byte("Hello"))
	fmt.Fprintf(&buffer, " World!")
	io.Copy(os.Stdout, &buffer)
}

type notifier interface {
	notify()
}

type admin struct {
	name  string
	email string
}

func (u *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", u.name,
		u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func Test_interface_variable(t *testing.T) {
	u1 := admin{
		name:  "zjj",
		email: "zjjblue@gmail.com",
	}
	//下面的的代码编译通不过,因为对notify的实现方法的接收者类型是admin的指针
	//对于接口的方法:接收者类型是指针的方法只能通过指针来调用,而接收者是值的方法既可以通过指针也可以通过值来调用
	//sendNotification(u1)
	sendNotification(&u1)
	//真正的原因是编译器并不是总能自动获得一个值的地址,下面的代码会报错
	//Duration(32).pretty()

	//多态
	u2 := user{
		"jjzi",
		"zjjblue@126.com",
		22,
		true,
	}
	sendNotification(u2)
}

func (d *Duration) pretty() string {
	return fmt.Sprintf("Duration:%d", d)
}
