package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	part1 := Day10("input.txt")
	assert.Equal(t, 17840, part1)
}
