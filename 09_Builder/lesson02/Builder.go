package main

import "fmt"

/*
    建造者模式  ---简单例子
 */

 type Pen struct {
 	Name string
 }

 type Graphics struct {
 	Content string
 }


 type ThinBuilder struct {
	 P *Pen
	 G *Graphics
 }

 func NewThinBuilder(p *Pen, g *Graphics) *ThinBuilder{
	 return &ThinBuilder{P:p,G:g}
 }


 func (this *ThinBuilder) Build(){
 	this.G.DrawHead(this.P)
 	this.G.DrawBody(this.P)
 	this.G.DrawLeftHand(this.P)
 	this.G.DrawRightHand(this.P)
 	this.G.DrawLeftFoot(this.P)
 	this.G.DrawRightFoot(this.P)
 }



 func NewGraphics(content string) *Graphics{
 	return &Graphics{Content:content}
 }
 
 func(this *Graphics)DrawHead(p *Pen){
 	fmt.Printf("%s: %s 画头。。。\n",this.Content ,p.Name)
 }

 func(this *Graphics)DrawBody(p *Pen){
	fmt.Printf("%s: %s 画身体。。。\n",this.Content ,p.Name)
}

 func(this *Graphics)DrawLeftHand(p *Pen){
	fmt.Printf("%s: %s 画左手。。。\n",this.Content ,p.Name)
}

func(this *Graphics)DrawRightHand(p *Pen){
	fmt.Printf("%s: %s 画右手。。。\n",this.Content ,p.Name)
}

func(this *Graphics)DrawLeftFoot(p *Pen){
	fmt.Printf("%s: %s 画左脚。。。\n",this.Content ,p.Name)
}

func(this *Graphics)DrawRightFoot(p *Pen){
	fmt.Printf("%s: %s 画右脚。。。\n",this.Content ,p.Name)
}

func main() {
	p := &Pen{Name:"马克笔"}
	g := NewGraphics("瘦子")
	PersonBuilder := NewThinBuilder(p,g)
	PersonBuilder.Build()






	
}
