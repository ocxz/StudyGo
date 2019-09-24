package main
import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string) {

	// 判断用户是否存在
	if users[name] != nil {
		// 用户存在
		users[name]["pwd"] = "888888"
	} else {
		// 用户不存在
		users[name] = map[string]string{
			"nickName" : "昵称~" + name,
			"pwd" : "888888",
		}
	}
}
func main() {

	// 使用map[string]map[string]string 数据类型
	// key表示用户名，唯一，不可重复
	// 如果key已存在，则将其密码修改成"88888"，不存在则添加用户，包括昵称nickName、密码pwd
	// 编写一个函数 modifyUser(users map[string]map[string]string, name string)完成上述功能

	users := make(map[string]map[string]string)
	users["smith"] = map[string]string {
		"nickName" : "小花猫",
		"pwd" : "123456",
	}

	modifyUser(users, "tom")
	modifyUser(users, "mary")
	fmt.Println(users)

	for name, info := range users {
		fmt.Println("用户：" + name)
		for key, value := range info {
			fmt.Printf("\t %v : %v \n", key, value)
		}
	}
}