package main

import (
	"fmt"
)

func main() {

	// 编写一个程序，输入成绩，打出评分 [90-100] 优秀  [70-90) 良 [60,70) 及格  小于60 不及格
	var score float64
	fmt.Print("请输入你的成绩：")
	fmt.Scanln(&score)

	var gender int = (int)(score/10)
	switch gender {
		case 10.0, 9.0 :
			fmt.Println("优秀")
		case 7.0,8.0 :
			fmt.Println("良好")
		case 6.0 :
			fmt.Println("及格")
		default :
			fmt.Println("不及格")
	}
}