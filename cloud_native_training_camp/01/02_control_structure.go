package _1

import (
	"fmt"
	"strconv"
)

func If(a int) {
	if a < 10 {
		fmt.Println("a < 10")
	} else if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a > 10")
	}
}

func IfSimple(a int) {
	if b := a - 100; b < 100 {
		fmt.Println("b < 100")
	} else {
		fmt.Println("b >= 100")
	}
}

func Switch(a int) {
	switch a {
	case 1:
	case 2:
		fallthrough // 执行下一分支的func
	case 3:
		fmt.Println("a == 2or3")
	default:
		fmt.Println("not match")
	}
}

func For(a int) int {
	var sum int
	// 从0加到a
	for i := 0; i <= a; i++ {
		sum += i
		fmt.Println("sum=" + strconv.Itoa(sum))
	}
	return sum
}
func ForSimple(a int) int {
	var sum = 1
	// 比a大的最小2的幂值
	for ; sum < a; {
		sum += sum
		fmt.Println("sum=" + strconv.Itoa(sum))
	}
	return sum
}

func ForForever(a int) int {
	var sum int
	// 递加
	for {
		sum += 1
		fmt.Println("sum=" + strconv.Itoa(sum))
		if sum > a {
			break
		}
	}
	return sum
}

func ForRangeString(a string) {
	for index, c := range a {
		fmt.Println(index, string(c))
	}
}

func ForRangeMap(a map[string]interface{}) {
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func ForRangeArray(a [3]int) {
	for index, value := range a {
		fmt.Println(index, value)
	}
}
