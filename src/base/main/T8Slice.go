package main

import (
	"fmt"
)

//内置类型切片[动态数组]
func main() {
	slice_base()
}

func slice_base() {
	//声明
	var sl []int
	fmt.Println(sl)

	//make创建
	//make([]type, len,capacity)   //len是数组的长度也是切片的初始长度,capacity指定容量？可以不写与-默认len一致
	ints := make([]int, 10)
	var sl2 []int = make([]int, 5, 10)

	fmt.Println(ints, "----", sl2)
	//切片截取
	i := []int{1, 2, 3, 4, 5, 6}
	i2 := i[:]
	i3 := i[2:5]
	i4 := i[2:]
	i5 := i[:4]
	fmt.Println(i, "-->默认：", i2, "-->头尾：", i3, "-->头：", i4, "-->尾：", i5)
	//计算切片的长度和容量
	fmt.Println("切片长度：", len(i3), "-容量：", cap(i3), "-切片：", i3)
	//初始切片为nil长度为0
	var sl_test []int
	fmt.Println(sl_test == nil, sl_test, len(sl_test), cap(sl_test))
	//append() and copy()
	i6 := append(sl_test, 0, 1)
	fmt.Println(i6, "-->原切片sl_test：", sl_test)
	//copy(dest,source)   新数组的容量得指定
	copy(sl_test, i)
	fmt.Println(sl_test)
	//var newSlice = make([]int, len(i), cap(i)*2)
	var newSlice = make([]int, len(i), cap(i))
	//var newSlice = make([]int, len(i), 5) //panic: runtime error: makeslice: cap out of range ##### when cap<len
	//var newSlice = make([]int, 5, 5) //当长度较小时数据拷贝不全
	copy(newSlice, i)
	fmt.Println(newSlice, "——长度：", len(newSlice), "——容量：", cap(newSlice))
}
