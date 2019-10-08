package main
import (
	"github.com/garyburd/redigo/redis"   // 引入redis包
	"time"
)

// 定义一个全局pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func initPool(connStr string, maxIdle int, maxActive int, idelTimeout time.Duration) {

	pool = &redis.Pool{

		MaxIdle : maxIdle,   // 最大空闲连接数
		MaxActive : maxActive,  // 最大活跃连接数，0表示不受限制
		IdleTimeout : idelTimeout,  // 最大空余时间
		Dial : func() (redis.Conn, error) {
			return redis.Dial("tcp", connStr)
		},
	}
}

type RedisManage struct {}
func CreateRM *RedisManage {
	return &RedisManage{}
}

func (this *RedisManage) GetReidsConn() redis.Conn {
	return pool.Get()
}