package main

import (
	"fmt"
)

// 根据年份和月份，判断该月天数 年份：1960-2050 月份：1-12 月
// 1 3 5 7 8 10 12 31天
// 2  闰年：29  平年 28
// 其他 30 天
func GetDay(year int, month int) (int, bool) {
	if month < 1 || month > 12 {
		return 0, false
	} else {
		switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				return 31, true
			case 2:
				if year % 400 == 0 || (year % 100 !=0 && year % 4 ==0) {   // 闰年
					return 29, true
				} else {
					return 28, true  // 平年
				}
			default:
				return 30, true
		}
	}
}
func main() {
	
	for {
		var year int
		fmt.Print("请输入年份（1960-2050）：")
		fmt.Scanln(&year)
		if year < 1960 || year > 2050 {   // 判断年份是否符合，不符合直接跳过下面代码，执行下一次输入年份循环
			fmt.Printf("%d 年份不符合范围，请重新输入 \n", year)
			continue
		} else {  // 年份符合，循环输入月份

			for {
				var month int
				fmt.Print("请输入月份")
				fmt.Scanln(&month)
				if month < 1 || month > 12 {   // 判断月份是否符合，不符合跳过，执行一次输入月份循环
					fmt.Printf("%d 月份不符合范围，请重新输入 \n", month)
					continue
				} else {
					day, _ := GetDay(year, month)
					fmt.Printf("%d年的%d月有%d天 \n", year, month, day)
					break   // 输出正确，结束月份输入，重新输入年份
				}
			}
		}
	}
}