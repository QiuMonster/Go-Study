package main

import "fmt"

func go_arr2() {
	a := [5]int{1, 2, 3, 4, 5}
	avgnum:= get_avg(a)
	fmt.Println("数组的平均值是:",avgnum)

}

//向函数传递 指定大小的数组参数
func get_avg(param [5]int) float32 {
	var sum, avg float32
	l := len(param)
	for _, val := range param {
		sum += float32(val)
	}
	avg = sum / float32(l)
	return avg
}

//传递不指定大小的数组参数
func ad_arr1(param []int) {

}

// func getAverage(arr []int, size int) float32{
//    var i int
//    var avg, sum float32

//    for i = 0; i < size; ++i {
//       sum += arr[i]
//    }

//    avg = sum / size

//    return avg;
// }
