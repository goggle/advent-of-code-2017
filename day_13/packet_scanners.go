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
	input := map[int]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputRow := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		re := regexp.MustCompile(`(\d+)\s*:\s*(\d+)`)
		match := re.FindStringSubmatch(inputRow)
		key, errKey := strconv.Atoi(match[1])
		if errKey != nil {
			fmt.Fprintln(os.Stderr, "parsing input:", errKey)
		}
		value, errValue := strconv.Atoi(match[2])
		if errValue != nil {
			fmt.Fprintln(os.Stderr, "parsing input:", errValue)
		}
		input[key] = value
	}

	states := map[int]*state{}
	for key := range input {
		states[key] = &state{0, true}
	}

	fmt.Println("Solution to part 1:", simulate(states, input))
	fmt.Println("Solution to part 2:", fewestPicoseconds(input))
}

func fewestPicoseconds(input map[int]int) int {
	picoseconds := 0
	states := map[int]*state{}
	for key := range input {
		states[key] = &state{0, true}
	}
	for key := range input {
		for i := 0; i < key; i++ {
			updateSingleState(states, input, key)
		}
	}

	for {
		if checkPass(states) {
			break
		}
		update(states, input)
		picoseconds++
	}

	return picoseconds
}

func checkPass(states map[int]*state) bool {
	for key := range states {
		if states[key].position == 0 {
			return false
		}
	}
	return true
}

// The  naive brute force method seems to take to long to complete, so use the
// better approach `func fewestPicoseconds` instead.
func fewestPicosecondsBruteForce(input map[int]int) int {
	picoseconds := 0
	maxKey := 0
	for key := range input {
		if key > maxKey {
			maxKey = key
		}
	}

	states := map[int]*state{}
	for key := range input {
		states[key] = &state{0, true}
	}
	var collided bool
	for {
		collided = false
		for key := range input {
			states[key].position = 0
			states[key].increasing = true
		}
		for i := 0; i < picoseconds; i++ {
			update(states, input)
		}
		for i := 0; i <= maxKey; i++ {
			if collides(states, i) {
				collided = true
				break
			}
			update(states, input)
		}
		if collided {
			picoseconds++
		} else {
			break
		}
	}
	return picoseconds
}

func simulate(states map[int]*state, input map[int]int) int {
	severity := 0
	maxKey := 0
	for key := range input {
		if key > maxKey {
			maxKey = key
		}
	}

	for i := 0; i <= maxKey; i++ {
		if collides(states, i) {
			severity += i * input[i]
		}
		update(states, input)
	}
	return severity
}

func update(states map[int]*state, input map[int]int) {
	for key := range states {
		if states[key].increasing {
			if states[key].position < input[key]-1 {
				states[key].position++
			}
			if states[key].position == input[key]-1 {
				states[key].increasing = false
			}
		} else {
			if states[key].position > 0 {
				states[key].position--
			}
			if states[key].position == 0 {
				states[key].increasing = true
			}
		}
	}
}

func updateSingleState(states map[int]*state, input map[int]int, key int) {
	if _, ok := states[key]; ok {
		if states[key].increasing {
			if states[key].position < input[key]-1 {
				states[key].position++
			}
			if states[key].position == input[key]-1 {
				states[key].increasing = false
			}
		} else {
			if states[key].position > 0 {
				states[key].position--
			}
			if states[key].position == 0 {
				states[key].increasing = true
			}
		}
	}
}

func collides(states map[int]*state, slot int) bool {
	if val, ok := states[slot]; ok {
		if val.position == 0 {
			return true
		}
	}
	return false
}
