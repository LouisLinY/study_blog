package main

import (
	"fmt"
)

//结构体指针类型匿名字段

type Person struct {
	name string    //名字
	sex byte       //性别，字符类型
	age int        //年龄
}

type Student struct {
	*Person   //指针类型
	id int
	addr string
}

func main(){
	s1 := Student{&Person{"yoyo", 'm', 5}, 666, "shenzheng"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	//先定义变量
	var s2 Student
	s2.Person = new(Person)
	s2.name = "louis"
	s2.sex = 'm'
	s2.age = 29
	fmt.Printf("%+v\n", s2)
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)


}