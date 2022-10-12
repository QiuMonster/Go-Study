package main

import "fmt"

//定义一个结构体
type Circle struct {
	radius float64 //半径
}

//定义结构体圆的专属方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

//方法的使用  一个包含接收者的函数
func go_func3() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积是:", c1.getArea())
}
