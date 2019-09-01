package main

// import "fmt"
// import "unsafe"
import(
	"fmt"
)

// 演示golang中字符类型
func main(){

	// var c1 byte = 'a'
	// var c2 byte = '0'

	// // 直接输出byte值时，就是输出对字符对应的acsii码值
	// fmt.Println("c1=", c1, "c2=", c2)
	// fmt.Printf("c1=%c, c2=%c\n", c1, c2)

	// var r rune = '北'
	// var r2 rune = '京'
	// fmt.Printf("%c对应的码值是%d\n", r, r)
	// fmt.Printf("%d, %c", r+r2, r+r2)

	var n int = 4E00
	fmt.Printf("%d", n)
}