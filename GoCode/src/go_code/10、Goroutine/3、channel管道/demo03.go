package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}
func main(){

	// 声明一个可以存放任意类型的管道
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	allChan <- Cat {
		Name : "tom",
		Age : 4,
	}

	// 我们只希望获取channel中第三个数据，可以将前两个数据扔掉
	<- allChan
	<- allChan
	cat := <- allChan  // 从管道中取出猫猫
	
	// 因为已经到了运行阶段，会自动将interface{}类型的猫猫强转为Cat类型
	fmt.Printf("cat = %v，cat的类型 = %T\n", cat, cat)  // 结果：cat = {tom 4}，cat的类型 = main.Cat

	// 试图取出猫猫的名字
	//直接cat.Name->抛异常：cat.Name undefined (type interface {} is interface with no methods)
	// 编译通不过
	// 使用类型断言，将interface{}类型的猫猫转换成猫猫类型
	newCat := cat.(Cat)
	fmt.Println("cat的名字是：", newCat.Name) 
	
}