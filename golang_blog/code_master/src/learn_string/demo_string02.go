package main


import (
	"fmt"
	"strconv"
)

/*
字符串转换
Append系列函数讲整数等转换为字符串，添加到现有的字节数组中

Format系列函数吧其他类型的转换为字符串

*/


func main(){
	//转换为字符串后追加到字节数组
	var str []byte
	str = strconv.AppendBool(str, true)

	str = strconv.AppendInt(str, 1234, 10)

	str = strconv.AppendQuote(str, "abcdefghijklmnopqrstuvwxyz")
	fmt.Println("str = ", string(str))

	//其他类型转换为字符串
	var str1 string
	str1 = strconv.FormatBool(true)
	fmt.Println("str1 = ", str1)
	str1 = strconv.FormatFloat(3.14, 'f', -1, 64)

	fmt.Println("str1 = ", str1)

	str1 = strconv.Itoa(6666)
	fmt.Println("str1 = ", str1)

	//字符串转换其他类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	}else{
		fmt.Println(err)
	}

	//字符串转换为整型

	a, err := strconv.Atoi("1314")
	fmt.Println("a = ", a)


}

