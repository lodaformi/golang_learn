package main

import (
	"fmt"
	"os"
	"time"
)

var sms stuManSys

//学生管理系统（结构体版）
//实现学生的增删改查
func stuManSysFun(s *stuManSys) {
	//执行死循环，直到用户输入退出选项
	for {
		//输出菜单，并提示用户选择
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
				1：查看所有学生信息
				2：增加学生
				3：删除学生
				4：更改学生信息
				5：退出
			`)

		//提示用户选择
		fmt.Print("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)

		//根据选项做出判断
		switch choice {
		case 1:
			s.printAllStudent()
		case 2:
			s.addStudent()
		case 3:
			s.delStudent()
		case 4:
			s.alterStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入有误，请重新输入")
			time.Sleep(1000)
		}
	}
}

func main() {
	sms = stuManSys{make([]student, 0, 50)}
	stuManSysFun(&sms)
}
