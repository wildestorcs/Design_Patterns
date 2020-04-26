package lession02

import (
	"fmt"
	"strconv"
)

/*
	何为简单工厂,拿计算器来说

 */


func main() {
	//获取输入的写法
	var (
		A float64
		B string
		C float64
		Result string
	)

	fmt.Println("请输入数字 A:")
	fmt.Scan(&A)
	fmt.Println("请选择运算符 B:")
	fmt.Scan(&B)
	fmt.Println("请输入数字 C:")
	fmt.Scan(&C)


	//获取相关数字
	if(B == "+"){
		Result = strconv.FormatFloat(A+C, 'f', -1,64)
	}else if (B == "-") {
		Result = strconv.FormatFloat(A-C, 'f', -1,64)
	}else if(B == "*") {
		Result = strconv.FormatFloat(A*C, 'f', -1,64)
	}else if(B == "/"){
		Result = strconv.FormatFloat(A/C, 'f', -1,64)
	}

	fmt.Printf("Result is %s", Result)

}
