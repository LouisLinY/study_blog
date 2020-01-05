package main

import (
	"bytes"
	"fmt"
	"io"
)


/*
接口是抽象的，是一组接口方法，并不知道它内部实现，
不用和具体的实现绑定在一起

fmt.Fprint函数的第一个参数是io.Writer这个接口，所以只要实现了这个接口的具体类型都可以作为参数
传递给fmt.Fprint函数，而bytes.Buffer恰恰实现了io.Writer接口，所以可以作为参数传递给fmt.Fprint函数。

接口是定义行为的类型，是抽象的，这些定义的行为不是由接口直接实现的，而是同方法由用户定义的类型实现。
如果用户定义的类型，实现了接口类型声明的所有方法，那么这个用户定义的类型就实现了这个借口，所以这个用户
定义类型的值可以赋值给接口类型的值。

我们看下接口的值被赋值后，接口值内部的布局。接口的值是一个两个字长度的数据结构，第一个字包含一个指向内部表结构的指针，
这个内部表里存储的有实体类型的信息以及相关联的方法集；第二个字包含的是一个指向存储的实体类型值的指针。所以接口的值结构其实是两个指针，
这也可以说明接口其实一个引用类型。
*/

type cat string
type dog string

func (c cat) printInfo() {
	fmt.Println("cat method :")
}

func (d dog) printInfo() {
	fmt.Println("dog method :")
}

type animal interface {
	printInfo()
}



func main(){
	var b bytes.Buffer
	fmt.Fprint(&b, "hello world")
	fmt.Println(b.String())

	var w io.Writer
	w = &b
	fmt.Println(w)

	var a animal
	var c cat
	a = c
	a.printInfo()

	var d dog
	a = d
	a.printInfo()
}


