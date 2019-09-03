package main

import (
	"fmt"
)

func main(){

	// 97天  多少星期 多少天
	var day int8 = 97

	fmt.Printf("97天是%d星期，%d天 \n", day/7, day%7)q1

	// 摄氏温度=5/9*(华氏温度-100)，求
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)

	fmt.Printf("%v对应的摄氏温度是%v \n", huashi, sheshi)

}