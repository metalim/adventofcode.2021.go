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
		var digits [][]int
		for _, line := range lines {
			for i, v := range line {
				if len(digits) <= i {
					digits = append(digits, []int{0, 0})
				}
				switch v {
				case '0':
					digits[i][0]++
				case '1':
					digits[i][1]++
				default:
					panic("invalid digit")
				}
			}
		}
		var gamma, epsilon int
		for _, freq := range digits {
			var d int
			if freq[0] < freq[1] {
				d = 1
			}
			gamma = gamma*2 + d
			epsilon = epsilon*2 + 1 - d
		}
		fmt.Printf("Part 1: %d\n", gamma*epsilon)
	}

	// Part 2
	{
		// fmt.Printf("Part 2: %d\n", ints[0])
	}
}
