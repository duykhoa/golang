package main

import (
  "fmt"
)

func main() {
  a, b := 1, 2
  swap(&a, &b)

  fmt.Println(a, b)
}

func swap(x *int, y *int) {
  *x, *y = *y, *x
}
