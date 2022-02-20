// Package router
// @Author      : lilinzhen
// @Time        : 2022/2/20 20:44:55
// @Description :
package router

import (
	"chapter04project/internal/app/project-api/api/handler/student_handler"
	"chapter04project/pkg/repo/mysql"
	"go.uber.org/zap"
	"net/http"
)

func RegistryAPIRouter(logger *zap.Logger, db mysql.DbRepo, mux *http.ServeMux) {
	studentHandler := student_handler.New(logger, db)
	mux.HandleFunc("/student", studentHandler.Get)
}
