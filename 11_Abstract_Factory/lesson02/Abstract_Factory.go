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

type IFactory interface {
	CreatUser() Iuser
}

type SqlserverFactory struct {
	IFactory
}

func (s *SqlserverFactory) CreateUser() *SqlserverUser{
	fmt.Println("Sqlserver CreateUser...")
	return new(SqlserverUser)
}

type AccessFactory struct {
	IFactory
}

func (a *AccessFactory) CreateUser() *AccessUser{
	fmt.Println("Access CreateUser...")
	return new(AccessUser)
}


func main()  {
	var user1 = new(User)
	var factory = new(SqlserverFactory)
	var iu Iuser= factory.CreateUser()

	iu.Insert(user1)
	iu.GetUser(1)

}
