package main

import (
	"aoc/utils"
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

type Monkey struct {
	startingItems  []int
	operator       func(int) int
	test           func(int) int
	itemsInspected int
	part2          int
}

const rounds = 10000

func main() {
	defer utils.MeasureTime(time.Now())
	monkeys := []Monkey{
		{
			startingItems: []int{75, 75, 98, 97, 79, 97, 64},
			operator:      func(i int) int { return i * 13 },
			test: func(i int) int {
				if i%19 == 0 {
					return 2
				}
				return 7
			},
			part2: 19,
		},
		{
			startingItems: []int{50, 99, 80, 84, 65, 95},
			operator:      func(i int) int { return i + 2 },
			test: func(i int) int {
				if i%3 == 0 {
					return 4
				}
				return 5
			},
			part2: 3,
		},
		{
			startingItems: []int{96, 74, 68, 96, 56, 71, 75, 53},
			operator:      func(i int) int { return i + 1 },
			test: func(i int) int {
				if i%11 == 0 {
					return 7
				}
				return 3
			},
			part2: 11,
		},
		{
			startingItems: []int{83, 96, 86, 58, 92},
			operator:      func(i int) int { return i + 8 },
			test: func(i int) int {
				if i%17 == 0 {
					return 6
				}
				return 1
			},
			part2: 17,
		},
		{
			startingItems: []int{99},
			operator:      func(i int) int { return i * i },
			test: func(i int) int {
				if i%5 == 0 {
					return 0
				}
				return 5
			},
			part2: 5,
		},
		{
			startingItems: []int{60, 54, 83},
			operator:      func(i int) int { return i + 4 },
			test: func(i int) int {
				if i%2 == 0 {
					return 2
				}
				return 0
			},
			part2: 2,
		},
		{
			startingItems: []int{77, 67},
			operator:      func(i int) int { return i * 17 },
			test: func(i int) int {
				if i%13 == 0 {
					return 4
				}
				return 1
			},
			part2: 13,
		},
		{
			startingItems: []int{95, 65, 58, 76},
			operator:      func(i int) int { return i + 5 },
			test: func(i int) int {
				if i%7 == 0 {
					return 3
				}
				return 6
			},
			part2: 7,
		},
	}

	// part 2 module worry coff
	var a int = 1
	for _, v := range monkeys {
		a *= v.part2
	}

	for i := 0; i < rounds; i++ {
		for j, monka := range monkeys {
			for _, item := range monka.startingItems {
				monkeys[j].itemsInspected++
				// Part 1
				// worry := monka.operator(item) / 3
				// Part 2
				worry := monka.operator(item) % a
				next := monka.test(worry)
				monkeys[next].startingItems = append(monkeys[next].startingItems, worry)
			}
			monkeys[j].startingItems = []int{}
		}
	}
	slices.SortFunc(monkeys, func(a, b Monkey) bool {
		return a.itemsInspected > b.itemsInspected
	})

	fmt.Println("Answer:", monkeys[0].itemsInspected*monkeys[1].itemsInspected)
}
