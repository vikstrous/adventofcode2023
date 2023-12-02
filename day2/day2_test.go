package day2

import (
	"os"
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
	possible := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	input := readInput()
	total := 0
	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, ":")
		gameNumStr := strings.TrimPrefix(lineParts[0], "Game ")
		gameNum, err := strconv.ParseInt(gameNumStr, 10, 64)
		checkErr(t, err)

		setsStr := strings.Split(lineParts[1], ";")
		impossible := false
		for _, setStr := range setsStr {
			ballsStr := strings.Split(setStr, ",")
			for _, ballStr := range ballsStr {
				// ballStr looks like "3 blue", or "4 red"
				ballStr = strings.TrimSpace(ballStr)
				ballParts := strings.Split(ballStr, " ")
				ballNumStr := ballParts[0]
				ballColor := ballParts[1]
				ballNum, err := strconv.ParseInt(ballNumStr, 10, 64)
				checkErr(t, err)

				if possible[ballColor] < ballNum {
					impossible = true
					break
				}
			}
		}
		if !impossible {
			total += int(gameNum)
		}
	}
	t.Log(total)
}

func TestPart2(t *testing.T) {
	input := readInput()
	total := int64(0)
	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, ":")

		setsStr := strings.Split(lineParts[1], ";")
		fewestPossible := map[string]int64{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, setStr := range setsStr {
			ballsStr := strings.Split(setStr, ",")
			for _, ballStr := range ballsStr {
				// ballStr looks like "3 blue", or "4 red"
				ballStr = strings.TrimSpace(ballStr)
				ballParts := strings.Split(ballStr, " ")
				ballNumStr := ballParts[0]
				ballColor := ballParts[1]
				ballNum, err := strconv.ParseInt(ballNumStr, 10, 64)
				checkErr(t, err)

				if fewestPossible[ballColor] < ballNum {
					fewestPossible[ballColor] = ballNum
				}
			}
		}

		total += fewestPossible["red"] * fewestPossible["green"] * fewestPossible["blue"]
	}
	t.Log(total)
}
