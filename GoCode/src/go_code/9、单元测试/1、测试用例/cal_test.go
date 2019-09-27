package main
import (
	_ "fmt"
	"testing"   // 引入go的testing框架包
)

// 编写一个测试用来，测试AddUpper()函数是否正确

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("AddUpper(10)执行错误，期望值=%v，实际值=%v\n", 55, res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v，实际值=%v\n", 55, res)
	}

	// 如果正确
	t.Logf("AddUpper(10)执行正确")
}

func TestGetSub(t *testing.T) {

	res := getSub(10, 5) + 1
	if res != 5 {
		t.Fatalf("getSub(10, 5)执行错误，期望值=%v，实际值=%v", 5, res)
	}

	// 执行正确
	t.Logf("getSub(10, 5)执行正确")
}