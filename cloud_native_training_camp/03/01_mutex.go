package _3

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

// 线程安全
// 线程1和线程2都加载内存中的变量a到线程缓存中，a=1
// 如果线程1将a的值修改为2，此时线程2读取缓存中的a，值还是1
// 此时就是线程不安全的
// 对对象加锁，可以保证线程安全

// 典型的例子是，多线程不能同时修改同一个map对象，除非加锁
// sync.Mutex的特点是，任何加锁都需要等待无锁的状态，从而避免了同时读写的场景
type syncMap struct {
	sm map[int]int
	sync.Mutex
}

func (s *syncMap) Read(key int) (value int) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	value = s.sm[key]
	return
}

func (s *syncMap) Write(key, value int) {
	s.Mutex.Lock()
	defer s.Unlock()
	s.sm[key] = value
}

func safeWrite() {
	s := syncMap{
		sm:    map[int]int{},
		Mutex: sync.Mutex{},
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
