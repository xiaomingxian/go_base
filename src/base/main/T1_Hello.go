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
	caseYu() //select 未写

}
func caseYu() {
	//if else  if嵌套
	//没有三目运算
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
		xa = 1 << iota //左移0位-->1
		j  = 3 << iota //左移1--->6
		k              //6 左移 2位-->12
		i              //12 左移 3位-->24
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
