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
	size         int
	current      *Entry
	start        *Entry
	currentIndex int
}

func (s *Spinlock) add(number int, steps int) {
	entry := Entry{number, nil, nil}
	if s.current == nil {
		entry.previous = &entry
		entry.next = &entry
		s.current = &entry
		s.size = 1
		s.start = &entry
		s.currentIndex = 0
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
	s.currentIndex = (s.currentIndex + steps) % s.size
	s.size++
	s.currentIndex = (s.currentIndex + 1) % s.size
	curr.next = &entry
	next.previous = &entry
}

func (s *Spinlock) simulateAdd(number int, steps int, afterZeroList []int) []int {
	index := (s.currentIndex + steps) % s.size

	if index == 0 {
		afterZeroList = append(afterZeroList, number)
	}
	s.currentIndex = index
	s.size++
	s.currentIndex = (s.currentIndex + 1) % s.size

	return afterZeroList

}

func (s *Spinlock) print() {
	start := s.start
	curr := start
	if s.size == 0 {
		return
	}
	for {
		if curr == s.current {
			fmt.Printf("(%d) ", curr.number)
		} else {
			fmt.Printf("%d ", curr.number)
		}
		curr = curr.next
		if curr == start {
			break
		}
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		spinlock := Spinlock{}
		spinlock.add(0, 0)

		for i := 1; i <= 2017; i++ {
			spinlock.add(i, input)
		}
		fmt.Println("Solution of part 1:", spinlock.current.next.number)
		afterZeroList := []int{spinlock.start.next.number}
		for i := 2018; i <= 50000000; i++ {
			afterZeroList = spinlock.simulateAdd(i, input, afterZeroList)
		}
		fmt.Println("Solution of part 2:", afterZeroList[len(afterZeroList)-1])
	}
}
