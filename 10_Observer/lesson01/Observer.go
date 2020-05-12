package main

import "fmt"

/*
    观察者模式的例子
 */

//前台秘书类
type Secretary struct {
	observers []*StockObserver
	action string
}

func NewSecretary() *Secretary{
	return &Secretary{}
}

//添加同事
func (this *Secretary) Attach(oberser *StockObserver){
	this.observers = append(this.observers, oberser)
}

//通知
func (this *Secretary) Notify(){
	for _,person :=range this.observers{
		person.Update()
	}
}

func (this *Secretary) SecretaryAction(action string){
	this.action = action
}



//看股票的同事类
type StockObserver struct {
	Name string
	Sub *Secretary
}

func NewStockObserver(name string, sub *Secretary) *StockObserver{
	return &StockObserver{Name:name, Sub:sub}
}

func (this *StockObserver) Update(){
	fmt.Printf("%s,%s 关闭股票行情,继续工作！\n", this.Sub.action, this.Name)
}



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
