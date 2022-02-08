package main

import "fmt"

//type i interface {
//
//}
//
//等同于
//var i interface{}

func main() {
	//空接口
	var i interface{} //空接口
	i = "2"
	if _, ok := i.(int); ok {
		println("这是一个int类型")
	}

	fmt.Println(i)
}
