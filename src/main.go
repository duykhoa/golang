package main

import (
  "fmt"
)

func main() {
  a := 8
  test(&a)

  fmt.Println(a)
}

func test(xPtr *int) {
  *xPtr = *xPtr * *xPtr
}
