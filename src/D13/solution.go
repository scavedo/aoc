package main

import (
  _ "embed"
  "fmt"
  "strings"
)

//go:embed input.txt
var input string
var lines []string
var patterns map[int][][]string

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n\n")
  patterns = make(map[int][][]string)
  for i, line := range lines {
    patterns[i] = [][]string{}
    for _, rule := range strings.Split(line, "\n") {
      patterns[i] = append(patterns[i], strings.Split(rule, ""))
    }
  }
}

func main() {
  part1()
  part2()
}


func part1() {
  output := 0

  for _, pattern := range patterns {
    horizontal := findHorizontalReflection(pattern)
    if horizontal != -1 {
      output += (horizontal) * 100
    }
    vertical := findVerticalReflection(pattern)
    if vertical != -1 {
      output += vertical
    }
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  fmt.Println("Part 2:", output)
}

func findHorizontalReflection(pattern [][]string) int {

  return -1
}

func findVerticalReflection(pattern [][]string) int {
  rotated := rotatePattern(pattern)
  return findHorizontalReflection(rotated)
}

func rotatePattern(pattern [][]string) [][]string {
	if len(pattern) == 0 {
		return pattern
	}

	rows, cols := len(pattern), len(pattern[0])

	newMatrix := make([][]string, cols)
	for i := range newMatrix {
		newMatrix[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newMatrix[j][rows-i-1] = pattern[i][j]
		}
	}

	return newMatrix
}
