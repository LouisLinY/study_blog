package main

import "fmt"
/*
函数类型和map、slice、chan一样，实际函数类型变量和函数名都可以当作指针变量，
该指针指向函数代码的开始位置。通常说函数类型变量是一种引用类型，未初始化的函数
类型的变量的默认值是nil。
有名函数的函数名可以看作函数类型的常量，可以直接使用函数名调用函数，也可以直接
赋值给函数类型的变量，后续通过该变量来调用该函数。
*/
func add(a, b int) int{
	return a+b
}

func sub(x, y int) int{
	return x-y
}

type op func(int, int) int   //定义一个函数类型，输入的是两个int类型，返回值是一个int类型

func do(f op, a, b int) int {  //定义一个函数，第一个参数是函数类型op，
	return f(a, b)             //函数类型变量可以直接用来进行函数调用
}

//匿名函数直接赋值给函数变量，作为函数字面量
var sum = func(a, b int) int{
	return a+b
}

//匿名函数作为函数形参
func doinput(f func(int, int) int, a, b int) int{
	return f(a, b)
}

//匿名函数作为返回值
func wrap(op string) func(int, int) int{
	switch op {
		case "add":
			return func(a, b int) int {
				return a+b
				}
		case "sub":
			return func(a, b int ) int{
				return a-b
	               }
	     default:
	     	return nil
	}
}


func main(){
	a := do(sub, 5,1)  //函数名add可以当作相同函数类型形参，不需要强制类型转换
	fmt.Println(a)
	s := do(add, 3, 5)
	fmt.Println(s)

	//直接调用函数
	fmt.Println(add(5, 9))
	f := sub //有名函数可以直接赋值给函数变量
	fmt.Println(f(10, 4))
    //打印函数类型
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)

	//匿名函数
	//Go提供两种函数：有名函数和匿名函数。匿名函数可以作为函数字面量，所有直接使用函数类型变量的
	//地方都可以由匿名好事代替。匿名函数可以直接赋值给函数变量，可以当作实参，也可以作为返回值，
	//还可以直接被调用
    defer func() {
    	if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	sum(1, 2)

	doinput(func(x, y int) int {
		return x+y
	}, 1, 2)

	//opFunc := wrap("aa")
	opFunc := wrap("add")
	re := opFunc(2, 3)
	fmt.Printf("%d\n", re)

   //defer，注册延时调用，这些调用以先进后出（FIFO）的顺序在函数返回前被执行。这有点类似java中的异常处理中的finaly子句
   //defer常用于保证一些资源最终一定能够得到回收和释放
   //defer后面的匿名函数或方法的调用，不能是语句，否则会报expression in defer must be function call
   //defer函数的实参在注册时通过值拷贝传递进去，实参a的值在defer注册时通过值拷贝传递进去，
   //后续语句a++并不会影响defer语句最后的输出结果
   //defer必须先注册后才能执行，如果defer位于return之后，则defer因为没有注册，不会被执行。
   //主动调用os.Exit(int)退出进程时，defer将不再被执行（即使defer已经提前注册）
   //defer的好处是可以在一定程度上避免资源泄露，特别是在有很多return语句，有多个资源需要关闭的场景中，
   //很容易漏掉资源的关闭操作。
    defer func() {
    	fmt.Println("first")
	}()

	defer func() {
		fmt.Println("second")
	}()
	fmt.Println("function body")
	//func f() int{
	//a := 0
	//defer func(i int) {
	//	fmt.Println("defer i=", i)
	//}(a)
	//
	//a++
	//return a
	//}()

}