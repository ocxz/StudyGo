package main

// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)

func main(){

	// var i int8 = -129   // 超出范围，会报错
	// var i uint8 = -2    // 也是会提示超出范围，报错

	// int 其他类型的使用
	 
	// byte/uint8 可用来表示单字符的ascii值  一般用byte存储unicode字符(0~255)
	// var b byte ='a'	
	// fmt.Println(b)

	// rune/int32 可用来表示中文的单个字的ascii值
	// var r rune = '程'
	// fmt.Println(r)

	// var r byte = 'a'
	// fmt.Printf("r 的类型是：%T，占用的字节数：%d", r, unsafe.Sizeof(r))

	var s = "程校长"
	fmt.Printf("s的类型是：%T，占用的字节数：%d", s, unsafe.Sizeof(s))
}