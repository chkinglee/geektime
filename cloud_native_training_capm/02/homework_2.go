package _2

import (
	"fmt"
	"math/rand"
	"time"
)

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
		time.Sleep(1 * time.Second)
		n := <-ch
		fmt.Println("receiving: ", n)
	}

}

