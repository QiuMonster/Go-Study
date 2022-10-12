package main

import "fmt"

type Student struct {
	age     int
	name    string
	height  float32
	subject string
}

func go_struct() {
	//基本类型快速初始化不需要写类型  会自动匹配
	//使用 {} 进行结构体的赋值
	s := Student{1, "qiu", 163.5, "清华大学"}
	fmt.Println("重新赋值前的年龄:",s.age)
	//直接赋值
	s.age=20
	fmt.Println("赋值之后的年龄:",s.age)
	fmt.Println(s.name)
	fmt.Println(s)

	//定义一个结构体类型的指针
	var sp *Student=&s
	//通过结构体指针访问成员属性
	fmt.Println(sp.age)

}
