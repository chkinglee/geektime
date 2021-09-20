package _1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	t.Log("test If")
	If(50)
	If(10)
	If(20)
}

func TestIfSimple(t *testing.T) {
	t.Log("test IfSimple")
	IfSimple(100)
	IfSimple(200)
}

func TestSwitch(t *testing.T) {
	Switch(1)
	Switch(2)
	Switch(3)
	Switch(4)
}

func TestFor(t *testing.T) {
	assert.Equal(t, For(100), 5050)
}

func TestForSimple(t *testing.T) {
	assert.Equal(t, ForSimple(200), 256)
}

func TestForForever(t *testing.T) {
	assert.Equal(t, ForForever(200), 201)
}

func TestForRangeString(t *testing.T) {
	ForRangeString("abcd")
}

func TestForRangeMap(t *testing.T) {
	ForRangeMap(map[string]interface{}{
		"a": "a",
		"b": 10,
		"c": 100.11,
		"d": true,
	})
}

func TestForRangeArray(t *testing.T) {
	ForRangeArray([3]int{1, 2, 3})
}
