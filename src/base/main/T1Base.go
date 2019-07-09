package main

import (
	"fmt"
)

func main() {
	//base()
	//常量
	//changliang()
	//运算符
	//yunSuan()
	//条件语句
	//caseYu()
	//select与并发
	//seAndBf()
	//循环
	//xunHuan()

}

func xunHuan() {
	//method1
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}
	//	method2
	fmt.Println("------------------------method2----------------")
	c := 1
	d := 10
	//	跳过本次循环并回到循环的开始语句 LOOP 处
LOOP:
	for c < d {
		fmt.Println(c)
		c++
		if c == 8 {
			fmt.Printf("break跳出for循环")
			break
		}
		if c == 2 {
			fmt.Printf("continue----\n")
			continue
		}
		if c == 6 {
			fmt.Printf("goto----\n")
			goto LOOP
		}
	}
	//	method3

	nums := [5]int{1, 2, 3, 4, 5}
	for i, j := range nums {
		fmt.Printf("索引%d,元素%d\n", i, j)
	}
}

func seAndBf() {
	//	send、receive操作是针对channel的
	//	send ： ch <- a
	//receive: a := <- ch
	//go开启并发执行
	//就是在做goroutine之间的内存共享---宏观上来讲，信道有点像其他语言中的队列（queue），遵循先进先出的规则。
	messages := make(chan string)
	go func(message string) {
		messages <- message // 存消息
	}("Ping!")

	fmt.Println(<-messages) // 取消息
	//--------------------------------------------------------------
	//select
	//得初始化
	var c1, c2, c3 chan int = make(chan int), make(chan int), make(chan int)
	var i1, i2 int
	go func(value int) {
		c1 <- value
	}(10)
	//fmt.Println(<-c1)
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := <-c3:
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
func caseYu() {
	//if else  if嵌套
	//没有三目运算
	//switch
	var a = 1
	switch a {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(a)
	case 3:
		fmt.Println(a)
	}

}
func yunSuan() {
	/**
	加减乘除 比较 位运算
	*/

	fmt.Println("+：", 1+2)
	fmt.Println("-：", 1-2)
	fmt.Println("x：", 1*2)
	fmt.Println("/：", 1/2)
	var a = 1
	a++
	fmt.Println("++:", a)
	a--
	fmt.Println("--:", a)

}

func changliang() {
	//隐士定义
	const name = "abc"
	//显示定义
	const name2 string = "abc"
	//多个类型
	const a, b, c = "A", "B", "C"

	fmt.Printf("name %s", name)
	//枚举
	const (
		m1 = "M1"
		m2 = "M2"
		m3 = "M3"
	)
	fmt.Println("\n枚举", m1, m2, m3)

	//iota---//默认0，每使用一次+1
	//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	// const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
	const (
		ax = iota
		ac = iota
		ad = iota
	)

	fmt.Println("ioat: ", ax, ac, ad)
	fmt.Println("------------- ioat 计数实验  -----------")
	//可以用作枚举内计数
	const (
		ca = iota
		cb
		cc = 100 //设置值后面的枚举也都是这个值
		cd
		ce
		cf = iota //恢复计数---还是原来的
		cg
	)
	fmt.Println(ca, cb, cc, cd, ce, cf, cg)

	fmt.Println("------------- 二进制测试 -------------")
	const (
		xa = 1 << iota //左移0位-->1====>1*2^0
		j  = 3 << iota //左移1--->6===>3*2^1
		k              //左移2位-->3<<2==>3*2^2
		i              //左移3位-->3<<3==>3*2^3
	)
	fmt.Println(xa, j, k, i)
}

//有返回值的函数，测试匿名变量
func haveR() (a, b, c int) {
	return 5, 8, 10
}
func base() {
	fmt.Println("hello go")
	//	基本数据类型var i int
	var name string
	var x int
	var b bool
	var bb byte
	c := 1 //自动推导类型
	b2 := true
	// 结构体
	var e struct {
		x int
	}
	//返回值为bool
	var d func() bool
	fmt.Println("字符串："+name, ",int:", x, ",bool:", b, ",byte:", bb, c, b2, e)
	//32位浮点切片
	var cf []float32
	//多重赋值
	a2, b3, c2 := 1, 12, 3
	fmt.Println(d, cf, a2, b3, c2)
	//	数值交换
	var (
		x1 = 10
		x2 = 20
	)
	x1, x2 = x2, x1
	//printf format
	fmt.Printf("变量1 = %d ,变量2 = %d\n", x1, x2)
	//	 匿名变量测试 _还是原来的值未做处理
	var aa, bbb, cc = 1, 2, 3
	_, bbb, cc = haveR()
	fmt.Println(aa, bbb, cc)

}
