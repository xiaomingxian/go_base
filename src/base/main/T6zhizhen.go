package main

import "fmt"

func main() {
	//	指针
	zhizhen()

	//指针作为参数[引用传递]
	a := 1
	b := 2
	yinyong2(&a, &b)
	fmt.Println(a, b)

}

//改变指针指向的变量
func yinyong2(a, b *int) {
	fmt.Println("----->", a, b)
	var temp int
	temp = *a
	*a = *b
	*b = temp

}

func zhizhen() {
	//前加& 取出地址值
	var a = 1
	fmt.Println("--->a的地址值：", &a)
	//声明指针  var 指针名称 *指针类型  #*表示是指针
	var z1 *int
	z1 = &a
	fmt.Println("--->z1的地址值：", z1)
	fmt.Println("--->z1指向的变量：", *z1)
	//	----------空指针
	var z2 *int
	fmt.Println(z2)
	fmt.Println("------------------------指针数组-----------------------------")
	zhizhenSz()
	fmt.Println("-----------------------指向指针的指针-------------------------")
	zzhizhen()
}

//指针数组
func zhizhenSz() {
	var sz [5]*int
	fmt.Println(sz)
	for i := 0; i < 5; i++ {
		var x = i
		sz[i] = &x
	}
	fmt.Println(sz)
	for i := 0; i < 5; i++ {
		fmt.Println("--->地址值对应的数值：", *sz[i])
	}

}

//指向指针的指针
func zzhizhen() {
	a := 1
	var pptr **int
	var ptr *int
	ptr = &a
	pptr = &ptr
	fmt.Println(pptr)

}
