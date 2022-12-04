package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	part1, part2 := Day4("input.txt")
	assert.Equal(t, 542, part1)
	assert.Equal(t, 900, part2)
}
