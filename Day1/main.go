package main

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func main() {
	utils.MeasureTime(time.Now())
	part1, part2 := Day1("input.txt")
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func Day1(inputFile string) (int, int) {
	elfs := make([]int, 0)
	for _, str := range utils.ReadAndSplit(inputFile, "\n\n") {
		total := 0
		for _, ss := range strings.Split(str, "\n") {
			i, err := strconv.Atoi(ss)
			utils.PanicOnErr(err)
			total += i
		}
		elfs = append(elfs, total)
	}

	sort.Ints(elfs)
	slices.SortFunc(elfs, func(a, b int) bool {
		return a > b
	})

	return utils.SumSlice(elfs[:1]), utils.SumSlice(elfs[:3])
}
