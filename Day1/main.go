package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"golang.org/x/exp/maps"
)

func main() {
	f, _ := os.Open("input.txt")

	scan := bufio.NewScanner(f)
	elfs := make(map[int]int)
	elfIndex := 0
	for scan.Scan() {
		t := scan.Text()
		if t == "" {
			elfIndex++
		}
		nr, _ := strconv.Atoi(t)
		elfs[elfIndex] += nr
	}

	vals := maps.Values(elfs)
	sort.Ints(vals)

	fmt.Println("Part 1", vals[len(vals)-1])
	fmt.Println("Part 2", vals[len(vals)-1]+vals[len(vals)-2]+vals[len(vals)-3])
}
