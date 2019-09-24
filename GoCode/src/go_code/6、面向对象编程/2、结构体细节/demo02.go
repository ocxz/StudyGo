package main
import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {

	var a A
	var b B
	b = B(a)   // 可以转换，但要求结构体的字段完全一样（名字、类型、个数）
	fmt.Println(a, b)
}