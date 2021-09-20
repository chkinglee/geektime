package _1

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Gender int
	Age    int
}

func GetName(s *Student) string {
	return s.Name
}

type StudentWithTag struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    int    `json:"age"`
}

func GetNameTag(s StudentWithTag) string {
	student := reflect.TypeOf(s)
	name := student.Field(0)
	fmt.Println(reflect.ValueOf(s).Field(1))
	return name.Tag.Get("json")
}
