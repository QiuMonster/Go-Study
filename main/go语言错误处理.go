package main

import (
	"errors"
	"fmt"
	"math"
)

//go的错误处理是一个接口类型
type error interface{
	Error()string
}
func go_error(){
	 fmt.Println(Sqrt(-2.0))
}

func Sqrt(f float64)(float64,error){
	if f<0{
		return 0,errors.New("不能给负数开方")
	}
	//当没有错误时 error是nil
	return math.Sqrt(f),nil
}