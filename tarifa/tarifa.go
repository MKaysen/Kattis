package main

import "fmt"

func main() {
  var x, n int
  fmt.Scanf("%d", &x)
  fmt.Scanf("%d", &n)

  amount := x * (n + 1)

  for i := 0; i < n; i++ {
    var tmp int
    fmt.Scanf("%d", &tmp)
    amount = amount - tmp
  }

  fmt.Println(amount)
}
