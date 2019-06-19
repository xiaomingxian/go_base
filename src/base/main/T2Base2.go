package main

import "fmt"

//函数
func main() {
	res := fun1(1, 2)
	fmt.Println("--------->普通函数:", res)
	//	多个返回值
	a, b := "Google", "FaceBook"
	a, b = resMore(a, b)
	fmt.Println("--------->多个返回值:", a, b)
	//	引用传递[一般为值传递]
	/**
	*&a 指向 a 指针，a 变量的地址
	*&b 指向 b 指针，b 变量的地址
	 */
	n1, n2 := 1, 2
	yinYong(&n1, &n2)
	fmt.Println("--------->引用传递：", n1, n2)

	//闭包[匿名函数]
	ming := niMing(5, 1)
	m1, m2 := ming()
	fmt.Println("--------->匿名函数：", m1, m2)
	//参数类型为数组
	shuZCan()
}

//参数是数组
func shuZCan() {

}
func niMing(x, y int) func() (int, int) {
	return func() (i int, i2 int) {
		return y, x
	}
}

//	引用传递[一般为值传递]
func yinYong(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

//	多个返回值
func resMore(a, b string) (string, string) {
	return b, a
}

//参数列表-返回类型
func fun1(num1, num2 int) int {
	return num1 + num2
}
