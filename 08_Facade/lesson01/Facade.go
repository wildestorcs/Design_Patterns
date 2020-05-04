package main

import "fmt"

/*
	外观模式

 */


type Stock struct {
}

func(this *Stock)Sell(){
	fmt.Println("股票卖出。。。")
}

func(this *Stock)Buy(){
	fmt.Println("股票购买。。。")
}

type NationalDebt struct {
}

func(this *NationalDebt)Sell(){
	fmt.Println("国债卖出。。。")
}

func(this *NationalDebt)Buy(){
	fmt.Println("国债购买。。。")
}

type Realty struct {
}

func(this *Realty)Sell(){
	fmt.Println("房地产卖出。。。")
}

func(this *Realty)Buy(){
	fmt.Println("房地产卖出。。。")
}

func main() {
	var gupiao = &Stock{}
	var national_debt = &NationalDebt{}
	var realty = &Realty{}

	gupiao.Buy()
	gupiao.Sell()

	national_debt.Buy()
	national_debt.Sell()

	realty.Buy()
	realty.Sell()
}
