package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1、len 统计字符串长度
	var strLen int = len("hello 北京")
	fmt.Println("strLen=", strLen)

	// 2、遍历函数 将字符串转换为切片，然后进行遍历
	r := []rune("hello 北京")
	for i := 0; i < len(r); i++ {
		fmt.Printf("r[%v] = %c\n", i, r[i])
	}

	// 3、字符串转整数 需要 n, err := strconv.Atoi(str)
	// err不等于nil说明有错
	n, err := strconv.Atoi("123")
	if err == nil {
		fmt.Printf("n=%v，n的类型是%T", n, n)
	} else {
		fmt.Println("转换失败，错误信息为：", err)
	}

	// 4、整数转字符串 str := strconv.Itoa(123456)
	str := strconv.Itoa(123456)
	fmt.Println("str=", str)

	// 5、字符串转byte[]：bytes := []byte("hello 中国")
	bytes := []byte("hello 中国")
	fmt.Println("bytes=", bytes)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%c", bytes[i])
	}
	fmt.Println()

	// 6、十进制转2、8、16进制 str := strconv.FormatInt(数, 进制)
	str2 := strconv.FormatInt(361, 2)
	str3 := strconv.FormatInt(19, 16)
	fmt.Printf("361的二进制数是：%v\n", str2)
	fmt.Printf("19的十六进制数是：%v\n", str3)

	// 7、查找子串是否包含在父串中 b := strings.Contains("seafood", "foo")  // true
	b := strings.Contains("seafood", "foo")  // true
	fmt.Println("b=", b)

	// 8、统计子串的个数 var count int = strings.Count("ceheeseeeee", "ee")  //3
	var count int = strings.Count("ceheeseeeee", "f")  //3
	fmt.Println("count=", count)

	// 9、字符串比较（不区分大小写） b ：= strings.EqualFold("abc", "AbC")  // true
	b2 := strings.EqualFold("abc", "AbC")  // true  不区分大小写
	b3 := "abc" == "Abc"   // false 区分大小写
	fmt.Println("b2=", b2)
	fmt.Println("b3=", b3)

	// 10、子串第一次出现的位置 var index int = strings.Index("NLT_abc", "abc")  // 4
	var index int = strings.Index("NLT_abc", "abc")  // 4
	fmt.Println("index=", index)

	// 11、子串最后一次出现的位置 var index int = strings.LastIndex("go goLang", "go") // 3
	var index2 int = strings.LastIndex("go goLang", "go") // 3
	fmt.Println("index2=", index2)

	// 12、字符串替换 var str string = strings.Replace("go go hello", 要替换, 替换成, int n)
	var str12 string = strings.Replace("go go hello", "go", "中国", 3)
	fmt.Println("str12=", str12)

	// 13、分割字符串 var []string strs = strings.Splitt("Hello, world：欢迎来到中国",",：")
	var strs = strings.Split("Hello, world, 欢迎来到中国",",")
	for i := 0; i < len(strs); i++ {
		fmt.Printf("strs[%v] = %v \n", i, strs[i])
	}

	// 14、去空格或者字符串两边
	// 1、去左右两边空格 var str string = string.TrimSpace(" 这是一个 美丽 的  误会   "）
	var strTrimSpace string = strings.TrimSpace(" 这是一个 美丽 的  误会    ")
	fmt.Printf(strTrimSpace+"%v", "啦啦啦\n")
	// 2、去左字符 var str string = strings.TrimLeft("!hello!", "!")  // 结果hello!
	var strTrimLeft string = strings.TrimLeft("!hello!", "!")  // 结果hello!
	fmt.Printf("中国%v \n", strTrimLeft)
	// 3、去右字符 var str string = strings.TrimRight("!hello!", "!")   // 结果!hello
	var strTrimRight string = strings.TrimRight("!hello!", "!")   // 结果!hello
	fmt.Printf("%v中国 \n", strTrimRight)
	// 4、去左右两边字符 var str string = strings.Trim("!hello!", "!")  // 结果hello
	var strTrim string = strings.Trim("!hello!!", "!")  // 结果hello
	fmt.Printf("中%v国\n", strTrim)

	// 17、判断是否指定字符开头 var b bool = strings.HasPrefix("ftp://127.0.0.1")  // true
	var hasFtpPre bool = strings.HasPrefix("ftp://127.0.0.1", "ftp")  // true
	fmt.Println(hasFtpPre)

	// 18、判断是否指定字符结尾 var b bool = strings.HasSuffix("NLT_abc.jpg", "abc"0)  // false
	var hasAbcSuf bool = strings.HasSuffix("NLT_abc.jpg", "abc")  // false
	fmt.Println(hasAbcSuf)
}