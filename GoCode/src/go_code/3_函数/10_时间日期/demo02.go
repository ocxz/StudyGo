package main
import (
	"fmt"
	"time"
)

func main() {

	// 时间的常量
	// sleep结合实际常量，完成每个1秒打印一个数字，打印100时就退出

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println(i)

	// 	// 休眠 1 秒钟
	// 	time.Sleep(100 * time.Microsecond)
	// }

	// Unix和UnixNano的使用
	var nowUnix int64 = time.Now().Unix()
	var nowUnixNano int64 = time.Now().UnixNano()
	fmt.Printf("当前秒时间戳：%v，纳秒时间戳：%v", nowUnix, nowUnixNano)
}