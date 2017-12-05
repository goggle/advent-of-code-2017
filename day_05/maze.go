package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "Input must be integers.")
			os.Exit(1)
		}
		input = append(input, number)
	}
	cpy := make([]int, len(input))
	copy(cpy, input)
	fmt.Println("Solution of part 1:", countSteps(cpy))
	copy(cpy, input)
	fmt.Println("Solution of part 2:", countStepsPartTwo(cpy))
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

func countSteps(input []int) int {
	nSteps := 0
	currentIndex := 0
	for {
		if currentIndex < 0 || currentIndex >= len(input) {
			break
		}
		currentInstruction := input[currentIndex]
		input[currentIndex]++
		currentIndex += currentInstruction
		nSteps++
	}

	return nSteps
}

func countStepsPartTwo(input []int) int {
	nSteps := 0
	currentIndex := 0
	for {
		if currentIndex < 0 || currentIndex >= len(input) {
			break
		}
		currentInstruction := input[currentIndex]
		if currentInstruction >= 3 {
			input[currentIndex]--
		} else {
			input[currentIndex]++
		}
		currentIndex += currentInstruction
		nSteps++
	}
	return nSteps
}
