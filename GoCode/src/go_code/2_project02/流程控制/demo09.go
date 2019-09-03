package main

import (
	"fmt"
)

func main() {

	// // switch小练习
	// // 1、使用switch 将小写char 类型转为大写（键盘输入） 只转换a-z 其他的输出other a->97 A->65
	// var myChar byte
	// fmt.Print("请输入a-z小写字母：")
	// // fmt.Scanln(&myChar)   不要使用这个，默认不是以char来接收的
	// fmt.Scanf("%c", &myChar)
	// switch {
	// case myChar >=97 && myChar <= (97+65):
	// 	fmt.Printf("%c 对应的大写字母是：%c", myChar, myChar-32)
	// default:
	// 	fmt.Printf("other")
	// }

	// // 2、成绩输入大于60合格、小于60不合格、不能大于100或小于0
	// var score byte
	// fmt.Print("\n请输入成绩：")
	// fmt.Scanf("%d", &score)
	// switch {
	// 	case score >100 || score <0 :
	// 		fmt.Printf("成绩%v输入不合法", score)
	// 	case score >=60 :
	// 		fmt.Printf("成绩%v分，合格", score)
	// 	default:
	// 		fmt.Printf("成绩%v分，不合格", score)
	// }

	// // 3、根据月份，打印季节  3、4、5春季  6、7、8 夏季   9、10、11 秋季  12、1、2 冬季
	// var month byte
	// fmt.Print("请输入月份（1-12）：")
	// fmt.Scanf("%d", &month)
	// switch month {
	// 	case 3, 4, 5 :
	// 		fmt.Printf("%v月是春季", month)
	// 	case 6, 7, 8 :
	// 		fmt.Printf("%v月是夏季", month)
	// 	case 9, 10, 11 :
	// 		fmt.Printf("%v月是秋季", month)
	// 	case 12, 1, 2 :
	// 		fmt.Printf("%v月是冬季", month)
	// 	default:
	// 		fmt.Printf("输入的月份不准确，无法判断")
	// }

	// 根据用户输入，显示对应的星期时间(string) 
	//如果是星期一：干煸豆角 星期二：酸醋土豆 星期三：红烧狮子头 星期四：油炸花生米
	// 星期五：蒜蓉扇贝  星期六：东北乱炖  星期日：大盘鸡
	var day string
	fmt.Print("请输入星期：")
	fmt.Scanf("%s", &day)
	switch day {
		case "一", "1", "星期一", "周一", "礼拜一":
			fmt.Printf("今天是周一，好吃的是干煸豆角")
		case "二", "2", "星期二", "周二", "礼拜二":
			fmt.Printf("今天是周二，好吃的是酸醋土豆")
		case "三", "3", "星期三", "周三", "礼拜三":
			fmt.Printf("今天是周三，好吃的是红烧狮子头")
		case "四", "4", "星期四", "周四", "礼拜四":
			fmt.Printf("今天是周四，好吃的是油炸花生米")
		case "五", "5", "星期五", "周五", "礼拜五":
			fmt.Printf("今天是周五，好吃的是蒜蓉扇贝")
		case "六", "6", "星期六", "周六", "礼拜六":
			fmt.Printf("今天是周六，好吃的是东北乱炖")
		case "七", "7", "天", "日", "星期七", "星期日", "星期天", "周日", "周天", "礼拜天":
			fmt.Printf("今天是周日，好吃的是大盘鸡")
		default:
			fmt.Printf("输入有误，无法为您推荐健康美味的食材~")
	}
}