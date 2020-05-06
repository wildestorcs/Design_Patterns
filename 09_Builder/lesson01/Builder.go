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
	var p = &Pen{Name:"马克笔"}
	Thin := NewGraphics("瘦子")
	Thin.DrawHead(p)
	Thin.DrawBody(p)
	Thin.DrawLeftHand(p)
	Thin.DrawRightHand(p)
	Thin.DrawLeftFoot(p)
	Thin.DrawRightFoot(p)



	
}
