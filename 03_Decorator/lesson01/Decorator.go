package main

import (
	"fmt"
)

/*
	装饰器模式
 */


 type Person struct {
 	Name string
 }

func (p Person) WearTShirts(){
	fmt.Println("大T恤")
}

func (p Person) WearBigTrouser(){
	fmt.Println("垮裤")
}

func (p Person) WearSneakers(){
	fmt.Println("破球鞋")
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

func (p Person) Show(){
	fmt.Printf("装扮的%s\n", p.Name)
}

func main() {
	var p1 = Person{Name:"JayChou"}
	fmt.Println("\n第一种装扮: ")
	p1.WearTShirts()
	p1.WearBigTrouser()
	p1.WearSneakers()
	p1.Show()

	fmt.Println("\n第二种装扮: ")
	p1.WearSuit()
	p1.WearTie()
	p1.WearLeatherShoes()
	p1.Show()

	
}
