package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	part1, part2 := Day2("input.txt")
	assert.Equal(t, 11666, part1)
	assert.Equal(t, 12767, part2)
}
