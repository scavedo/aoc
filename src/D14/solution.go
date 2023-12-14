package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var lines []string
var pattern [][]string

var dirs = []string{"N", "W", "S", "E"}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  pattern = buildPattern(lines)
}

func main() {
  part1()
  part2()
}

func part1() {
  output := 0

  x := rollStonesUp(pattern)
  for i, line := range x {
    s := indexStones(line)
    output += (len(x) - i) * len(s)
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  cache := map[string]int{}
  i := 0
  repeatStart := 0
  repeatLength := 0
  cycles := 1000000000
  for i < cycles {
    pattern = spinPattern(pattern)

    y := stringifyPattern(pattern)
    if _, ok := cache[y]; ok {
      repeatStart = i
      repeatLength = i - cache[y]
      break
    } else {
      cache[y] = i
    }
    i++
  }

  offset := repeatStart - repeatLength
  idx := (cycles - 1 - offset) % repeatLength + offset
  for str, i := range cache {
    if i == idx {
      patternLine := strings.TrimRight(str, "\n")
      pattern = buildPattern(strings.Split(patternLine, "\n"))
    }
  }

  for i, line := range pattern {
    s := indexStones(line)
    output += (len(pattern) - i) * len(s)
  }

  fmt.Println("Part 2:", output)
}

func spinPattern(in [][]string) [][]string {
  out := append([][]string{}, in...)
  for range dirs {
    out = rollStonesUp(out)
    out = rotatePattern(out)
  }
  return out
}

func buildPattern(in []string) [][]string {
  out := [][]string{}
  for _, line := range in {
    line = strings.TrimSpace(line)
    out = append(out, strings.Split(line, ""))
  }
  return out
}

func stringifyPattern(in [][]string) string {
  out := ""
  for _, line := range in {
    x := strings.Join(line, "")
    x = strings.TrimSpace(x)
    out += x + "\n"
  }
  return out
}

func rotatePattern(pattern [][]string) [][]string {
rows := len(pattern)
	cols := len(pattern[0])

	rotated := make([][]string, cols)
	for i := range rotated {
		rotated[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-i-1] = pattern[i][j]
		}
	}

	return rotated
}

func indexBlocks(line []string, blocks []string) []int {
  output := []int{}
  for i, s := range line {
    for _, block := range blocks {
      if s == block {
        output = append(output, i)
      }
    }
  }
  return output
}

func indexStones(line []string) []int {
  output := []int{}
  for i, s := range line {
    if s == "O" {
      output = append(output, i)
    }
  }
  return output
}

func rollStonesUp(in [][]string) [][]string {
  out := append([][]string{}, in...)
  j := 2
  for j <= len(out) {
    l := out[:j]
    i := len(l) - 1
    for i > 0 {
      currentLine := l[i]
      nextLine := l[i-1]
      currentLine, nextLine = rollStones(currentLine, nextLine)
      l[i] = currentLine
      l[i-1] = nextLine
      i--
    }
    out = append(l, out[j:]...)
    j++
  }
  return out
}

func rollStones(current []string, compare []string) ([]string, []string) {
  blocks := indexBlocks(compare, []string{"#", "O"})
  stones := indexStones(current)
  for _, stone := range stones {
    if !slices.Contains(blocks, stone) {
      compare = replaceAt(compare, "O", stone)
      current = replaceAt(current, ".", stone)
    }
  }
  return current, compare
}

func replaceAt(input []string, replacement string, index int) []string {
  input[index] = replacement
  return input
}
