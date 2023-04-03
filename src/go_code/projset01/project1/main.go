package main

import (
	"fmt"
)

func main() {
	key := ""    //接收用户的输入
	loop := true //跳出循环
	//账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//收支的明细
	details := "收支\t账户金额\t收支金额\t说明"

	//主菜单
	for {
		fmt.Println("\n--------------家庭收支记账软件--------------")

		fmt.Println("		1 收入明细")
		fmt.Println("		2 登记收入")
		fmt.Println("		3 登记支出")
		fmt.Println("		4 退出")

		fmt.Println("请输入1-4： ")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("--------------当前收支明细--------------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t %v\t  	%v\t  	 %v", balance, money, note)

		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t	%v\t  	%v\t  	 %v", balance, money, note)
		case "4":
			fmt.Println("确认退出 输入 y 或者 n：")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("输入错误 请重新输入")
			}
			if choice == "y" {
				loop = false
			}
			loop = false
		default:
			fmt.Println("输入错误请从新输入")
		}
		if !loop {
			break
		}
	}
	fmt.Println("退出家庭收支软件")
}
