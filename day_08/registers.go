package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registers := map[string]int{}
	lval := 0
	for scanner.Scan() {
		inputRow := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		re := regexp.MustCompile(`([[:alpha:]]+)\s+((?:inc|dec))\s+(-?\s*\d+)\s+if\s+([[:alpha:]]+)\s*((?:==|!=|<|>|<=|>=))\s*(-?\d+)`)
		match := re.FindStringSubmatch(inputRow)
		changeVal, _ := strconv.Atoi(match[3])
		compareVal, _ := strconv.Atoi(match[6])
		if checkCondition(registers, match[4], match[5], compareVal) {
			if match[2] == "inc" {
				lval = increase(registers, match[1], changeVal, lval)
			} else if match[2] == "dec" {
				lval = decrease(registers, match[1], changeVal, lval)
			}
		}
	}
	fmt.Println("Solution to part 1:", findLargestValue(registers))
	fmt.Println("Solution to part 2:", lval)
}

func findLargestValue(reg map[string]int) int {
	largestValue := 0
	for _, val := range reg {
		if val > largestValue {
			largestValue = val
		}
	}
	return largestValue
}

func checkCondition(reg map[string]int, name string, operator string, value int) bool {
	cVal := reg[name]
	if cVal == 0 {
		reg[name] = 0
	}
	switch operator {
	case ">":
		if cVal > value {
			return true
		}
		return false
	case "<":
		if cVal < value {
			return true
		}
		return false
	case "<=":
		if cVal <= value {
			return true
		}
		return false
	case ">=":
		if cVal >= value {
			return true
		}
		return false
	case "==":
		if cVal == value {
			return true
		}
		return false
	case "!=":
		if cVal != value {
			return true
		}
		return false
	}
	return false
}

func increase(reg map[string]int, name string, value int, lval int) int {
	reg[name] += value
	if reg[name] > lval {
		return reg[name]
	}
	return lval
}

func decrease(reg map[string]int, name string, value int, lval int) int {
	reg[name] -= value
	if reg[name] > lval {
		return reg[name]
	}
	return lval
}
