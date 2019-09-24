package main
import (
	"fmt"
	"sort"
)

func main() {

	// map的排序
	map1 := map[int]int {
		10 : 100,
		1 : 13,
		2 : 56,
		58 : 39,
		8 : 90,
		11 : 98,
	}
	fmt.Println(map1)

	map2 := map[string]int {
		"10" : 100,
		"1" : 13,
		"2" : 56,
		"58" : 39,
		"8" : 90,
		"11" : 98,
		"16" : 98,
	}
	fmt.Println(map2)

	map3 := map[string]int {
		"name" : 100,
		"age" : 13,
		"sex" : 56,
		"gender" : 39,
		"class" : 90,
		"hello" : 98,
		"ok" : 98,
	}
	fmt.Println(map3)
	fmt.Println()
	for key, value := range map3 {
		fmt.Printf("map3[%v] = %v \n", key, value)
	}
	fmt.Println()

	// 解决方法，先将map的key放入切片中，对切片进行排序，遍历切片，按照key来输出map的值
	var keys[] int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map[%v] = %v \n", k, map1[k])
	}
	fmt.Println()
	for key, value := range map1 {
		fmt.Printf("map[%v] = %v \n", key, value)
	}
}