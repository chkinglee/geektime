package _2

import (
	"fmt"
	"math/rand"
	"time"
)

func Channel1() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			//rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10000) // n will be between 0 and 10
			fmt.Println("putting: ", n)
			ch <- n
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()
	fmt.Println("hello from main")
	for v := range ch {
		fmt.Println("receiving: ", v)
	}
}

func Channel2() {
	ch := make(chan int, 10)
	singleChan := make(chan bool)
	go prod(ch)

	fmt.Println("hello from main")
	go rec(ch)
	<- singleChan
}

func prod(ch chan<- int) {
	for {
		time.Sleep(1 * time.Second)
		n := rand.Intn(10000) // n will be between 0 and 10
		fmt.Println("putting: ", n)
		ch <- n

	}
}

func rec(ch <-chan int) {
	for {
		time.Sleep(2 * time.Second)
		n := <-ch
		fmt.Println("receiving: ", n)
	}

}

// ChannelBase 在协程中，完成业务后，向channel发送信号，在主协程里接收该信号，以识别该让进程退出
func ChannelBase() {
	c := make(chan int)
	go func() {
		// do your business
		fmt.Println("do your business")
		n := rand.Intn(30) // n will be between 0 and 10
		time.Sleep(time.Duration(n) * time.Second)
		c <- 0
	}()
	fmt.Println("main go on.")
	loc, _ := time.LoadLocation("Local") //获取时区

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	fmt.Println("start:", startTime)
	select {
	case <-c:
		timeStr = time.Now().Format("2006-01-02 15:04:05")
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
		fmt.Println("end:", endTime)
	}

}
