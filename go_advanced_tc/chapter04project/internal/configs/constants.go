// Package configs
// @Author      : lilinzhen
// @Time        : 2022/2/20 20:26:31
// @Description :
package configs

import "time"

const (
	ProjectName          = "chapter04project"
	ProjectAccessLogFile = "./logs/" + ProjectName + "-access.log"
	ShutdownTimeout      = time.Second * 5
)
