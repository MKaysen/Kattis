package main

import "fmt"

func main() {
  var n, k, p, q int

  fmt.Scanf("%d", &n)

  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d/%d", &k, &p, &q)
    fmt.Println(Max(p, q))
  }
}

func Max(x, y int) int {
  if y > x {
    return y
  }
  return x
}
