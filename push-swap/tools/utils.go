package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AreAllDifferent(numbers []int) bool {
	seen := make(map[int]bool)
	for _, number := range numbers {
		if _, exists := seen[number]; exists {
			return false
		}
		seen[number] = true
	}
	return true
}
func EstTrie(tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			return false
		}
	}
	return true
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
