package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func MeasureTime(t time.Time) {
	fmt.Println("Program took:", time.Since(t))
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

func ReadLineByLine(fileName string, fn func(string) error) error {
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
