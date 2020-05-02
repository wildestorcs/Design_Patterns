package main

import "fmt"

/*
   简历代码初步实现
 */

 type Resume struct {
 	Name 		string
 	Sex  		string
 	Age  		string
 	TimeArea	string
 	Company     string
 }

 func NewResume(name string) *Resume{
 	return &Resume{Name:name}
 }

 //个人经历
 func (this *Resume)SetPersonalInfo(sex, age string){
 	this.Sex = sex
 	this.Age = age
 }

 //设置工作经历
 func (this *Resume)SetWorkExperice(timearea, company string){
 	this.TimeArea = timearea
 	this.Company = company
 }

 //显示
 func (this *Resume)DisPlay(){
 	fmt.Printf("\n%s  %s  %s\n",this.Name, this.Sex, this.Age)
 	fmt.Printf("工作经历: %s %s\n", this.TimeArea, this.Company)
 }

func main() {
	a := NewResume("大鸟")
	a.SetPersonalInfo("男","29")
	a.SetWorkExperice("1998-2000", "XXX公司")

	b := NewResume("大鸟")
	b.SetPersonalInfo("男","29")
	b.SetWorkExperice("1998-2000", "XXX公司")

	c := NewResume("大鸟")
	c.SetPersonalInfo("男","29")
	c.SetWorkExperice("1998-2000", "XXX公司")

	a.DisPlay()
	b.DisPlay()
	c.DisPlay()
}
