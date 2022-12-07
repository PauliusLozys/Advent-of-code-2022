package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MeasureTime(t time.Time) {
	fmt.Println("Program took:", time.Since(t))
}

func Max[K constraints.Ordered](a, b K) K {
	if a < b {
		return b
	}
	return a
}
func Min[K constraints.Ordered](a, b K) K {
	if a < b {
		return a
	}
	return b
}

func ReadAndSplit(fileName, splitPattern string) []string {
	f, err := os.Open(fileName)
	PanicOnErr(err)

	file, err := io.ReadAll(f)
	PanicOnErr(err)

	return strings.Split(string(file), splitPattern)
}

func SumSlice(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func ReadLineByLine(fileName string, fn func(line string) error) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		if err := fn(scan.Text()); err != nil {
			return err
		}
	}
	return nil
}

func ReadGroupedLines(fileName string, numOfLines int, fn func(lines []string) error) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(f)
	linesScanned := 0
	buffer := make([]string, numOfLines)
	for scan.Scan() {
		buffer[linesScanned] = scan.Text()
		linesScanned++

		if linesScanned == numOfLines {
			if err := fn(buffer); err != nil {
				return err
			}
			linesScanned = 0
		}
	}
	return nil
}

func MustParseInt(str string) int {
	v, err := strconv.Atoi(str)
	PanicOnErr(err)
	return v
}
