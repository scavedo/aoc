package main

import (
  _ "embed"
  "fmt"
  "strings"
)

//go:embed input.txt
var input string

func init() {
  input = strings.TrimRight(input, "\n")
}

func main() {
  part1()
  part2()
}


func part1() {
  fmt.Println("Part 1:", 0)
}

func part2() {
  fmt.Println("Part 2:", 0)
}
