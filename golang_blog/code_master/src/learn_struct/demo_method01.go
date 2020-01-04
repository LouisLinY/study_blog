package main

import (
	"fmt"
)

/*
这种带有接受者的函数，称为方法（method）
一个方法则是一个和特殊类型关联的函数。
一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象
而是借助方法来做这些事情。
在go语言中，可以给任意自定义类型（包括内置类型、但不包括指针类型）添加相应的方法
方法总是绑定对象实例，并隐式将实例作为第一实参（receiver），
只要接收者类型不一样，就算方法名同名，也不会出现同名方法
*/

func Add01(a, b int) int {
	return a+b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

//tmp叫接收者，接收者就是传递的一个参数
func (tmp long) Add02(other long) long{
	return tmp + other
}




func main(){
	var result int
	result = Add01(1, 2)
	fmt.Println("result = ", result)

	var a long = 2
	r := a.Add02(3)
	fmt.Println("result = ", r)
}


