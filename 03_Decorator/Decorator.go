package main

import (
	"fmt"
)

/*
	模式特点：动态地为对象增加额外的职责
	程序实例：展示一个人一件一件穿衣服的过程。 装饰器内可以添加各种操作

	装饰器模式
	人和服装类分开
 */

//抽象构件 (Component)
type AbsPerson interface {
	Show()
}


//具体构件 (ConcreteComponent)
type Person struct {
	AbsPerson
	Name string
}

func NewPerson(name string) *Person{
	return &Person{Name:name}
}

func (p *Person) Show(){
	fmt.Printf("装扮的%s\n", p.Name)
}


//抽象装饰 (Decorator)
type Finery struct {
	*Person
	absperson AbsPerson
}

func (this *Finery) Show(){
	this.absperson.Show()
}


//具体装饰 (ConcreteDecorator)
type TShirts struct {
	*Finery
}

func NewTShirts(absperson AbsPerson) *TShirts{
	return &TShirts{&Finery{absperson:absperson}}
}

func (this *TShirts) Show(){
	fmt.Println("大T恤")
	this.absperson.Show()
}

type BigTrouser struct{
	*Finery
}

func NewBigTrouser(absperson AbsPerson) *BigTrouser{
	return &BigTrouser{&Finery{absperson:absperson}}
}

func (this *BigTrouser) Show(){
	fmt.Println("垮裤")
	this.absperson.Show()
}


type Sneakers struct {
	*Finery
}

func NewSneakers(absperson AbsPerson) *Sneakers{
	return &Sneakers{&Finery{absperson:absperson}}
}

func (this *Sneakers) Show(){
	fmt.Println("破球鞋")
	this.absperson.Show()
}

type Suit struct {
	*Finery
}

func NewSuit(absperson AbsPerson) *Suit{
	return &Suit{&Finery{absperson:absperson}}
}

func (this *Suit) Show(){
	fmt.Println("西装")
	this.absperson.Show()
}

type Tie struct {
	*Finery
}

func NewTie(absperson AbsPerson) *Tie{
	return &Tie{&Finery{absperson:absperson}}
}


func (this *Tie) Show(){
	fmt.Println("领带")
	this.absperson.Show()
}

type LeatherShoes struct {
	*Finery
}

func NewLeatherShoes(absperson AbsPerson) *LeatherShoes{
	return &LeatherShoes{&Finery{absperson:absperson}}
}

func (this *LeatherShoes) Show(){
	fmt.Println("皮鞋")
	this.absperson.Show()
}


func main() {
	var person = NewPerson("JayChou")
	fmt.Println("\n第一种装扮: ")
	ts := NewTShirts(person)
	bt := NewBigTrouser(ts)
	sn := NewSneakers(bt)
	sn.Show()

	//===============================
	fmt.Println("\n第二种装扮: ")
	suit := NewSuit(person)
	tie := NewTie(suit)
	ls := NewLeatherShoes(tie)
	ls.Show()
}
