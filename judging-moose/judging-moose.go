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
	for scanner.Scan() {
		var input = strings.Split(scanner.Text(), " ")
		l, _ := strconv.Atoi(input[0])
		r, _ := strconv.Atoi(input[1])
		if l == r {
			if l == 0 {
				fmt.Printf("Not a moose\n")
			} else {
				fmt.Printf("Even %d\n", l*2)
			}
		} else {
			fmt.Printf("Odd %d\n", max(l, r)*2)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
