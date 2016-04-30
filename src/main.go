package main

import (
  "fmt"
)

func main() {
  scores := []int{1, 2, 3, 4, 5}
  scores = scores[0:4]
  fmt.Println(scores)
}
