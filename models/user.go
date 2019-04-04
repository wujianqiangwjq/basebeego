package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int `orm:"pk"`
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

func DeleteUser(user *User) bool {
	o := orm.NewOrm()
	_, err := o.Delete(user)
	return err == nil
}

func GetUser(user *User) (*User, error) {
	o := orm.NewOrm()
	err := o.Read(user)
	return user, err

}

func UpdateUser(user *User) bool {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err == nil

}

func GetUsers() ([]*User, error) {
	var users []*User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").All(&users)
	return users, err

}
