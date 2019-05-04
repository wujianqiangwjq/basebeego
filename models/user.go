package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Name       string
	Paswd      string
	CreateTime int64
	LastTime   int64
	ShowName   string
	Aphosism   string
	Status     int
	Role       int
}

func AddUser(user *User) bool {
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return err == nil

}

func (user *User) DeleteUser() bool {
	o := orm.NewOrm()
	_, err := o.Delete(user)
	return err == nil
}

func (user *User) GetUser(col string) (*User, error) {
	o := orm.NewOrm()
	err := o.Read(user, col)
	return user, err

}

func (user *User) UpdateUser() bool {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err == nil

}

func (user *User) CheckUserExistByName() bool {

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	return qs.Filter("name", user.Name).Exist()
}
func GetUsers() ([]*User, error) {
	var users []*User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&users)
	return users, err

}
