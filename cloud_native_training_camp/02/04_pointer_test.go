package _2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetNameWithPointer(t *testing.T) {
	s := Student{
		Name: "a",
		Age: 21,
	}
	SetNameWithPointer(&s)
	assert.Equal(t, s.Name, "chkinglee")
}

func TestSetNameWithStruct(t *testing.T) {
	s := Student{
		Name: "a",
		Age: 21,
	}
	SetNameWithStruct(s)
	assert.Equal(t, s.Name, "a")
}
