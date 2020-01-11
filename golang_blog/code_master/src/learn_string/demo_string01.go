package main

import (
	"fmt"
	"strings"
)


func main(){
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
