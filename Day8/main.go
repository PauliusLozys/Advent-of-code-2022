package main

import (
	"aoc/utils"
	"fmt"
	"time"
)

func Day8(input string) (int, int) {
	lines := utils.ReadAndSplit(input, "\n")
	forest := make([][]int, len(lines))
	for i, l := range lines {
		for _, v := range l {
			forest[i] = append(forest[i], utils.MustParseInt(string(v)))
		}
	}

	invisibleCount := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			tree := forest[i][j]
			maxUp := tree - 1
			maxDown := tree - 1
			maxLeft := tree - 1
			maxRight := tree - 1

			//look up
			for h := i - 1; h >= 0; h-- {
				maxUp = utils.Max(forest[h][j], maxUp)
			}
			//look down
			for h := i + 1; h < len(lines); h++ {
				maxDown = utils.Max(forest[h][j], maxDown)
			}

			//look right
			for h := j + 1; h < len(lines[0]); h++ {
				maxRight = utils.Max(forest[i][h], maxRight)
			}
			//look left
			for h := j - 1; h >= 0; h-- {
				maxLeft = utils.Max(forest[i][h], maxLeft)
			}
			if maxUp >= tree && maxDown >= tree && maxLeft >= tree && maxRight >= tree {
				invisibleCount++
			}
		}
	}

	// part 2
	maxScenic := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			tree := forest[i][j]
			seenTreesUp := 0
			seenTreesDown := 0
			seenTreesLeft := 0
			seenTreesRight := 0

			//look up
			for h := i - 1; h >= 0; h-- {
				seenTreesUp++
				if forest[h][j] >= tree {
					break
				}
			}
			//look down
			for h := i + 1; h < len(lines); h++ {
				seenTreesDown++
				if forest[h][j] >= tree {
					break
				}
			}

			//look right
			for h := j + 1; h < len(lines[0]); h++ {
				seenTreesRight++
				if forest[i][h] >= tree {
					break
				}
			}
			//look left
			for h := j - 1; h >= 0; h-- {
				seenTreesLeft++
				if forest[i][h] >= tree {
					break
				}
			}
			maxScenic = utils.Max(maxScenic, seenTreesUp*seenTreesDown*seenTreesLeft*seenTreesRight)
		}
	}

	total := len(lines) * len(lines[0])
	return total - invisibleCount, maxScenic
}

func main() {
	defer utils.MeasureTime(time.Now())
	part1, part2 := Day8("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
