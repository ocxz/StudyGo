package main
import (
	"fmt"
)

func main() {

	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	hens[3] = 6.0
	hens[4] = 7.0
	hens[5] = 8.0

	total := 0.0
	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}

	// 平均
	avg := fmt.Sprintf("%.2f", total / float64(len(hens)))
	fmt.Printf("总数为：%v，平均数为：%v", total, avg)
}