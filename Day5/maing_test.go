package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	part1, part2 := Day5("input.txt")
	assert.Equal(t, "BZLVHBWQF", part1)
	assert.Equal(t, "TDGJQTZSL", part2)
}
