package regiorm
import (
	"github.com/garyburd/redigo/redis"   // 引入redis包
	"fmt"
	"errors"
)
// 声明定义全局的连接池对象
var pool *redis.Pool
const (
	ptl = "tcp"   // 协议
	ip = "127.0.0.1"
	port = "6379"
)
type Orm struct {

	addr string      // 要连接的redis服务器ip
	conn redis.Conn    // 持有的连接对象
}

func init() {
	pool = &redis.Pool{

		MaxIdle : 8,   // 最大空闲连接数
		MaxActive : 0,  // 最大活跃连接数，0表示不受限制
		IdleTimeout : 100,  // 最大空余时间
		Dial : func() (redis.Conn, error) {
			return redis.Dial(ptl, ip + ":" + port)
		},
	}
	fmt.Println("连接池对象初始化成功")
}

// 创建一个orm对象
func NewOrm() *Orm {

	orm := &Orm{addr : ip}
	orm.conn = pool.Get()
	return orm
}

// 执行一条redis命令
func (orm *Orm) Run(cmd string, args ...interface{}) (interface{}, error) {

	r, err := orm.conn.Do(cmd, args...)
	if err != nil {
		fmt.Println("执行redis命令错误，错误原因：", err)
		return nil, err
	}
	return r, err
}

// 将对象存入Hash中
func (orm *Orm) SaveHash(key string, val interface{}, conver bool) (interface{}, error) {

	//判断key是否已存在
	keys, _ := redis.Strings(orm.Run("keys", key))
	if ExistSlice(key, keys) > -1 && !conver {     // 已存在，不覆盖，那就不运行
		return nil, errors.New("该对象已存在，添加失败")
	}

	args := make([]interface{}, 0)
	args = append(GetParams(val)[:1], GetParams(val)[:]...)
	args[0] = key
	return orm.Run("hmset", args...)
}

// 传入的是指针类型|结构体类型
// func (orm *Orm) GetHash(key string, s interface{}) {

// 	rValue := s.Elemt()
// }

// 关闭orm，即关闭orm中的conn对象
func (orm *Orm) Close() {
	orm.conn.Close()
}