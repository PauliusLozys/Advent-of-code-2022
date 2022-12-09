package main

import (
	"aoc/utils"
	"fmt"
	"time"
)

type Point struct {
	x int
	y int
}

func main() {
	defer utils.MeasureTime(time.Now())
	part1, part2 := Day9("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func Day9(input string) (int, int) {
	visitedPart1 := make(map[Point]struct{})
	visitedPart2 := make(map[Point]struct{})

	lines := utils.ReadAndSplit(input, "\n")

	head := Point{
		x: 0,
		y: 0,
	}
	tails := []Point{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	var direction string
	var steps int
	for _, line := range lines {
		_, err := fmt.Sscanf(line, "%s %d", &direction, &steps)
		utils.PanicOnErr(err)

		for i := 0; i < steps; i++ {
			switch direction {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}

			for j := 0; j < len(tails); j++ {
				if j == 0 {
					tails[j] = calculateNewTailPosition(tails[j], head)
				} else {
					tails[j] = calculateNewTailPosition(tails[j], tails[j-1])
				}
			}

			visitedPart1[tails[0]] = struct{}{}
			visitedPart2[tails[len(tails)-1]] = struct{}{}
		}
	}

	return len(visitedPart1), len(visitedPart2)
}

func calculateNewTailPosition(t, h Point) Point {
	if headNextToIt(t, h) {
		return t
	}
	if t.x == h.x {
		if t.y < h.y {
			t.y++
		} else {
			t.y--
		}
		return t
	}
	if t.y == h.y {
		if t.x < h.x {
			t.x++
		} else {
			t.x--
		}
		return t
	}
	if t.x < h.x {
		if t.y < h.y {
			t.x++
			t.y++
		} else {
			t.x++
			t.y--
		}
		return t
	}
	if t.x > h.x {
		if t.y < h.y {
			t.x--
			t.y++
		} else {
			t.x--
			t.y--
		}
		return t
	}
	panic("unreachable")
}

func headNextToIt(t, h Point) bool {
	return (t.x == h.x && (t.y+1 == h.y || t.y-1 == h.y)) ||
		(t.y == h.y && (t.x+1 == h.x || t.x-1 == h.x)) ||
		(t.x+1 == h.x && t.y+1 == h.y) ||
		(t.x-1 == h.x && t.y-1 == h.y) ||
		(t.x+1 == h.x && t.y-1 == h.y) ||
		(t.x-1 == h.x && t.y+1 == h.y) ||
		(t == h)
}
