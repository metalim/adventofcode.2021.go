package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func catch(err error, args ...interface{}) {
	if err == nil {
		return
	}
	if len(args) > 0 {
		panic(errors.Wrap(err, fmt.Sprintf(args[0].(string), args[1:]...)))
	}
	panic(err)
}

func Input() string {
	file := "input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	input, err := os.ReadFile(file)
	catch(err)
	return strings.TrimSpace(string(input))
}

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	lines := Lines(Input())

	// Part 1
	{
		var total int
		for _, line := range lines {
			parts := strings.Split(line, " | ")
			digits := strings.Split(parts[1], " ")
			for _, digit := range digits {
				switch len(digit) {
				case 2, 3, 4, 7: // for digits 1, 7, 4 and 8
					total++
				}
			}
		}
		fmt.Printf("Part 1: %d\n", total)
	}

	digitSegments := []string{
		"abcefg",  // 0 -> 6
		"cf",      // 1 -> 2
		"acdeg",   // 2 -> 5
		"acdfg",   // 3 -> 5
		"bcdf",    // 4 -> 4
		"abdfg",   // 5 -> 5
		"abdefg",  // 6 -> 6
		"acf",     // 7 -> 2
		"abcdefg", // 8 -> 7
		"abcdfg",  // 9 -> 6
	}

	lengthDigits := map[int][]int{
		2: {1},
		3: {7},
		4: {4},
		5: {2, 3, 5},
		6: {0, 6, 9},
		7: {8},
	}
	segmentDigits := [][]int{
		{0, 2, 3, 5, 6, 7, 8, 9},
	}

	// Part 2
	{
		_ = lengthDigits
		_ = segmentDigits
		_ = digitSegments
		// fmt.Printf("Part 2: %d\n", ints[0])
	}
}
