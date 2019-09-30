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

	// 向Redis写入数据 string key-val
	_, err = conn.Do("Set", "name", "张三")
	if err != nil {
		fmt.Println("向redis服务器写入数据失败，失败原因：", err)
		return
	}
	fmt.Println("写入string key-val成功")

	// 从Redis获取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("从redis服务器取出数据失败，失败原因：", err)
		return
	}
	// 返回的r是interface{}空空接口，我们需要转换
	fmt.Println(r)
}