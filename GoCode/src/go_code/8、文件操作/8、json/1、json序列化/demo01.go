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

// 对结构体的演示
func testStruct() {

	monster := Monster {
		Name : "牛魔王",
		Age : 500,
		Birthday : "2011-11-11",
		Sales : 8000.0,
		Skill : "牛磨拳",
	}

	// 将monster结构体序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败，错误信息：", err)
		return
	}

	// 输出序列化后的结果
	fmt.Println("monster序列化后的结果是：", string(data))
}

// 将map进行序列化
func testMap() {

	// 定义一个map
	m := map[string]interface{} {
		"name" : "红孩儿",
		"age" : 3,
		"address" : "火云洞",
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败，失败原因：\n", err)
		return
	}

	fmt.Println("map序列化后的结果是：", string(data))
}

// 对切片进行序列化，
func testSlice() {

	slice := []map[string]interface{} {
		map[string]interface{} {
			"name" : "jack",
			"age" : 7,
			"address" : "北京",
		},
		map[string]interface{} {
			"name" : "tom",
			"age" : 20,
			"address" : [...]string {
				"北京",
				"上海",
				"天津",
			},
		},
	}

	// 将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败，失败原因：\n", err)
		return
	}

	fmt.Println("切片序列化后的结果是：", string(data))
}

// 对基本数据类型，如：int、string进行序列化
func testFloat64() {

	num1 := 3.14159
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化错误，错误原因：\n", err)
		return
	}

	fmt.Println("Float64序列化后的结果是：", string(data))
}
func main() {

	// 演示 将结构体、map、切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}