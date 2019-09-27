package main
import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	
)

/*
	要求：1、创建一个Person结构体（Name, Age, Address）
		 2、 使用rand方法，随机创建10个Person实例，并放入到channel中
		 3、遍历channel，将各个Person实例的信息显示在终端
*/

type Person struct {
	Name string
	Age int
	Address string
}

// 随机生成一个Person实例，将其指针返回
func getRandPerson() *Person {
	age := rand.Intn(100) + 1
	name := "Person" + strconv.Itoa(age)
	address := fmt.Sprintf("北京，三环，%v号", age)
	return &Person{
		Name : name,
		Age : age,
		Address : address,
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())  // 设置随机种子

	// 随机创建10个Person实例，并放入到channel中
	cap := 10
	personChan := make(chan *Person, cap)
	for i := 0; i < cap; i++ {
		personChan <- getRandPerson()
	}

	// 循环取出，personChan的内容，并打印出来
	for i := 0; i < 10; i++ {
		fmt.Println(*( <- personChan))
	}
}