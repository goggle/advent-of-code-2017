package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type hexPoint struct {
	x int
	y int
	z int
}

// func (p *hexPoint) euklideanDistanceSquare() float64 {
// 	return math.Abs(float64(p.x)) + math.Abs(float64(p.y))
// y := float64(p.y) * math.Sqrt(3.0) / 2.0
// x := float64(p.x) * 3.0 / 2.0
// return x*x + y*y
// }

func (p *hexPoint) distance() int {
	return int((math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z))) / 2)
}

// func (p *hexPoint) countSteps() int {
// 	steps := 0
// 	copyPoint := hexPoint{p.x, p.y}
// 	for copyPoint.x != 0 && copyPoint.y != 0 {
// 		neighbours := []hexPoint{
// 			hexPoint{copyPoint.x, copyPoint.y + 2},
// 			hexPoint{copyPoint.x, copyPoint.y - 2},
// 			hexPoint{copyPoint.x + 1, copyPoint.y + 1},
// 			hexPoint{copyPoint.x + 1, copyPoint.y - 1},
// 			hexPoint{copyPoint.x - 1, copyPoint.y + 1},
// 			hexPoint{copyPoint.x - 1, copyPoint.y - 1},
// 		}
// 		distances := []float64{}
// 		for _, neighbour := range neighbours {
// 			distances = append(distances, neighbour.euklideanDistanceSquare())
// 		}
// 		minIndex := getMinIndex(distances)
// 		copyPoint = neighbours[minIndex]
// 		fmt.Println(copyPoint)
// 		steps++
// 	}
// 	return steps
// }

// func getMinIndex(arr []float64) int {
// 	minIndex := 0
// 	for i, d := range arr {
// 		if d < arr[minIndex] {
// 			minIndex = i
// 		}
// 	}
// 	return minIndex
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rowElements := strings.Split(scanner.Text(), ",")
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		endPoint := hexPoint{0, 0, 0}
		distances := []int{}
		for _, instruction := range rowElements {
			// fmt.Println(instruction)
			switch instruction {
			case "n":
				endPoint.y++
				endPoint.z--
			case "s":
				endPoint.y--
				endPoint.z++
			case "ne":
				endPoint.x++
				endPoint.z--
			case "sw":
				endPoint.x--
				endPoint.z++
			case "nw":
				endPoint.x--
				endPoint.y++
			case "se":
				endPoint.x++
				endPoint.y--
			}
			distances = append(distances, endPoint.distance())
		}
		// fmt.Println(endPoint)
		fmt.Println("Solution to part 1:", endPoint.distance())
		maxDistance := 0
		for _, d := range distances {
			if d > maxDistance {
				maxDistance = d
			}
		}

		fmt.Println("Solution to part 2:", maxDistance)

	}
	// fmt.Println("Solution of part 1:", sum1)
	// fmt.Println("Solution of part 2:", sum2)
}
