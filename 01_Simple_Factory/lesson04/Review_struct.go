package main

import (
	"fmt"
	"unsafe"
)

/*
   类型定义
 */

func typeTest(){
	//类型别名  给原有的数据类型起别名
	//类型定义  新建了一个数据类型

	//类型测试
	type Newint int
	type Myint = int

	//定义类型 --声明变量
	var a Newint
	var b Myint
	fmt.Printf("Newint is the type of: %T\n", a)
	fmt.Printf("Myint is the type of: %T\n", b)
}

/*
   定义结构体  //结构体是一组数据类型
 */

type Person struct {
	name string
	city string
	age  int8
}

//================= 实例化 ====================

//结构体实例  传统方式 var+声明的变量
func structTest(){
	var p1 Person
	p1.name = "LiuWei"
	p1.city = "Beijing"
	p1.age = 20

	//fmt.Println(p1)
	fmt.Printf("P1 is %v\n", p1)
	fmt.Printf("P1 is %#v\n", p1)
}

//匿名结构体
func structTest2(){
	var p2 struct{Name string; Age int}
	p2.Name = "JayChou"
	p2.Age = 32
	fmt.Printf("The P2 Value is %#v", p2)
}


//结构体实例3用new关键字实例化
func structTest3(){
	var p3 = new(Person)
	p3.name = "Jaychou"
	p3.city = "Shanghai"
	p3.age = 20
	//fmt.Println(p1)
	fmt.Printf("P2 type is %T\n", p3)
	fmt.Printf("P2 is %v\n", p3)
	fmt.Printf("P2 is %#v\n", p3)
}

//取结构体地址实例化
func structTest4(){
	p4 := &Person{}
	fmt.Printf("P4 type is %T\n", p4)
	fmt.Printf("P4 is %#v\n", p4)
	p4.name = "Hello1"
	p4.age = 29
	p4.city = "Shenzhen"
	fmt.Printf("P4 is %#v\n", p4)
}


//================= 初始化 ==================
//如果没有初始化  则对应各类型其零值

func structInit(){
	var p5 Person
	fmt.Printf("P5 is %#v\n", p5)
}

//键值对的初始化
func structInit2(){
	p6 := Person{
		name:"person6",
		age: 30,
		city:"Anyang",
	}
	fmt.Printf("P6 is %#v\n", p6)
}

//结构体内内存变量的内容  这个排兵布阵法
func structInit3(){
	p7 := &Person{
		name:"person7",
		age: 30,
		city:"Anyang",
	}
	fmt.Printf("P7 is %#v\n", p7)
	fmt.Printf("P7.name is %p\n", &p7.name)
	fmt.Printf("P7.age is %p\n", &p7.age)
	fmt.Printf("P7.city is %p\n", &p7.city)
}

func structInit4(){
	p8 := &Person{
		name:"person8",
		age: 30,
	}
	fmt.Printf("P8 is %#v\n", p8)
}

//使用值的列表初始化  !!! 注意顺序不能变
func structInit5(){
	p9 := &Person{
		"person9",
		"guangzhou",
		 30,

	}
	fmt.Printf("P9 is %#v\n", p9)
}

//空结构体不占空间
func structEmpty(){
	var hello struct{}
	fmt.Println("The struct space is ", unsafe.Sizeof(hello))
}

//=================  构造函数 ==================

func newPerson(name, city string, age int8) *Person{
	return &Person{
		name:name,
		city:city,
		age:age,
	}
}

//值接收者
func (p Person) Sing() {
	fmt.Printf("I'm %s,I can sing well\n", p.name)
}

//指针接收者
func (p *Person) SetAge(newAge int8){
	p.age = newAge
	fmt.Printf("Now, I'm %d years old,I can sing better than before\n", p.age)
}

//任意类型都可以添加方法
type myInt int

func (m myInt) SayHello(){
	fmt.Println("hello, I'm a int")
}

func RandFuncTest(){
	var m1 myInt
	m1.SayHello()
	m1 = 200
	fmt.Printf("I'm a new number, %#v", m1)
}

//结构体的匿名字段
type AoteMan struct {
	string
	int8
}

func NiMingTest(){
	aoteman1 := AoteMan{
		"LeiOu",
		30,
	}

	fmt.Printf("The value is %#v\n", aoteman1)
	fmt.Printf("The value 2 is %#v\n", aoteman1.int8)
}

//结构体的嵌套
type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

type User2 struct {
	Name string
	Gender string
	Address
}

func NestFuncTest(){
	newUser := User{
		Name:"FengGong",
		Gender:"Man",
		Address: Address{
			Province:"天津武清",
			City:"天津",
		},
	}

	fmt.Printf("New User is %#v\n", newUser)

}

//嵌套匿名结构体
func NestFuncTest2(){
	var user_2 User2
	user_2.Name = "Maji"
	user_2.Gender = "Man"
	user_2.Address.Province = "北京朝阳"
	user_2.Address.City = "北京"
	fmt.Printf("New User_2 is %#v\n", user_2)
}

//嵌套的结构体内的字段冲突
type Address2 struct {
	Province string
	City string
	CreateTime string
}

//Email的结构体
type Email struct {
	Account string
	CreateTime string
}

//User用户的结构体
type User3 struct {
	Name string
	Gender string
	Address2
	Email
}

func NestFuncTest3(){
	var user_3 User3
	user_3.Name = "Dazhangwei"
	user_3.Gender = "Man"
	user_3.Address2.Province = "朝阳"
	user_3.Address2.City = "北京"
	user_3.Address2.CreateTime = "20200427"
	user_3.Email.Account = "11223344@qq.com"
	user_3.Email.CreateTime = "20200506"
	fmt.Printf("New User_3 is %#v\n", user_3)
}

//=================  结构体的继承和可见性  ==================
type Animal struct {
	Name string
}

type Dog struct {
	Feet int8
	*Animal
}

func(d *Dog) wang(){
	fmt.Printf("I'm %s, I can wangwang...\n", d.Name)
}

func (a *Animal) run(){
	fmt.Printf("I'm %s, I can run...\n",a.Name)
}

func inheritTest(){
	//var samoye Dog
	samoye := &Dog{
		Feet:4,
		Animal: &Animal{
			Name:"hahha",
		},
	}

	samoye.run()
	samoye.wang()

}

//=================  结构体和Json之间的转换  ==================


func main() {
	//typeTest()

	//structTest()

	//structTest2()

	//structTest3()

	//structTest4()

	//structInit()

	//structInit2()

	//structInit3()

	//structInit4()

	//structInit5()

	//structEmpty()

	//调用构造函数
	//newP100 := newPerson("Huachenyu", "changsha", 30)
	////fmt.Printf("The P100 is %#v",newP100)
	//newP100.Sing()
	//newP100.SetAge(32)
	//fmt.Println(newP100)

	//RandFuncTest()

	//NiMingTest()

	//NestFuncTest()

	//NestFuncTest2()

	//NestFuncTest3()

	inheritTest()

	}
