package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// open file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 函数返回时，关闭文件
	defer file.Close()

	// 将文件解码到一个切片里
	// 这个切片的每一项是一个指向Feed类型值的指针

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//不需要做检查错误，调用者会做
	return feeds, err
}
