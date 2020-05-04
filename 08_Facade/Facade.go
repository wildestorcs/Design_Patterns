package main

import "fmt"

/*
	外观模式

	模式特点：外观模式为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得子系统更加容易使用。
	程序实例：经常使用三层结构：controller 控制器层，service 服务层，dao 数据访问层。本例为买卖股票

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
	fmt.Println("房地产购买。。。")
}


type Fund struct {
   S 	*Stock
   N	*NationalDebt
   R	*Realty
}

func(this *Fund)BuyFund(){
	this.S.Buy()
	this.N.Buy()
	this.R.Buy()
}

func(this *Fund)SellFund(){
	this.S.Sell()
	this.N.Sell()
	this.R.Sell()
}


func main() {
	var jijin = &Fund{}

	jijin.BuyFund()

	jijin.SellFund()

}
