package main
import (
	"errors"
	"fmt"
)


/*
Go 语言引入了一个关于错误处理的标准模式，即error接口，它是Go语言内建的接口类型。
该接口定义如下：
type error interface {
    Error() string
}

Go语言的标准库代码包errors为用户提供如下方法：
package errors
type errorString struct {
    text string
}

func New(text string) error{
    return &errorString{text}
}

func (e *errorstring) Error() string {
    return e.text
}

另外一个可以生成error类型值得方法是调用fmt包中的Errorf函数
package fmt
import "errors"

func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args))
}

例子：
var tmp error   //定义接口变量，给接口变量赋值需要对象实现该接口下的方法的具体实现
tmp = &errorString{text}

 */

func Div(a, b int) (result int, err error) {
	if b == 0 {
		return 0, errors.New("b buneng wei ling")
	}else{
		return a/b, nil
	}
}

func testa(){
	fmt.Println("first aaaaaaaaa")
}

func testb(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(recover())
		}
	}()
	var a [10] int
	a[x] = 10
}

func testc(){
	fmt.Println("three cccccccccc")
}




func main(){
	err1 := fmt.Errorf("%s", "this is normal err")
	fmt.Printf("err1 type = %T\n", err1)
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Printf("err2 type = %T \n", err2)
	fmt.Println("err2 = ", err2)

	result , err := Div(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	}else{
		fmt.Println("result = ", result)
	}

	testa()
	testb(3)
    testc()

}