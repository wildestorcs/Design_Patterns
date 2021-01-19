package main

import "fmt"

type User struct {
	ID int
	Name string
}

func(u *User) get_id() int{
	return u.ID
}

func(u *User) set_id(id int){
	u.ID = id
}

func(u *User) get_name() string{
	return u.Name
}

func(u *User) set_name(name string){
	u.Name = name
}

type SqlserverUser struct {
}

func (s *SqlserverUser) Insert(u *User){
	fmt.Println("在SQL Server中给User表增加一条记录")
}

func (s *SqlserverUser) GetUser(id int){
	fmt.Println("在SQL Server中根据ID得到User表一条记录")
}




func main()  {
	var user1 = new(User)
	var sqlServerUser = new(SqlserverUser)

	sqlServerUser.Insert(user1)
	sqlServerUser.GetUser(1)


}
