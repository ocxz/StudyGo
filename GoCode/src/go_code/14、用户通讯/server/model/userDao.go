package model

import (
	"github.com/garyburd/redigo/redis"   // 引入redis包
	"encoding/json"
)

// 全局变量，定义一个userDao
var (
	UserDao *UserDao
)

// 定义一个UserDao 结构体
// 完成对user 结构体的各种操作 增删改查
type UserDao struct {
	Pool *redis.Pool
}

func CreateUserDao(pool *redis.Pool) {
	return &UserDao{ Pool : pool}
}

func InitUserDao(pool *redis.Pool) {
	userDao = CreateUserDao(pool)
}

// 根据用户userName，返回用户实例或error
func (this *UserDao) GetUserByName(name string) (user *User, err error){

	// 通过给定的name，去redis查询用户
	conn := this.Pool.Get()   // 从连接池中取出一个连接
	defer conn.Close()
	res, err := conn.Do("hget", "users", name)
	if err != nil {
		if err == redis.ErrNil {   // 表示不存在
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	
	// 将res反序列化
	err = json.Unmarshal(byte[](res), &user)
	if err != nil {
		fmt.Println("反序列化失败，失败原因：", err)
		return
	}
	return 
}

// 完成登录校验，用户存在且密码正确
// 如果用户名和密码都正确，则返回user实例，否则返回对应的错误信息
func (this *UserDao) Login(userName string, userPwd string) (user *User, err error) {
	
	user, err = this.GetUserByName(userName)
	if err != nil {   // 用户不正确
		return
	}
	if user.userPwd != userPwd {   // 密码不正确
		err = ERROR_USER_PWD
		return
	}
	return
}