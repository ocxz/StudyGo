package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {

	var p Person
	fmt.Println(p)
}