package main
import (
	"fmt"
	"../lib"
)

var userId int
var pwd int


func main(){
	fmt.Println()

	var key int
	var loop = true

	for{
		fmt.Println("-------------欢迎使用多人聊天室-------------")
		fmt.Println("\t\t\t1 登陆聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")

		fmt.Scanf("%d\n", &key)
		switch key{
			case 1:
				fmt.Println("登陆聊天室")
				loop = false
			case 2:
				fmt.Println("注册用户")
				loop = false
			case 3:
				fmt.Println("退出系统")
				loop = false
			default:
				fmt.Println("你的输入有误，请重新输入")
		}
		if loop == false{
			break
		}
	}
	if key == 1{
		//说明用户要登陆了
		fmt.Println("请输入用户的id号")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%d\n", &pwd)
		err := lib.Login(userId, pwd)
		if err != nil{
			fmt.Println("登陆失败")
		}else{
			fmt.Println("登陆成功")
		}
	}else if key == 2{

	}
}