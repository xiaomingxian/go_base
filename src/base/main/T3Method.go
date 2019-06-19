package main

import "fmt"

func main() {
	//方法-类似与java中的对象的方法
	p := Person{}
	p.show()
}
func (p Person) show() {
	fmt.Println("=====>方法调用")
}

type Person struct {
	Name string
}
