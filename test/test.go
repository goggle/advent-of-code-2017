package main

import (
	"fmt"
	"sync"
)

type Command struct {
	name  string
	value int
}

func main() {
	commandsP1 := []Command{
		Command{"snd", 1},
		Command{"snd", 2},
		Command{"snd", 3},
		Command{"rcv", 0},
		Command{"rcv", 0},
		Command{"snd", 4},
		Command{"rcv", 0},
		Command{"rcv", 0},
	}
	commandsP2 := []Command{
		Command{"snd", 10},
		Command{"snd", 11},
		Command{"rcv", 0},
		Command{"rcv", 0},
		Command{"snd", 12},
		Command{"snd", 13},
		Command{"rcv", 0},
		Command{"rcv", 0},
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	out1 := make(chan int)
	out2 := make(chan int)
	go program(commandsP1, ch1, ch2, out1)
	go program(commandsP2, ch2, ch1, out2)
	fmt.Println(<-out1)
	fmt.Println(<-out2)
}

func program(commands []Command, sender chan int, receiver chan int, out chan int) {
	var mux sync.Mutex
	count := 0
	sendQueue := []int{}
	for _, command := range commands {
		switch command.name {
		case "snd":
			mux.Lock()
			sendQueue = append(sendQueue, command.value)
			mux.Unlock()
			count++
		case "rcv":
			canSend := false
			value := 0
			mux.Lock()
			if len(sendQueue) > 0 {
				canSend = true
				value, sendQueue = sendQueue[0], sendQueue[1:]
			}
			mux.Unlock()

			value := <-receiver
			fmt.Println("Received value", value)
		}
	}
	out <- count

}
