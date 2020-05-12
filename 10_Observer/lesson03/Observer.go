package lesson03

import "fmt"

/*
    观察者模式的例子
 */

//观察者
type Subject interface {
	 Attach(oberser *Observer)
	 Notify()
	 SecretaryAction(action string)
 }

//定义 被观察者
type IObserver interface {
	Update()
}


//前台秘书类
//定义 观察者
type Secretary struct {
	observers []*Observer
	action string
}

func NewSecretary() *Secretary{
	return &Secretary{}
}

//添加同事
func (this *Secretary) Attach(oberser *Observer){
	this.observers = append(this.observers, oberser)
}

//通知
func (this *Secretary) Notify(){
	for _,person :=range this.observers{
		//fmt.Println(person.Name)
		//fmt.Println(person.Update)
		person.Update()
	}
}

func (this *Secretary) SecretaryAction(action string){
	this.action = action
}




type Observer struct {
	Name string
	Sub *Secretary
}

func (this *Observer)Update(){

}

//看股票的同事类
type StockObserver struct {
	Observer
}

func NewStockObserver(name string, sub *Secretary) *Observer{
	return &Observer{Name:name, Sub:sub}
}

func (this *StockObserver) Update(){
	fmt.Printf("%s,%s 关闭股票行情,继续工作！\n", this.Sub.action, this.Name)
}


//看NBA的同事类
type NBAObserver struct {
	Observer
}

func NewNBAObserver(name string, sub *Secretary) *Observer{
	return &Observer{Name:name, Sub:sub}
}

func (this *NBAObserver) Update(){
	fmt.Printf("%s,%s 关闭NBA直播,继续工作！\n", this.Sub.action, this.Name)
}


func main() {
	var tongzizhe = NewSecretary()

	//同事1
	tongshi1 := NewStockObserver("魏关姹", tongzizhe)
	tongshi2 := NewNBAObserver("易管查", tongzizhe)


	fmt.Println(111111)
	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)
	fmt.Println(222222)
	tongzizhe.SecretaryAction("老板回来了!")
	tongzizhe.Notify()
}
