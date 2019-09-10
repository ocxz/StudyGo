package main
import (
	"fmt"
	"math"
)

// 输出100以内所有素数，每行显示五个，并求和
// 只能被1和自己整除的数，叫做素数 
// 做法：两层循环，第一次循环取出1-100数
// 第二层：判断是否为素数

// 这个函数用来判断，一个数是否为素数
func IsPrimeNum(num int) bool {
	if num == 1 || num == 2 || num == 3{   // 1、2、3直接返回true
		return true
	}else {
		sqrtNum := int(math.Sqrt(float64(num)))   // 求开方的，化为整数
		for i := 2; i <= sqrtNum; i++ {    // 循环判断，取余
			if num % i == 0 {       // 有一个为整数，不是素数
				return false
			}
		}
		return true
	}
}

func main() {
	count := 0   // 记录素数总个数
	sum := 0    // 记录素数和
	num := 112
	for i := 1; i <= num; i++ {
		if IsPrimeNum(i) {   // 如果i是素数
			count++
			sum += i
			fmt.Printf("%v\t", i)  
			if count % 5 == 0 {    // 如果素数的个数是5的倍数，则换行
				fmt.Println()
			}
		}

		if count % 5 != 0 && i == num {    // 最后，如果不足五个，前打印完成，则换行
			fmt.Println()
		}
	}
	fmt.Printf("1-%v间的素数总和为：%v", num, sum)
}

