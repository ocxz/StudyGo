package testModel
import (
	"testing"
	"fmt"
	"go_code/9、单元测试/2、练习/model"
)

// 对NewMonster0() *Monster函数的测试
func TestNewMonster0(t *testing.T) {
	monster := model.NewMonster0()
	fmt.Println(monster)
}

// 对NewMonster() *Monster函数的测试
func TestNewMonster(t *testing.T) {
	monster := model.NewMonster("牛魔王", 18, "牛家拳")
	fmt.Println(monster)
}

// 对(m *Monster) Store(path string) bool 序列化方法的测试
func TestStore(t *testing.T) {

	m := model.NewMonster("牛魔王", 18, "牛家拳")
	path := `C:\Users\22527\Desktop/monster.txt`
	ok := m.Store(path)
	if !ok {
		t.Fatalf("序列化失败，期望值%v，实际值%v", 
		true, ok)
	}
	t.Logf("Monster序列化成功")
}

// 对(m *Monster) ReStore(path string) *Monster反序列化的测试
func TestReStore(t *testing.T) {

	path := `C:\Users\22527\Desktop/monster.txt`
	// 比较标准
	m := model.NewMonster("牛魔王", 18, "牛家拳")
	rm := m.ReStore(path)
	if m != rm {
		t.Fatalf("反序列化失败，期望值=%v，实际值=%v", m, rm)
	}
	t.Logf("Monster反序列化成功%v", rm)
}