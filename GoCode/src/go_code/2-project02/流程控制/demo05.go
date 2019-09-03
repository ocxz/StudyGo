package main

import (
	"fmt"
	"math"
)

// 求 一元二次方程解  ax^2 + bx +c = 0 的方程解
// b^2 - 4ac >0 方程有两个解
// b^2 - 4ac =0 方程有一个解
// b^2 -4ac < 0 方程没有解
// x1 = (-b + Sqrt(b^2 - 4ac)) / 2a
// x2 = (-b - Sqrt(b^2 - 4ac)) / 2a
func Test(a float64, b float64, c float64) (float64, float64, int) {
	m := b * b - 4 *a * c
	if m > 0 {
		return (-b + math.Sqrt(m)) / (2*a), (-b - math.Sqrt(m)) / (2*a), 2
	} else if m == 0 {
		return (-b + math.Sqrt(m)) / (2*a), (-b - math.Sqrt(m)) / (2*a), 1
	} else {
		return 0.0, 0.0, 0
	}	
}
func main() {

	// 求方程 x^2 + 2x + 1 = 0 的解
	var a float64
	var b float64
	var c float64
	var strB string
	var strC string
	fmt.Println("请输入ax^2 + bx + c =0 方程的参数：")
	fmt.Print("a=")
	fmt.Scanln(&a)
	fmt.Print("b=")
	fmt.Scanln(&b)
	fmt.Print("c=")
	fmt.Scanln(&c)
	x1, x2, result := Test(a, b, c)
	
	if b>= 0 {
		strB = "+" + fmt.Sprintf("%v",b)
	} else {
		strB = fmt.Sprintf("%v",b)
	}
	if c >=0 {
		strC = "+" + fmt.Sprintf("%v",c)
	} else {
		strC = fmt.Sprintf("%v",c)
	}

	if result == 2 {
		fmt.Printf("方程%vx^2 %sx %s = 0 有两个解：x1=%v, x2=%v", a, strB, strC, x1, x2)
	} else if result == 1 {
		fmt.Printf("方程%vx^2 %sx %s = 0 有一个解：x1=%v, x2=%v", a, strB, strC, x1, x2)
	} else {
		fmt.Printf("方程%vx^2 %sx %s = 0 没有解", a, strB, strC)
	}
}