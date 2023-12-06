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
  power := 1
  times := strings.Fields(strings.Split(lines[0], ":")[1])
  distances := strings.Fields(strings.Split(lines[1], ":")[1])
  for i, t := range times {
    numWinner := 0
    t, _ := strconv.Atoi(t)
    distance, _ := strconv.Atoi(distances[i])
    for j := 1; t - j > 0; j++ {
      timeLeft := t - j
      speed := j
      if speed * timeLeft > distance {
        numWinner++
      }
    }
    power *= numWinner
  }
  fmt.Println("Part 1:", power)
}

func part2() {
  numWinner := 0
  times := strings.Fields(strings.Split(lines[0], ":")[1])
  distances := strings.Fields(strings.Split(lines[1], ":")[1])
  time, _ := strconv.Atoi(strings.Join(times, ""))
  distance, _ := strconv.Atoi(strings.Join(distances, ""))
  for j := 1; time - j > 0; j++ {
    timeLeft := time - j
    speed := j
    if speed * timeLeft > distance {
      numWinner++
    }
  }
  fmt.Println("Part 2:", numWinner)
}
