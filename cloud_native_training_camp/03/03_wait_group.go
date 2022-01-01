package _3

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

// 在03/02_rw_mutex的例子中容易发现
// 所有的线程并没有运行完毕，主线程就已经结束了
// 在之前02/08_channel 以及 02/homework_2 中，使用channel作为等待所有线程运行完毕的通信工具
// 不免有些繁琐
// sync.WaitGroup的特点是支持等待一组goroutine返回，而不需要使用channel

func rwSeparationWaitGroup() {
	s := syncMapRW{
		sm:      map[int]int{},
		RWMutex: sync.RWMutex{},
	}
	rWg := sync.WaitGroup{} // 等待写
	wWg := sync.WaitGroup{} // 等待读
	for i := 0; i < 10; i++ {
		wWg.Add(1)
		go func() {
			defer wWg.Done()
			value := rand.Int()
			fmt.Println("new value:" + strconv.Itoa(value))
			s.Write(1, value)
		}()

		rWg.Add(1)
		go func() {
			defer rWg.Done()
			fmt.Println(s.Read(1))
		}()
	}
	rWg.Wait()
	wWg.Wait()

}
