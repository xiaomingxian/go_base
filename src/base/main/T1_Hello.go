package main

import "fmt"

func main() {
	fmt.Println("hello go")
	//	基本数据类型var i int
	var name string
	var x int
	var b bool
	var bb byte
	c := 1//自动推导类型
	fmt.Print("字符串："+name, ",int:", x, ",bool:", b, ",byte:", bb, c)
}
