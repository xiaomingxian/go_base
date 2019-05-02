package main

import (
	"fmt"
)

//有返回值的函数，测试匿名变量
func haveR() (a, b, c int) {
	return 5, 8, 10
}

func main() {
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
