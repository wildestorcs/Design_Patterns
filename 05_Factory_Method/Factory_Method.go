package main

import (
	"fmt"
	"strconv"
)

/*
	工厂模式
	模式特点：定义一个用于创建对象的接口，让子类决定实例化哪一个类。这使得一个类的实例化延迟到其子类。
	程序实例：计算器。
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

type IFactory interface {
	CreatOperation() Operator
}


type AddFactory struct {
}

type SubFactory struct {
}

type MulFactory struct {
}

type DivFactory struct {
}

//定义接口
type Operator interface {
	getResult() string
	Init(a float64, b float64)
}

func(this *AddFactory) CreatOperation() (operator Operator){
	return &Add{}
}

func(this *SubFactory) CreatOperation() (operator Operator){
	return &Sub{}
}

func(this *MulFactory) CreatOperation() (operator Operator){
	return &Mul{}
}

func(this *DivFactory) CreatOperation() (operator Operator){
	return &Div{}
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

	operFactory := new(AddFactory) //此处可以替换为其他运算
	oper := operFactory.CreatOperation()
	oper.Init(NumberA, NumberB)
	Result := oper.getResult()
	fmt.Printf("Result is %s", Result)

}





func main()  {
	getInput()
}