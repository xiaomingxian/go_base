package main

import (
	"fmt"
)

func main() {
	//函数作为参数传递
	h1(1, callback)
	//匿名函数作为参数-test
	h1(2, func(i int) int {
		fmt.Println("---------->匿名函数作为参数测试")
		return i + 10
	})
	//方法
	fangFa()

}

//-----------------------------------------------函数相关测试
//声明一个函数类型--类型名为cb 参数为int 返回值为int
type cb func(int) int

//接受函数参数的函数
func h1(x int, f cb) {
	f(x)
}

//接受函数参数的函数

//定义符合自定义函数类型的函数-作为回调函数
func callback(x int) int {
	fmt.Println("回调函数：", x)
	return x + 5
}

//-----------------------------------------------方法相关测试
func fangFa() {
	//方法-类似与java中的对象的方法
	p := Person{} //实体实例化
	p.show()      //实例化对象调用方法
}

//Person对象的方法
func (p Person) show() {
	fmt.Println("=====>方法调用")
}

//实体
type Person struct {
	Name string
}
