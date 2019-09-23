package main
import (
	"fmt"
)

func main() {

	// 顺序查找
	names := [...]string{"1", "2", "3", "4", "5"}
	var input = ""
	fmt.Println("请输入一个数字")
	fmt.Scanln(&input)

	// 第一种顺序查找的方式
	for i :=0; i < len(names); i++ {
		if names[i] == input {
			fmt.Printf("找到了...，下标为%v，值为%v \n", i, names[i])
			break;
		} else if i == len(names) - 1 {
			fmt.Println("没有找到");
			break;
		}
	}

	// 顺序查找第二种方式，定义index初始值为-1，找到了则将下标赋给index
	index := -1
	for i :=0; i < len(names); i++ {
		if names[i] == input {
			index = i
			break;
		} 
	}
	if index != -1 {
		fmt.Printf("找到了...，下标为%v，值为%v \n", index, names[index])
	} else {
		fmt.Println("没有找到");
	}
}