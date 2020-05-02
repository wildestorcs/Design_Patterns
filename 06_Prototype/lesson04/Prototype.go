package main

import "fmt"

/*
   简历代码 原型模式浅复制
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

 func (this *Resume)Clone() *Resume{
 	resume := *this
 	return &resume
 }


 //显示
 func (this *Resume)DisPlay(){
 	fmt.Printf("\n%s  %s  %s\n",this.Name, this.Sex, this.Age)
 	fmt.Printf("工作经历: %s %s\n", this.TimeArea, this.Company)
 }

func main() {
	aa := NewResume("小菜")
	aa.SetPersonalInfo("男", "24")
	aa.SetWorkExperice("2016-2020","XXXX公司")
	bb := aa.Clone()
	bb.SetWorkExperice("2012-2016", "YYY企业")
	cc := aa.Clone()
	cc.SetPersonalInfo("男", "26")
	aa.DisPlay()
	bb.DisPlay()
	cc.DisPlay()

}
