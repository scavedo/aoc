package main

import (
  _ "embed"
  "fmt"
  "strings"
)

//go:embed input.txt
var input string
var lines []string

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
}

func main() {
  part1()
  part2()
}


func part1() {
  output := 0
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  fmt.Println("Part 2:", output)
}
