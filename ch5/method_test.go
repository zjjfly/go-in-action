package ch5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//定义一个user结构体的方法,接收者是user类型的值,这个值是副本,所以不会修改调用方法的那个变量
//这种类型的方法实用于原始类型和引用类型
func (u user) notify() {
	fmt.Printf("Send Email to %s<%s>\n", u.name, u.email)
}

//定义一个接收者是user类型的指针,实际传入方法的是user类型的值,这个值是就是调用这个方法的指针所指向的值,不是副本
func (u *user) changeEmail(email string) {
	u.email = email
}

func Test_invoke(t *testing.T) {
	zjj := user{"zjj", "zjjblue@126.com", 12, false}
	//user类型的值调用方法
	zjj.notify()

	//使用user类型的值可以调用使用指针接收者声明的方法
	zjj.changeEmail("zjjblue@gmail.com")
	assert.Equal(t, "zjjblue@gmail.com", zjj.email)

	james := &user{"james", "james@126.com", 22, true}
	//使用user类型的指针可以调用使用值接收者声明的方法
	james.notify()

	//user类型的指针调用方法
	james.changeEmail("james@gmail.com")
	assert.Equal(t, "james@gmail.com", james.email)
}
