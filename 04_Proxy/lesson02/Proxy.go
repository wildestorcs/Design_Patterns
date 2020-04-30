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
 	Proxy
  */

 type Proxy struct {
 	Mm *SchoolGirl
 }

 func NewProxy(mm *SchoolGirl) *Proxy{
 	return &Proxy{Mm:mm}
 }

 func (this *Proxy) GiveDolls(){
 	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
 }

func (this *Proxy) GiveFlowers(){
	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
}

func (this *Proxy) GiveChocolate(){
	fmt.Printf("%s, 送你洋娃娃\n",this.Mm.Name)
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
