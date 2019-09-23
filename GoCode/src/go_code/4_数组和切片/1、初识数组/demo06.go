package main
import (
	"fmt"
)

func test(arr *[3]int) {
	for i, v := range(*arr){
		(*arr)[i] = v + 1
	}
}
func main() {
	arr := [3]int {1, 2, 3}
	fmt.Println(arr) 
	test(&arr)
	fmt.Println(arr)

}