package _2

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	dogStr := "{\"Name\":\"花花\", \"Gender\":1}"
	fmt.Println(unmarshal2Struct(dogStr))

	dog := dog{
		Name:   "黑黑",
		Gender: 2,
	}
	fmt.Println(marshal2String(dog))
}
