package main
import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[10] = 900
}
func main() {

	// map是引用类型，遵守引用传递机制，函数接收map后，修改，可直接影响原来的map
	map1 := make(map[int]int)
	map1[1] = 90
	map1[2] = 87
	map1[10] = 1
	map1[11] = 12
	modify(map1)   // 函数接收map后，修改，可直接影响原来的map
	fmt.Println(map1) 
}