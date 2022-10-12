package main

import "fmt"

type BlackDog struct {
}
type WhiteDog struct {
}

//定义一个Dog接口
type Dog interface {
	//无实现方法
	eat(age int)string //可以使用形参和 返回值
}

//BlackDog继承接口 并重写方法
func (blackDog BlackDog) eat(age int)string {
	fmt.Println(a,"岁的小黑狗正在吃东西")
	return ""
}
func (whiteDog WhiteDog) eat(age int) string{
	fmt.Println(a,"岁的小白狗正在吃东西")
	return ""
}
func go_interface() {
	var dog Dog
	//使用接口对象创建时一定要初始化才行
	dog = BlackDog{}
	dog.eat(4)
	//可以使用new关键词创建对象
	dog = new(BlackDog)
	dog.eat(5)

	//并没有初始化
	var bddog BlackDog
	bddog.eat(6)
}
