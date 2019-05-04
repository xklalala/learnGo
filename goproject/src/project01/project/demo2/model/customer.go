package model
import (
	"fmt"
)

//声明一个Customer结构体，表示客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//编写一个工厂模式，返回一个Customer的实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户的信息
func (this Customer) GetInfo() string{
	info := fmt.Sprintf("\t%-v\t|\t%-v\t|\t%-v\t|\t%-v\t|%-v\t|%-v", this.Id, this.Name, this.Age, this.Gender, this.Phone, this.Email)
	return info
}