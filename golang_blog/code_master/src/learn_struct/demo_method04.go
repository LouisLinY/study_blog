package main

import (
	"fmt"
)

type Person struct {
	name string
	sex byte
	age int
}

//同名方法，就近调用
func (p Person) PrintInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d \n", p.name, p.sex, p.age)
}

type Student struct{
	Person
	id int
	addr string
}

func (s Student) PrintInfo(){
	fmt.Printf("name = %s, sex = %c, age = %d, id = %d, addr = %s\n", s.name, s.sex, s.age, s.id, s.addr)
}

func main(){
	s := Student{Person{"louis", 'm', 32}, 1001, "shenzheng"}
	s.PrintInfo()
	//方法值
	f1 := s.PrintInfo
	f1()
	//方法表达式
	f2 := (Student).PrintInfo
	f2(s)
}
