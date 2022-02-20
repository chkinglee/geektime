// Package chapter04project
// @Author      : lilinzhen
// @Time        : 2022/2/20 19:51:24
// @Description :
package main

import (
	"chapter04project/internal/app/project-api/router"
	"chapter04project/internal/configs"
	"chapter04project/pkg/logger"
	"chapter04project/pkg/repo/mysql"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFilePath(configs.ProjectAccessLogFile),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	accessLogger.Info("Init logger success.")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
		configs.Get().MySQL.Read.User,
		configs.Get().MySQL.Read.Pass,
		configs.Get().MySQL.Read.Addr,
		configs.Get().MySQL.Read.Name,
	)
	dbRepo := mysql.New(dsn)
	defer func(repo mysql.DbRepo) {
		err := repo.DbClose()
		if err != nil {
			panic(err)
		}
	}(dbRepo)

	stop := make(chan struct{})
	done := make(chan error, 4)
	go signalListen(done)
	go stopListen(stop, done)
	eg := &errgroup.Group{}
	eg.Go(func() error {
		httpServer := router.NewHTTPServer(accessLogger, dbRepo, configs.Get().Project.Api.Addr)
		return httpServer.Start(stop, done)
	})
	eg.Go(func() error {
		return serverFork(done)
	})
	err = eg.Wait()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// 模拟一个server，最终返回一个error
func serverFork(done chan<- error) error {
	time.Sleep(time.Second * 10)
	err := fmt.Errorf("fork err")
	done <- err
	return err
}

func signalListen(done chan<- error) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-c:
		fmt.Println("get signal:", s)
	}
	signal.Stop(c)
	err := fmt.Errorf("get signal")
	done <- err
}

func stopListen(stop chan<- struct{}, done <-chan error) {
	select {
	case <-done:
		close(stop)
	}
}
