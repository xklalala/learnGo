package main
import (
	"fmt"
)

func main(){
	key := ""
	
	loop := true//控制循环
	
	balance := 0.0 //定义余额

	money := 0.0 //每次收支的金额

	note := "" //收支说明

	details := "收支\t账户金额\t收支金额\t说    明"

	for{
		fmt.Println("\n---------------家庭收支记账软件----------------")
		fmt.Println("               1.收支明细")
		fmt.Println("               2.登记收入")
		fmt.Println("               3.登记支出")
		fmt.Println("               4.退出软件")

		fmt.Scanln(&key)

		switch key{
			case "1":
				fmt.Println("---------------当前收支明细记录----------------")
				fmt.Println(details)
			case "2":
				fmt.Println("本次收入金额...")
				fmt.Scanln(&money)
				balance += money
				fmt.Println("本次收入说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n收入\t  %v\t\t%v\t\t%v", balance, money, note)
			case "3":
				fmt.Println("本次支出金额...")
				fmt.Scanln(&money)
				if balance < money{
					fmt.Println("你的余额不足")
					break
				}else{
					balance -= money
				}
				fmt.Println("本次支出说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n支出\t  %v\t\t%v\t\t%v", balance, money, note)
			case "4":
				fmt.Println("你确定要退出吗？y/n")
				choice := ""
				for{
					fmt.Scanln(&choice)
					if choice == "y" || choice == "n"{
						break
					}
					fmt.Println("你的输入有误，请重新输入 y/n")
				}
				if choice == "y"{
					loop = false
				}
				
			default:
				fmt.Println("请输入正确的选项...")
		}
		if !loop{
			break
		}
	}
	fmt.Println("你退出了这个软件")
}