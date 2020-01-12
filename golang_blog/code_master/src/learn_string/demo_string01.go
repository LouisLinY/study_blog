package main

import (
	"fmt"
	"strings"
)


func main(){
	//判断字符串是否以prefix开头
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "http"))
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "https"))

	//判断字符串是否以suffix结尾
	fmt.Println(strings.HasSuffix("http://www.baidu.com", "com"))
	fmt.Println(strings.HasSuffix("http://www.baidu.com", "cn"))



	fmt.Println(strings.Contains("hello world", "abc"))
	fmt.Println(strings.Contains("hello world", "worl"))
	fmt.Println(strings.Contains("hello 中国", "中"))
	fmt.Println(strings.Contains("hello 中国", "华"))

	fmt.Println(strings.Index("hello world", "abc"))
	fmt.Println(strings.Index("hello world", "worl"))
	fmt.Println(strings.Index("hello 中国", "国"))

	fmt.Println(strings.Count("abcddddabceeeeabc", "abc"))
	fmt.Println(strings.Count("abcddddabceeeeabc", "fhj"))




}
