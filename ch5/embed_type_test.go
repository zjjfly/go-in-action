package ch5

import (
	"fmt"
	"testing"
)

type myuser struct {
	name  string
	email string
}

func (u *myuser) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name,
		u.email)
}

type myadin struct {
	//把已有类型声明在类型内部,称为外部类型的内部类型
	myuser
	name  string
	level string
}

func Test_embed(t *testing.T) {
	admin := myadin{
		myuser{
			"jjzi",
			"zjjblue@126.com",
		},
		"zjj",
		"super",
	}
	//外部类型可以直接访问内部类型的方法
	admin.myuser.notify()
	//Sending user email to jjzi<zjjblue@126.com>
	//内部类型的字段和方法会提升到外部类型中,所以可以直接使用外部类型访问内部类型的方法
	admin.email = "zjjblue@gmail.com"
	//外部类型可以通过声明和内部类型相同的字段和方法来覆盖内部类型的字段方法
	admin.notify()
	//Sending admin email to zjj<zjjblue@gmail.com>

	//由于内部类型的提升,内部类型实现的接口,外部类型也一样实现了
	sendNotification(&admin)
	//如果外部类也实现了notify方法,会覆盖内部类的实现
}
