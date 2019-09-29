package main
import (
	"fmt"
)

func main() {

	// 常量的定义和使用

	const intConst int = 3
	fmt.Println(intConst)

	const (
		a = iota
		b
		c, d = iota, iota
		e = 10
		f 
		g
		h = iota
		i 
	)

	// 0 1 2 2 10 10 10 6 7
	fmt.Println(a, b, c, d, e, f, g, h, i)
}