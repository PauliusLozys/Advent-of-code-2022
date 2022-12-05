package main

import (
	"aoc/utils"
	"fmt"
)

var stacks = [][]rune{
	{'Q', 'S', 'W', 'C', 'Z', 'V', 'F', 'T'},
	{'Q', 'R', 'B'},
	{'B', 'Z', 'T', 'Q', 'P', 'M', 'S'},
	{'D', 'V', 'F', 'R', 'Q', 'H'},
	{'J', 'G', 'L', 'D', 'B', 'S', 'T', 'P'},
	{'W', 'R', 'T', 'Z'},
	{'H', 'Q', 'M', 'N', 'S', 'F', 'R', 'J'},
	{'R', 'N', 'F', 'H', 'W'},
	{'J', 'Z', 'T', 'Q', 'P', 'R', 'B'},
}

var stacks2 = [][]rune{
	{'Q', 'S', 'W', 'C', 'Z', 'V', 'F', 'T'},
	{'Q', 'R', 'B'},
	{'B', 'Z', 'T', 'Q', 'P', 'M', 'S'},
	{'D', 'V', 'F', 'R', 'Q', 'H'},
	{'J', 'G', 'L', 'D', 'B', 'S', 'T', 'P'},
	{'W', 'R', 'T', 'Z'},
	{'H', 'Q', 'M', 'N', 'S', 'F', 'R', 'J'},
	{'R', 'N', 'F', 'H', 'W'},
	{'J', 'Z', 'T', 'Q', 'P', 'R', 'B'},
}

func Day5(input string) (string, string) {
	utils.ReadLineByLine(input, func(line string) error {
		var move, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		// index offset
		from--
		to--
		for i := 0; i < move; i++ {
			f := stacks[from][len(stacks[from])-1]
			stacks[to] = append(stacks[to], f)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
		return nil
	})

	part1 := ""
	for _, s := range stacks {
		part1 += string(s[len(s)-1])
	}
	fmt.Println()

	utils.ReadLineByLine(input, func(line string) error {
		var move, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
		// index offset
		from--
		to--
		f := stacks2[from][len(stacks2[from])-move:]
		stacks2[to] = append(stacks2[to], f...)
		stacks2[from] = stacks2[from][:len(stacks2[from])-move]
		return nil
	})

	part2 := ""
	for _, s := range stacks2 {
		part2 += string(s[len(s)-1])
	}
	return part1, part2
}

func main() {
	part1, part2 := Day5("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
