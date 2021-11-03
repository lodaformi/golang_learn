package main

import "fmt"

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

type stuManSys struct {
	stuInfo []student
}

func (s *stuManSys) printAllStudent() {
	for _, v := range s.stuInfo {
		fmt.Println(v.id, "---", v.name, "---", v.score)
	}
}
func (s *stuManSys) addStudent() {
	fmt.Print("请输入学生的id：")
	var id int64
	fmt.Scanln(&id)
	fmt.Print("请输入学生的name：")
	var name string
	fmt.Scanln(&name)
	fmt.Print("请输入学生的score：")
	var score float32
	fmt.Scanln(&score)

	tmpStu := newStudent(id, name, score)
	s.stuInfo = append(s.stuInfo, *tmpStu)
}

func (s *stuManSys) delStudent() {
	fmt.Print("请输入学生的id：")
	var id int64
	fmt.Scanln(&id)

	var index int64 = -1
	for ind, val := range s.stuInfo {
		if val.id == id {
			index = int64(ind)
			s.stuInfo = append(s.stuInfo[:index], s.stuInfo[(index+1):]...)
			fmt.Printf("删除id为%d的学生\n", id)
			return
		}
	}
	fmt.Println("没有此学生")
}

//对于切片中某个元素信息的修改没有带回到stuInfo中，问题出在哪里？？？
func (s *stuManSys) alterStudent() {
	fmt.Print("请输入学生的id：")
	var id int64
	fmt.Scanln(&id)
	fmt.Print("请输入学生的score：")
	var score float32
	fmt.Scanln(&score)

	var index int64 = -1
	for ind, val := range s.stuInfo {
		if val.id == id {
			//修改信息
			val.score = score
			fmt.Printf("修改id为%d学生的信息\n", id)

			//重新添加该元素，使新值生效
			s.stuInfo = append(s.stuInfo, val)

			//将该元素的旧值删除
			index = int64(ind)
			s.stuInfo = append(s.stuInfo[:index], s.stuInfo[(index+1):]...)
			return
		}
	}
	fmt.Println("查无此人")
}
