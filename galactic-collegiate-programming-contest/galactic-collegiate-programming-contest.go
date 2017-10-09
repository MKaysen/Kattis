package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type team struct {
	id      int
	solved  int
	penalty int
}

type byBest []team

func (a byBest) Len() int      { return len(a) }
func (a byBest) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byBest) Less(i, j int) bool {
	if a[i].solved > a[j].solved {
		return true
	} else if a[i].solved == a[j].solved && a[i].penalty < a[j].penalty {
		return true
	} else if a[i].solved == a[j].solved && a[i].penalty == a[j].penalty {
		return a[i].id < a[j].id
	}
	return false
}

func main() {
	var better []team
	teams := make(map[int]*team)
	myRank := 1
	scanner := bufio.NewScanner(os.Stdin)
	var n, m int
	itr := 1
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if n == 0 {
			n, _ = strconv.Atoi(input[0])
			m, _ = strconv.Atoi(input[1])
			for i := 1; i <= n; i++ {
				teams[i] = &team{id: i}
			}
			continue
		}
		t, _ := strconv.Atoi(input[0])
		p, _ := strconv.Atoi(input[1])
		teams[t].solved++
		teams[t].penalty += p
		if t == 1 {
			better = append(better, *teams[t])
			sort.Sort(byBest(better))
			for idx, value := range better {
				if value.id == 1 {
					myRank = idx + 1
					better = better[:idx]
				}
			}
		} else {
			if teams[t].solved > teams[1].solved || (teams[t].solved == teams[1].solved && teams[t].penalty < teams[1].penalty) {
				addToList := true
				for idx, value := range better {
					if value.id == t {
						better[idx] = *teams[t]
						addToList = false
					}
				}
				if addToList {
					better = append(better, *teams[t])
					myRank++
				}
			}
		}
		fmt.Println(myRank)
		if itr == m {
			break
		}
		itr++
	}
}
