package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	i int
	j int
}

type Direction int

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type State int

const (
	CLEAN = iota
	WEAKEND
	INFECTED
	FLAGGED
)

func (p *Point) move(direction Direction) {
	switch direction {
	case UP:
		p.i--
	case RIGHT:
		p.j++
	case DOWN:
		p.i++
	case LEFT:
		p.j--
	}
}

func (d Direction) turnRight() Direction {
	return (d + 1) % 4
}

func (d Direction) turnLeft() Direction {
	return (d + 3) % 4
}

func main() {
	infectionMap := map[Point]bool{}
	stateMap := map[Point]State{}
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
	i := 0
	middleJ := 0
	for scanner.Scan() {
		row := scanner.Text()
		j := 0
		for _, r := range row {
			if r == '#' {
				infectionMap[Point{i, j}] = true
				stateMap[Point{i, j}] = INFECTED
			} else if r == '.' {
				stateMap[Point{i, j}] = CLEAN
			}
			j++
		}
		i++
		middleJ = (j - 1) / 2
	}
	middleI := (i - 1) / 2
	currentPoint := Point{middleI, middleJ}
	currentPoint2 := Point(currentPoint)

	nSwitches := simulate(infectionMap, currentPoint, 10000)
	fmt.Println("Solution to part 1:", nSwitches)
	fmt.Println("Solution to part 2:", simulateStates(infectionMap, stateMap, currentPoint2, 10000000))
}

func simulate(infectionMap map[Point]bool, point Point, nBursts int) int {
	switchedInfected := 0
	var direction Direction = UP

	nodeInfected := false
	for i := 0; i < nBursts; i++ {
		nodeInfected = infectionMap[point]
		if nodeInfected {
			direction = direction.turnRight()
		} else {
			direction = direction.turnLeft()
			switchedInfected++
		}
		infectionMap[point] = !infectionMap[point]
		point.move(direction)
	}

	return switchedInfected
}

func simulateStates(infectionMap map[Point]bool, stateMap map[Point]State, point Point, nBursts int) int {
	switchedInfected := 0
	var direction Direction = UP

	for i := 0; i < nBursts; i++ {
		state := stateMap[point]
		switch state {
		case CLEAN:
			direction = direction.turnLeft()
			stateMap[point] = WEAKEND
		case WEAKEND:
			stateMap[point] = INFECTED
			switchedInfected++
		case INFECTED:
			direction = direction.turnRight()
			stateMap[point] = FLAGGED
		case FLAGGED:
			direction = direction.turnRight().turnRight()
			stateMap[point] = CLEAN
		}
		point.move(direction)
	}

	return switchedInfected
}
