package _9_Builder

import "fmt"

/*
    建造者模式
	Builder: 抽象建造者
	ConcreteBuilder: 具体建造者
	Director: 指挥者
	Production: 产品(大型产品以及拆分成的小型产品)

	模式特点：将一个复杂对象的构建(Director)与它的表示(Builder)分离，使得同样的构建过程可以创建不同的表示(ConcreteBuilder)。
	程序实例：“画”出一个四肢健全（头身手腿）的小人

 */


 //产品类Production
type PersonBuilding struct {
	Head		string
	Body		string
	LeftHand 	string
	RightHand   string
	LeftFoot    string
	RightFoot   string
}

 //建造类Builder
 //建造者接口
type PersonBuilder interface {
	DrawHead()
	DrawBody()
	DrawLeftHand()
	DrawRightHand()
	DrawLeftFoot()
	DrawRightFoot()
	GetResult() *PersonBuilding
}

 //具体建造类ConcreteBuilder
type PersonThinBuilder struct {
	building *PersonBuilding
}

func NewPersonThinBuilder() *PersonThinBuilder {
	return &PersonThinBuilder{new(PersonBuilding)}
}

func(this *PersonThinBuilder)DrawHead(){
	this.building.Head = "draw thin head ..."
}

func(this *PersonThinBuilder)DrawBody(){
	this.building.Body = "draw thin body ..."
}

func(this *PersonThinBuilder)DrawLeftHand(){
	this.building.LeftHand = "draw thin left hand ..."
}

func(this *PersonThinBuilder)DrawRightHand(){
	this.building.RightHand = "draw thin right hand ..."
}

func(this *PersonThinBuilder)DrawLeftFoot(){
	this.building.LeftFoot =  "draw thin left foot ..."
}

func(this *PersonThinBuilder)DrawRightFoot(){
	this.building.RightFoot =  "draw thin right foot ..."
}

func(this *PersonThinBuilder)GetResult() *PersonBuilding{
	return this.building
}

//指挥者类Director
//========================================================================
 type PersonDirector struct {
 	pb PersonBuilder
 }

//生成
func(this *PersonDirector)CreatePerson(){
	this.pb.DrawHead()
	this.pb.DrawBody()
	this.pb.DrawLeftHand()
	this.pb.DrawRightHand()
	this.pb.DrawLeftFoot()
	this.pb.DrawRightFoot()
}

func (this *PersonDirector)GetResult() *PersonBuilding{
	return this.pb.GetResult()
}


func main() {

	ptb := NewPersonThinBuilder()
	pdThin := &PersonDirector{ptb}
	pdThin.CreatePerson()
	pdRes:= pdThin.GetResult()
	fmt.Println(pdRes)





	
}
