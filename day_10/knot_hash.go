package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CircularList defines the data type used for the
// circular list.
type CircularList []int

func (cl *CircularList) init(length int) {
	for i := 0; i < length; i++ {
		*cl = append(*cl, i)
	}
}

func (cl *CircularList) nextIndex(currentIndex int, jump int) int {
	return (currentIndex + jump) % len(*cl)
}

func (cl *CircularList) invert(start int, length int) {
	for i := 0; i < length/2; i++ {
		k := cl.nextIndex(start, i)
		l := cl.nextIndex(start+length-1, -i)
		(*cl)[k], (*cl)[l] = (*cl)[l], (*cl)[k]
	}
}

func (cl *CircularList) getResultPartOne() int {
	return (*cl)[0] * (*cl)[1]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputString := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		numberStrings := strings.Split(inputString, ",")
		input := []int{}
		for _, numberString := range numberStrings {
			number, err := strconv.Atoi(strings.Trim(numberString, " "))
			if err != nil {
				fmt.Fprintln(os.Stderr, "converting input to integer", err)
			}
			input = append(input, number)
		}
		var circularList CircularList
		circularList.init(256)
		currentIndex := 0
		skipSize := 0
		for _, length := range input {
			circularList.invert(currentIndex, length)
			currentIndex = circularList.nextIndex(currentIndex, length+skipSize)
			skipSize++
		}

		var circularList2 CircularList
		bytes := []byte(inputString)
		commands := []int{}
		for _, b := range bytes {
			i := int(b)
			commands = append(commands, i)
		}
		commands = append(commands, 17, 31, 73, 47, 23)

		currentIndex = 0
		skipSize = 0
		circularList2.init(256)
		for i := 0; i < 64; i++ {
			for _, length := range commands {
				circularList2.invert(currentIndex, length)
				currentIndex = circularList2.nextIndex(currentIndex, length+skipSize)
				skipSize++
			}
		}
		denseHash := []int{}
		for i := 0; i < 16; i++ {
			block := circularList2[i*16]
			for j := 1; j < 16; j++ {
				block ^= circularList2[i*16+j]
			}
			denseHash = append(denseHash, block)
		}
		byteList := []byte{}
		for _, elem := range denseHash {
			byteList = append(byteList, byte(elem))
		}
		hexadecimalString := hex.EncodeToString(byteList)
		fmt.Println("Solution to part 1:", circularList.getResultPartOne())
		fmt.Println("Solution to part 2:", hexadecimalString)
	}
}
