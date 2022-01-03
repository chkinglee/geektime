// Package metrics
// @Author      : lilinzhen
// @Time        : 2022/1/3 15:51:57
// @Description :
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "geektime_cloud_native_training_camp"
	subsystem = "httpserver"
)

var RequestsCost = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_cost_seconds",
		Help:      "request(ms) cost seconds",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(RequestsCost)
}
