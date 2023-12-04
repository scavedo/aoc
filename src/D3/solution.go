package main

import (
  _ "embed"
  "fmt"
  // "slices"
  "strconv"
  "strings"
)

//go:embed input.txt
var input string
var lines []string
var dirs = [][]int{
  {1, 1}, {1, -1}, {-1, -1}, {-1, 1},
  {-1, 0}, {1, 0},
  {0, 1}, {0, -1},
}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
}

func main() {
  part1()
  part2()
}


func part1() {
  total := 0
  for i, line := range lines {
    var lineNums [][]int
    for j := 0; j < len(line); j++ {
      var idx []int
      for isNumber(string(line[j])) {
        idx = append(idx, j)
        if j + 1 < len(line) {
          j++
        } else {
          break
        }
      }
      if len(idx) > 0 {
        a := anySurroundingSymbols(i, idx)
        if a {
          lineNums = append(lineNums, idx)
        }
      }
    }

    for _, l := range lineNums {
      part := getPart(i, l)
      partNumber, _ := strconv.Atoi(part)
      total += partNumber
    }
  }
  fmt.Println("Part 1:", total)
}

func part2() {
  fmt.Println("Part 2:", 0)
}

func isNumber(s string) bool {
  _, err := strconv.Atoi(s)
  return err == nil
}

func isDot(s string) bool {
  return s == "."
}

func isSymbol(s string) bool {
  return !isNumber(s) && !isDot(s)
}

func anySurroundingSymbols(x int, y []int) bool {
  for _, i := range y {
    if anyAdjacentSymbols(x, i) {
      return true
    }
  }
  return false
}

func anyAdjacentSymbols(x int, y int) bool {
  for _, dir := range dirs {
    xi := 0
    yi := 0
    if x + dir[0] < 0 || x + dir[0] >= len(lines) {
      xi = 0
    } else {
      xi = x + dir[0]
    }
    if y + dir[1] < 0 || y + dir[1] >= len(lines[x]) {
      yi = 0
    } else {
      yi = y + dir[1]
    }
    if isSymbol(string(lines[xi][yi])) {
      return true
    }
  }
  return false
}

func getPart(x int, ys []int) string {
  start := ys[0]
  end := ys[len(ys) - 1]
  line := lines[x]
  return string(line[start:end + 1])
}
