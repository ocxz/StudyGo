package main

import (
	"fmt"

	// 引入Utils包  本来是要写F:\Go语言\GoCode\src\go_code\project02\标识符\Utils 但是我们之前配置了GOPATH=F:\Go语言\GoCode
	// 而它默认会添加一个src目录，因此我们只需要写go_code/project02/标识符/Utils
	// 也可以使用相对路径  如：../Utils
	"go_code/project02/标识符/Utils"   // 使用绝对路径 
	// "../Utils"        // 使用相对路径
)

func main(){

	// 使用Utils包中的变量  包名.变量

	fmt.Println(My Utils.HeroName)
}