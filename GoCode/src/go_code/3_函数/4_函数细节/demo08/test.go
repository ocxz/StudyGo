package main

import (
	"fmt"
	"go_code/3_函数/4_函数细节/utils"
)

func init() {
	fmt.Println("执行了Test中的函数")
}

func main() {
	fmt.Println("执行来Test中的main函数")

	utils.Test()
}