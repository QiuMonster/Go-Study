package main

import (
	"fmt"
	"time"
)

//并发使用关键字 go进行线程的开启

func go_go() {
	// go say("hello")
	// say("world")
	// go_chan()
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//原始定义方式
	// var c chan int=make(chan int)
	ch := make(chan int)
	go chan_sum(s[1:4], ch) //值将会保存在 ch的通道中
	go chan_sum(s[3:], ch)
	s1, s2 := <-ch, <-ch
	//将值从通道中取出来
	fmt.Println(s1, s2)

	// go chan_1()
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//通道 chan
// ch <- v    // 把 v 发送到通道 ch
// v := <-ch  // 从 ch 接收数据
//            // 并把值赋给 v
func go_chan() {
	//声明一个通道 通道在使用前必须要先创建
	//默认情况下通道是不带缓冲区的
	ch := make(chan int)
	fmt.Println(ch)

	//创建一个带有1000字节的缓冲区通道
	ch1 := make(chan int, 1000)
	fmt.Println(ch1)
}

//将值放到chan中 后使用变量进行接收
func chan_sum(s []int, ch chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	ch <- sum //将sum的值赋给通道 ch
}
//可以通过通道中的返回值 ok来判断通道是否闲置  可以使用 close函数将通道关闭
func chan_1() {
	fmt.Println("hfiudasf")
	ch := make(chan int)
	ch <- 1
	// fmt.Println(<-ch)
	//通过range进行遍历值
	val, ok := <-ch
	if ok {
		fmt.Println(val, ok)
		//将通道关闭
		close(ch)
	}

}
