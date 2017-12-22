package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	_, frequency := execute(commands)
	fmt.Println("Solution of part 1:", frequency)
	// fmt.Println("solution of part 2:", sum2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	out1 := make(chan int)
	out2 := make(chan int)
	go executeParallel(commands, 0, ch1, ch2, out1)
	go executeParallel(commands, 1, ch2, ch1, out2)
	<-out1
	fmt.Println("Solution of part 2:", <-out2)
}

ype SyncQueue struct {
	mu    sync.Mutex
	queue []int
}

func (sq *SyncQueue) Active() (a bool) {
	sq.mu.Lock()
	if len(sq.queue) > 0 {
		a = true
	} else {
		a = false
	}
	sq.mu.Unlock()
	return
}

func (sq *SyncQueue) Add(value int) {
	sq.mu.Lock()
	sq.queue = append(sq.queue, value)
	sq.mu.Unlock()
}

func (sq *SyncQueue) Pop() (value int) {
	sq.mu.Lock()
	value, sq.queue = sq.queue[0], sq.queue[1:]
	sq.mu.Unlock()
	return
}

func (sq *SyncQueue) Length() (length int) {
	sq.mu.Lock()
	length = len(sq.queue)
	sq.mu.Unlock()
	return
}

func executeParallel(commands [][]string, initValue int, sender chan int, receiver chan int, out chan int) {
	count := 0
	results := map[string]int{}
	// sendQueue := []int{}
	results["p"] = initValue

	receiverQueue := []int{}
	go func() {
		for {
			value := <-receiver
			receiverQueue = append(receiverQueue, value)
		}
	}()

	i := 0
	for i < len(commands) {
		// fmt.Println(initValue)
		switch commands[i][0] {
		case "snd":
			count++
			val, err := strconv.Atoi(commands[i][1])
			if err != nil {
				val = results[commands[i][1]]
			}
			// fmt.Println(initValue, count)
			// sendQueue = append(sendQueue, val)
			// ch <- val // Go-Routine is blocked now, until the value is received.
			sender <- val
		case "rcv":
			// if results[commands[i][1]] != 0 {
			// 	return count
			// }

			// If our sending queue is not empty, we send to the
			// other go routine
			// if len(sendQueue) > 0 {
			// 	sendVal := sendQueue[0]
			// 	sendQueue = sendQueue[1:]
			// 	sender <- sendVal
			// }
			if len(receiverQueue) == 0 {
				// fmt.Println(initValue)
				out <- count
			} else {
				results[commands[i][1]] = receiverQueue[0]
				receiverQueue = receiverQueue[1:]
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
			if results[commands[i][1]] != 0 {
				val, err := strconv.Atoi(commands[i][2])
				if err != nil {
					val = results[commands[i][2]]
				}
				i += val
				continue
			}
		}
		i++
	}
}

func execute(commands [][]string) (map[string]int, int) {
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
			if results[commands[i][1]] != 0 {
				val, err := strconv.Atoi(commands[i][2])
				if err != nil {
					val = results[commands[i][2]]
				}
				i += val
				continue
			}
		}
		i++
	}

	return results, freq
}
