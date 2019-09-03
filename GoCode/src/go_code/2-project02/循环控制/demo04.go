package main

import (
	"fmt"
)

func main() {

	// 多重循环控制

	// 1、统计3个班的成绩情况，班中有5名同学，求每个班的平均分和所有班级平均分 学生成绩键盘输入
	// 定义两个变量，表示班级的个数和班级的人数
	totalAvg := 0.0
	allPassCount :=0
	totalClass := 0
	fmt.Print("请输入班级个数：")
	fmt.Scanln(&totalClass)
	for i := 1; i <= totalClass; i++ {
		sum := 0.0
		passCount := 0
		numberOfClass := 0
		fmt.Printf("请输入%v班同学的人数", i)
		fmt.Scanln(&numberOfClass)
		fmt.Printf("请输入%v班同学的成绩", i)
		fmt.Println()
		for j := 1; j <= numberOfClass; j++ {
			score := 0.0
			fmt.Printf("请输入%v班第%v位同学的成绩：", i, j)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		allPassCount += passCount
		fmt.Printf("%v班同学的成绩输入完毕，平均成绩为%.2f, 合格人数为：%d \n", i, sum / float64(numberOfClass), passCount)
		totalAvg += sum / float64(numberOfClass)
	}
	fmt.Printf("三个班同学的成绩输入完毕，平均成绩为%.2f, 合格人数为：%d", totalAvg / float64(totalClass), allPassCount)
}