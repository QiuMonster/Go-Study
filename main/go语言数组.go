package main

import "fmt"

func go_arr() {
	one_arr()
	two_arr()
	add_onearr()
}

//一维数组
func one_arr() {
	//声明一个数组
	// var strarr[5]string
	// strarr=[5]string{"","","",""}
	// fmt.Println(strarr)

	//声明并初始化一个数组
	var intarr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(intarr)
	//通过字面量直接初始化数组
	intarr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intarr1)
	//当数组长度不确定时 写成 ...
	intarr2 := [...]int{1, 2, 3, 4}
	fmt.Println(intarr2)
	//指定索引下标元素初始化
	intarr3 := [...]int{1: 123, 3: 3, 6: 90} //初始化 下标 1 3 6 其余为0  index:value
	intarr3[2] = 100                         //给数组进行赋值
	fmt.Println(intarr3)
}

//二维数组
func two_arr() {
	//定义一个二维数组
	intarr := [3][3]int{{1, 2, 3}}
	fmt.Println(intarr)
	//遍历二维数组  使用 range遍历
	for i, value := range intarr {
		fmt.Println(i, value)
		for j, val := range value {
			fmt.Println(i, j, val)
		}
	}
	//原始遍历
	for i := 0; i < len(intarr); i++ {
		for j := 0; j < len(intarr[i]); j++ {
			fmt.Println(i, j, intarr[i][j])
		}

	}
}

//向二维数组中添加一维数组
func add_onearr() {
	//定义空的二维数组  不指定大小
	var intarr [][]string
	//使用的一维数组也不需要定义大小
	a := []string{"a", "ab", "abc"}
	intarr = append(intarr, a)
	b := []string{"b", "bb", "bbc"}
	intarr = append(intarr, b)
	fmt.Println(intarr)

	//使用一个参数时  返回的是数组中的下标  使用占位符_
	for _,val := range intarr {
		fmt.Println(val)
		for _,va := range val {
			fmt.Println(va)
		}
	}
}
