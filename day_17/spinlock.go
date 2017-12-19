package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Entry struct {
	number   int
	previous *Entry
	next     *Entry
}

type Spinlock struct {
	size    int
	current *Entry
	start   *Entry
}

func (s *Spinlock) add(number int, steps int) {
	entry := Entry{number, nil, nil}
	if s.current == nil {
		entry.previous = &entry
		entry.next = &entry
		s.current = &entry
		s.size = 1
		s.start = &entry
		return
	}
	for i := 0; i < steps%s.size; i++ {
		s.current = s.current.next
	}
	curr := s.current
	next := curr.next
	s.current = &entry
	entry.next = next
	entry.previous = curr
	s.size++
	curr.next = &entry
	next.previous = &entry
}

func (s *Spinlock) print() {
	curr := s.current
	for i := 0; i < s.size; i++ {
		fmt.Printf("%d ", curr.number)
		curr = curr.next
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		fmt.Println(input)
		spinlock := Spinlock{}
		spinlock.add(0, 0)
		for i := 1; i <= 2017; i++ {
			spinlock.add(i, input)
		}
		fmt.Println("Solution of part 1:", spinlock.current.next.number)
		for i := 2018; i <= 50000000; i++ {
			if i%1000000 == 0 {
				fmt.Println(i)
			}
			spinlock.add(i, input)
		}
		fmt.Println("Solution of part 2:", spinlock.start.next)
	}
	// fmt.Println("Solution of part 1:", sum1)
	// fmt.Println("solution of part 2:", sum2)
}
