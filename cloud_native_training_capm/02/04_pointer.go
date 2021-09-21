package _2

type Student struct {
	Name string
	Age  int
}

// SetNameWithPointer 当需要修改外部变量的属性时，传指针地址
func SetNameWithPointer(t *Student) {
	t.Name = "chkinglee"
}

// SetNameWithStruct 当不需要修改外部变量的属性时，传对象
func SetNameWithStruct(t Student) {
	t.Name = "chkinglee"
}
