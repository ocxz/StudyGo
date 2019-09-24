package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	Scores [5]float64
	prt *int   // 指针
	slice []int  // 切片
	map1 map[string]string  // map
}

func main() {

	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)
	p1.slice = make([]int, 10)
	p1.slice[0] = 1
	p1.slice[1] = 2
	fmt.Println(p1)

	// 
	p2 := p1
	p2.slice[0] = 500
	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
}