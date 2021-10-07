package _3

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

// sync.RWMutex 实现了读写分离，其特点是读加锁时不需要等待读锁解锁（但仍需要等待写锁解锁）
// 由此提高了读的性能，适合读多写少的场景
type syncMapRW struct {
	sm map[int]int
	sync.RWMutex
}

func (s *syncMapRW) Read(key int) (value int) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()
	value = s.sm[key]
	return
}

func (s *syncMapRW) Write(key, value int) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()
	s.sm[key] = value
}

func rwSeparation() {
	s := syncMapRW{
		sm:      map[int]int{},
		RWMutex: sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			value := rand.Int()
			fmt.Println("new value:" + strconv.Itoa(value))
			s.Write(1, value)
		}()

		go func() {
			fmt.Println(s.Read(1))
		}()
	}

}
