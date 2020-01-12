package main

import (
	"fmt"
	"regexp"
)

func main(){
	src := "abc azc a7c aac 888 a9c tac"
	//1、解释规则
	reg := regexp.MustCompile(`a[0-9]c`)
	//2、根据规则提取关键信息
	if reg == nil {
		fmt.Println("err = ", reg)
		return
	}
	result := reg.FindAllStringSubmatch(src, -1)
	fmt.Println("result = ", result)

	buf := "3.14 567 agsdg 1.23 7. 8.99 lsdljgl 6.666 7.8      "

	reg1 := regexp.MustCompile(`\d+.\d+`)
	if reg1 == nil {
		fmt.Println("MustCompile fail")
		return
	}

	//result1 := reg1.FindAllString(buf, -1)
	//result1 =  [3.1 4 5 1.2 3 7 8.9 6.6 6 7]
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)

}