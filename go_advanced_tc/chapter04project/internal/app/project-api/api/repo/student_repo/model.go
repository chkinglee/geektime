// Package student_repo
// @Author      : lilinzhen
// @Time        : 2022/2/20 21:52:14
// @Description :
package student_repo

import "fmt"

// Student 学生实体
type Student struct {
	Id   int32  // 主键
	Name string // 姓名
	Age  int32  // 年龄
}

func (s *Student) ToString() {
	if s == nil {
		fmt.Println("stu not found")
		return
	}
	fmt.Println(fmt.Sprintf("Id: %d, Name: %s, Age: %d", s.Id, s.Name, s.Age))
}

