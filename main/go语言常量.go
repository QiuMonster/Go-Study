//go语言常量
package main

import "fmt"
import "unsafe"

//常量用作枚举
const (
	unknow = 0
	female = 1
	male   = 2
)

//常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
	//const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
	ta = iota //iota是特殊常量   在const出现后 值将变为0
	ta1
	ta2
	abc     = "abc"
	alen    = len(abc)
	csizeof = unsafe.Sizeof(abc)
)

const (
	a = iota    //0
	b           //1
	c = "hello" //hello
	d           //hello
	e = 100     //100
	f           //100
	g = iota    //6   要恢复计数才能够继续使用
	h           //7
	//期间iota一直在计数
)

func go_const() {
	const LENGTH int = 10 //显式定义 变量有数据类型
	const WIDTH = 5       //隐式定义
	var area int
	const ab, bc, cd = 1, false, "sotre"
	area = WIDTH * LENGTH
	fmt.Printf("面积是:%d", area)
	fmt.Println(ab, bc, cd)
	//使用常量的内置函数
	var tyu = "dsds"
	var io int16 = 1
	fmt.Println(unsafe.Sizeof(io))
	fmt.Println(unsafe.Sizeof(tyu))
	//使用iota特殊常量
	fmt.Println(ta, ta1, ta2) //iota 0 1 2   在遇到 const时自动加1

	fmt.Println(a, b, c, d, e, f, g, h)
	go_const_test1()

}

//iota配合移位运算<<
const (
	i = 1 << iota //iota=0 i=1
	j = 3 << iota //iota=1 j=6  011->110
	k // 3<<2  011->1100 12
	l // 3<<3  011->11000 24
)

func go_const_test1() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
