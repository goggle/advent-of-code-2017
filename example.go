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
	// fmt.Println("Hello World.")
	commandsP1 := []Command{
		Command{"snd", 2},
		Command{"snd", 3},
		Command{"rcv", 0},
		Command{"rcv", 0},
		Command{"snd", 4},
		Command{"rcv", 0},
		// Command{"rcv", 0},
	}
	commandsP2 := []Command{
		Command{"snd", 10},
		Command{"rcv", 0},
		Command{"snd", 10},
		Command{"snd", 10},
		Command{"rcv", 0},
		Command{"rcv", 0},
		// Command{"rcv", 0},
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	out1 := make(chan int)
	out2 := make(chan int)
	ready := make(chan bool, 2)
	quit := make(chan struct{})
	// chb1 := make(chan bool)
	// chb2 := make(chan bool)
	go func(ch chan bool) {
		for {
			v1 := <-ch
			v2 := <-ch
			if !v1 || !v2 {
				break
			}
		}
		fmt.Println("Shutting down supervisor...")
	}(ready)
	go program(commandsP1, ch1, ch2, out1, ready, quit)
	go program(commandsP2, ch2, ch1, out2, ready, quit)
	fmt.Println(<-out1)
	fmt.Println(<-out2)
}

type Counter struct {
	mu sync.Mutex
	x  int
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.x++
	c.mu.Unlock()
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	c.x--
	c.mu.Unlock()
}

func (c *Counter) Value() (x int) {
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}

type SyncQueue struct {
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

func program(commands []Command, sender chan int, receiver chan int, out chan int, ready chan bool, quit chan struct{}) {
	count := 0
	sendQueue := SyncQueue{}
	for _, command := range commands {
		if command.name == "snd" {
			sendQueue.Add(command.value)
			count++
		} else if command.name == "rcv" {
			ready <- true

			go func() {
				sender <- sendQueue.Length()
			}()
			nPossibleReceives := <-receiver

			if sendQueue.Active() {
				go func() {
					sender <- sendQueue.Pop()
				}()
			} else if nPossibleReceives == 0 {
				fmt.Println("Deadlock...")
				out <- count
			}

			for nPossibleReceives == 0 {
				ready <- true
				// select {
				// case <-quit:
				// 	return
				// case nPossibleReceives = <-receiver:
				// }
				nPossibleReceives = <-receiver

			}
			value := <-receiver
			fmt.Println("Received value", value)

		}
	}
	// ready <- false
	fmt.Println("Regular exit.")
	out <- count
	// quit <- struct{}{}
}
