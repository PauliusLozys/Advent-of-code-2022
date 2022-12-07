package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	rootFolder := buildFS("input.txt")
	needToFreeUpSpace := updateSize - (maxFileSystemSize - rootFolder.Size())
	part1, part2 := walkDirectories(*rootFolder, needToFreeUpSpace, 0)
	assert.Equal(t, 1581595, part1)
	assert.Equal(t, 1544176, part2)
}
