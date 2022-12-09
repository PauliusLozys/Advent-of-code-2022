package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	part1, part2 := Day9("input.txt")
	assert.Equal(t, 6067, part1)
	assert.Equal(t, 2471, part2)
}
