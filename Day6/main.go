package main

import (
	"aoc/utils"
	"fmt"
	"time"
)

func checkIfUnique(str string) bool {
	set := make(map[rune]struct{}, len(str))
	for _, ch := range str {
		if _, ok := set[ch]; ok {
			return false
		}
		set[ch] = struct{}{}
	}
	return true
}

func getMessageStart(line string, startLenght int) int {
	for i := 0; i+startLenght < len(line); i++ {
		if checkIfUnique(line[i : i+startLenght]) {
			return i + startLenght
		}
	}
	return -1
}

func Day6(input string) (int, int) {
	line := utils.ReadAndSplit(input, "\n")[0]
	return getMessageStart(line, 4), getMessageStart(line, 14)
}

func main() {
	defer utils.MeasureTime(time.Now())
	part1, part2 := Day6("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
