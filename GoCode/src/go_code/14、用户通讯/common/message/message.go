package message

// 定义一下常量，保存消息类型
const (
	LoginMsgType = "LoginMsg"      // 登录信息
	ResultMsgType = "ResultMsg"      // 登录结果信息
	RegisterMsgType = "RegisterMsg"
)


// 消息结构体
type Message struct {

	Type string `json:"type"`    // 消息类型
	Data string `json:"data"`   // 消息内容|具体消息的反序列化串
}

// 结果信息结构体
type ResultMsg struct {

	Code int `json:"code"`   // 状态码  500 表示该用户不存在  200 表示登录成功 400表示密码错误
	Error string `json:"error"`  // 返回错误信息

}

// 登录信息结构体
type LoginMsg struct {

	UserName string `json:"userName"`   // 用户名
	UserPwd string  `json:"userPwd"`  // 用户密码
}
// 注册信息结构体
type RegisterMsg struct {

}

