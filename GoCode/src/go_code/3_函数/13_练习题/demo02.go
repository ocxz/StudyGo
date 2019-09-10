package main
import (
	"fmt"
	"errors"
)

// 循环输入年、月、日打印日期，使用continue实现
// 要判断输入的月份是否错误的语句
// 年份：1970-2100
// 月份：1-12
// 日期：1-29、1-30、1-31 根据闰年判断
// 1 3 5 7 8 10 12 31天
// 2  闰年：29  平年 28
// 其他 30 天
var  isContinue bool = false
func inputYear() (year int) {
	defer func() {
		err := recover();
		if err != nil {
			fmt.Println(err)
			isContinue = true;
		}
	}()
	fmt.Print("请输入年份（1970-2100）：")
	fmt.Scanln(&year)
	if year < 1970 || year > 2100 {
		err := errors.New("年份范围不准确，应该是1970年-2100年")
		panic(err)
	}
	isContinue = false
	return year
}

func inputMonth() (month int) {
	defer func() {
		err := recover();
		if err != nil {
			fmt.Println(err)
			isContinue = true;
		}
	}()
	fmt.Print("请输入月份（1-12）：")
	fmt.Scanln(&month)
	if month < 1 || month > 12 {
		err := errors.New("月份范围不准确，应该是1月-12月")
		panic(err)
	}
	isContinue = false
	return month
}

// 日期：1-29、1-30、1-31 根据闰年判断
// 1 3 5 7 8 10 12 31天
// 2  闰年：29  平年 28
// 其他 30 天
func inputDay(year int, month int) (day int) {
	defer func() {
		err := recover();
		if err != nil {
			fmt.Println(err)
			isContinue = true;
		}
	}()
	 isLeapYear := year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
	 var maxDay int    // 当前年月的最大天数
	 switch month{
		case 1, 3, 5, 7, 8, 10, 12:
			maxDay = 31
		case 2:
			if isLeapYear {
				maxDay = 29
			} else {
				maxDay = 28
			}
		default:
			maxDay = 30
	 }
	 fmt.Printf("请输入日期（1-%v）：", maxDay)
	 fmt.Scanln(&day)
	 if day < 1 || day > maxDay {
		 err := errors.New(fmt.Sprintf("日期范围不准确，应该为1日-%v日", maxDay))
		 panic(err)
	 }
	 isContinue = false
	 return day
}


// 打印时间，通过异常处理
func PrintTime2() {
	reback:
	for {   // 循环打印输出年
		year := inputYear()
		if isContinue {
			continue
		}
		for {
			month := inputMonth()
			if isContinue {
				continue
			}
			for {
				day := inputDay(year, month)
				if isContinue {
					continue
				}
				fmt.Printf("输入的日期：%v年%v月%v日 \n", year, month, day)
				continue reback
			}
		}
	}
}

func main() {
	
	PrintTime2()
}