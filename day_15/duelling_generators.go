package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type state struct {
	position   int
	increasing bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	startValues := []int{}
	for scanner.Scan() {
		inputRow := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		re := regexp.MustCompile(`Generator\s+([[:alpha:]])\s+starts\s+with\s+(\d+)`)
		match := re.FindStringSubmatch(inputRow)
		value, _ := strconv.Atoi(match[2])
		startValues = append(startValues, value)
	}

	values := []int{startValues[0], startValues[1]}
	multValues := []int{16807, 48271}
	divisor := 2147483647

	judgeCount := 0
	for i := 0; i < 40000000; i++ {
		updateValues(values, multValues, divisor)
		if checkMatch(values) {
			judgeCount++
		}
	}
	fmt.Println("Solution to part 1:", judgeCount)

	judgeCount = 0
	for i := 0; i < 5000000; i++ {
		updateValues2(startValues, multValues, divisor)
		if checkMatch(startValues) {
			judgeCount++
		}
	}
	fmt.Println("Solution to part 2:", judgeCount)

}

func updateValues(values []int, multValues []int, divisor int) {
	for i, v := range values {
		v *= multValues[i]
		values[i] = v % divisor
	}
}

func updateValues2(values []int, multValues []int, divisor int) {
	for {
		v := values[0]
		v *= multValues[0]
		values[0] = v % divisor
		if values[0]%4 == 0 {
			break
		}
	}
	for {
		v := values[1]
		v *= multValues[1]
		values[1] = v % divisor
		if values[1]%8 == 0 {
			break
		}
	}
}

func checkMatch(values []int) bool {
	var i uint
	for i = 0; i < 16; i++ {
		if (values[0]>>i)%2 != (values[1]>>i)%2 {
			return false
		}
	}
	return true
}
