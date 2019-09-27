package main
import (
	"fmt"
	"runtime"
)

func main() {

	// 设置多个cpu执行此程序
	numCpu := runtime.NumCPU()
	fmt.Printf("本机上有%v个cpu\n", numCpu)

	// 自己设置本程序使用多少个cpu，我这个地方预留一个
	runtime.GOMAXPROCS(numCpu - 1)
	fmt.Println("ok")
}