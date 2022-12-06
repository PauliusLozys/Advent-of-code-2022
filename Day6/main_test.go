package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	part1, part2 := Day6("input.txt")
	assert.Equal(t, 1912, part1)
	assert.Equal(t, 2122, part2)
}
