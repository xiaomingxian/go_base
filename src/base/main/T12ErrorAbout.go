package main

//error顶层接口
//type error interface {
//	Error() string
//}

import (
	"errors"
	"fmt"
)

func main() {
	error_base()
}

func error_base() {
	result, error_my := easyYestOferror(1, 0)
	if error_my != nil {
		fmt.Println(error_my)
	} else {
		fmt.Println(result)
	}

}

func easyYestOferror(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("被除数不能为0")
	} else {
		return x / y, nil
	}

}
