package regiorm
import (
	"go_code/13、Redis的使用/2、orm"
	"testing"
	"fmt"
	"github.com/garyburd/redigo/redis"   // 引入redis包
)

// 测试创建orm函数
func TestNewOrm(t *testing.T) {
	orm := regiorm.NewOrm()
	defer orm.Close()
	fmt.Println(*orm)
}

// 测试执行redis命令的Run方法
func TestRun(t *testing.T) {
	orm := regiorm.NewOrm()
	defer orm.Close()
	orm.Run("hmset", "stu1", "name", "张三", "age", 18, "sex", "男")
	r, _ := redis.String(orm.Run("hget", "stu1", "name"))
	fmt.Println("name = ", r)
	TestSaveHash(t)
}

// func GetStructInfos(s interface{}) res []interface{}
func TestGetParams(t *testing.T) {

	stu := Student{
		Name : "张三",
		Age : 18,
		Sex : "男",
	}
	fmt.Println(regiorm.GetParams(&stu))
	// fmt.Println(regiorm.GetParams(3))	
}

// 测试Hash存储
func TestSaveHash(t *testing.T) {

	stu := Student{
		Name : "张三",
		Age : 18,
		Sex : "男",
	}

	orm := regiorm.NewOrm()
	defer orm.Close()
	_, err := orm.SaveHash("stu1", &stu, false)
	if err != nil {
		fmt.Println("存储失败，失败原因：", err)
	} else {
		fmt.Println("存储成功")
	}
}

// 测试ExistSlice(val interface{}, slice []interface{}) int
func TestExistSlice(t *testing.T) {

	key := "stu"
	keys := []string{
		"stu",
		"stu2",
	}
	fmt.Println(regiorm.ExistSlice(key, keys))
}