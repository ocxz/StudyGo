package main
import (
	"github.com/garyburd/redigo/redis"   // 引入redis包
	"fmt"
)

// 定义一个全局pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{

		MaxIdle : 8,   // 最大空闲连接数
		MaxActive : 0,  // 最大活跃连接数，0表示不受限制
		IdleTimeout : 100,  // 最大空余时间
		Dial : func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}


func main() {

	conn := pool.Get()   // 从连接池中取出一个连接
	defer conn.Close()

	_, err := conn.Do("set", "name", "汤姆~")
	if err != nil {
		fmt.Println("set错误，", err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get错误，", err)
		return
	}

	fmt.Println("r = ", r)
}