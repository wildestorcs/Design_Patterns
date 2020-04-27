package main

import "fmt"

/*
    题目:商场收银软件 计算单价和总价,并展示
 */

const (discount = iota
       discount8
	   discount7
)



var total float64

func addGoods(txtPrice, txtNum float64, discount int){
	totalPrice := 0.0
	switch discount {
		case 0:
			totalPrice = txtPrice * txtNum
			break
		case 1:
			totalPrice = txtPrice * txtNum * 0.8
			break
		case 2:
			totalPrice = txtPrice * txtNum * 0.7
			break
	}
	fmt.Printf("单价:%f, 数量:%f, 合计:%f\n", txtPrice, txtNum, totalPrice)
	total += totalPrice

}

func showTotalPrice(){
	fmt.Printf("总价:%f\n",total)
}


func main() {
	addGoods(23.2, 1.5, discount8)
	addGoods(23.2, 2, discount7)
	showTotalPrice()
}
