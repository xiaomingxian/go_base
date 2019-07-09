package main

import (
	"fmt"
)

func main() {
	//匿名形式
	//bingfaOfNM()
	//普通形式
	//bingfaNormal()
	//select
	selectdemo()

}

func selectdemo() {
	/**
	每个 case 都必须是一个通信
	所有 channel 表达式都会被求值
	所有被发送的表达式都会被求值
	如果任意某个通信可以进行，它就执行，其他被忽略。
	如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
	否则：
		如果有 default 子句，则执行该语句。
		如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
	*/
	//创建信道
	channel, channel2 := make(chan string), make(chan string)
	//go producer(channel, "通道1的信息")
	//fmt.Println(<-channel)

	select {
	//接受消息
	case data, ok := <-channel:
		fmt.Println(data, ok)
	//发送---不知道咋测试 没法写go关键字-无法建立信道
	case channel2 <- "通道2发送消息":
		println("----->通道2发送消息")
	//默认-如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
	default:
		fmt.Println("----->无结果")
	}

	//可以注释掉default 用以下代码测试阻塞
	//time.Sleep(10)
	//fmt.Println(<-channel)

}

//类似与switch-只不过元素是信道
func bingfaOfNM() {

	//1.创建信道---指定类型/自动类型推导
	//var messages chan string=make(chan string)
	channel := make(chan string)

	//2.存消息 匿名函数-----用go修饰表示一个携程[goroutine]
	go func(message string) {
		channel <- message // 存消息
	}("Ping!")

	//3.接收消息
	data := <-channel
	fmt.Print(data)
	//或者
	//fmt.Print(<-messages)
}

func bingfaNormal() {
	//创建信道
	channel := make(chan string)
	//	发送消息
	go producer(channel, "哈哈哈")
	//	接受消息
	data := consumer(channel)
	fmt.Print("接收到的消息：", data)
}

//channel[变量] chan[信道类型] string[信道内元素类型]
func producer(channel chan string, message string) {
	channel <- message
}
func consumer(channel chan string) string {
	data := <-channel
	return data
}
