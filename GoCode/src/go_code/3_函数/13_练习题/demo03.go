package main
import (
	"fmt"
	"math/rand"
	"time"
)

// 编写一个函数，生成一个随机1-100 的整数
// 利用random和UnixNano()
// 猜数字10发 一发入魂是天才 2-3发很聪明  4-9发一般般  最后一发可对来，一发未中说啥嘞

// 获得随机整数的方法 获得区间内随机一整数
// 左闭右闭
func GetRandInt(start int, end int) int {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end - start + 1) + start    // 随机生成1-100的整数
}

// 
func GuessNumber() {
	time := 0;        // 初始化次数，开始为0
	targetNum := GetRandInt(1, 100)   // 获得1-100之间的一个整数
	var num int;
	for {
		time++; // 次数加一
		fmt.Printf("目标数是（内测）：%v\n", targetNum)
		fmt.Print("请输入你猜的数字：")
		fmt.Scanln(&num)
		switch {
		case num == targetNum && time == 1:
			fmt.Println("一发入魂，真是个天才")
			return
		case num == targetNum && time >=2 && time <= 3:
			fmt.Println("两三发就中了，和爹一样聪明")
			return
		case num == targetNum && time >= 4 && time <= 9:
			fmt.Printf("%v发中，还说的过去吧\n", time)
			return
		case num == targetNum && time == 10:
			fmt.Println("唉……可算是中咯")
			return
		case num != targetNum && time == 10:
			fmt.Println("整整十发，一发未中，真是气死老子了")
			return
		}
	}
}

func main() {
	GuessNumber()
}