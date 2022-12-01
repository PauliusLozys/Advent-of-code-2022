package utils

import (
	"io"
	"os"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
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
