package main
import (
	"fmt"
)

type Person struct {
	Age int   // 联系存储，地址相差8
	Gender int
}

func main() {

	person := Person { 1, 2}
	fmt.Println(person)
	fmt.Printf("person的地址%p\n", &person)
	fmt.Printf("person中Age的地址%p\n", &person.Age)
	fmt.Printf("person中Gender的地址%p\n", &person.Gender)
}