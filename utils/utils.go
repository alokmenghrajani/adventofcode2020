package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Readfile(day int) []string {
	filename := fmt.Sprintf("day%02d/input.txt", day)
	file, err := os.Open(filename)
	PanicOnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, _ := ioutil.ReadAll(reader)

	lines := strings.Split(string(contents), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func StringsToInts(input []string) []int {
	numbers := make([]int, len(input))
	for i, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			n = 0
		}
		numbers[i] = n
	}
	return numbers
}
