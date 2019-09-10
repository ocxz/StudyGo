package main
import (
	"fmt"
)

func PrintAllChar(start byte, count int, args...bool) {

	if len(args) > 0 && args[0] {
		for i := count - 1; i >= 0; i-- {
			fmt.Printf("%c ", start + byte(i))
		}
	} else {
		for i := 0; i < count; i++ {
			fmt.Printf("%c ", start + byte(i))
		}
	}
	fmt.Println()
}

func main() {
	
	fmt.Print("打印a-z：")
	PrintAllChar('a', 26, false)

	fmt.Print("打印Z-A：")
	PrintAllChar('A', 26, true)
}