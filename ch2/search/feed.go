package search

import (
	"encoding/json"
	"os"
)

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

const dataFile = "/Users/zjjfly/GoProjects/src/GoInAction/ch2/data/data.json"

func RetrieveFeed() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	//defer会让后面的函数调用在函数返回的时候执行,类似java的finally
	//直接写在打开文件和错误处理之后,可以提高代码可读性
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
