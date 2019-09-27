package main

// 需要测试的函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

// 求两个数的差
func getSub(n1 int, n2 int) int {
	return n1 - n2 
}