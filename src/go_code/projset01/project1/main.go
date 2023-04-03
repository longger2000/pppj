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

		if !loop {
			break
		}
	}
	fmt.Println("退出家庭收支软件")
}
