package main

import (
	"aoc/utils"
	"fmt"
	"strings"
	"time"
)

var lookupMap = func() map[rune]int {
	m := make(map[rune]int)
	count := 1
	for i := 'a'; i <= 'z'; i++ {
		m[rune(i)] = count
		count++
	}
	for i := 'A'; i <= 'Z'; i++ {
		m[rune(i)] = count
		count++
	}
	return m
}()

func Day3(input string) (int, int) {
	part1 := 0
	utils.ReadLineByLine(input, func(line string) error {
		h := len(line) / 2
		half1 := line[:h]
		half2 := line[h:]

		for _, r := range half1 {
			if strings.ContainsRune(half2, r) {
				part1 += lookupMap[r]
				break
			}
		}
		return nil
	})

	part2 := 0
	utils.ReadGroupedLines(input, 3, func(lines []string) error {
		if len(lines) != 3 {
			panic("not enough lines")
		}
		for _, r := range lines[0] {
			if strings.ContainsRune(lines[1], r) && strings.ContainsRune(lines[2], r) {
				part2 += lookupMap[r]
				break
			}
		}
		return nil
	})

	return part1, part2
}

func main() {
	defer utils.MeasureTime(time.Now())
	part1, part2 := Day3("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
