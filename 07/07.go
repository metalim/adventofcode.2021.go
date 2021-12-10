package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

func Ints(lines []string) []int {
	var err error
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], err = strconv.Atoi(line)
		catch(err, "line %d", i)
	}
	return ints
}

func main() {
	ints := Ints(strings.Split((Input()), ","))

	// Part 1
	{
		right := len(ints)
		crabCounts := map[int]int{}
		var minFuel int
		for _, pos := range ints {
			crabCounts[pos]++
			minFuel += pos
		}

		positions := make([]int, 0, len(crabCounts))
		for pos := range crabCounts {
			positions = append(positions, pos)
		}
		sort.Ints(positions)
		minFuel -= positions[0] * len(ints)

		fmt.Printf("positions: %d, crabs: %d\n", len(positions), len(ints))
		var split, left int
		fuel := minFuel
		for x := positions[0] + 1; x <= positions[len(positions)-1]; x++ {
			if positions[split] < x {
				left += crabCounts[positions[split]]
				right -= crabCounts[positions[split]]
				split++
			}
			fuel += left - right
			if minFuel > fuel {
				minFuel = fuel
			}
		}

		fmt.Printf("Part 1: %d\n", minFuel)
	}

	// Part 2
	{
		minPos := ints[0]
		maxPos := ints[0]

		var minFuel int
		for _, pos := range ints {
			if minPos > pos {
				minPos = pos
			}
			if maxPos < pos {
				maxPos = pos
			}
			minFuel += getFuel(pos)
		}

		for x := minPos; x <= maxPos; x++ {
			var fuel int
			for _, pos := range ints {
				df := getFuel(pos - x)
				fuel += df
			}
			if minFuel > fuel {
				minFuel = fuel
			}
		}
		fmt.Printf("Part 2: %d\n", minFuel)
	}
}

func getFuel(dist int) int {
	if dist < 0 {
		dist = -dist
	}
	return dist * (dist + 1) / 2
}
