package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
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

var hexToBinTable = map[rune]string{}
var countFreeTable = map[rune]int{}

// HexToBinTable['0'] = "0000"

// 	'0': "0000",
// }

func main() {
	initHexToBinTable()
	initCountFreeTable()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputString := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}

		knotHashes := []string{}
		for i := 0; i < 128; i++ {
			input := fmt.Sprintf("%s-%d", inputString, i)
			knotHashes = append(knotHashes, createKnotHash(input))
		}
		fmt.Println("Solution to part 1:", countUsedEntries(knotHashes))
		table := createTable(knotHashes)
		fmt.Println("Solution to part 2:", setBlocks(table))
	}
}

func createKnotHash(input string) string {
	var circularList CircularList
	bytes := []byte(input)
	commands := []int{}
	for _, b := range bytes {
		i := int(b)
		commands = append(commands, i)
	}
	commands = append(commands, 17, 31, 73, 47, 23)

	currentIndex := 0
	skipSize := 0
	circularList.init(256)
	for i := 0; i < 64; i++ {
		for _, length := range commands {
			circularList.invert(currentIndex, length)
			currentIndex = circularList.nextIndex(currentIndex, length+skipSize)
			skipSize++
		}
	}
	denseHash := []int{}
	for i := 0; i < 16; i++ {
		block := circularList[i*16]
		for j := 1; j < 16; j++ {
			block ^= circularList[i*16+j]
		}
		denseHash = append(denseHash, block)
	}
	byteList := []byte{}
	for _, elem := range denseHash {
		byteList = append(byteList, byte(elem))
	}
	hexadecimalString := hex.EncodeToString(byteList)
	return hexadecimalString
}

func initHexToBinTable() {
	hexToBinTable['0'] = "0000"
	hexToBinTable['1'] = "0001"
	hexToBinTable['2'] = "0010"
	hexToBinTable['3'] = "0011"
	hexToBinTable['4'] = "0100"
	hexToBinTable['5'] = "0101"
	hexToBinTable['6'] = "0110"
	hexToBinTable['7'] = "0111"
	hexToBinTable['8'] = "1000"
	hexToBinTable['9'] = "1001"
	hexToBinTable['a'] = "1010"
	hexToBinTable['b'] = "1011"
	hexToBinTable['c'] = "1100"
	hexToBinTable['d'] = "1101"
	hexToBinTable['e'] = "1110"
	hexToBinTable['f'] = "1111"
}

func initCountFreeTable() {
	for k, v := range hexToBinTable {
		count := 0
		for _, c := range v {
			if c == '0' {
				count++
			}
		}
		countFreeTable[k] = count
	}
}

func countUsedEntries(hashes []string) int {
	count := 0
	for _, hash := range hashes {
		for _, c := range hash {
			count += (4 - countFreeTable[c])
		}
	}
	return count
}

func createTable(hashes []string) [128][128]int {
	table := [128][128]int{}
	for i, hash := range hashes {
		ind := 0
		for _, hexC := range hash {
			bin := hexToBinTable[hexC]
			for _, b := range bin {
				if b == '0' {
					table[i][ind] = 0
				} else {
					table[i][ind] = -1
				}
				ind++
			}
		}
	}
	return table
}

func setBlocks(table [128][128]int) int {
	blockCounter := 0

	type tuple struct {
		i int
		j int
	}

	found := true
	for found {
		found = false
		for rowIndex, row := range table {
			for colIndex := range row {
				if table[rowIndex][colIndex] == -1 {
					found = true
					blockCounter++
					indexQueue := []tuple{tuple{rowIndex, colIndex}}
					index := tuple{}
					for len(indexQueue) > 0 {
						index, indexQueue = indexQueue[0], indexQueue[1:]
						i, j := index.i, index.j
						table[i][j] = blockCounter
						if i+1 < 128 && table[i+1][j] == -1 {
							indexQueue = append(indexQueue, tuple{i + 1, j})
						}
						if i-1 >= 0 && table[i-1][j] == -1 {
							indexQueue = append(indexQueue, tuple{i - 1, j})
						}
						if j+1 < 128 && table[i][j+1] == -1 {
							indexQueue = append(indexQueue, tuple{i, j + 1})
						}
						if j-1 >= 0 && table[i][j-1] == -1 {
							indexQueue = append(indexQueue, tuple{i, j - 1})
						}
					}
				}
			}
		}
	}
	return blockCounter
}
