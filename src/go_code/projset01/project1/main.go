package main

import "fmt"

type FamilyAccount struct {
	key  string //接收用户的输入
	loop bool   //跳出循环
	//账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//收支的明细
	details string
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------当前收支明细--------------")
	fmt.Println(this.details)
}
func (this *FamilyAccount) Revenue() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t %v\t  	%v\t  	 %v", this.balance, this.money, this.note)
}
func (this *FamilyAccount) expenditures() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t	%v\t  	%v\t  	 %v", this.balance, this.money, this.note)

}
func (this *FamilyAccount) quit() {
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
		this.loop = false
	}
	this.loop = false
}
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		details: "收支\t账户金额\t收支金额\t说   明",
	}

}

// 主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------家庭收支记账软件--------------")

		fmt.Println("		1 收入明细")
		fmt.Println("		2 登记收入")
		fmt.Println("		3 登记支出")
		fmt.Println("		4 退出")

		fmt.Println("请输入1-4： ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.Revenue()
		case "3":
			this.expenditures()
		case "4":
			this.quit()
		default:
			fmt.Println("输入错误请从新输入")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出家庭收支软件")
}
