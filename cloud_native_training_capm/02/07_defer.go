package _2

import "fmt"

func Defer() {
	var sum int
	defer func() {
		fmt.Println("Done, sum=", sum)
	}()

	for i:=0;i<100;i++ {
		sum += i
	}


}
