package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := map[int][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputRow := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		re := regexp.MustCompile(`(\d+)\s*<->\s*(.+)`)
		match := re.FindStringSubmatch(inputRow)
		door, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting input to integer:", err)
		}
		fields := strings.Split(match[2], ",")
		fieldInts := []int{}
		for _, f := range fields {
			fi, err := strconv.Atoi(strings.Trim(f, " "))
			if err != nil {
				fmt.Fprintln(os.Stderr, "converting input to integer:", err)
			}
			fieldInts = append(fieldInts, fi)
		}
		input[door] = fieldInts
	}
	groupZero := generateGroup(input, 0)
	fmt.Println("Solution to part 1:", len(groupZero))

	groups := [][]int{}
	for key := range input {
		inGroup := false
		for _, group := range groups {
			if isElem(group, key) {
				inGroup = true
				break
			}
		}
		if !inGroup {
			groups = append(groups, generateGroup(input, key))
		}
	}

	fmt.Println("Solution to part 2:", len(groups))

}

func generateGroup(input map[int][]int, groupNumber int) []int {
	groupZero := []int{}
	stack := []int{groupNumber}
	var x int

	for len(stack) > 0 {
		x, stack = stack[0], stack[1:]
		if !isElem(groupZero, x) {
			groupZero = append(groupZero, x)
		}
		for _, y := range input[x] {
			if !isElem(stack, y) && !isElem(groupZero, y) {
				stack = append(stack, y)
			}
		}
	}

	return groupZero
}

func isElem(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
