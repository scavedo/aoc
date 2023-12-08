package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

type Node struct {
  left string
  right string
}
var nodes map[string]Node
var directions []string

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  nodes = make(map[string]Node)
  directions = strings.Split(lines[0], "")
  for _, line := range lines[2:] {
    l := strings.Split(line, "=")
    current := strings.TrimSpace(l[0])
    re := regexp.MustCompile(`\(([^,]+),\s*([^)]+)\)`)
    coords := re.FindStringSubmatch(l[1])[1:]
    nodes[current] = Node{coords[0], coords[1]}
  }
}

func main() {
  part1()
  part2()
}


func part1() {
  node := nodes["AAA"]
  count := 0
  for node != nodes["ZZZ"] {
    for _, direction := range directions {
      if direction == "L" {
        node = nodes[node.left]
      } else if direction == "R" {
        node = nodes[node.right]
      }
      count++
    }
  }
  fmt.Println("Part 1:", count)
}

func part2() {
  count := 0
  startingNodes := []string{}
  for k, _ := range nodes {
    if strings.HasSuffix(k, "A") {
      startingNodes = append(startingNodes, k)
    }
  }

  nodeCounts := []int{}
  for _, n := range startingNodes {
    nCount := 0
    for !strings.HasSuffix(n, "Z") {
      for _, direction := range directions {
        if direction == "L" {
          n = nodes[n].left
        } else if direction == "R" {
          n = nodes[n].right
        }
        nCount++
      }
    }
    nodeCounts = append(nodeCounts, nCount)
  }

  count = LCM(nodeCounts[0], nodeCounts[1], nodeCounts[2:]...)
  fmt.Println("Part 2:", count)
}

func GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

func LCM(a, b int, integers ...int) int {
  result := a * b / GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}
