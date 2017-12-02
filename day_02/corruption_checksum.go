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
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		rowElements := strings.Fields(scanner.Text())
		min, max := getMinMax(rowElements)
		sum1 += max - min
		sum2 += getDivisionResult(rowElements)
	}
	fmt.Println("Solution of part 1:", sum1)
	fmt.Println("solution of part 2:", sum2)
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

func getDivisionResult(row []string) int {
	var numbers []int
	for _, elem := range row {
		number, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		numbers = append(numbers, number)
	}
	for i, number1 := range numbers {
		for _, number2 := range numbers[i+1:] {
			sm := number1
			bg := number2
			if bg < sm {
				sm, bg = bg, sm
			}
			if bg%sm == 0 {
				return bg / sm
			}
		}
	}
	return 0
}
