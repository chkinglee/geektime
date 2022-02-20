// Package router
// @Author      : lilinzhen
// @Time        : 2022/2/20 20:44:44
// @Description :
package router

import (
	"chapter04project/pkg/repo/mysql"
	"context"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type httpServer struct {
	logger  *zap.Logger
	server  http.Server
	handler *httpHandler
}

func (s *httpServer) Start(stop <-chan struct{}, done chan<- error) error {
	go func() {
		<-stop // goroutine会在此阻塞，直到stop内有消息，或stop被关闭
		s.logger.Info("receive a stop signal or channel(stop) is closed")

		s.handler.BeforeShutdown()

		s.server.Shutdown(context.Background()) // 在此主动关闭server，该goroutine会结束
	}()

	err := s.server.ListenAndServe()
	done <- err
	return err
}

func NewHTTPServer(logger *zap.Logger, db mysql.DbRepo, addr string) *httpServer {
	h := httpHandler{wg: sync.WaitGroup{}, logger: logger}
	mux := http.NewServeMux()
	mux.Handle("/", &h)

	RegistryAPIRouter(logger, db, mux)

	return &httpServer{
		logger: logger,
		server: http.Server{
			Addr:    addr,
			Handler: mux,
		},
		handler: &h,
	}
}
