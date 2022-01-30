// Package mysql
// @Author      : lilinzhen
// @Time        : 2022/1/30 12:30:43
// @Description :
package mysql

import (
	"database/sql"
)

type dbRepo struct {
	Db *sql.DB
}

func (d *dbRepo) GetDb() *sql.DB {
	return d.Db
}

func (d *dbRepo) DbClose() error {
	return d.Db.Close()
}

type DbRepo interface {
	GetDb() *sql.DB
	DbClose() error
}

func New(dsn string) DbRepo {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return &dbRepo{
		Db: db,
	}
}

