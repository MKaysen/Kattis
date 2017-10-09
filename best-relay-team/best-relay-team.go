package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type runner struct {
	name  string
	first float64
	other float64
}

func (r runner) String() string {
	return fmt.Sprintf("%s", r.name)
}

type byFirst []runner

func (a byFirst) Len() int           { return len(a) }
func (a byFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFirst) Less(i, j int) bool { return a[i].first < a[j].first }

type byOther []runner

func (a byOther) Len() int           { return len(a) }
func (a byOther) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byOther) Less(i, j int) bool { return a[i].other < a[j].other }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var runners []runner
	var n int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		if len(input) == 1 {
			n, _ = strconv.Atoi(input[0])
		} else if len(input) == 3 {
			runner := parseRunner(input)
			runners = append(runners, runner)
		}
		if len(runners) == n {
			break
		}
	}
	sort.Sort(byFirst(runners))
	firstRunners := append([]runner{}, runners[:4]...)
	sort.Sort(byOther(runners))
	otherRunners := append([]runner{}, runners[:4]...)
	bestRunners, best := findBest(firstRunners, otherRunners)
	fmt.Println(best)
	for _, r := range bestRunners {
		fmt.Println(r)
	}
}

func parseRunner(input []string) runner {
	firstVal, _ := strconv.ParseFloat(input[1], 64)
	secondVal, _ := strconv.ParseFloat(input[2], 64)
	runner := runner{
		name:  input[0],
		first: firstVal,
		other: secondVal,
	}
	return runner
}

func findBest(first, other []runner) ([4]runner, float64) {
	best := math.MaxFloat64
	bestRunners := [4]runner{}
	for _, firstR := range first {
		itr := 1
		tmpBest := firstR.first
		tmpBestL := [4]runner{}
		tmpBestL[0] = firstR
		for _, otherR := range other {
			if otherR.name == firstR.name {
				continue
			} else if itr != 4 {
				tmpBestL[itr] = otherR
				tmpBest += otherR.other
				itr++
			}
		}
		if tmpBest < best {
			best = tmpBest
			bestRunners = tmpBestL
		}
	}
	return bestRunners, best
}
