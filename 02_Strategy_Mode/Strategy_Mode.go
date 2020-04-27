package main

import (
	"fmt"
)

/*
    模式特点：定义算法家族并且分别封装，它们之间可以相互替换而不影响客户端，客户只知道一个类就可以了。
	程序实例：商场促销。

    题目:商场收银软件 计算单价和总价,并展示
	//抽象三个类

	现金收费抽象类

	正常收费子类

	打折收费子类

	返利收费子类

	现金收费工厂类

	客户端类

	上下文管理策略类

 */

type CashSuper interface {
	AcceptCash(money float64) float64
}

type CashNormal struct {
}

func (this *CashNormal) AcceptCash(money float64) float64{
	if money < 0 {
		panic("不允许负数作为参数")
	}
	return money
}

type CashRebate struct {
	moneyRebate float64
}

func (this *CashRebate) AcceptCash(money float64) float64{
	if money < 0 || this.moneyRebate < 0 {
		panic("不允许负数作为参数")
	}
	if this.moneyRebate >= 1 {
		panic("折扣无效")
	}
	return money * this.moneyRebate
}

type CashReturn struct {
	moneyCondition float64
	moneyReturn float64
}

func (this *CashReturn) AcceptCash(money float64) float64 {
	if money < 0 || this.moneyCondition < 0 || this.moneyReturn < 0 {
		panic("不允许负数作为参数")
	}
	if money >= this.moneyCondition {
		return money - float64(int(money/this.moneyCondition))*this.moneyReturn
	} else {
		return money
	}
}


//策略类
type CashContext struct {
	CashSuper CashSuper
}

func NewCashSuper(strategyType string) *CashContext{
	context := &CashContext{}
	switch strategyType {
	case "正常收费":
		context.CashSuper = &CashNormal{}
	case "8折":
		context.CashSuper = &CashRebate{0.8}
	case "满300返100":
		context.CashSuper = &CashReturn{300, 100}
	default:
		panic("计价方式错误！")
	}

	return context
}

func (this *CashContext)GetResult(money float64) float64 {
	return this.CashSuper.AcceptCash(money)
}


func main() {
	var total float64
	cash1 := NewCashSuper("正常收费")
	total += cash1.GetResult(1 * 10000)
	cash2 := NewCashSuper("满300返100")
	total += cash2.GetResult(1 * 10000)
	cash3 := NewCashSuper("8折")
	total += cash3.GetResult(1 * 10000)
	fmt.Println("total:", total)
}
