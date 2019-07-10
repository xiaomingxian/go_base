package main

import "fmt"

func main() {
	range_base()
}

func range_base() {

	array := []int{1, 2, 3, 4, 5}
	sum := 0
	//_表示空白符，此处用不到index
	for _, num := range array {
		sum += num
	}
	fmt.Println("--->sum:", sum)
	for i, _ := range array {
		fmt.Println("-->索引：", i)
	}
	//在map上的应用  ###map[key_type]value_type
	map_ := map[string]int{"k1": 1, "k2": 2}
	for k, v := range map_ {
		fmt.Println("key:", k, "-value:", v)
	}

}
