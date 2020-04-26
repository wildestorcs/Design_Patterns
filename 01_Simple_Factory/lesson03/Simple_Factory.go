package main

import (
	"fmt"
	"strconv"
)

/*
	计算和显示分开 ---封装
 */

func getInput(){
	//获取输入的写法
	var (
		NumberA float64
		Operate string
		NumberB float64
	)

	fmt.Println("请输入数字 NumberA:")
	fmt.Scan(&NumberA)
	fmt.Println("请选择运算符 Operate:")
	fmt.Scan(&Operate)
	fmt.Println("请输入数字 NumberB:")
	fmt.Scan(&NumberB)
	Result := getResult(NumberA, NumberB, Operate)
	fmt.Printf("Result is %s", Result)

}

func getResult(NumberA , NumberB float64, Operate string) string{
	var Result string
	switch Operate {
	case "+":
		Result = strconv.FormatFloat(NumberA+NumberB, 'f', -1,64);
		break;
	case "-":
		Result = strconv.FormatFloat(NumberA-NumberB, 'f', -1,64)
		break;
	case "*":
		Result = strconv.FormatFloat(NumberA*NumberB, 'f', -1,64)
		break;
	case "/":
		if (NumberB == 0){
			Result = "除数不能为0"
		}else {
			Result = strconv.FormatFloat(NumberA/NumberB, 'f', -1,64)
		}
		break;
	}
	return Result
}

func main() {
	getInput()
}
