package main
import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Rect struct {
	LeftUp Point
	RightDown Point
}

type Rect2 struct {
	LeftUp *Point
	RightDown *Point
}

func main() {

	r1 := Rect{Point{1, 2}, Point{3, 4}}

	// r1中，有四个int，在内存中是联系分布的
	fmt.Printf("r1.leftUp.X的地址%p\n", &r1.LeftUp.X)
	fmt.Printf("r1.leftUp.Y的地址%p\n", &r1.LeftUp.Y)
	fmt.Printf("r1.RightDown.X的地址%p\n", &r1.RightDown.X)
	fmt.Printf("r1.RightDown.Y的地址%p\n", &r1.RightDown.Y)
	fmt.Println()

	// r2有两个*Point类型，这两个*Point类型的本身地址是连续的，但其指定的地址不一定连续
	// 也就是说存放两个指针变量的空间是连续的，指针变量的值即不一定连续
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp本身的地址%p\n", r2.LeftUp)
	fmt.Printf("r2.leftUp.X的地址%p\n", &r2.LeftUp.X)
	fmt.Printf("r2.leftUp.Y的地址%p\n", &r2.LeftUp.Y)
	fmt.Printf("r2.RightDown%p\n", r2.RightDown)
	fmt.Printf("r2.RightDown.X的地址%p\n", &r2.RightDown.X)
	fmt.Printf("r2.RightDown.Y的地址%p\n", &r2.RightDown.Y)
}