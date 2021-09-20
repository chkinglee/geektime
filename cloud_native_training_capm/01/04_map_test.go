package _1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	myMap := Map()
	fmt.Println(Map())
	_, exist := myMap["string"]
	assert.True(t, exist)
	_, exist = myMap["string1"]
	assert.False(t, exist)
}
