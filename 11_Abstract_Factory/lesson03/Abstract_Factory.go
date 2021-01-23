package main

import "fmt"

type User struct {
	ID int
	Name string
}

type Iuser interface {
	Insert(user *User)
	GetUser(id int)
}

type SqlserverUser struct {
	Iuser
}

func (s *SqlserverUser) Insert(u *User){
	fmt.Println("在SQL Server中给User表增加一条记录")
}

func (s *SqlserverUser) GetUser(id int){
	fmt.Println("在SQL Server中根据ID得到User表一条记录")
}


type AccessUser struct {
	Iuser
}


func (a *AccessUser) Insert(u *User){
	fmt.Println("在Access中给User表增加一条记录")
}

func (a *AccessUser) GetUser(id int){
	fmt.Println("在Access中根据ID得到User表一条记录")
}

type Department struct {

}

type IDepartment interface {
	Insert(department *Department)
	GetDepartment(id int)
}

type SqlserverDepartment struct {
	IDepartment
}

func (s *SqlserverDepartment) Insert(d *Department){
	fmt.Println("在SqlserverDepartment中插入一条数据 ...")
}

func (s *SqlserverDepartment) GetDepartment(id int){
	fmt.Println("在SqlserverDepartment中根据ID得到Department表一条记录 ...")
}

type AccessDepartment struct {

}

func (a *AccessDepartment) Insert(d *Department) {
	fmt.Println("在AccessDepartment中插入一条数据 ...")
}

func (a *AccessDepartment) GetDepartment(id int) {
	fmt.Println("在AccessDepartment中根据ID得到Department表一条记录 ...")
}


type IFactory interface {
	CreatUser() Iuser
	CreatDepartment() IDepartment
}

type SqlserverFactory struct {
	IFactory
}

func (s *SqlserverFactory) CreateUser() *SqlserverUser{
	fmt.Println("Sqlserver CreatUser...")
	return new(SqlserverUser)
}

func (s *SqlserverFactory) CreateDepartment() *SqlserverDepartment{
	fmt.Println("Sqlserver CreatDepartment...")
	return new(SqlserverDepartment)
}

type AccessFactory struct {
	IFactory
}

func (a *AccessFactory) CreateUser() *AccessUser{
	fmt.Println("Access CreateUser...")
	return new(AccessUser)
}

func (a *AccessFactory) CreateDepartment() *AccessDepartment{
	fmt.Println("Access CreateDepartment...")
	return new(AccessDepartment)
}

func main()  {
	var user1 = new(User)
	var department1 = new(Department)

	var factory = new(AccessFactory)
	var iu Iuser= factory.CreateUser()

	iu.Insert(user1)
	iu.GetUser(1)

	var id IDepartment = factory.CreateDepartment()
	id.Insert(department1)
	id.GetDepartment(1)


}
