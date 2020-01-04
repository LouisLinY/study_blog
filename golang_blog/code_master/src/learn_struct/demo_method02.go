package main

import (
	"fmt"
)

type Person struct {
	name string
	sex byte
	age int
}

func (p Person) PrintInfo() {
	fmt.Println("name : ", p.name)
	fmt.Println("sex : ", p.sex)
	fmt.Println("age : ", p.age)
}

func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}


func main(){
	p := Person{"louis", 'm', 18}
	p.PrintInfo()
	p.SetInfo("yoyo", 'f', 28)
	p.PrintInfo()

	var p2 Person
	(&p2).SetInfo("golang", 'm', 6)
	p2.PrintInfo()
}



