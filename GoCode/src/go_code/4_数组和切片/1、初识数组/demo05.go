package main
import (
	"fmt"
)

func main() {

	// for-range 的使用演示
	heros := [...]string {"宋江", "武松", "卢俊义"}

	fmt.Println(heros)

	for i, v := range(heros) {
		fmt.Println(i,v)
	}

	myheros := heros 
	
	myheros[0] = "吴用"
	fmt.Println(heros)
	fmt.Println(myheros)
}
