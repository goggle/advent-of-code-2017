package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Solution to part 1:", calculateManhattenDistance(input))
		fmt.Println("Solution to part 2:", calculateFirstNumberBiggerThanInput(input))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}

// Position holds the row and column indices (i and j) of the position
// in the matrix.
type Position struct {
	i, j int
}

func calculateManhattenDistance(input int) int {
	if input < 1 {
		fmt.Println("Input must be >= 1")
		os.Exit(1)
	}
	if input == 1 {
		return 0
	}
	squareLength := 1
	for squareLength*squareLength < input {
		squareLength += 2
	}
	var centerPosition Position
	centerPosition.i = squareLength / 2
	centerPosition.j = squareLength / 2

	const (
		UP int = iota
		LEFT
		DOWN
		RIGHT
	)

	var currentPosition Position
	currentPosition.i = squareLength - 2
	currentPosition.j = squareLength - 1

	currentNumber := (squareLength-2)*(squareLength-2) + 1
	direction := UP

	for currentNumber != input {
		currentNumber++

		if direction == UP {
			currentPosition.i--
			if currentPosition.i == 0 {
				direction = LEFT
			}
		} else if direction == LEFT {
			currentPosition.j--
			if currentPosition.j == 0 {
				direction = DOWN
			}
		} else if direction == DOWN {
			currentPosition.i++
			if currentPosition.i == squareLength-1 {
				direction = RIGHT
			}
		} else {
			currentPosition.j++
		}
	}

	a := currentPosition.i - centerPosition.i
	b := currentPosition.j - centerPosition.j
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	return a + b
}

type entry struct {
	i, j  int
	value int
}

func calculateSumOfAdjacentNeighbours(pos Position, list []entry) int {
	sum := 0
	for _, en := range list {
		d1 := en.i - pos.i
		d2 := en.j - pos.j
		if d1 < 0 {
			d1 *= -1
		}
		if d2 < 0 {
			d2 *= -1
		}
		if d1 <= 1 && d2 <= 1 {
			sum += en.value
		}
	}
	return sum
}

func calculateFirstNumberBiggerThanInput(input int) int {
	entryList := []entry{entry{0, 0, 1}}
	currentPosition := Position{0, 0}
	const (
		RIGHT int = iota
		UP
		LEFT
		DOWN
	)
	direction := RIGHT
	steps := 1
	stepCounter := 0
	currentValue := 1

	for currentValue <= input {
		stepCounter++
		if direction == RIGHT {
			currentPosition.j++
			if stepCounter == steps {
				stepCounter = 0
				direction = UP
			}
		} else if direction == UP {
			currentPosition.i--
			if stepCounter == steps {
				stepCounter = 0
				direction = LEFT
				steps++
			}
		} else if direction == LEFT {
			currentPosition.j--
			if stepCounter == steps {
				stepCounter = 0
				direction = DOWN
			}
		} else if direction == DOWN {
			currentPosition.i++
			if stepCounter == steps {
				stepCounter = 0
				direction = RIGHT
				steps++
			}
		}
		currentValue = calculateSumOfAdjacentNeighbours(currentPosition, entryList)
		entryList = append(entryList, entry{currentPosition.i, currentPosition.j, currentValue})
	}

	return currentValue
}
