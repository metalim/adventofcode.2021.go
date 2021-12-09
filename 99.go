package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile(os.Args[1])
	noError(err)
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		fmt.Println(line)
		break
	}
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}
