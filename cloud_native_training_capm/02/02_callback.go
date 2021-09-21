package _2

func increase(a, b int) int {
	return a + b
}

func decrease(a, b int) int {
	return a - b
}

func DoOperation(c int, f func(a int, b int) int) int {
	return f(c, 1)
}