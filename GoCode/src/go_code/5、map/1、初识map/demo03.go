package main
import (
	"fmt"
)

func main() {

	citys := map[string]string {
		"no1" : "北京",
		"no2" : "上海",
		"no3" : "天津",
	}

	fmt.Println(citys)

	// 演示删除
	delete(citys, "no2")
	fmt.Println(citys)
	// 演示删除不存在，指定key不存在时，不操作也不报错
	delete(citys, "no10")
	fmt.Println(citys)

	// 清空map
	// 1、遍历所有key，然后逐一删除
	// 2、直接make一个新map，将原来的覆盖，原来的被gc回收
	citys = make(map[string]string)
	fmt.Println(citys)

	// map查找
	city, findRes := citys["no1"]
	fmt.Println(city)
	fmt.Println(findRes)
}