package main

import (
	"fmt"
)

// import "unsafe"

func go_tiaojian() {
	go_swich()
	go_select()
	go_type_switch()
	go_fallthrough_switch()
}

//if else    switch  select
func go_swich() {
	var grade string = "B"
	var mark int = 90
	//存在条件的swich
	switch mark {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	//无条件的switch
	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "C" || grade == "B":
		// case grade=="C",grade=="B":
		fmt.Println("良好")
	case grade == "F":
		fmt.Println("及格")
	default:
		fmt.Println("差")
	}
	fmt.Println("你的等级是:", grade)
}

//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
func go_type_switch() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型是:%T", i)
	case int:
		fmt.Printf("x 是int型")
	case float64:
		fmt.Printf("x是float64类型")
	case func(int) float64:
		fmt.Printf("x是func(int)类型")
	default:
		fmt.Printf("x是未知类型")
	}
}

//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func go_fallthrough_switch() {
	switch {
	case false:
		fmt.Println("1 case的条件是 false")
		fallthrough
	case true:
		fmt.Println("2 case的条件是 false")
		fallthrough
	case false:
		fmt.Println("3 case的条件是 false")
		fallthrough
	default:
		fmt.Println("fallthrough后的条件无论对错都会执行")
	}
}

//什么意思不理解
func go_select() {
	var c1, c2, c3 chan int //chan使用管道
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received", i1, "from c1\n")
	case c2 <- i2:
		fmt.Printf("sent", i1, "to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")

	}
}
