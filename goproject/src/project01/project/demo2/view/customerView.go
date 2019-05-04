package main
import (
	"fmt"
	"../service"
)

type customerView struct{
	key		string	//接收用户输入
	loop	bool	//表示是否循环显示菜单
	CustomerService *service.CustomerService
}

//显示主菜单
func (this *customerView) mainMenu(){
	for{
		fmt.Println("------------------客户信息管理软件-----------------")
		fmt.Println("                  1.添加客户")
		fmt.Println("                  2.修改客户")
		fmt.Println("                  3.删除客户")
		fmt.Println("                  4.客户列表")
		fmt.Println("                  5.退    出")

		fmt.Scanln(&this.key)
		switch this.key{
			case "1":
				fmt.Println("添加客户")
			case "2":
				fmt.Println("修改客户")
			case "3":
				fmt.Println("删除客户")
			case "4":
				this.list()
			case "5":
				this.loop = false
			default:
				fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("你退出了软件的事使用")
}

//显示所有要显示的客户的信息
func (this *customerView) list(){
	customers := this.CustomerService.List()
	fmt.Println("\n************************************客户列表******************************************************")
	fmt.Println("\t编号\t|\t姓名\t|\t性别\t|\t年龄\t|\t电话\t|\t邮箱")
	for _, v := range customers{
		fmt.Println("-------------------------------------------------------------------------------------------------")
		fmt.Println(v.GetInfo())
	}
	fmt.Println("**********************************客户列表完成****************************************************\n")
}

func main(){
	//在主函数中创建一个customer实例
	customerView := customerView{
		key		: "",
		loop	: true,
	}
	customerView.CustomerService = service.NewCustomerService()
	customerView.mainMenu()
}