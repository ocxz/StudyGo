package main
import (
	"fmt"
)

/*
	要求：编写一个方法，判断是奇数还是偶数(even)
*/

type Integer int

func (integer *Integer) IsEven() bool {
	return *integer % 2 == 0
}

func main() {

	var i Integer = 0
	fmt.Println(i.IsEven())
}