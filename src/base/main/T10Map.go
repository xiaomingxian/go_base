package main

import "fmt"

func main() {
	//map_base()
	type_convert()
}

//类型转换
func type_convert() {
	a := 1
	b := 2
	c := float32(a) / float32(b)
	c_ := a / b
	fmt.Println(c, c_)

}

func map_base() {
	//var声明方式创建
	var map1 map[string]string
	//赋值
	fmt.Println(map1, map1 == nil)
	map1 = map[string]string{}
	map1["k1"] = "v1"
	map1["k2"] = "v2"
	map1["k3"] = "v3"
	map1["k3"] = "v3"
	fmt.Println(map1, map1 == nil)
	//make方式创建
	map2 := make(map[string]string)
	fmt.Println(map2, map2 == nil)

	//遍历
	bianli(map1)

	//查看元素是否存在
	data, ok := map1["kkk"]
	fmt.Println(data, ok)

	//delete()函数
	map_2 := map[string]string{"k1": "v1", "k2": "v2"}
	delete(map_2, "k1")
	fmt.Println(map_2)

}

func bianli(map_ map[string]string) {
	for i := range map_ {
		fmt.Println("-->key:", i, "--value:", map_[i])
	}

}
