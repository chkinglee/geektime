package _3

import (
	"fmt"
	"testing"
	"time"
)

func TestRwSeparationWaitGroup(t *testing.T) {
	loc, _ := time.LoadLocation("Local") //获取时区

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Println("start:", startTime)

	rwSeparationWaitGroup()

	timeStr = time.Now().Format("2006-01-02 15:04:05")
	endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Println("end:", endTime)

}
