package _2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosure1(t *testing.T) {
	assert.Equal(t, Closure1(1, 2), 3)
}

func TestClosure2(t *testing.T) {
	assert.Equal(t, Closure2(1, 2), 3)
}

func TestClosure3(t *testing.T) {
	// 注意跟上面的区别
	assert.Equal(t, Closure3()(1, 2), 3)
}
