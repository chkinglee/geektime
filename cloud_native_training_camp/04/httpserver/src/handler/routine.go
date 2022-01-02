// Package handler
// @Author      : lilinzhen
// @Time        : 2022/1/2 10:34:08
// @Description :
package handler

import "net/http"

func RegistryHttpRoutine()  {
	http.HandleFunc("/healthz", CommonHandler(Healthz))
	http.HandleFunc("/a", CommonHandler(A))

}
