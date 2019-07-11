package main

import "fmt"

func main() {
	interface_base()

}

func interface_base() {
	var phone Phone
	phone = new(IPhone)
	phone.call()

	phone = new(HuaWei)
	phone.call()

}

//接口
type Phone interface {
	call()
}

//iphone 对象
type IPhone struct {
}

//iphone对象的方法
func (iPhone IPhone) call() {
	fmt.Println("----->use iphone call someone")
}

//华为对象
type HuaWei struct {
}

//华为对象的方法
func (huaWei HuaWei) call() {
	fmt.Println("===>use HuaWei Super Call")
}
