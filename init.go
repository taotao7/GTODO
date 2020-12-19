package main

import (
	"fmt"
	"io/ioutil"
	//"reflect"

	"gopkg.in/yaml.v2"
)

type TodoList struct {
	Lists    []string `yaml:"TodoLists"`
	CloudKey string   `yaml:"CloudKey"`
}

func importConfig() {
	//TODO 更改为用户目录下的.config/gotodo.yaml
	todolists := TodoList{}
	filePath := "./config.yaml"
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(buffer, &todolists)
	//循环输出
	for _, value := range todolists.Lists {
		fmt.Println(value)
	}

}

//TODO 通过网络数据库获取数据并入
func getDatas() {}
