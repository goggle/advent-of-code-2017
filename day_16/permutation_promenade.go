package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Exchange struct {
	p1 int
	p2 int
}

type Partners struct {
	r1 rune
	r2 rune
}

type Spin struct {
	n int
}

type DanceMove interface {
	Move([]rune)
}

func (e Exchange) Move(programs []rune) {
	exchange(programs, e.p1, e.p2)
}

func (p Partners) Move(programs []rune) {
	partners(programs, p.r1, p.r2)
}

func (s Spin) Move(programs []rune) {
	spin(programs, s.n)
}

func main() {
	programs := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	programsCopy := make([]rune, len(programs))
	danceMoves := []DanceMove{}
	copy(programsCopy, programs)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rowElements := strings.Split(scanner.Text(), ",")
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		// fmt.Println(rowElements)
		re := regexp.MustCompile(`([[:alpha:]])(?:(\d+|[[:alpha:]])/(\d+|[[:alpha:]])|(\d+))`)
		for _, command := range rowElements {
			match := re.FindStringSubmatch(command)
			if match[1] == "x" {
				i, _ := strconv.Atoi(match[2])
				j, _ := strconv.Atoi(match[3])
				exchange(programs, i, j)
				danceMoves = append(danceMoves, Exchange{i, j})
			} else if match[1] == "p" {
				r1 := rune(match[2][0])
				r2 := rune(match[3][0])
				partners(programs, r1, r2)
				danceMoves = append(danceMoves, Partners{r1, r2})
			} else if match[1] == "s" {
				n, _ := strconv.Atoi(match[4])
				spin(programs, n)
				danceMoves = append(danceMoves, Spin{n})
			}
		}
		fmt.Println("Solution to part 1:", string(programs))

		dance(programsCopy, danceMoves, 1000000000)
		fmt.Println("Solution to part 2:", string(programsCopy))
	}
}

func dance(programs []rune, moves []DanceMove, nTimes int) {
	output := string(programs)
	loopLength := 0
	for i := 0; i < nTimes; i++ {
		for _, move := range moves {
			move.Move(programs)
		}
		if output == string(programs) {
			loopLength = i
			break
		}
	}
	if loopLength > 0 {
		nTimes %= (loopLength + 1)
		dance(programs, moves, nTimes)
	}
}

func danceOld(programs []rune, moves []DanceMove, nTimes int) {
	for i := 0; i < nTimes; i++ {
		for _, move := range moves {
			move.Move(programs)
		}
	}
}

func exchange(programs []rune, p1 int, p2 int) {
	programs[p1], programs[p2] = programs[p2], programs[p1]
}

func partners(programs []rune, p1 rune, p2 rune) {
	i, j := 0, 0
	for i = range programs {
		if programs[i] == p1 {
			break
		}
	}
	for j = range programs {
		if programs[j] == p2 {
			break
		}
	}
	exchange(programs, i, j)
}

func spin(programs []rune, n int) {
	tmp := make([]rune, len(programs))
	copy(tmp, programs)
	for i := 0; i < n; i++ {
		programs[i] = tmp[len(programs)-n+i]
	}
	for i := n; i < len(programs); i++ {
		programs[i] = tmp[i-n]
	}
}
