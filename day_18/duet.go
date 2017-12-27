package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := [][]string{}
	for scanner.Scan() {
		rowElements := strings.Fields(scanner.Text())

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		commands = append(commands, rowElements)

	}

	_, freq := programSingle(commands)
	fmt.Println("Solution to part 1:", freq)

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch1Dead := make(chan bool)
	ch2Dead := make(chan bool)
	out1 := make(chan int)
	out2 := make(chan int)
	go programPair(commands, 0, ch1, ch2, ch1Dead, ch2Dead, out1)
	go programPair(commands, 1, ch2, ch1, ch2Dead, ch1Dead, out2)
	<-out1
	fmt.Println("Solution to part 2:", <-out2)
}

func programSingle(commands [][]string) (map[string]int, int) {
	results := map[string]int{}
	freq := 0

	i := 0
	for i < len(commands) {
		switch commands[i][0] {
		case "snd":
			freq = results[commands[i][1]]
		case "set":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] = val
		case "add":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] += val
		case "mul":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] *= val
		case "mod":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] %= val
		case "rcv":
			if results[commands[i][1]] != 0 {
				return results, freq
			}
		case "jgz":
			reg := commands[i][1]
			regVal, err1 := strconv.Atoi(reg)
			if err1 != nil {
				regVal = results[commands[i][1]]
			}
			jmp := commands[i][2]
			jmpVal, err2 := strconv.Atoi(jmp)
			if err2 != nil {
				jmpVal = results[commands[i][2]]
			}
			if regVal > 0 {
				i += jmpVal
				if i >= 0 && i < len(commands) {
					continue
				} else {
					return results, freq
				}
			}
		}
		i++
	}

	return results, freq
}

func programPair(commands [][]string, initValue int, sender chan int, receiver chan int, senderDead chan bool, receiverDead chan bool, out chan int) {
	results := map[string]int{}
	results["p"] = initValue
	otherRoutineDead := false
	var mux sync.Mutex
	count := 0
	commandQueue := []int{}
	go func() {
		for {
			value := <-receiver
			mux.Lock()
			commandQueue = append(commandQueue, value)
			mux.Unlock()
			senderDead <- false
		}
	}()

	// Checker for possible deadlock:
	go func() {
		for {
			dead := <-receiverDead
			mux.Lock()
			otherRoutineDead = dead
			mux.Unlock()
		}
	}()

	for i := 0; i < len(commands); {
		switch commands[i][0] {
		case "snd":
			val, err := strconv.Atoi(commands[i][1])
			if err != nil {
				val = results[commands[i][1]]
			}
			sender <- val
			count++
		case "rcv":
		receive:
			mux.Lock()
			var ok bool
			var value int
			if len(commandQueue) > 0 {
				value, commandQueue = commandQueue[0], commandQueue[1:]
				ok = true
			}
			mux.Unlock()
			if ok {
				results[commands[i][1]] = value
			} else {
				senderDead <- true
				mux.Lock()
				deadLock := otherRoutineDead
				mux.Unlock()
				if deadLock {
					goto stop
				}
				goto receive

			}
		case "set":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] = val
		case "add":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] += val
		case "mul":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] *= val
		case "mod":
			val, err := strconv.Atoi(commands[i][2])
			if err != nil {
				val = results[commands[i][2]]
			}
			results[commands[i][1]] %= val

		case "jgz":
			reg := commands[i][1]
			regVal, err1 := strconv.Atoi(reg)
			if err1 != nil {
				regVal = results[commands[i][1]]
			}
			jmp := commands[i][2]
			jmpVal, err2 := strconv.Atoi(jmp)
			if err2 != nil {
				jmpVal = results[commands[i][2]]
			}
			if regVal > 0 {
				i += jmpVal
				if i >= 0 && i < len(commands) {
					continue
				} else {
					goto stop
				}
			}
		}
		i++
	}
stop:
	select {
	case senderDead <- true:
	}
	out <- count
}
