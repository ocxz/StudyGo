package main

import (
	"fmt"
)

// 斐波那契数的求取，规律：1，1，2，3，5…… 等于前两数之和
// 根据fibonacci数列规律，求第n个数对应的值
// 分析：当n==1 || n==2 return 1
//  	 当n>=3 时为前两项之和
func Getfibonacci(n int) int {
	if n == 1 || n== 2 {
		return 1
	} else {
		return Getfibonacci(n-1) + Getfibonacci(n-2)
	}
}

// 求函数f(1)=3; f(n)=2*f(n-1)+1
// 分析：当n=1时 返回3 否则 返回前一项的2倍加1
func MyFunc(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2 * MyFunc(n - 1) + 1
	}
}

// 猴子吃桃子，第一天吃掉其中一半+1，以后每天吃掉其中一半+1，第十天只有一个，求最初共多少个桃子
// 分析  当第十天，及n==10 时，返回1个
//      当不足十天时，前一天的桃子数，是后一天的桃子数两倍+2
func Peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不正确")
		return 0
	} else if n == 10 {
		return 1
	} else {
		return 2 * Peach(n + 1) + 2
	}
}

func main() {

	// 递归的小案例
	// 斐波那契数的求取，规律：1，1，2，3，5…… 等于前两数之和
	fmt.Println(Getfibonacci(5))

	// 求函数f(1)=3; f(n)=2*f(n-1)+1 的值
	i := 5
	fmt.Printf("f(%d)=%d \n", i, MyFunc(i))

	// 猴子吃桃子，第一天吃掉其中一半+1，以后每天吃掉其中一半+1，第十天只有一个，求最初共多少个桃子
	fmt.Printf("最初共有%d个桃子", Peach())  // 1534
} 