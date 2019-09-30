package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"   // 引入redis包
)

func main() {

	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接至redis服务器失败，失败原因：", err)
		return
	}

	defer conn.Close()

	// 向Redis写入数据 Hset key-val
	_, err = conn.Do("Hset","user01", "name", "张三丰")
	if err != nil {
		fmt.Println("向redis服务器写入数据失败，失败原因：", err)
		return
	}

	_, err = conn.Do("Hset","user01", "age", 18)
	if err != nil {
		fmt.Println("向redis服务器写入数据失败，失败原因：", err)
		return
	}

	// 从Redis获取Hset数据
	name, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("从redis服务器取出数据失败，失败原因：", err)
		return
	}
	// 返回的r是interface{}空空接口，我们需要转换
	fmt.Println("name = ", name)

		// 从Redis获取Hset数据
	age, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("从redis服务器取出数据失败，失败原因：", err)
		return
	}
	// 返回的r是interface{}空空接口，我们需要转换
	fmt.Println("age = ", age)


}