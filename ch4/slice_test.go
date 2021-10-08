package ch4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//切片
func Test_slice_create(t *testing.T) {
	//声明一个切片
	//第二个参数是切片长度和容量
	slice := make([]string, 5)
	//第二个参数是长度,第三个参数是容量
	slice = make([]string, 3, 5)
	//使用字面量初始化,[]中不能有数字,有的话就是数组了
	slice = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//使用索引初始化,可以指定长度和容量
	slice = []string{99: ""}
	assert.Equal(t, 100, len(slice))
	assert.Equal(t, 100, cap(slice))
	//nil切片,用于描述一个不存在的切片
	var nilSlice []int
	assert.Equal(t, 0, len(nilSlice))
	//空切片,用于表示空集合
	emptySlice := []int{}
	assert.Equal(t, 0, len(emptySlice))

	intSlice := []int{10, 20, 30, 40, 50}
	intSlice[0] = 5
	newSlice := intSlice[1:3]
	assert.Equal(t, []int{20, 30}, newSlice)
	//新切片的容量是5-1=4
	assert.Equal(t, 4, cap(newSlice))
	//newSlice和intSlice的底层数组是共享的,一个切片修改这部分数据会让另一个切片感知到
	intSlice[2] = 35
	assert.Equal(t, 35, newSlice[1])
}

func Test_slice_append(t *testing.T) {
	intSlice := []int{10, 20, 30, 40, 50}
	newSlice := intSlice[1:3]
	//只能访问切片长度返回内的元素.超出长度的元素只能在append之后才能访问.
	//append就是在切片末尾添加元素,如果容量足够,append操作不会增加容量
	assert.Equal(t, 2, len(newSlice))
	newSlice = append(newSlice, 60)
	assert.Equal(t, 60, intSlice[3])
	//如果容量不够,则会增加容量,如果切片的容量是1000以下的,则新容量是原来的两倍
	intSlice = append(intSlice, 70)
	assert.Equal(t, 10, cap(intSlice))

	slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//使用3个索引来控制切片的容量,第三个索引表示容量的结束位置,它不能大于可用容量
	colorSlice := slice[2:3:4]
	assert.Equal(t, 2, cap(colorSlice))
	//指定新切片的容量可以让第一个append操作不会修改底层数组而是创建新的底层数组,与原来的底层数组分离
	colorSlice = slice[2:3:3]
	colorSlice = append(colorSlice, "Kiwi")
	assert.NotEqual(t, colorSlice[1], slice[3])

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	//append最后一个是可变参数,和...结合使用可以把一个切片的所有元素加到另一个切片中
	assert.Equal(t, []int{1, 2, 3, 4}, append(s1, s2...))

}

func Test_slice_iterate(t *testing.T) {
	slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	//range每次迭代返回两个值,一个是index,一个是元素的副本
	for _, color := range slice {
		t.Log(color)
	}

	//range创建元素的副本,而不是直接返回元素的引用,所以每次打印出的
	intSlice := []int{10, 20, 30, 40, 50}
	for index, value := range intSlice {
		t.Logf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &intSlice[index])
	}

	//使用传统的for可以进行更多的控制
	for i := 2; i < len(intSlice); i++ {
		t.Logf("Index: %d Value: %d\n", i, intSlice[i])
	}
}

func Test_multi_dimension(t *testing.T) {
	//多维切片的每一行的长度不需要一样
	slice := [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
}

func Test_func_arg(t *testing.T) {
	//切片可以直接做参数,因为它的大小和底层数组的大小无关
	foo2([]int{1, 2})
}

func foo2(slice []int) []int {
	return slice
}
