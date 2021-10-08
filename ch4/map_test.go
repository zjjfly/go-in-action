package ch4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_init(t *testing.T) {
	//映射的键的类型可以使任何可以使用==进行相等性判断的类型,除了切片,函数以及含有切片的结构体
	dict1 := make(map[string]int)
	dict1 = map[string]int{"Red": 1, "Orange": 2}
	assert.Equal(t, 2, len(dict1))
	//映射的值的类型没有限制
	dict2 := map[int][]string{1: {"1", "2"}}
	assert.Equal(t, 2, len(dict2[1]))
	//nil映射,不能存储键值对
	var colors map[string]string
	assert.Nil(t, colors)
}

func Test_operate(t *testing.T) {
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	colors["Blue"] = "#f0f8ff"
	//判断一个键是否存在于映射中
	value, exist := colors["Green"]
	if exist {
		t.Log(value)
	}
	//如果键不存在,也会返回一个value,是值类型的零值,所以最好不要使用value来判断键是否存在
	assert.Equal(t, "", value)
	//删除键值对
	delete(colors, "Blue")
	assert.Equal(t, "", colors["Blue"])
}

func Test_iterate(t *testing.T) {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range colors {
		t.Logf("Key: %s Value: %s\n", key, value)
	}
}

func Test_arg(t *testing.T) {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	removeColor(colors, "Coral")
	assert.Equal(t, "", colors["Coral"])
}

//映射作为参数传递的不是副本,对它的修改会影响到其他使用它的代码
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
