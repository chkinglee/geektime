// Package student
// @Author      : lilinzhen
// @Time        : 2022/1/30 12:32:03
// @Description :
package student

import (
	"chapter02error/src/mysql"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// model 学生实体
type model struct {
	Id   int32  // 主键
	Name string // 姓名
	Age  int32  // 年龄
}

// dao 持久层对象
type dao struct {
	dbRepo mysql.DbRepo
}

// Dao 持久层接口
type Dao interface {
	QueryAll() ([]*model, error)
	QueryRow() (*model, error)
}

func (m *model) ToString() {
	if m == nil {
		fmt.Println("stu not found")
		return
	}
	fmt.Println(fmt.Sprintf("Id: %d, Name: %s, Age: %d", m.Id, m.Name, m.Age))
}

func (d dao) QueryAll() ([]*model, error) {
	sqlString := "select * from student"
	rows, err := d.dbRepo.GetDb().Query(sqlString)
	if err != nil {
		return nil, errors.Wrap(err, "query err")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(errors.Wrap(err, "rows close err"))
		}
	}(rows)

	var students []*model
	for rows.Next() {
		var stu model
		if err = rows.Scan(&stu.Id, &stu.Name, &stu.Age); err != nil {
			return nil, errors.Wrap(err, "scan err")
		}
		students = append(students, &stu)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "next err")
	}
	return students, nil
}

func (d dao) QueryRow() (*model, error) {
	sqlString := "select * from student limit 1"
	var stu model
	err := d.dbRepo.GetDb().QueryRow(sqlString).Scan(&stu.Id, &stu.Name, &stu.Age)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil :
		return nil, err
	}
	return &stu, nil
}

func New(dbRepo mysql.DbRepo) Dao {
	return &dao{
		dbRepo: dbRepo,
	}
}
