package main

import (
	"fmt"
)

func main() {

	// 判断一个年份是否位闰年
	// 闰年条件：能被400整除 | 能被4整除不能被100整除

	var year int = 2020
	
	if year % 400 == 0 || (year % 4 ==0 && year % 100 !=0){
		fmt.Printf("%v年是闰年", year)
	} else {
		fmt.Printf("%v年不是闰年", year)
	}
}