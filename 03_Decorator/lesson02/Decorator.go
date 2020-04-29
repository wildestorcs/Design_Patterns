package main

import (
	"fmt"
)

/*
	装饰器模式

	人和服装类分开

 */


 type Person struct {
 	Name string
 }

func (p Person) Show(){
	fmt.Printf("装扮的%s\n", p.Name)
}

type Finery struct {

}

type AbsFinery interface {
	Show()
}

type TShirts struct {
	*Finery
}

func (this *TShirts) Show(){
	fmt.Println("大T恤")
}

type BigTrouser struct{
	*Finery
}

func (this *BigTrouser) Show(){
	fmt.Println("垮裤")
}

type Sneakers struct {
	*Finery
}

func (this *Sneakers) Show(){
	fmt.Println("破球鞋")
}

type Suit struct {
	*Finery
}

func (this *Suit) Show(){
	fmt.Println("西装")
}

type Tie struct {
	*Finery
}

func (this *Tie) Show(){
	fmt.Println("领带")
}

type LeatherShoes struct {
	*Finery
}

func (this *LeatherShoes) Show(){
	fmt.Println("皮鞋")
}

func (p Person) WearSuit(){
	fmt.Println("西装")
}

func (p Person) WearTie(){
	fmt.Println("领带")
}

func (p Person) WearLeatherShoes(){
	fmt.Println("皮鞋")
}



func main() {
	var p1 = Person{Name:"JayChou"}
	fmt.Println("\n第一种装扮: ")
	ts := new(TShirts)
	bt := new(BigTrouser)
	sn := new(Sneakers)
	ts.Show()
	bt.Show()
	sn.Show()
	p1.Show()
	//p1.WearTShirts()
	//p1.WearBigTrouser()
	//p1.WearSneakers()
	//p1.Show()

	fmt.Println("\n第二种装扮: ")
	p1.WearSuit()
	p1.WearTie()
	p1.WearLeatherShoes()
	p1.Show()

	
}
