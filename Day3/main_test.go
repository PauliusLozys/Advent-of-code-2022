package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	part1, part2 := Day3("input.txt")
	assert.Equal(t, 7817, part1)
	assert.Equal(t, 2444, part2)
}
