package main

import (
	"fmt"
	"math"
)

/*
    题目:商场收银软件 计算单价和总价,并展示
	//抽象三个类

	现金收费抽象类

	正常收费子类

	打折收费子类

	返利收费子类

	现金收费工厂类

	客户端类

 */

type CashSuper interface {
	AcceptCash(money float64) float64
}

type CashNormal struct {
}

func (this *CashNormal) AcceptCash(money float64) float64{
	return money
}

type CashRebate struct {
	moneyRebate float64
}

func (this *CashRebate) AcceptCash(money float64) float64{
	money = money * this.moneyRebate
	return money
}

type CashReturn struct {
	moneyCondition float64
	moneyReturn float64
}

func (this *CashReturn) AcceptCash(money float64) float64 {
	result := money
	if(money > this.moneyCondition){
		result = money - math.Floor(money/this.moneyCondition)*this.moneyReturn
	}
	return result
}

//创建一个工厂类
type CashFactory struct {
}

func(this *CashFactory) CreateCashAccept(oper string) (cashsuper CashSuper){
	switch oper {
	case "正常收费":
		cashsuper = &CashNormal{}
		break;
	case "打8折":
		cashsuper = &CashRebate{0.8}
		break;
	case "满300返100":
		cashsuper = &CashReturn{300,100}
		break;
	default:
		panic("运算符号错误！")
	}
	return
}


//客户端代码
var total float64

func AddGoods(txtPrice, txtNum float64){
	CF := &CashFactory{}
	csuper := CF.CreateCashAccept("满300返100")
	//执行相关策略
	totalPrice := csuper.AcceptCash(txtPrice * txtNum)
	fmt.Printf("单价:%f, 数量:%f, 合计:%f\n", txtPrice, txtNum, totalPrice)
	total += totalPrice
}

func ShowTotalPrice(){
	fmt.Printf("总价:%f\n",total)
}


func main() {

	AddGoods(230.2, 2)
	ShowTotalPrice()
}
