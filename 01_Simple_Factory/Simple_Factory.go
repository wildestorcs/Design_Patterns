package main

import (
	"fmt"
	"strconv"
)


/*
 	模式特点: 工厂类根据条件产生不同功能的运算对象,客户端不需要具体的运算类。

	程序实例: 四则运算计算器,根据用户的输入产生相应的运算类,用这个运算类处理具体的运算。


	计算和显示分开 --继承/多态

 */

//=========== 运算类 ============

type Operation struct{
	NumberA float64
	NumberB float64
}

//=========== 运算类方法 ============
//初始化变量
func (this *Operation) Init(a float64, b float64){
	this.NumberA = a
	this.NumberB = b
}

//定义接口
type Operator interface {
	getResult() string
	Init(a float64, b float64)
}

//定义加减乘除方法
type Add struct {
	Operation
}

func(this *Add) getResult() string{
	return strconv.FormatFloat(this.NumberA+this.NumberB, 'f', -1,64);
}

type Sub struct {
	Operation
}

func(this *Sub) getResult() string{
	return strconv.FormatFloat(this.NumberA-this.NumberB, 'f', -1,64)
}

type Mul struct {
	Operation
}

func(this *Mul) getResult() string{
	return strconv.FormatFloat(this.NumberA*this.NumberB, 'f', -1,64)
}

type Div struct {
	Operation
}

func(this *Div) getResult() string{
	if(this.NumberB == 0){
		return "除数不能为0"
	}
	return strconv.FormatFloat(this.NumberA/this.NumberB, 'f', -1,64)
}


type OpeateFactory struct {
}

func(this *OpeateFactory) CreateOperator(oper string) (operator Operator){
	switch oper {
	case "+":
		operator = &Add{}
		break;
	case "-":
		operator = &Sub{}
		break;
	case "*":
		operator = &Mul{}
		break;
	case "/":
		operator = &Div{}
		break;
	default:
		panic("运算符号错误！")
	}
	return
}



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

	OF := &OpeateFactory{}
	add := OF.CreateOperator(Operate)
	add.Init(NumberA, NumberB)
	Result := add.getResult()
	fmt.Printf("Result is %s", Result)

}

func main() {
	getInput()
}
