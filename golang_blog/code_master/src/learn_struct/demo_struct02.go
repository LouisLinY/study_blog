package main

import (
	"fmt"
)
//同名字段

type Person struct {
	name string    //名字
	sex byte       //性别，字符类型
	age int        //年龄
}

type Student struct {
	Person   //匿名字段，那么默认Student就包含了Person的所有字段
	id int
	addr string
	name string   //和Person同名
}

func main(){
	var s1 Student
	s1.name = "louis"    //默认操作的是Student中的name（就近原则），如果能在作用域找到此成员，就操作此成员，如果没有找到，就会找继承的字段
	s1.sex = 'm'
	s1.age = 18
	s1.addr = "sz"
	s1.Person.name = "lin"
	fmt.Printf("%+v\n", s1)
}

