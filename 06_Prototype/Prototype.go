package main

import "fmt"

/*
   模式特点：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。
   程序实例：从简历原型，生成新的简历，简历类Resume提供Clone()方法。

   简历代码 原型模式 深复制

 */

type PersonalInfo struct {
	Name 		string
	Sex  		string
	Age  		string
}


 type Resume struct {
	 PersonalInfo
	 WorkPrice
 }

 //个人经历
 func (this *Resume)SetPersonalInfo(name,sex, age string){
 	this.Name = name
 	this.Sex = sex
 	this.Age = age
 }

 //设置工作经历
 type WorkPrice struct {
	 TimeArea	string
	 Company     string
 }


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
	aa := &Resume{}
	aa.SetPersonalInfo("小菜","男", "24")
	aa.SetWorkExperice("2016-2020","XXXX公司")
	bb := aa.Clone()
	bb.SetWorkExperice("2012-2016", "YYY1企业")
	bb.SetWorkExperice("2012-2016", "YYY2企业")
	cc := aa.Clone()
	cc.SetPersonalInfo("小菜","男", "26")
	aa.DisPlay()
	bb.DisPlay()
	cc.DisPlay()

}
