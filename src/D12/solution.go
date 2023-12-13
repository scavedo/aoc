package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string
var records []string
var lengths []int

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

  for _, line := range lines {
    parseLine(line)
    result := process(0, 0)
    output += result
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  fmt.Println("Part 2:", output)
}

func parseLine(line string) {
  split := strings.Split(line, " ")
  records = strings.Split(split[0], "")
  lengths = []int{}
  for _, i := range strings.Split(split[1], ",") {
    i, _ := strconv.Atoi(i)
    lengths = append(lengths, i)
  }
}

func process(r int, l int) int {
  if r == len(records) {
    if l == len(lengths) {
      return 1
    } else {
      return 0
    }
  }
  record := records[r]
  switch record {
  case ".": return processWorking(r, l)
  case "#": return processBroken(r, l)
  case "?": return processUnknown(r, l)
  }
  return 0
}

func processWorking(r int, l int) int {
  return process(r + 1, l)
}

func processBroken(r int, l int) int {
  if l == len(lengths) {
    return 0
  }
  size := lengths[l]
  endIdx := r + size
  if !willFit(r, endIdx) {
    return 0
  }
  if endIdx == len(records) {
    if l == len(lengths) - 1 {
      return 1
    } else {
      return 0
    }
  }
  return process(endIdx + 1, l + 1)
}

func processUnknown(r int, l int) int {
  return processWorking(r, l) + processBroken(r, l)
}

func willFit(start int, end int) bool {
  if end > len(records) {
    return false
  } else if end == len(records) {
    sub := strings.Join(records[start:end], "")
    return !strings.Contains(sub, ".")
  } else {
    sub := strings.Join(records[start:end], "")
    return !strings.Contains(sub, ".") && records[end] != "#"
  }
}
