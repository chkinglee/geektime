package _1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModifyStringSlice(t *testing.T) {
	srcSlice := []string{"I", "am", "stupid", "and", "weak"}
	dstSlice := []string{"I", "am", "smart", "and", "strong"}
	assert.Equal(t, ModifyStringSlice(srcSlice, dstSlice), dstSlice)
}
