package main

import (
	"aoc/utils"
	"fmt"
)

const (
	rock     = "A"
	paper    = "B"
	scissors = "C"

	win  = "Z"
	draw = "Y"
	lose = "X"
)

var scores = map[string]int{
	rock:     1,
	paper:    2,
	scissors: 3,

	lose: 0,
	draw: 3,
	win:  6,
}

var outcomesToWin = map[string]string{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

func convertPart1(choice string) string {
	switch choice {
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	default:
		panic("cannot match to any choice")
	}
}

func outcome(opponent, me string) int {
	switch {
	case opponent == me: // draw
		return scores[draw]
	case outcomesToWin[me] == opponent:
		return scores[win]
	default:
		return scores[lose]
	}
}

func reverseOutcome(opponent, outcome string) int {
	switch outcome {
	case draw: // draw
		return scores[opponent]
	case lose: // lose
		if opponent == rock {
			return scores[scissors]
		}
		if opponent == paper {
			return scores[rock]
		}
		return scores[paper]
	case win: // win
		if opponent == rock {
			return scores[paper]
		}
		if opponent == paper {
			return scores[scissors]
		}
		return scores[rock]
	default:
		panic("unknown outcome")
	}
}

func Day2(input string) (int, int) {
	var (
		part1, part2     int
		column1, column2 string
	)

	err := utils.ReadLineByLine(input, func(line string) error {
		if _, err := fmt.Sscanf(line, "%s %s", &column1, &column2); err != nil {
			return err
		}

		part1 += (outcome(column1, convertPart1(column2)) + scores[convertPart1(column2)])
		part2 += (reverseOutcome(column1, column2) + scores[column2])
		return nil
	})
	utils.PanicOnErr(err)

	return part1, part2
}

func main() {
	part1, part2 := Day2("input.txt")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
