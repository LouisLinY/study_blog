package main

import (
	"fmt"
)


type Person struct {
	name string
	sex byte
	age int
}

type Student struct {
	Person   //匿名字段，那么默认Student就包含了Person的所有字段
	id int
	addr string
}

func main(){
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 25}, 1, "fujian"}
	fmt.Printf("%+v\n", s1)  //%+v 显示更详细

	//自动推导类型
	s2 := Student{Person{"louis", 'f', 28}, 2, "shenzheng"}
	fmt.Println("s2 = ", s2)

	//指定成员初始化，没有初始化的常用自动赋值为0
	s3 := Student{Person:Person{name:"sijie"}, id:3}
	fmt.Println("s3 = ", s3)

	//指定所有成员初始化
	s4 := Student{Person:Person{name:"mayun", sex:'m', age: 56}, id:4, addr:"hangzhou"}
	fmt.Println("s4 = ", s4)

	//成员操作
	s4.Person = Person{"mahuateng", 'm', 58}
	fmt.Println(s4.name, s4.sex, s4.age, s4.id, s4.addr)




}