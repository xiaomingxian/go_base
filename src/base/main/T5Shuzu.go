package main

import "fmt"

func main() {
	//数组基本
	shuzu_base()

	//参数类型为数组
	ints := [2]int{1, 2}
	ints2 := []int{1, 2, 3, 4}
	//指定大小
	shuZCan(ints)
	//不限制大小
	shuZCan2(ints2)

}

//参数是数组
func shuZCan(s1 [2]int) {
	fmt.Println(s1)
}

//参数是数组-无大小限制
func shuZCan2(s1 []int) {
	fmt.Println(s1)

}

func shuzu_base() {
	//初始值0
	var sz [10]int
	fmt.Println(sz)
	var sz2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(sz2)

	//	多维数组
	var msz [2][3][4]int
	fmt.Println(msz)
	var msz2 [2][3]int
	msz2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(msz2)
}
