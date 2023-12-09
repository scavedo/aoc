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
  for _, line := range lines {
    entries := []int{}
    for _, e := range strings.Fields(line) {
      entry, _ := strconv.Atoi(e)
      entries = append(entries, entry)
    }

    _, right := buildTree(entries)
    total += right
  }
  fmt.Println("Part 1:", total)
}

func part2() {
  total := 0
  for _, line := range lines {
    entries := []int{}
    for _, e := range strings.Fields(line) {
      entry, _ := strconv.Atoi(e)
      entries = append(entries, entry)
    }

    left, _ := buildTree(entries)
    total += left
  }
  fmt.Println("Part 2:", total)
}

func buildTree(nums []int) (int, int) {
  combinedNums := []int{}
  zeroCount := 0
  for i := 0; i < len(nums); i++ {
    if i + 1 == len(nums) {
      break
    }

    value := nums[i + 1] - nums[i]
    if value == 0 {
      zeroCount++
    }
    combinedNums = append(combinedNums, value)
  }

  if zeroCount == len(combinedNums) {
    return nums[0], nums[len(nums) - 1]
  }

  left, right := buildTree(combinedNums)

  return nums[0] - left, nums[len(nums) - 1] + right
}
