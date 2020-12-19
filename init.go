package main

import (
	//"fmt"
	"io/ioutil"
	//"reflect"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	TodoLists     []string `yaml:"TodoLists"`
	CompleteLists []string `yaml:"CompleteLists"`
	CloudKey      string   `yaml:"CloudKey"`
}

func getConfig() Configs {
	//TODO 更改为用户目录下的.config/gotodo.yaml
	//TODO 如果没有就创建gotodo.yaml
	configs := Configs{}
	filePath := "./config.yaml"
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(buffer, &configs)
	return configs
}

//TODO 列表信息
func getListDatas() [][]string {
	var allDatas [][]string
	todoList := getConfig()
	allDatas = append(allDatas, todoList.TodoLists)
	allDatas = append(allDatas, todoList.CompleteLists)
	return allDatas
}

//TODO 从网络同步数据
func getSyncConfig() {}
