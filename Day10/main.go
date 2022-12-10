package main

import (
	"aoc/utils"
	"fmt"
	"time"
)

func main() {
	defer utils.MeasureTime(time.Now())
	part1 := Day10("input.txt")
	fmt.Println("Part 1:", part1)
}

func Day10(input string) int {
	cycle := 0
	regX := 1
	part1 := 0
	fmt.Print("#")
	utils.ReadLineByLine(input, func(line string) error {
		for i := 0; i < 2; i++ {
			cycle++
			if cycle%40 == 20 {
				part1 += cycle * regX
			}
			if line == "noop" {
				drawPixel(cycle, regX)
				break
			} else if i == 1 {
				var val int
				_, err := fmt.Sscanf(line, "addx %d", &val)
				utils.PanicOnErr(err)
				regX += val
			}
			drawPixel(cycle, regX)
		}
		return nil
	})

	return part1
}

func drawPixel(index, regX int) {
	idx := index % 40
	if idx == 0 {
		fmt.Println()
	}
	if idx == regX || idx == regX+1 || idx == regX-1 {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}
