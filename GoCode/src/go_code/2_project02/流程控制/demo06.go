package main

import (
	"fmt"
)

func main() {

	// 百米赛跑 8秒内 可进入决赛 否则淘汰
	// 进入决赛时，根据性别提示 进入男子队|女子队

	fmt.Println("请输入你百米赛跑的成绩（单位：秒）")
	var score float64
	fmt.Scanln(&score)
	if score <=8 {

		// 进入决赛
		var gender byte
		fmt.Println("请输入您的性别（0：男，1：女）")
		fmt.Scanln(&gender)
		if gender == 0 {
			fmt.Println("恭喜你，你被分到了男子决赛组")
		} else if gender == 1 {
			fmt.Println("恭喜你，你被分到了女子决赛组")
		}
	} else {
		fmt.Println("抱歉，成绩不足8秒，您被淘汰了")
	}
}