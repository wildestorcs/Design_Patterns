package main

import "fmt"

/*
    观察者模式的例子
 */

//被观察者接口
type Subject interface {
	 Attach(observer *Observer)
	 Detach(observer *Observer)
	 Notify()
	 SubjectSetAction(action string)
	 SubjectGetAction() string

 }


//观察者
type Observer interface {
	Update(string)
}


//被观察者实例
type Boss struct {
	observers []Observer
	action string
}

func NewBoss(action string) *Boss{
	return &Boss{action:action, observers:[]Observer{}}
}


func(this *Boss)Attach(ob Observer){
	this.observers = append(this.observers, ob)
}

func(this *Boss)Detach(ob Observer){
	for i,sv := range this.observers{
		if sv == ob{
			this.observers = append(this.observers[:i],this.observers[i+1:]...)
		}
	}
}

func(this *Boss)Notify(){
	for _,person := range this.observers{
		person.Update(this.action)
	}
}

func(this *Boss)SubjectSetAction(action string){
	this.action = action
}

func(this *Boss)SubjectGetAction() string{
	return this.action
}

//观察者实例
type StockObserver struct {
	Name string
	SelfAction string
}

//创建观察者实例
func NewStockObserver(name string) *StockObserver{
	return &StockObserver{Name:name}
}

func (this *StockObserver) Update(action string){
	this.SelfAction = action
	fmt.Printf("%s,%s 关闭股票行情,继续工作！\n", this.SelfAction, this.Name)
}

type NBAObserver struct {
	Name string
	SelfAction string
}

//创建观察者实例
func NewNBAObserver(name string) *NBAObserver{
	return &NBAObserver{Name:name}
}

func (this *NBAObserver) Update(action string){
	this.SelfAction = action

	fmt.Printf("%s,%s 关闭NBA,继续工作！\n", this.SelfAction, this.Name)
}

func main() {
	huhansan := NewBoss("老板走了")
	tongshi1 := NewStockObserver("魏关姹")
	tongshi2 := NewNBAObserver("易管查")

	huhansan.Attach(tongshi1)
	huhansan.Attach(tongshi2)
	//huhansan.Detach(tongshi2)


	huhansan.action ="我胡汉三回来了"

	huhansan.Notify()

}
