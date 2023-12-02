package day1

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func readInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}
func checkErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPart1(t *testing.T) {
	input := readInput()
	total := 0
	for _, line := range strings.Split(input, "\n") {
		numbers := regexp.MustCompile(`\d`).FindAllString(line, -1)
		numStr := numbers[0] + numbers[len(numbers)-1]
		num, err := strconv.ParseInt(numStr, 10, 64)
		checkErr(t, err)
		total += int(num)
	}
	t.Log(total)
}

var remap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var remapReverse = map[string]string{
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func TestPart2(t *testing.T) {
	input := readInput()
	total := 0
	for _, line := range strings.Split(input, "\n") {
		numbers := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`).FindAllString(line, -1)
		first := numbers[0]
		remapped, ok := remap[first]
		if ok {
			first = remapped
		}
		reverse := Reverse(line)
		numbersReverse := regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`).FindAllString(reverse, -1)
		last := numbersReverse[0]
		remapped, ok = remapReverse[last]
		if ok {
			last = remapped
		}

		numStr := first + last
		num, err := strconv.ParseInt(numStr, 10, 64)
		checkErr(t, err)
		total += int(num)
	}
	t.Log(total)
}
