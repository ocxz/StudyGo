package main
import (
	"fmt"
	"strconv"
	"strings"
)

// 输入数字，等到中文数字

func getCNindex(index int) string {
	cindex := [...]string {0:"零", 1:"一", 2:"二", 3:"三", 4:"四", 5:"五", 6:"六", 
					7:"七", 8:"八", 9:"九", 10:"十", 11:"百", 12:"千", 13:"万", 14:"亿"}
	// 将index各个元素解析，即取出index的每个数字，放在数组中
	// 123456789 == > [1 亿, 2, 3, 4 , 5, 6, 7, 8, 9]
	// 步骤：
	// 第一步：将数组转为string类型
	// 第二步：遍历字符串，将每个值转为in数组元素
	// 目的：一亿 二千 三百 四十 五万 六千 七百 八十 九
	// 逻辑上拆成：6，7，8，9    2， 3， 4， 9  万   1 亿
	strIndex := strings.Split(strconv.Itoa(index), "")  
	cnIndex := make([]int, len(strIndex))

	for index, value := range strIndex {
		cnIndex[index], _ = strconv.Atoi(value)
	}
	fmt.Println(cnIndex)
	var result string
	// len = 9   (len - index) / 3  ==> 3、2、1
	// 9 / 4.1 = 2  1 ==> 分在第二组  12 + 2
	// 8、7、6、5 / 4.1 ==> 分在第一组 12 + 1
	// 4、3、2、1  ==> 分在第0组   12 + 0
	for i, v := range cnIndex {
		// 获得当前数分在第几组
		// 组内处理
		// 12、11、10、9    取余：10
		// 8、7、6、5   第1组  加万 ==> 8后面加千 取余分别是：2、1、0  取余6
		// 4、3、2、1   第0组   取余：2  分别得到：
		// 六十七百八十九千
		index := len(cnIndex) - i   // 位置
		group:= int(float64(index) / 4.1)

		if index % (2 + 4 * group) == index {
			if group != 0 {
				result += cindex[v] + cindex[group + 12]
			} else {
				result += cindex[v]
			}
		} else {
			if index == 4 {
				result += cindex[v] + cindex[12]
			} else {
				result += cindex[v] + cindex[index % (2 + 4 * group) + 10]
			}
		}
		// switch{
		// case index % (2 + 4 * group) == index:
		// 	result += cindex[v] + cindex[group + 12]
		// case 
		// }
		
		// result += cindex[v] + cindex[group + 12]
	}
	return result
}

func main() {

	// 使用for-range结构遍历数组

	// cindex := [...]string {"一", "二", "三", "四", "五", "六", "七", "八", "九", 
	// 					   "十", "百", "千", "万", "亿"}
	// subjects := [...]string {"语文", "数学", "英语", "物理", "化学", "生物"}

	// for index, value := range subjects {
	// 	fmt.Printf("第%v门课是%v \n", cindex[index], value)
	// }
	fmt.Println(getCNindex(111126789))
}