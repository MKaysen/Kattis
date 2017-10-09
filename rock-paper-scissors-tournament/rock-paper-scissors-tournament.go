package main

import (
	"bufio"
	"os"
)

type struct GameStat {
  w int
  l int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

  }
}

func ScanGame(s *Scanner, n, k, gameNr int) {

}

func PrintGame(gameData map[string]GameStat, gameNr int) {
  if gameNr > 0 {
    fmt.Printf("\n")
  }
  for _, value := range gameData {
    if gameData.w == 0 && gameData.l == 0 {

    } else {
      winrate := value.w / (value.w + value.l)
    fmt.Printf("%d\n", winrate)
    }
  }
}
