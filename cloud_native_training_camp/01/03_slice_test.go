package _1

import (
	"fmt"
	"testing"
)

func TestMakeAndNew(t *testing.T) {
	MakeAndNew()
}

func TestAppendItem(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := AppendItem(&s, 4)
	fmt.Println(&s, s)
	fmt.Println(&s1, s1)
}

func TestModifyItem(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := ModifyItem(s, 1, 4)
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s1)
}

func TestDeleteItem(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := DeleteItem(s, 2)
	fmt.Printf("%p\n", s)
	fmt.Println(s, len(s), cap(s))
	fmt.Printf("%p\n", s1)
	fmt.Println(s1, len(s1), cap(s1))
}
