package main

import (
	"fmt"
	"math"
)

// 根据身高，检测体重是否合适
// 公式：(身高 - 108) * 2 = 体重 上下可以有10公斤的浮动  0 正好 -1 偏胖 +1 偏瘦
func TestWeight(height float64, weight float64) (int, bool){
	if height <= 108 {  // 身高小于108，返回错误false, false
		return 0, false
	} else {
		temp := (height - 108) *2 - weight
		if math.Abs(temp) < 10 {
			return 0, true
		} else if temp < 0{
			return -1, true
		} else {
			return 1, true
		}
	}
}
func main() {

	for {
		var height float64
		fmt.Printf("请输入身高（108-300）cm：")
		fmt.Scanln(&height)
		if height < 108 || height > 300{
			fmt.Printf("身高%.2fcm 不符合范围，请重新输入 \n", height)
			continue
		} else {
			for {
				var weight float64
				fmt.Printf("请输入体重（>0）斤：")
				fmt.Scanln(&weight)
				if weight <= 0 {
					fmt.Printf("体重输入有误，请重新输入 \n")
					continue
				} else {
					isOk, _ := TestWeight(height, weight)
					switch isOk {
						case 0 :
							fmt.Printf("恭喜您，您的身高为：%.2fcm，体重为：%.1f斤，身高体重合格", height, weight)
						case -1:
							fmt.Printf("您的身高为：%.2fcm，体重为：%.1f斤，体重偏胖", height, weight)
						case 1:
							fmt.Printf("您的身高为：%.2fcm，体重为：%.1f斤，体重偏瘦", height, weight)	
					}
					fmt.Println()
					break   // 跳出体重循环，进行下一次循环判断
				}
			}
		}
	}
}