package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

const MaxInt = int(^uint(0) >> 1)

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Readfile(day int) string {
	filename := fmt.Sprintf("day%02d/input.txt", day)
	file, err := os.Open(filename)
	PanicOnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	PanicOnErr(err)

	return strings.TrimSuffix(string(contents), "\n")
}

func Grid(input []string) [][]byte {
	rows := len(input)

	r := make([][]byte, rows)
	for i := range r {
		r[i] = []byte(input[i])
	}

	return r
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

func ParseToStruct(re *regexp.Regexp, input string, target interface{}) {
	m := re.FindStringSubmatch(input)
	if m == nil {
		return
	}

	var useOffset bool

	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		var field reflect.Value
		if name == "" {
			// use offset
			if i == 1 {
				useOffset = true
			} else if !useOffset {
				panic("can't mix named and unnamed subexpressions")
			}
			field = reflect.ValueOf(target).Elem().Field(i - 1)
		} else {
			// use name
			if i == 1 {
				useOffset = false
			} else if useOffset {
				panic("can't mix named and unnamed subexpressions")
			}
			field = reflect.ValueOf(target).Elem().FieldByName(name)
		}
		if field.Kind() == reflect.String {
			field.SetString(m[i])
		} else if field.Kind() == reflect.Int {
			v, err := strconv.Atoi(m[i])
			PanicOnErr(err)
			field.SetInt(int64(v))
		} else if field.Kind() == reflect.Uint8 {
			if len(m[i]) != 1 {
				panic(fmt.Sprintf("expecting 1 char, got: %s", m[i]))
			}
			field.SetUint(uint64(m[i][0]))
		} else {
			panic(fmt.Sprintf("unknown kind: %s", field.Kind()))
		}
	}
}
