package main
import (
	"fmt"
	"encoding/json"
)


// 定义一个结构体
type Monster struct {
	Name string `json:"name"`   // 用的是反射机制
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sales float64 `json:"sales"`
	Skill string `json:"skill"`
}

// 演示将json字符串，反序列化成struct结构体
func unmarshalStruct() {

	// 在项目开发中，str通过网络传输、文件获取到的
	str := `{"age":500,"birthday":"2011-11-11","sales":8000,"skill":"牛磨拳","name":"牛魔王"}`

	// 定义一个Monster实例
	var monster Monster
	// 反序列化
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("结构体反序列化失败，失败原因：\n", err)
		return
	}
	fmt.Println(monster)
}

// 将json字符串，反序列化成map
func unmarshalMap() {

	str := `{"address":"火云洞","age":3,"name":"红孩儿"}`
	var m map[string]interface{}

	// 反序列化
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("map反序列化失败，失败原因：\n", err)
		return
	}
	fmt.Println(m)
}

// 将字符串反序列化成切片
func unmarshalSlice() {

	str := ` [{"address":"北京","age":7,"name":"jack"},{"address":["北京","上海","天津"],"age":20,"name":"tom"}]`
	var slice []map[string]interface{}

	// 反序列化
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("切片反序列化失败，失败原因：\n", err)
		return
	}
	fmt.Println(slice)
}



func main() {

	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}