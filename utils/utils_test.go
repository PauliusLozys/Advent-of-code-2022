package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicOnErr(t *testing.T) {
	assert.Panics(t, func() {
		PanicOnErr(errors.New("test error"))
	})
	assert.NotPanics(t, func() {
		PanicOnErr(nil)
	})
}

func TestReadAndSplit(t *testing.T) {
	expected := []string{"10", "11", "12", "13"}
	actual := ReadAndSplit("test_file.txt", ",")
	assert.Equal(t, expected, actual)
}

func TestSumSlice(t *testing.T) {
	actual := SumSlice([]int{1, 2, 3, 10})
	assert.Equal(t, 16, actual)
}
