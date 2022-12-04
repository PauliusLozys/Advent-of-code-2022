package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func containsOther(b1, e1, b2, e2 int) bool {
	return b1 <= b2 && e1 >= e2
}

func containsAtAll(b1, e1, b2, e2 int) bool {
	if b1 == b2 || b1 == e2 || e1 == e2 || b2 == e1 {
			return true
		}
	if b1 < b2 { // b1 chain starts first
		return b2 < e1
	}
	return b1 < e2
}

func splitPairs(line string) (int, int, int, int) {
	pairs := strings.FieldsFunc(line, func(r rune) bool {
		return r == ',' || r == '-'
	})
	b1, _ := strconv.Atoi(pairs[0])
	e1, _ := strconv.Atoi(pairs[1])
	b2, _ := strconv.Atoi(pairs[2])
	e2, _ := strconv.Atoi(pairs[3])

	return b1, e1, b2, e2
}

func Day4(input string) (int, int) {
	part1 := 0
	utils.ReadLineByLine(input, func(line string) error {
		b1, e1, b2, e2 := splitPairs(line)
		if containsOther(b1, e1, b2, e2) || containsOther(b2, e2, b1, e1) {
			part1++
		}
		return nil
	})

	part2 := 0
	utils.ReadLineByLine(input, func(line string) error {
		if containsAtAll(splitPairs(line)) {
			part2++
		}
		return nil
	})
	return part1, part2
}

func main() {
	defer utils.MeasureTime(time.Now())
	part1, part2 := Day4("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
