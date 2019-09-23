package main
import (
	"fmt"
)

func main() {

	// 演示切片的make创建使用
	var slice []float64 = make([]float64, 5, 10)
	fmt.Println(slice)

	// 演示创建切片时，直接指定数组的方式
	var strSlice []string = []string {"tom", "jack", "marry"}
	fmt.Println(strSlice)
	fmt.Println("len=", len(strSlice))
	fmt.Println("cap=", cap(strSlice))
}