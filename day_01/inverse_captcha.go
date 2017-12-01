package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Solution to part 1:", InverseCaptchaPartOne(scanner.Text()))
		fmt.Println("Solution to part 2:", InverseCaptchaPartTwo(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

// InverseCaptchaPartOne takes an input string and returns the sum of the
// character values as described in the description of the first part of
// the challange.
func InverseCaptchaPartOne(input string) int64 {
	var sum int64

	firstRune, firstRuneSize := utf8.DecodeRuneInString(input)
	prevRune := firstRune

	var w int
	for i := firstRuneSize; i < len(input); i += w {
		currRune, width := utf8.DecodeRuneInString(input[i:])
		if prevRune == currRune {
			sum += int64(currRune) - '0'
		}
		w = width
		prevRune = currRune
	}

	lastRune, _ := utf8.DecodeLastRuneInString(input)
	if utf8.RuneCountInString(input) >= 2 && firstRune == lastRune {
		sum += int64(lastRune) - '0'
	}

	return sum
}

// InverseCaptchaPartTwo takes an input string and returns the sum of the
// character values as described in the description of the second part of
// the challange.
func InverseCaptchaPartTwo(input string) int64 {
	var sum int64

	runes := []rune(input)
	length := utf8.RuneCountInString(input)

	for i, r := range runes {
		j := (i + length/2) % length
		if r == runes[j] {
			sum += int64(r) - '0'
		}
	}

	return sum
}
