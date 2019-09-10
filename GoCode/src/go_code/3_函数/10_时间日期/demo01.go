package main
import (
	"fmt"
	"time"
)

func main() {

	// 日期和时间相关函数和方法的使用
	// 1、获取当前时间 time.Now()
	now := time.Now()
	fmt.Printf("now=%v，now的类型是%T \n", now, now)

	// 2、年月日、时分秒的获取
	year := now.Year()
	fmt.Printf("year=%v，year的类型是%T \n", year, year)

	var month time.Month = now.Month()
	fmt.Printf("month=%v，month的类型是%T \n", month, month)

	day := now.Day()
	fmt.Printf("day=%v，day的类型是%T \n", day, day)

	hour := now.Hour()
	fmt.Printf("hour=%v，hour的类型是%T \n", hour, hour)

	minute := now.Minute();
	fmt.Printf("minute=%v，minute的类型是%T \n", minute, minute)

	second := now.Second()
	fmt.Printf("second=%v，second的类型是%T \n", second, second)

	// 3、日期的格式化 利用日期.Format("2006/01/02 15:04:05")进行格式化
	var nowString string
	nowString = now.Format("2006/01/02 :15:04:05")
	fmt.Printf("当前时间为：%v \n", nowString)

	nowString = now.Format("06/01/02 :15:04:05")
	fmt.Printf("当前时间为：%v \n", nowString)

	nowString = now.Format("006/01/02 :15:04:05")
	fmt.Printf("当前时间为：%v \n", nowString)

	nowString = now.Format("2006-01-02")
	fmt.Printf("当前时间为：%v \n", nowString)

	nowString = now.Format("2006-1-2")
	fmt.Printf("当前时间为：%v \n", nowString)
}