package service

import (
	"../model"
)

//该CustomerService完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers   []model.Customer
	customerNum int //表示当前切片含有多少个客户，还可以作为新客户的id+1
}

//返回*CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 3
	customer1 := model.NewCustomer(1, "张三", "男", 20, "13845852151", "zhangsan@qq.com")
	customer2 := model.NewCustomer(2, "李四", "男", 22, "15236954625", "lisi@sina.com")
	customer3 := model.NewCustomer(3, "王二麻", "女", 18, "18745262514", "wangermazi@aliyun.com")

	customerService.customers = append(customerService.customers, customer1)
	customerService.customers = append(customerService.customers, customer2)
	customerService.customers = append(customerService.customers, customer3)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
