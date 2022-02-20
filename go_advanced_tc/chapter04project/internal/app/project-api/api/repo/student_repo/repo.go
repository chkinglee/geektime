// Package student_repo
// @Author      : lilinzhen
// @Time        : 2022/2/20 21:52:08
// @Description :
package student_repo

import (
	"chapter04project/pkg/repo/mysql"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// repo 持久层对象
type repo struct {
	dbRepo mysql.DbRepo
}

// Repo 持久层接口
type Repo interface {
	QueryAll() ([]*Student, error)
	QueryRow() (*Student, error)
}

func (d repo) QueryAll() ([]*Student, error) {
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

	var students []*Student
	for rows.Next() {
		var stu Student
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

func (d repo) QueryRow() (*Student, error) {
	sqlString := "select * from student limit 1"
	var stu Student
	err := d.dbRepo.GetDb().QueryRow(sqlString).Scan(&stu.Id, &stu.Name, &stu.Age)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &stu, nil
}

func New(dbRepo mysql.DbRepo) Repo {
	return &repo{
		dbRepo: dbRepo,
	}
}
