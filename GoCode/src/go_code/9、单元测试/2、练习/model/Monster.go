package model
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {

	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

// 工厂模式初始化结构体变量
func NewMonster0() *Monster {
	return &Monster{}
}

func NewMonster(name string, age int, skill string) *Monster {
	return &Monster{
		Name : name,
		Age : age,
		Skill : skill,
	}
}

// 序列化
func (m *Monster) Store(path string) bool {

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Monster结构体序列化失败，失败原因：\n", err)
		return false
	}
	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		fmt.Println("存储文件打开失败，失败原因：\n", err)
		return false
	}
	return true
}

// 反序列化
func (m *Monster) ReStore(path string) *Monster {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：\n", err)
		return nil
	}
	err = json.Unmarshal(data, &m)

	if err != nil {
		fmt.Println("Monster反序列化失败，失败原因：\n", err)
		return nil
	}
	return m
}