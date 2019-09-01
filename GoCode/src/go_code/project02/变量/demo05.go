package main

// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main(){

	// var price float32 = 89.12
	// var num1 float64 = 89.12

	// fmt.Println(price)
	// fmt.Println(num1)

	// 默认声明为float64
	var num = 89.63
	fmt.Printf("num的数据类型是：%T，所占用的字节数是：%d\n", num, unsafe.Sizeof(num))

	num, num2, num3 := 5.12, .123, 5e2
	fmt.Println("num=", num, "num2=", num2, "num3=", num3)
}