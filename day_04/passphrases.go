package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	validPassphrases := 0
	validPassphrasesExtended := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		if checkValid(scanner.Text()) {
			validPassphrases++
		}
		if checkExtended(scanner.Text()) {
			validPassphrasesExtended++
		}
	}
	fmt.Println("Solution to part 1:", validPassphrases)
	fmt.Println("Solution to part 2:", validPassphrasesExtended)
}

func checkValid(input string) bool {
	words := strings.Fields(input)
	for i, word := range words {
		for _, compareWord := range words[i+1:] {
			if word == compareWord {
				return false
			}
		}
	}
	return true
}

func checkExtended(input string) bool {
	if !checkValid(input) {
		return false
	}
	words := strings.Fields(input)
	for i, word := range words {
		for _, compareWord := range words[i+1:] {
			if isAnagram(word, compareWord) {
				return false
			}
		}
	}
	return true
}

func isAnagram(word1, word2 string) bool {
	map1 := map[rune]int{}
	map2 := map[rune]int{}
	for _, r := range word1 {
		map1[r]++
	}
	for _, r := range word2 {
		map2[r]++
	}
	if reflect.DeepEqual(map1, map2) {
		return true
	}
	return false
}
