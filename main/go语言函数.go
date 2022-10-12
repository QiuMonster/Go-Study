package main

import (
	"fmt"
)

func go_func() {
	result := max(33, 42)
	fmt.Println("结果是:", result)
	a, b := swap("world", "hello")
	fmt.Println("函数返回的结果是:", a, b)

	a1 := 10
	b1 := 20
	swap1(&a1, &b1)
	fmt.Println("函数处理后的a1和b1的值分别为:", a1, b1)

	

}

//含参函数  函数返回值类型写在后面 返回一个值
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//函数返回多个值
func swap(a, b string) (string, string) {
	// var p *string //定义一个字符指针
	// p = &a //指向a所在的地址
	// a=b
	// b=*p  //将p所指向地址空间的值赋给b
	return b, a

}

//函数使用引用传递参数会影响到实际的参数
func swap1(a, b *int) {
	//实现a和b中所指向值的交换
	var temp int
	temp = *a //保存a地址中的值
	*a = *b   //将b所指向的值赋给a所指向的值  并不是指针的指向同一个地址空间
	*b = temp
}
