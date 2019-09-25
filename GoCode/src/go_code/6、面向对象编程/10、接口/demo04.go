package main
import (
	"fmt"
	"sort"
	"math/rand"
)

// 声明Hero结构体
type Hero struct {
	Name string
	Age int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

// HeroSlice实现三个方法，Len()、Less()、Swap()
// 确定要排序的长度
func (heros HeroSlice) Len() int {
	return len(heros)
}

// 排序规则，如可以按年龄从小到大排序，返回值为真不交换，为假交换
func (heros HeroSlice) Less(i, j int) bool {
	return heros[i].Age < heros[j].Age
}

// 交换
func (heros HeroSlice) Swap(i, j int) {
	// temp := heros[i]
	// heros[i] = heros[j]
	// heros[j] = temp
	// 等价于
	heroes[i], heroes[j] = heroes[j], heroes[i]
}


func main() {

	// 排序
	intSlice := []int { 0, -1, 10, 7, 90}

	// 1.冒泡排序
	// 2.系统定义的排序方法
	sort.Ints(intSlice)   // 升序
	fmt.Println(intSlice)

	// 对结构体进行排序
	// 测试，对结构体类型进行排序
	var heroes HeroSlice
	for i :=0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		// 将hero append到heroes切片中
		heroes = append(heroes, hero)
	}

	for _, v := range heroes {
		fmt.Println(v)
	}

	fmt.Println("------------排序后------------")
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}