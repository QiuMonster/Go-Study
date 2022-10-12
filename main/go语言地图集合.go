package main

import "fmt"

func go_map() {
	go_m1()
	go_m2()
}
func go_m1() {
	//原始方式创建
	var map1 map[int]string
	//必须使用make进行创建 可以不指定map的大小
	map1 = make(map[int]string)

	fmt.Println(len(map1))
	//向map中添加元素
	map1[0] = "hello 1"

	fmt.Println(map1)
}
func go_m2() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India "] = "新德里"

	//使用键值输出map中的值
	for key:=range map1{
		fmt.Println(key,"对应的首都为:",map1[key])
	}

	//检查元素在集合中是否存在
	capital,ok:=map1["France"] //存在两个返回值 一个是map对应的值 一个是存在标志
	if ok{
		fmt.Println("France的首都是:",capital)
	}else{
		fmt.Println("无对应的值",capital) //nil
	}

	//根据key进行map元素的删除
	delete(map1,"France")
	fmt.Println(map1)

}
