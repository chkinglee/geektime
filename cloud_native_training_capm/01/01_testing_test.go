package _1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrease(t *testing.T) {
	t.Log("Starting Testing")
	result := Increase(1, 2)
	assert.Equal(t, result, 3)
}

// assert.Equal(t, result, 4)
//=== RUN   TestIncrease
//01_testing_test.go:9: Starting Testing
//01_testing_test.go:11:
//Error Trace:    01_testing_test.go:11
//Error:          Not equal:
//expected: 3
//actual  : 4
//Test:           TestIncrease
//--- FAIL: TestIncrease (0.00s)
//FAIL
//FAIL    01/01   0.223s
//FAIL

// assert.Equal(t, result, 3)
//=== RUN   TestIncrease
//    01_testing_test.go:9: Starting Testing
//--- PASS: TestIncrease (0.00s)
//PASS
//ok      01/01   0.268s
