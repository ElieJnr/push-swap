package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *Stack) Shift(item int) {
	temp := []int{item}
	temp = append(temp, s.items...)
	s.items = temp
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}
func TakeArgs() []int {
	args := os.Args[1]
	tab := strings.Fields(args)
	var tabInt []int
	for i := range tab {
		nb, err := strconv.Atoi(tab[i])
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		tabInt = append(tabInt, nb)
	}
	return tabInt
}
