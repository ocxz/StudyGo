package main
import (
	"fmt"
	"time"
)

var formTime string = "2006-1-2"
var startDateStr string = "1990-1-1"
func FishOrNet() {
	startDate, _ := time.Parse(formTime, startDateStr)  // 开始执行的日期
	startDateUnix := startDate.Unix()
	var checkDateUnix int64
	var inputDateStr string    // 输入的日期字符串
	for {
		fmt.Print("请按照 "+startDateStr+" 日期格式输入日期：")
		fmt.Scanln(&inputDateStr)
		inputDate, err := time.Parse(formTime, inputDateStr)
		if err != nil {
			fmt.Println("日期输出格式不正确，请重新输入")
		} else if inputDate.Unix() < startDateUnix {
			fmt.Println("日期必须是"+startDateStr+"之后，请重新输入")
		} else {
			checkDateUnix = inputDate.Unix()
			break   // 跳出循环
		}
	}
	
	durDateUnix := (checkDateUnix - startDateUnix) / (24 * 3600) + 1
	switch durDateUnix % 5 {
	case 1, 2, 3:
		fmt.Println(inputDateStr + "打了渔")
	case 0,4:
		fmt.Println(inputDateStr + "晒了网")
	}
}
// 三天打鱼，两天晒网，从1990年1月1日开始执行，判断以后某天是打鱼还是晒网
func main() {

	for {
		FishOrNet()
	}
}