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
			display := strings.Split(parts[1], " ")
			for _, digit := range display {
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

	// Part 2
	{
		var sum int
		for _, line := range lines {
			// maps code rune to possible segments
			mappings := map[rune]map[rune]bool{}
			for _, char := range "abcdefg" {
				mappings[char] = map[rune]bool{}
				for _, char2 := range "abcdefg" {
					mappings[char][char2] = true
				}
			}

			/*
				1. all mappings are possible.
				2. for each code, z.
				code(runes) -> digits -> segments
			*/
			parts := strings.Split(line, " | ")
			codes := strings.Split(parts[0], " ")
			for _, code := range codes {
				runes := map[rune]bool{}
				for _, r := range code {
					runes[r] = true
				}

				possibleDigits := lengthDigits[len(code)]
				possibleSegments := map[rune]bool{}
				for _, digit := range possibleDigits {
					for _, segment := range digitSegments[digit] {
						possibleSegments[segment] = true
					}
				}

				// remove impossible mappings
				for char := range runes {
					for char2 := range mappings[char] {
						if !possibleSegments[char2] {
							delete(mappings[char], char2)
						}
					}
				}
			}
			for char, mapping := range mappings {
				charsOut := []string{}
				for char2 := range mapping {
					charsOut = append(charsOut, string(char2))
				}
				fmt.Printf("%c -> %s\n", char, strings.Join(charsOut, ""))
			}
			break
			sum += 0
		}
		fmt.Printf("Part 2: %d\n", sum)
	}
}

/*
g -> a

e -> cf
a -> cf

b -> bd
c -> bd

f -> eg
d -> eg


a   e   -> 1
a   e g -> 7
abc e   -> 4
 bc efg -> 5
ab  efg -> 3
ab d fg -> 2
abc efg -> 09
 bcdefg -> 6
a cdefg -> 09
abcdefg -> 8

*/
