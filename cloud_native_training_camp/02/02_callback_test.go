package _2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoOperation(t *testing.T) {
	assert.Equal(t, DoOperation(5, increase), 6)
	assert.Equal(t, DoOperation(5, decrease), 4)
}
