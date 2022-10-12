package main

import "fmt"

type Stu struct {
	age  int
	name string
}
//使用指针的形参与结构体的形参基本没有什么区别 
func go_obj_point() {
	s := Stu{20, "qiu"}
	var point *Stu //定义一个结构体指针
	point=&s
	p:=&s //直接初始化
	go_p(point)
	go_p(p)
	go_obj(s)
	
}

//函数形参是一个结构体指针
func go_p(s *Stu) {
	fmt.Println("使用指针参数")
	fmt.Println(s)
	fmt.Println(*s)
	fmt.Println(s.age)
}

//函数的额形参是一个结构体对象
func go_obj(s Stu) {
	fmt.Println(s)
	fmt.Println(s.name)
}
