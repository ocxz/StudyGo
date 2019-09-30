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

	// 向Redis写入批量数据 Hmset key-val
	_, err = conn.Do("Hmset","user02", "name", "张三丰", "age", 18, "sex", "男")
	if err != nil {
		fmt.Println("向redis服务器写入数据失败，失败原因：", err)
		return
	}
	
	// 从Redis获取Hmset数据，获得的是[]string
	r, err := redis.Strings(conn.Do("Hmget", "user02", "name", "age", "sex"))
	if err != nil {
		fmt.Println("从redis服务器取出数据失败，失败原因：", err)
		return
	}

	fmt.Printf("r = %v，r的类型是：%T\n", r, r)
	for i, v := range r {
		fmt.Printf("r[%v] = %v\n", i, v)
	}
}