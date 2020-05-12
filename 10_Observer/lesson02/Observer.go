package main

import "fmt"

/*
    观察者模式的例子
 */

//通知者接口
type Subject interface {
	Attach(oberser *Observer)
	Notify()
	SecretaryAction(action string)
}

//观察者抽象
type Observer interface {
	Update()
}




//BOSS实现
type Boss struct {
	observers []Observer
	action string
}

func(this *Boss)Attach(oberser Observer){
	this.observers = append(this.observers, oberser)
}

func (this *Boss) Notify(){
	for _,person :=range this.observers{
		person.Update()
		//数组种类的问题  数组里面为一类数据类型
	}
}

func (this *Boss) SecretaryAction(action string){
	this.action = action
}


////前台秘书类
//type Secretary struct {
//	observers []*Observer
//	action string
//}
//
//func NewSecretary() *Secretary{
//	return &Secretary{}
//}
//
////添加同事
//func (this *Secretary) Attach(oberser *Observer){
//	this.observers = append(this.observers, oberser)
//}
//
////通知
//func (this *Secretary) Notify(){
//	for _,person :=range this.observers{
//		person.Update()
//		//数组种类的问题  数组里面为一类数据类型
//	}
//}

//func (this *Secretary) SecretaryAction(action string){
//	this.action = action
//}




//type IObserver interface {
//	Update()
//}
//
////看股票的同事类
//type StockObserver struct {
//	Observer
//}
//
//func NewStockObserver(name string, sub *Secretary) *StockObserver{
//	return &StockObserver{Observer{Name:name,Sub:sub}}
//}
//
//func (this *StockObserver) Update(){
//	fmt.Printf("%s,%s 关闭股票行情,继续工作！\n", this.Sub.action, this.Name)
//}
//
////看NBA的同事类
//type NBAObserver struct {
//	Observer
//}
//
//func NewNBAObserver(name string, sub *Secretary) *NBAObserver{
//	return &NBAObserver{Observer{Name:name,Sub:sub}}
//}
//
//func (this *NBAObserver) Update(){
//	fmt.Printf("%s,%s 关闭NBA直播,继续工作！\n", this.Sub.action, this.Name)
//}





func main() {
	var tongzizhe = NewSecretary()

	//同事1
	tongshi1 := NewStockObserver("魏关姹", tongzizhe)
	tongshi2 := NewStockObserver("易管查", tongzizhe)

	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)

	tongzizhe.SecretaryAction("老板回来了!")
	tongzizhe.Notify()
}
