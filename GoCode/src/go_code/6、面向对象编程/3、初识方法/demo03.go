package main
import (
	"fmt"
)

type Integer int

func (i *Integer) print() {
	fmt.Println("i=",*i)
}
func (i *Integer) change() {
	*i += 1
}

func main() {

	var i Integer = 10
	i.print()
	i.change()
	i.print()
}