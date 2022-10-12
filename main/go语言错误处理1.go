package main

import "fmt"

type Num struct {
	dee int //被除数
	der int //除数
}

func go_error1() {
	rest1, error1 := D(1, 0)
	fmt.Println(rest1, error1)

	rest2, error2 := D(1, 1)
	fmt.Println(rest2, error2)
}

//指定函数的所属对象 继承error接口
func (num *Num) Error() string {
	// fmt.Println("数据信息:", num)

	return "除数不能是0"

}

//返回多个返回值时 数据类型使用逗号进行分割
func D(a, b int) (int, string) {
	if b == 0 {
		//将数据放到结构体中
		num := Num{a, b}

		return 0, num.Error()
	} else {
		return a / b, ""
	}
}
