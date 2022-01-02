// Package handler
// @Author      : lilinzhen
// @Time        : 2022/1/2 10:33:08
// @Description :
package handler

import "net/http"

// Healthz 4. 当访问localhost/healthz时，应返回200
func Healthz(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusOK
	writer.WriteHeader(*statusCode)
	_, err := writer.Write([]byte("I am running\n"))
	if err != nil {
		return
	}
}

func A(writer http.ResponseWriter, request *http.Request, statusCode *int) {
	*statusCode = http.StatusAccepted
	writer.WriteHeader(*statusCode)
}
