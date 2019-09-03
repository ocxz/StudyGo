package main

import (
	"fmt"
)

func main() {

	// 岳小鹏 参加Golang考试  与父亲岳不群达成承诺
	// 成绩为100分  奖励一辆BMW
	// 成绩为(80,99] 奖励一台iphone 7 plus
	// 成绩为[60,80] 奖励一个iPad
	// 否则没有奖励

	var score byte
	fmt.Print("请输入岳小鹏的Golang成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("哇，成绩为100分  奖励一辆BMW")
	} else if score> 80 {
		fmt.Println("成绩为(80,99] 奖励一台iphone 7 plus")
	} else if score >= 60 {
		fmt.Println("成绩为[60,80] 奖励一个iPad")
	} else {
		fmt.Println("成绩不及格，什么奖励都没有")
	}
}