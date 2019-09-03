package main

import (
	"fmt"
)

func main() {

	// continue小练习
	// 某人有100,000元，每经过一次路口需要缴费
	// 缴费规则：现金 > 50000 每次交5%
	//			现金 <= 50000 每次交1000
	// 求该人可经过多少次路口
	
	count, money := 0, 100000.0

	for {
		if money > 50000 {
			money *= 0.95
			count++
		} else if money >1000 && money < 50000{
			money -= 1000
			count++
		} else {
			break
		}
	}
	fmt.Printf("该人通过了%d个路口，还剩下%.2f元零头", count, money)
}