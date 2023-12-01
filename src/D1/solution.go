package main

import (
  _ "embed"
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

//go:embed input.txt
var input string

var wordToNum = map[string]string{
  "one": "1",
  "two": "2",
  "three": "3",
  "four": "4",
  "five": "5",
  "six": "6",
  "seven": "7",
  "eight": "8",
  "nine": "9",
}

func init() {
  input = strings.TrimRight(input, "\n")
}

func main() {
  // part1()
  part2()
}

func part1() {
  var digits []int
  for _, s := range strings.Split(input, "\n") {
    s = stripChars(s)
    first, last := firstLastDigits(s)
    num := first + last
    numInt, _ := strconv.Atoi(num)
    digits = append(digits, numInt)
  }
  sum := sum(digits...)
  fmt.Println("Part 1: ", sum)
}

func part2() {
  var digits []int
  for _, s := range strings.Split(input, "\n") {
    s = replaceNums(s)
    s = stripChars(s)
    first, last := firstLastDigits(s)
    num := first + last
    fmt.Println(s, first, last, num)
    numInt, _ := strconv.Atoi(num)
    digits = append(digits, numInt)
  }
  sum := sum(digits...)
  fmt.Println("Part2: ", sum)
}

func replaceNums(s string) string {
  for k, v := range wordToNum {
    re := regexp.MustCompile(k).FindAllStringIndex(s, -1)
    fmt.Println(k, re)
    for _, idx := range re {
      s = s[:idx[0] + 1] + v + s[idx[0] + 2:]
    }
  }
  return s
}

func stripChars(s string) string {
  var re = regexp.MustCompile(`[^0-9]`)
  return re.ReplaceAllString(s, "")
}

func firstLastDigits(s string) (string, string) {
  if s == "" {
    return "0", "0"
  }
  
  first := string(s[0])
  last := string(s[len(s)-1])
  return first, last
}

func sum(nums... int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}
