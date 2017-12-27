package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	down int = iota
	left
	up
	right
)

type IndexTuple struct {
	i int
	j int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tubes := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		tubes = append(tubes, runes)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
	}

	word, nSteps, _ := passTubes(tubes)
	fmt.Println("Solution to part 1:", word)
	fmt.Println("Solution to part 2:", nSteps)

}

func passTubes(tubes [][]rune) (string, int, error) {
	jEntry, err := findEntry(tubes)
	if err != nil {
		return "", 0, err
	}
	nSteps := 0
	current := IndexTuple{0, jEntry}
	direction := down
	letters := []rune{}
	for ok := true; ok; {
		i, j := current.i, current.j
		if tubes[i][j] != '|' && tubes[i][j] != '-' && tubes[i][j] != '+' {
			letters = append(letters, tubes[i][j])
		}
		current, direction, ok = step(tubes, current, direction)
		nSteps++
	}
	word := string(letters)
	return word, nSteps, nil
}

func findEntry(tubes [][]rune) (int, error) {
	for i := 0; i < len(tubes[0]); i++ {
		if tubes[0][i] == '|' {
			return i, nil
		}
	}
	return 0, errors.New("no entry point found")
}

func step(tubes [][]rune, current IndexTuple, direction int) (IndexTuple, int, bool) {
	directions := []int{down, left, up, right}
	opposites := []int{up, right, down, left}
	nextIndex := move(current, direction)
	if inRegion(nextIndex, tubes) {
		r := tubes[nextIndex.i][nextIndex.j]
		if r != ' ' {
			return nextIndex, direction, true
		}
	}
	for i, dir := range directions {
		if direction == dir || direction == opposites[i] {
			continue
		}
		nextIndex = move(current, dir)
		if inRegion(nextIndex, tubes) {
			r := tubes[nextIndex.i][nextIndex.j]
			if r != ' ' {
				return nextIndex, dir, true
			}
		}
	}
	return current, direction, false
}

func inRegion(it IndexTuple, t [][]rune) bool {
	iLen := len(t)
	if it.i >= 0 && it.i < iLen {
		jLen := len(t[it.i])
		if it.j >= 0 && it.j < jLen {
			return true
		}
	}
	return false
}

func move(it IndexTuple, direction int) IndexTuple {
	switch direction {
	case down:
		return IndexTuple{it.i + 1, it.j}
	case left:
		return IndexTuple{it.i, it.j - 1}
	case up:
		return IndexTuple{it.i - 1, it.j}
	case right:
		return IndexTuple{it.i, it.j + 1}
	}
	return IndexTuple{it.i, it.j}
}
