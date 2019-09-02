package main

import (
	"fmt"
)

func main() {

	// 编写一个程序，输入成绩，打出评分 [90-100] 优秀  [70-90) 良 [60,70) 及格  小于60 不及格
	var score float64
	fmt.Print("请输入你的成绩：")
	fmt.Scanln(&score)

	var grade int = (int)(score/10)
	switch grade {
		case 10.0, 9.0 :
			fmt.Println("优秀")
		case 7.0,8.0 :
			fmt.Println("良好")
		case 6.0 :
			fmt.Println("及格")
		default :
			fmt.Println("不及格")
	}

	// 类似于if-else的写法
	switch {
		case score >= 90:
			fmt.Println("优秀")
		case score >= 70 :
			fmt.Println("良好")
		case score >=60 :
			fmt.Println("及格")
		default :
			fmt.Println("不及格")
	}

	// switch 穿透 fallthrought  默认穿透一层
	switch  num := 10; num {
		case 10 :
			fmt.Println("ok1")
			fallthrough  // 默认穿透一层
		case 20 :
			fmt.Println("ok2")
			fallthrough
		case 30 :
			fmt.Println("ok3")
		default :
			fmt.Println("没有匹配成功")
	}

	// type switch 用来判断指向变量的类型
	var x interface {}    // 声明一个空接口变量
	y := 10.0
	x = y
	switch i := x.(type) {   // 获得x变量的数据类型，赋值给i
		case nil:
			fmt.Printf("x 的数据类型~ ：%T", i)
		case int:
			fmt.Printf("x 是int类型")
		case float64:
			fmt.Printf("x 是float64类型")
		case func(int) float64:
			fmt.Printf("x 是 func(int) 类型")
		case bool,string:
			fmt.Printf("x 是bool或者string类型")
		default:
			fmt.Printf("未知类型")
	}
}