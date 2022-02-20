// Package student_handler
// @Author      : lilinzhen
// @Time        : 2022/2/20 21:50:38
// @Description :
package student_handler

import (
	"chapter04project/internal/app/project-api/api/service/student_service"
	"chapter04project/pkg/repo/mysql"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

var _ Handler = (*handler)(nil)

type handler struct {
	logger         *zap.Logger
	studentService student_service.Service
}

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	info, err := h.studentService.GetOne()
	if err != nil {
		h.logger.Error("err",zap.Error(err))
		fmt.Fprintln(w, "internal server error")
		return
	}

	if info == nil {
		fmt.Fprintln(w, "not found")
		return
	}

	fmt.Fprintln(w, info)
	h.logger.Info("get a student", zap.Any("student", info))
	return
}

func New(logger *zap.Logger, db mysql.DbRepo) Handler {
	return &handler{
		logger:         logger,
		studentService: student_service.New(db),
	}
}
