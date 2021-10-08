package ch5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//一种声明类型的方法:struct
//定义一个struct
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func Test_struct(t *testing.T) {
	//声明一个user类型的变量
	var jjzi user
	//不初始化其中的字段,那么这些字段会使用零值初始化
	assert.Equal(t, "", jjzi.name)
	assert.Equal(t, "", jjzi.email)
	assert.Equal(t, 0, jjzi.ext)
	assert.Equal(t, false, jjzi.privileged)

	//结构字面量的第一种形式
	zjj := user{
		name:       "zjj",
		email:      "zjjblue@gmail.com",
		ext:        22,
		privileged: true,
	}
	t.Log(zjj)
	//结构字面量的第二种形式
	zjj = user{"zjj", "zjjblue@gmail.com", 12, false}
	t.Log(zjj)

	//结构体的字段类型可以是自定义类型
	type admin struct {
		person user
		level  string
	}

	james := admin{
		person: user{
			name:       "james",
			email:      "test@gmail.com",
			ext:        12,
			privileged: true,
		},
		level: "super",
	}
	t.Log(james)
}

//声明类型的另一种方法:基于已有类型,重新命名,类似Scala的关键字type的用法
//go编译器值只允许为用户定义的类型,所以这样的语法是很有用的
type Duration int64

func Test_alias(t *testing.T) {
	var dur Duration
	//dur=int64(1000) 这种写法无法编译通过,go编译器把Duration和int64看做两个不同的值,不会对类型做隐式转换,即使它们是兼容的
	dur = Duration(1000)
	t.Log(dur)
}
