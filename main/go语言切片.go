package main

import "fmt"

func go_slice() {
	go_s1()
	go_s2()
	go_s3()
	slince_delete()
}
func go_s1() {
	//使用make定义一个切片
	var arr []int = make([]int, 5, 10) //1 切片类型 2 数组长度 3 数组最大容量
	//使用len函数 和cap函数分别查看数组长度和最大容量
	fmt.Println("当前切片数组的大小:", len(arr), "当前数组的最大容量:", cap(arr))
	//简写的切片
	arr1 := make([]string, 5, 10)
	s := "hello"
	arr1 = append(arr1, s) //向数组中追加一个元素
	fmt.Println(arr1)
	fmt.Println(arr1[5])
	fmt.Println(len(arr1)) //6
	arr1Copy := make([]string, 2*len(arr1))
	copy(arr1Copy, arr1) //将数组arr1中数据复制到arr1Copy中，并改变arr1Copy的数组大小
	fmt.Println("新的数组的大小:", len(arr1Copy))

}
func go_s2() {
	//定义并初始化一个切片 其中 len=cap=6
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	//引用另一个数组实现切片的初始化
	//切片数组arr中arr[1]到arr[3]的全部元素
	s := arr[1:4] //1是数组开始位置包含 4是数组结束位置不包含
	fmt.Println("新的数组切片:", s)
	s1 := arr[1:] //1 - len-1
	fmt.Println("新的数组切片:", s1)
	s2 := arr[:4] // 0 - 3
	fmt.Println("新的数组切片:", s2)
}
func go_s3() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
//数组或切片中的删除
func slince_delete() {
	arr := []int{1, 2, 3, 4}
	//删除下标为2的元素
	arr = append(arr[:2], arr[3:]...)
	fmt.Println(arr)
}
