package main

import "fmt"

func go_circle() {
	sumNum()
	sumNum1()
	go_for_range()
	go_for_map()
	sushu()
	go_goto()
}

//计算 1到10的数字之和 for循环操作 带上范围
func sumNum() {
	// var sum int
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1到10相加的结果是:", sum)
}

//类似于while的for样式
func sumNum1() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println("sum循环相加的结果是:", sum)
}

//for each range 循环将数组中的数据连通下标一取出
func go_for_range() {
	//定义大小为5的字符数组
	strs := []string{"hello", "world"}
	for i, s := range strs {
		//循环在将数据取完自动结束
		fmt.Printf("下标%d对应的字符值是:%s\n", i, s)
	}
}

//在使用map时可以省略 key或 value
func go_for_map() {
	//使用make关键字创建一个map对象
	mymap := make(map[int]int) //int 是key的值 float32 是value的值
	mymap[0] = 10
	mymap[1] = 11
	mymap[2] = 12
	for key, value := range mymap {
		fmt.Printf("key值是:%d,对应的value值是:%d\n", key, value)
	}
	for key := range mymap {
		fmt.Printf("key值是:%d,不显示的value值\n", key)
	}
	//注意在省略key时要在前面添加一个占位符 _
	for _, value := range mymap {
		fmt.Printf("不显示key的值,value值是:%d\n", value)
	}
}

//判断前100个数是否是素数
func sushu() {
	var i, j int
	for i = 2; i <= 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break //结束不是素数 发现因子
			}
		}
		if j > (i / j) {
			fmt.Println(i, "是一个素数", j)
		}
	}
}

//循环中的goto使用 在循环前面添加一个标签 使用goto进入
func go_goto() {
	a := 10
	for a < 20 {
		a += 1
		if a == 15 {
			goto FLAG1 //会跳过a==15的输出
		}
		fmt.Println("a的值是:", a)
	}

FLAG1:
	for i := 0; i < 5; i++ {
		fmt.Println("i的值是:", i, "通过goto进行调用的")
	}
}
