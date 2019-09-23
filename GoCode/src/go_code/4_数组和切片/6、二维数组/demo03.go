package main
import (
	"fmt"
)

func main() {

	// 定义二维数组，保存三个班，每个班5名成绩
	// 求出每个班平均分、以及所有班平均分
	var grades [3][5]float64
	
	for i, v := range grades{
		fmt.Printf("请输入第%v个班级同学的成绩\n", i+1)
		for j, _ := range v {
			fmt.Printf("%v班第%v个同学的成绩：", i+1, j+1)
			fmt.Scanln(&grades[i][j])
		}
	}

	fmt.Println(grades)
}