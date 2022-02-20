// Package student_service
// @Author      : lilinzhen
// @Time        : 2022/2/20 22:05:07
// @Description :
package student_service

import (
	"chapter04project/internal/app/project-api/api/repo/student_repo"
	"chapter04project/pkg/repo/mysql"
)

var _ Service = (*service)(nil)

type service struct {
	dbRepo mysql.DbRepo
}

func (s *service) GetOne() (*student_repo.Student, error) {
	repo := student_repo.New(s.dbRepo)
	info, err := repo.QueryRow()
	if err != nil {
		return nil, err
	}
	return info, nil
}

type Service interface {
	GetOne() (*student_repo.Student, error)
}

func New(db mysql.DbRepo) Service {
	return &service{dbRepo: db}
}
