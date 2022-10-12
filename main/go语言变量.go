package main

import "fmt"

//go语言变量
func go_var() {
	//变量声明和定义分开操作
	var a = "helle world"
	var flag = true
	fmt.Println(a)
	fmt.Println(flag)
	var b, c int = 1, 2
	fmt.Println(b, c)
	//定义后没有初始化时 值一般是0 "" false
	var i int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)
	//使用 := 进行变量定义与赋值  只能在函数内部存在
	tom := "tom cat"
	name, age, height := "jack", 22, 180 //实现不同类型的多个变量同一声明初始化
	fmt.Println(tom + "is a cute cat")
	fmt.Println(name, age, height)

	//使用全局变量 MaxSize
	fmt.Printf("输出了全局变量MaxSize:%d", MaxSize)
	//一个变量对另一个变量进行赋值  实际上是对变量所指的地址空间的复制  这个空间指向了同一个值
	var num int = 10
	var num_copy = num //对值的复制
	if num == num_copy {
		fmt.Println(&num, &num_copy) //使用 &输出变量的地址空间
	}
	num_copy = 11
	fmt.Println(num)

}