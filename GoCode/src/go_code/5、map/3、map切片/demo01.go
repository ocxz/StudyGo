package main
import (
	"fmt"
)

func main() {

	// 需求：1个妖怪对应一个map
	//      2、可动态增加

	// monster := make([]map[string]string)

	monsters := []map[string]string{
		{"name" : "白骨精", "age" : "18"},
		{"name" : "牛魔王", "age" : "201"},
		{"name" : "玉兔精", "age" : "34"},
		{"name" : "红孩儿", "age" : "3"},
	}
	fmt.Println(monsters)
	//动态增加妖怪信息，这样做是行不通的
	if monsters[len(monsters)] == nil {
		monster := map[string]string{ "name" : "狐狸精", "age" : "207"}
		monsters[len(monsters)] = monster
	}

	// 可以使用切片append动态增加
	monster := map[string]string{ "name" : "狐狸精", "age" : "207"}
	monsters = append(monsters, monster)

	fmt.Println(monsters)
}