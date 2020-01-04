package main
import (
	"fmt"
)
//非结构体字段

type mystr string


type Person struct {
	name string    //名字
	sex byte       //性别，字符类型
	age int        //年龄
}

type Student struct {
	Person   //匿名字段，那么默认Student就包含了Person的所有字段
	int  //基础类型的匿名字段
	mystr
	id int
	addr string
	name string   //和Person同名
}


func main() {
	s := Student{Person{"louis", 'f', 5}, 666, "hello world", 23, "putian","yoyo"}
	fmt.Printf("s = %+v\n", s)

}