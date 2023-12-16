package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	items []int
}

func executeInstruction(a, b *Stack, instruction string) ([]int, []int, error) {
	switch instruction {
	case "sa":
		a.Sa()
		return a.items, b.items, nil
	case "sb":
		b.Sb()
		return a.items, b.items, nil
	case "ss":
		a.Ss(b)
		return a.items, b.items, nil
	case "ra":
		a.Ra()
		return a.items, b.items, nil
	case "rb":
		b.Rb()
		return a.items, b.items, nil
	case "rr":
		a.Rr(b)
		return a.items, b.items, nil
	case "rra":
		a.Rra()
		return a.items, b.items, nil
	case "rrb":
		b.Rrb()
		return a.items, b.items, nil
	case "rrr":
		a.Rrr(b)
		return a.items, b.items, nil
	case "pa":
		a.Pa(b)
		return a.items, b.items, nil
	case "pb":
		b.Pb(a)
		return a.items, b.items, nil
	case "Solution: " + Conv(a.items, 0):
		return nil, nil, nil
	default:
		return nil, nil, fmt.Errorf("unknown instruction: %s", instruction)
	}
}

func isSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	stackA := &Stack{}
	stackB := &Stack{}

	for _, num := range TakeArgs() {
		stackA.Push(num)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var instruction string
		if scanner.Text() == "" {
			continue
		} else {
			instruction = scanner.Text()
		}
		var err error
		stackA.items, stackB.items, err = executeInstruction(stackA, stackB, instruction)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	if isSorted(stackA.items) && len(stackB.items) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func Conv(tab []int, a int) string {
	if a != 0 {
		word := strconv.Itoa(a)
		return word
	} else {
		var word string
		for _, v := range tab {
			temp := strconv.Itoa(v)
			word += temp + " "
		}
		return word
	}
}

func IsEmpty() []string {
	var tab []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var instruction string
		if scanner.Text() == "" {
			continue
		} else {
			instruction = scanner.Text()
			tab = append(tab, instruction)
		}
	}

	return tab
}
