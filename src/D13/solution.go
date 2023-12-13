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
      output += (horizontal + 1) * 100
    }
    vertical := findVerticalReflection(pattern)
    if vertical != -1 {
      output += vertical + 1
    }
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  fmt.Println("Part 2:", output)
}

func findHorizontalReflection(pattern [][]string) int {
  for i := 0; i < len(pattern) - 1; i++ {
    current := strings.Join(pattern[i], "")
    next := strings.Join(pattern[i+1], "")
    if current == next {
      if checkAdjacent(pattern, i) {
        return i
      }
    }
  }

  return -1
}

func checkAdjacent(pattern [][]string, pos int) bool {
  for offset := 0; offset <= pos; offset++ {
    leftPos := pos - offset
    rightPos := pos + offset + 1
    if leftPos < 0 || rightPos >= len(pattern) {
      return true
    }
    left := strings.Join(pattern[leftPos], "")
    right := strings.Join(pattern[rightPos], "")
    if left == right {
      continue
    } else {
      return false
    }
  }
  return true
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
