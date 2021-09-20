package _1

import "fmt"

func MakeAndNew() {
	s1 := new([]int)           // 返回指针地址
	s2 := append(*s1, 1, 2, 3) // 通过 * 获取指针地址对应的对象
	s3 := &s2                  // 通过& 获取对象的指针地址
	s4 := make([]int, 0)       // 返回一个对象
	s4 = append(*s3, 2, 3, 4)
	fmt.Println(s4)

}

func AppendItem(s *[]int, a int) []int {
	*s = append(*s, a)
	return append(*s, a)
}

func ModifyItem(s []int, index int, a int) []int {
	s[index] = a
	return s
}

func DeleteItem(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
