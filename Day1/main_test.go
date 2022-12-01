package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	part1, part2 := Day1("input.txt")
	assert.Equal(t, 69289, part1)
	assert.Equal(t, 205615, part2)
}
