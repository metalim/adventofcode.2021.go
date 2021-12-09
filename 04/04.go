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
func Blocks(input string) []string {
	return strings.Split(input, "\n\n")
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

type Board struct {
	fields           [5][5]int
	markedX, markedY [5]int
}

func NewBoard(block string) *Board {
	out := &Board{}
	for y, line := range Lines(block) {
		for x, number := range Ints(strings.Fields(line)) {
			out.fields[y][x] = number
		}
	}
	return out
}

func (b *Board) Mark(n int) {
	for y := range b.fields {
		for x, val := range b.fields[y] {
			if val == n {
				b.markedX[x]++
				b.markedY[y]++
				b.fields[y][x] = -1
			}
		}
	}
}

func (b *Board) Won() bool {
	for _, marked := range b.markedX {
		if marked == 5 {
			return true
		}
	}
	for _, marked := range b.markedY {
		if marked == 5 {
			return true
		}
	}
	return false
}

func (b *Board) Sum() int {
	var sum int
	for _, line := range b.fields {
		for _, val := range line {
			if val >= 0 {
				sum += val
			}
		}
	}
	return sum
}

func main() {
	blocks := Blocks(Input())
	numbers := Ints(strings.Split(blocks[0], ","))

	// Part 1
	{
		boards := make([]*Board, len(blocks)-1)
		for i, block := range blocks[1:] {
			boards[i] = NewBoard(block)
		}

		func() {
			for _, number := range numbers {
				for _, board := range boards {
					board.Mark(number)
					if board.Won() {
						fmt.Printf("Part 1: %d * %d = %d\n", board.Sum(), number, board.Sum()*number)
						return
					}
				}
			}
			fmt.Println("Part 1: No solution found")
		}()
	}

	// Part 2
	{
		boards := map[int]*Board{}
		for i, block := range blocks[1:] {
			boards[i] = NewBoard(block)
		}

		func() {
			for _, number := range numbers {
				for i, board := range boards {
					board.Mark(number)
					if board.Won() {
						delete(boards, i)
						if len(boards) == 0 {
							fmt.Printf("Part 2: %d * %d = %d\n", board.Sum(), number, board.Sum()*number)
							return
						}
					}
				}
			}
			fmt.Println("Part 2: No solution found")
		}()
	}
}
