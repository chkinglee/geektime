package _1

import (
	"fmt"
	"testing"
)

func TestGetName(t *testing.T) {
	s := Student{
		Name:   "chkinglee",
		Gender: 1,
		Age:    12,
	}
	fmt.Println(GetName(&s))
}

func TestGetNameTag(t *testing.T) {
	swt := StudentWithTag{
		Name:   "chkinglee",
		Gender: 1,
		Age:    12,
	}
	fmt.Println(GetNameTag(swt))

}
