package main

import (
	"fmt"
	"os"
	"io"
)

func CopyFile(dst, src string) (w int64, err error) {
	src, err := os.Open(src)
	if err != nil {
		return
	}
	//defer src.Close() 使用defer改写后，在打开资源无报错后直接调用defer关闭资源，一旦养成这样的编程习惯，
	//则很难会忘记资源的释放
	dst, err := os.Open(dst)
	if err != nil {
		return
	}
	//defer dst.Close() defer语句的位置不当，有可能导致panic，一般defer语句放在检查语句之后。defer会推迟资源的释放，
	//defer尽量不要放在循环语句里，将大函数内的defer语句单独拆分成一个小函数是一种很好的实践。
	//defer相对于普通的函数调用需要间接的数据结构的支持，相对于普通函数调用有一定的性能损耗。
	//defer中最好不要对有名返回值参数进行操作，否则会引发匪夷所思的结果
	w, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}



