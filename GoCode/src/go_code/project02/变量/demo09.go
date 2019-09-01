package main

import (
	"fmt"
)

// 演示golang的数据类型默认值
func main() {

	var a int  // 0
	var b float32  //0
	var c float64  // 0
	var isMarrired bool  // false
	var name string   // ""

	fmt.Printf("a=%d, b=%f, c=%f, isMarried=%v, name=%v", a, b, c, isMarrired, name)
}
