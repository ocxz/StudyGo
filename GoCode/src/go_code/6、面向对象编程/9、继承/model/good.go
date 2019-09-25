package model
import (
	// "fmt"
)

type Goods struct {
	Name string
	Price int
}

type Book struct {
	Goods // 这里就是嵌套匿名结构体Goods，就拥有了Goods的字段和方法
	Writer string
}