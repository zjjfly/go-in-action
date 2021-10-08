package ch4

//一个单元测试需要导入testing包
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//单元测试方法必须以Test开头,并且参数必须是*testing.T类型
func Test_array(t *testing.T) {
	//声明数组,每个元素会使用对应类型的零值来初始化,对于int就是0
	var array [5]int
	//数组字面量,...意思是根据初始化时的元素个数来确定数组长度
	array = [...]int{1, 2, 3, 4, 5}
	assert.Equal(t, 5, len(array))
	if array[2] != 3 {
		t.Error("测试失败")
	}
	//指针数组
	pArray := [5]*int{0: new(int), 1: new(int)}
	//给指针赋值
	*pArray[0] = 10
	*pArray[1] = 20
	//同类型的数组之间可以互相赋值,同类型意味着元素类型和元素个数相同
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2
	assert.Equal(t, array1[1], array2[1])
	//把一个指针数组赋值给另一个
	var array3 [3]*string
	array4 := [3]*string{new(string), new(string), new(string)}
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"
	array3 = array4
	//两个数组中的指针是一样的,都指向相同的字符串
	assert.Equal(t, array3[0], array4[0])
	assert.Equal(t, *array3[0], *array4[0])
	//声明二维数组
	var arrays [4][2]int
	//使用字面量初始化
	arrays = [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	//使用索引来部分初始化
	arrays = [4][2]int{0: {0, 1}, 3: {5, 6}}
	assert.Equal(t, [2]int{0, 1}, arrays[0])
	assert.Equal(t, [2]int{5, 6}, arrays[3])

	//定义函数的时候最好不要接受数组,而是接受数组指针,这样就免去了复制数组的开销
	foo1(&[2]int{1, 2})
}

func foo1(array *[2]int) {
}
