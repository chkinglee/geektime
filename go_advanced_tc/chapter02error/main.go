// Package chapter02error
// @Author      : lilinzhen
// @Time        : 2022/1/30 10:22:47
// @Description :
package main

import (
	"chapter02error/src/mysql"
	"chapter02error/src/student"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	dsn := "root:Root1234@tcp(127.0.0.1:3306)/go_advanced_tc?charset=utf8mb4"
	dbRepo := mysql.New(dsn)
	defer func(repo mysql.DbRepo) {
		err := repo.DbClose()
		if err != nil {
			panic(err)
		}
	}(dbRepo)
	fmt.Println("new db repo success")

	studentDao := student.New(dbRepo)
	students, err := studentDao.QueryAll()
	if err != nil {
		fmt.Printf("err: %+v",err)
		return
	}

	for _, stu := range students {
		stu.ToString()
	}

	stuOne, err := studentDao.QueryRow()
	if err != nil {
		fmt.Printf("err: %+v",err)
		return
	}
	stuOne.ToString()
	return
}
