package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		rowElements := strings.Fields(scanner.Text())
		min, max := getMinMax(rowElements)
		sum += max - min
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

func getMinMax(row []string) (int, int) {
	var min, max int
	first := true
	for _, elem := range row {
		number, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if first {
			first = false
			min = number
			max = number
			continue
		}
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}
	return min, max
}
