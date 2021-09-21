package _2

func Closure1(m, n int) int {
	// 将匿名函数赋值给变量
	x := func(a, b int) int { return a + b }
	return x(m, n)
}

func Closure2(m, n int) int {
	// 不声明，直接调用匿名函数
	return func(a, b int) int {
		return a + b
	}(m, n)
}

func Closure3() func(a, b int) int {
	// 作为函数返回值
	return func(a, b int) int {
		return a + b
	}
}
