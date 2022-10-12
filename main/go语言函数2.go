package main

import (
	"fmt"
	"math"
)

func go_func2() {

	//将函数作为实参 有点js的味道 将函数返回值赋给一个变量 使用变量名函数
	//声明函数变量
	getSquareRoot := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println("getSquareRoot使用的函数返回值是:", getSquareRoot(5))
	
	//调用闭包函数
	nextNumber := getSquence()
	fmt.Println("闭包函数值为:", nextNumber())
	fmt.Println("闭包函数值为:", nextNumber())
	fmt.Println("闭包函数值为:", nextNumber())

}

//函数闭包  我们创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
func getSquence() func() int {
	i := 0
	//使用匿名的函数
	return func() int {
		i += 1
		return i
	}
}




