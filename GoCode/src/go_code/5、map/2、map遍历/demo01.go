package main
import (
	"fmt"
)

func main() {

	// 使用for-range遍历map
	citys := map[string]string {
		"no1" : "北京",
		"no2" : "上海",
		"no3" : "天津",
		"no4" : "武汉",
		"no5" : "长沙",
	}
	for key, value := range citys {
		fmt.Printf("key = %v，value = %v\n", key, value)
	}

	// for-range遍历多层map结构
	students := map[string]map[string]string {
		"stu01" : { "name" : "张三", "age" : "18", "sex" : "男" },
		"stu02" : { "name" : "李四", "age" : "19", "sex" : "男" },
		"stu03" : { "name" : "王五", "age" : "20", "sex" : "女" },
		"stu04" : { "name" : "赵六", "age" : "21", "sex" : "男" },
		"stu05" : { "name" : "田七", "age" : "22", "sex" : "男" },
	}
	fmt.Println(students) 
	for snum, student := range students{
		fmt.Println(snum, "：")
		for key, value := range student{
			fmt.Printf("\t%v：%v\n", key, value)
		}
		fmt.Println()
	}
}