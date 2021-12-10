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

func Ints(lines []string) []int {
	var err error
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], err = strconv.Atoi(line)
		catch(err, "line %d", i)
	}
	return ints
}

type Line struct {
	X1, Y1, X2, Y2 int
}

func NewLine(text string) *Line {
	points := strings.Split(text, " -> ")
	p1 := Ints(strings.Split(points[0], ","))
	p2 := Ints(strings.Split(points[1], ","))
	return &Line{p1[0], p1[1], p2[0], p2[1]}
}

func (l *Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.X1, l.Y1, l.X2, l.Y2)
}

func main() {
	textLines := Lines(Input())

	lines := make([]*Line, len(textLines))
	for i, line := range textLines {
		l := NewLine(line)
		lines[i] = l
	}

	// Part 1
	{
		var grid [1000][1000]int
		for _, line := range lines {
			var dx, dy int
			switch {
			case line.X1 < line.X2:
				dx = 1
			case line.X1 > line.X2:
				dx = -1
			}
			switch {
			case line.Y1 < line.Y2:
				dy = 1
			case line.Y1 > line.Y2:
				dy = -1
			}

			if dx != 0 && dy != 0 {
				continue
			}

			x := line.X1
			y := line.Y1
			for {
				grid[y][x]++
				if x == line.X2 && y == line.Y2 {
					break
				}
				x += dx
				y += dy
			}
		}

		var count int
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] > 1 {
					count++
				}
			}
		}
		fmt.Printf("Part 1: %d\n", count)
	}

	// Part 2
	{
		var grid [1000][1000]int
		for _, line := range lines {
			var dx, dy int
			switch {
			case line.X1 < line.X2:
				dx = 1
			case line.X1 > line.X2:
				dx = -1
			}
			switch {
			case line.Y1 < line.Y2:
				dy = 1
			case line.Y1 > line.Y2:
				dy = -1
			}
			x := line.X1
			y := line.Y1
			for {
				grid[y][x]++
				if x == line.X2 && y == line.Y2 {
					break
				}
				x += dx
				y += dy
			}
		}

		var count int
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] > 1 {
					count++
				}
			}
		}
		fmt.Printf("Part 2: %d\n", count)
	}
}
