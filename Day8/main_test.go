package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	part1, part2 := Day8("input.txt")
	assert.Equal(t, 1676, part1)
	assert.Equal(t, 313200, part2)
}
