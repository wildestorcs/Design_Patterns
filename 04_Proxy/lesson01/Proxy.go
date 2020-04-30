package main

import "fmt"

/*
   Not Proxy
 */


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
    客户端调用
  */

func main() {
	var jiaojiao = new(SchoolGirl)
	jiaojiao.Name = "李娇娇"

	var zhuojiayi = NewPursuit(jiaojiao)

	zhuojiayi.GiveDolls()
	zhuojiayi.GiveFlowers()
	zhuojiayi.GiveChocolate()
	
}
