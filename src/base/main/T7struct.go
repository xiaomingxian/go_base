package main

import "fmt"

//结构体[java-对象]
func main() {
	base_test()

	//作为方法参数-是值传递
	var book Book
	book.name = "java context"
	book.context = "java"
	book.price = 20
	stf(book)
	fmt.Println(book.context)
	fmt.Println(&book)
	//想改变函数的值-得传入指针
	stfc(&book)
	fmt.Println(book.context)

}

//结构体[对象]作为参数是值传递
func stf(book Book) {
	name := book.name
	context := book.context
	book.name = context
	book.context = name
}

//改变了指针所指向的内容
func stfc(book *Book) {
	//简单示例属性交换
	name := book.name
	context := book.context
	book.name = context
	book.context = name
}

func base_test() {
	//形式1--相当于java构造参数
	book := Book{"深入jvm", 12.2, "jvm-context"}
	fmt.Println(book)
	//形式二--创建对象-字段赋值
	var bok2 Book
	bok2.name = "go base"
	bok2.price = 10.5
	bok2.context = "go coontext"
	fmt.Println(bok2)

}

type Book struct {
	name    string
	price   float64
	context string
}
