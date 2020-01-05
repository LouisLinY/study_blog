package main

import "fmt"

type Humaner interface {
    //方法，只声明，没有实现，由具体的类型来实现
    say()
}

type Student struct {
	name string
	age int
}

func (s *Student) say() {
	fmt.Println("my name is ", s.name)
}

type Teacher struct {
	addr string
	group string
}

func (t Teacher)say() {
	fmt.Println("teacher addr ", t.addr)
}

func sayInfo(i Humaner) {
	i.say()
}



func main(){
	//定义接口类型的变量
	var i Humaner
	s := &Student{"YOYO", 23}
	i = s
	i.say()


}
