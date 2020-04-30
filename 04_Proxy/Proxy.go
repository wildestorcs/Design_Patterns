package main

import "fmt"

/*
   Proxy

   模式特点：为其他对象提供一种代理以控制对这个对象的访问。

   模式好处：对外部提供统一的接口方法，而代理类在接口中实现对真实类的附加操作行为，从而达到了“对修改关闭，对扩展开放”（OCP原则）

   应用场景：当我们需要使用的对象很复杂或者需要很长时间去构造，这时就可以使用代理模式。

 */

 type IGiveGift interface {
	 GiveDolls()
	 GiveFlowers()
	 GiveChocolate()
 }


/*
 	SchoolGirl
 */

 type SchoolGirl struct {
 	Name string
 }

/*
   Pursuit
*/

type Pursuit struct {
	Mm *SchoolGirl
}

func NewPursuit(mm *SchoolGirl) *Pursuit{
	return &Pursuit{Mm:mm}
}

func (this *Pursuit) GiveDolls(){
	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
}

func (this *Pursuit) GiveFlowers(){
	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
}

func (this *Pursuit) GiveChocolate(){
	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
}

/*
	Proxy
 */

 type Proxy struct {
	 IGiveGift
	 Gg *Pursuit

 }

 func NewProxy(mm *SchoolGirl) *Proxy{
 	return &Proxy{Gg:NewPursuit(mm)}
 }

 func (this *Proxy) GiveDolls(){
 	this.Gg.GiveDolls()
 }

func (this *Proxy) GiveFlowers(){
	this.Gg.GiveFlowers()
}

func (this *Proxy) GiveChocolate(){
	this.Gg.GiveChocolate()
}


 /*
    客户端调用
  */

func main() {
	var jiaojiao = new(SchoolGirl)
	jiaojiao.Name = "李娇娇"

	var daili = NewProxy(jiaojiao)

	daili.GiveDolls()
	daili.GiveFlowers()
	daili.GiveChocolate()
	
}
