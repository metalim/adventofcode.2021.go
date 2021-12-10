package main

import (
	"fmt"
	"os"
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
	ints := Ints(strings.Split(Input(), ","))

	// Part 1
	{
		counters := map[int]int{}
		for _, days := range ints {
			counters[days]++
		}

		counters = simulate(counters, 80)

		var total int
		for _, count := range counters {
			total += count
		}
		fmt.Printf("Part 1: %d\n", total)
	}

	// Part 2
	{
		counters := map[int]int{}
		for _, days := range ints {
			counters[days]++
		}

		counters = simulate(counters, 256)

		var total int
		for _, count := range counters {
			total += count
		}
		fmt.Printf("Part 2: %d\n", total)
	}
}

func simulate(counters map[int]int, daysTotal int) map[int]int {
	for i := 0; i < daysTotal; i++ {
		newCounters := map[int]int{}
		for days, count := range counters {
			if days == 0 {
				newCounters[6] += count
				newCounters[8] += count
				continue
			}
			newCounters[days-1] += count
		}
		counters = newCounters
	}
	return counters
}
