package ch5

import (
	"fmt"
	"github.com/zjjfly/go-in-action/ch5/counters"
	"github.com/zjjfly/go-in-action/ch5/entities"
	"testing"
)

func Test_export(t *testing.T) {
	counter := counters.New(12)
	fmt.Printf("Counter: %d\n", counter)

	//下面的代码会编译不过,因为age字段是非公开的
	//u := entities.User{
	//	"jjzi",
	//	"zjjblue@126.com",
	//	28,
	//}

	//下面的代码可以运行,因为外部类把内部类的字段提升了,而这些字段声明的是公开的,所以可以直接访问
	a := entities.Admin{
		Rights: 1,
	}
	a.Name = "jjzi"
	a.Email = "zjjblue@126.com"
}
