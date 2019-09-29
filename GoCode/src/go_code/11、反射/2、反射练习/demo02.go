package main
import (
	"fmt"
	"reflect"
)

func main() {

	str := "tom"
	rValue := reflect.ValueOf(str)
	fmt.Printf("rValue = %v，rValue的type = %T\n", rValue, rValue)

	rElem := reflect.ValueOf(&str).Elem()
	fmt.Printf("rElem = %v，rElem的type = %T\n", rElem, rElem)
	
	rElem.SetString("jace")
	fmt.Println(str)

	reflect.ValueOf(&str).Elem().SetString("马云")
	fmt.Println(str)
}