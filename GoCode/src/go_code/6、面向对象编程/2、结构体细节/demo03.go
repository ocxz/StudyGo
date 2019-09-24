package main
import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`   // `json:"name"` 就是 结构体标签
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	monster := Monster{
		"红孩儿", 10, "吐火",
	}

	// 我们希望将其进行json格式的序列化处理
	// json.Marshal 函数中使用反射
	josnStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误", err)
	} else {
		fmt.Println("josnStr", string(josnStr))
	}
}