package main

import (
	"fmt"
	"os"
	"time"
)

//学生管理系统（函数版）
//实现学生的增删改查

type student struct {
	id    int64
	name  string
	score float32
}

func newStudent(id int64, name string, score float32) *student {
	return &student{
		id:    id,
		name:  name,
		score: score,
	}
}

var stuInfo []student

func printAllStudent() {
	for _, v := range stuInfo {
		fmt.Println(v.id, "---", v.name, "---", v.score)
	}
}
func addStudent(id int64, name string, score float32) {
	tmpStu := newStudent(id, name, score)
	stuInfo = append(stuInfo, *tmpStu)
}

func delStudent(id int64) {
	var index int64 = -1
	for ind, val := range stuInfo {
		if val.id == id {
			index = int64(ind)
			stuInfo = append(stuInfo[:index], stuInfo[(index+1):]...)
			fmt.Printf("删除id为%d的学生\n", id)
			break
		}
	}
}

//对于切片中某个元素信息的修改没有带回到stuInfo中，问题出在哪里？？？
func alterStudent(id int64, score float32) {
	var index int64 = -1
	for ind, val := range stuInfo {
		if val.id == id {
			//修改信息
			val.score = score
			fmt.Printf("修改id为%d学生的信息\n", id)

			//重新添加该元素，使新值生效
			stuInfo = append(stuInfo, val)

			//将该元素的旧值删除
			index = int64(ind)
			stuInfo = append(stuInfo[:index], stuInfo[(index+1):]...)
			break
		}
	}
}

func stuManSys() {
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
			printAllStudent()
		case 2:
			fmt.Print("请输入学生的id：")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("请输入学生的name：")
			var name string
			fmt.Scanln(&name)
			fmt.Print("请输入学生的score：")
			var score float32
			fmt.Scanln(&score)
			addStudent(id, name, score)
		case 3:
			fmt.Print("请输入学生的id：")
			var id int64
			fmt.Scanln(&id)
			delStudent(id)
		case 4:
			fmt.Print("请输入学生的id：")
			var id int64
			fmt.Scanln(&id)
			fmt.Print("请输入学生的score：")
			var score float32
			fmt.Scanln(&score)
			alterStudent(id, score)
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入有误，请重新输入")
			time.Sleep(1000)
		}
	}
}

func main() {
	stuInfo = make([]student, 0, 50)

	fmt.Println(stuInfo)
	stu1 := newStudent(123, "tom", 76.5)
	stuInfo = append(stuInfo, *stu1)
	fmt.Println(stuInfo)
	//此处修改stu1的信息，但是stuInfo中的信息没有发生改变
	stu1.score = 88.2
	//再次将stu1追加到stuInfo后，发现信息是做了修改的，但是stuInfo还是旧值
	//目前想到的办法是将原来的stu1删除，再追加更改后的stu1
	stuInfo = append(stuInfo, *stu1)
	fmt.Println(stuInfo)

	stuManSys()
}
