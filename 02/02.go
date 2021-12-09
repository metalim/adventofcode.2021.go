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

func Lines(input string) []string {
	return strings.Split(input, "\n")
}

func main() {
	lines := Lines(Input())

	// Part 1
	{
		var x, y int
		for i, line := range lines {
			fields := strings.Fields(line)
			n, err := strconv.Atoi(fields[1])
			catch(err, "line %d", i)

			switch fields[0] {
			case "forward":
				x += n
			case "down":
				y += n
			case "up":
				y -= n
			default:
				panic("unknown direction")
			}
		}
		fmt.Printf("Part 1: %d\n", x*y)
	}

	// Part 2
	{
		var x, y, aim int
		for i, line := range lines {
			fields := strings.Fields(line)
			n, err := strconv.Atoi(fields[1])
			catch(err, "line %d", i)

			switch fields[0] {
			case "forward":
				x += n
				y += aim * n
			case "down":
				aim += n
			case "up":
				aim -= n
			default:
				panic("unknown direction")
			}
		}
		fmt.Printf("Part 2: %d\n", x*y)

	}
}
