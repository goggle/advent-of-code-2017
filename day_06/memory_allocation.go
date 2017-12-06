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
	for scanner.Scan() {
		inputFields := strings.Fields(scanner.Text())
		input := []int{}
		for _, inputField := range inputFields {
			inp, err := strconv.Atoi(inputField)
			if err != nil {
				fmt.Fprintln(os.Stderr, "converting to integer: ", err)
			}
			input = append(input, inp)
		}
		if len(input) > 0 {
			nIterations, nLoops := countRedistributions(input)
			fmt.Println("Solution of part 1:", nIterations)
			fmt.Println("Solution of part 2:", nLoops)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

func countRedistributions(input []int) (int, int) {
	count := 0
	tmp := make([]int, len(input))
	copy(tmp, input)
	knownConfigurations := [][]int{tmp}

	for {
		redistribute(input)
		count++
		known, nLoops := isKnownConfiguration(input, knownConfigurations)
		if known {
			return count, nLoops
		}
		tmp = make([]int, len(input))
		copy(tmp, input)
		knownConfigurations = append(knownConfigurations, tmp)
	}
}

func isKnownConfiguration(input []int, knownConfigurations [][]int) (bool, int) {
	for index, config := range knownConfigurations {
		equal := true
		for i, val := range config {
			if val != input[i] {
				equal = false
				break
			}
		}
		if equal {
			return true, len(knownConfigurations) - index
		}
	}
	return false, 0
}

func getHighestIndex(input []int) int {
	if len(input) == 0 {
		return -1
	}
	highestValue := input[0]
	highestIndex := 0
	for i, val := range input {
		if val > highestValue {
			highestIndex = i
			highestValue = val
		}
	}
	return highestIndex
}

func getNextIndex(currentIndex int, length int) int {
	if currentIndex+1 < length {
		return currentIndex + 1
	}
	return 0
}

func redistribute(input []int) {
	startIndex := getHighestIndex(input)
	nBlocks := input[startIndex]
	nDividedBlocks := nBlocks / len(input)
	nRestBlocks := nBlocks % len(input)
	input[startIndex] = nDividedBlocks
	currentIndex := getNextIndex(startIndex, len(input))
	for currentIndex != startIndex {
		input[currentIndex] += nDividedBlocks
		if nRestBlocks > 0 {
			input[currentIndex]++
			nRestBlocks--
		}
		currentIndex = getNextIndex(currentIndex, len(input))
	}
}
