package tools

import (
	"fmt"
	"os"
	"strconv"
)

// //-----if the size of the array is less than or equal to three this function will be executed-----////////
func ResolveThree(a *Stack) {
	if len(TakeArgs()) == 2 {
		if a.items[0] > a.items[1] {
			a.Sa()
			Count++
		}
		return
	} else {
		if a.items[0] > a.items[1] && a.items[0] > a.items[2] && a.items[1] > a.items[2] {
			a.Sa()
			a.Rra()
			Count += 2
			fmt.Println("sa")
			fmt.Println("rra")

		}
		if a.items[0] > a.items[1] && a.items[0] < a.items[2] && a.items[1] < a.items[2] {
			a.Sa()
			Count++
			fmt.Println("sa")

		}
		if a.items[0] > a.items[1] && a.items[0] > a.items[2] && a.items[1] < a.items[2] {
			a.Ra()
			Count++
			fmt.Println("ra")

		}
		if a.items[0] < a.items[1] && a.items[0] < a.items[2] && a.items[1] > a.items[2] {
			a.Sa()
			a.Ra()
			Count += 2
			fmt.Println("sa")
			fmt.Println("ra")

		}
		if a.items[0] < a.items[1] && a.items[0] > a.items[2] && a.items[1] > a.items[2] {
			a.Rra()
			Count++
			fmt.Println("rra")

		}
	}
}

///////////-----------------END--------------------/////////////

///----------------------Resolve all other length of number -----------------------------------///

func PivotWork(a *Stack) {
	min, pos := Min(a.items)
	for a.items[0] > a.items[1] && a.items[1] == min {
		a.Sa()
		Count++
		fmt.Println("sa")
	}
	pivot := len(a.items) / 2
	i := 0
	if pos > pivot {
		for pos != len(a.items) {
			a.Rra()
			Count++
			fmt.Println("rra")
			i++
			pos++
		}
	} else if pos <= pivot && pos != 1 {
		for i < pos {
			a.Ra()
			Count++
			fmt.Println("ra")
			i++
		}
	}
}

func Resolve(a *Stack, b *Stack) {
	PivotWork(a)
	for len(a.items) != 3 {
		PivotWork(a)
		b.Pb(a)
		Count++
		fmt.Println("pb")
	}
	ResolveThree(a)
	taille := len(b.items)
	for i := 0; i < taille; i++ {
		for EstTrie(a.items) && len(b.items) > 0 {
			a.Pa(b)
			Count++
			fmt.Println("pa")

		}
	}
}
func Min(tab []int) (int, int) {
	min := tab[0]
	pos := 0
	for i, v := range tab {
		if v < min {
			min = v
			pos = i
		}
	}
	return min, pos
}

// /----------------------END---------------------------/////////////////
// /----------------------Check length and exec the good function---------------///////////
func Exec() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	var stackA, stackB Stack
	for _, num := range TakeArgs() {
		stackA.push(num)
	}
	if !AreAllDifferent(stackA.items) {
		fmt.Println("Error")
		os.Exit(0)
	}
	if EstTrie(stackA.items) {
		os.Exit(0)
	}
	if len(stackA.items) <= 3 {
		ResolveThree(&stackA)
		fmt.Println("Solution: " + Conv(stackB.items, 0))
	} else if len(stackA.items) < 100 {
		Resolve(&stackA, &stackB)
		fmt.Println("Solution: " + Conv(stackA.items, 0))
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

/////-------------------END----------------------------------////////////////
