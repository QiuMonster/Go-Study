package main

import "fmt"

func go_point() {
	go_point1()
	nil_point()
	point_arr()
	point_point()
}
func go_point1() {
	a := 10
	fmt.Println("a的地址是:", &a)
	//指针的定义
	var p *int     //int型指针
	var p1 *string //string型指针
	b := 20
	str := "hello world"
	p = &b    //p指向b的地址空间
	p1 = &str //p1指向 str的地址空间
	fmt.Println("p所指向地址空间的值:", *p, "p1所指向的地址空间的值是:", *p1)
	fmt.Println("p所指向的地址为:", p, "b所在的地址空间为:", &b, "p所在的地址空间为:", &p)
}

//go的空指针  值为 nil
func nil_point() {
	var point *int
	fmt.Println("point所指向的空间地址为:", point)
}

//指针数组使用  指针暂时不能后移
func point_arr() {
	a := [...]int{1, 2, 3, 4}
	var p *int
	p=&a[0] //p指针指向a数组的第一个空间地址
	fmt.Println("a的第一个下标的值",*p,"所在的空间地址",p)
	
	// fmt.Println(p)
}
//指向指针的指针
func point_point(){
	//定义一个指向指针的指针
	var pp **int
	var p *int
	a:=10
	pp=&p
	p=&a
	fmt.Println("pp指向的是p的地址",pp,"p的地址是:",&p,"p指向的是a的地址",p,"a的地址是:",&a)
	fmt.Println(*(&pp),pp,&p)
	fmt.Println(*(&p),p,&a)
}
