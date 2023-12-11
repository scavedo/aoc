package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var lines []string
var universe [][]string
var expansionRows []int
var expansionColumns []int

type Galaxy struct {
  Id int
  X int
  Y int
}

var galaxyMap = make(map[int]Galaxy)

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  universe = make([][]string, len(lines))
  for i, line := range lines {
    universe[i] = strings.Split(line, "")
  }
  mapExpansions()
}

func main() {
  part1()
  part2()
}


func part1() {
  total := 0
  mapGalaxies()
  for _, galaxy := range galaxyMap {
    paths := galaxy.findShortestPaths(2)
    for _, path := range paths {
      total += path
    }
  }
  fmt.Println("Part 1:", total)
}

func part2() {
  total := 0
  mapGalaxies()
  for _, galaxy := range galaxyMap {
    paths := galaxy.findShortestPaths(1000000)
    for _, path := range paths {
      total += path
    }
  }
  fmt.Println("Part 2:", total)
}

func mapExpansions() {
  for i, line := range universe {
    if !slices.Contains(line, "#") {
      expansionRows = append(expansionRows, i)
    }
  }

  for j := 0; j < len(universe[0]); j++ {
    var column []string
    for _, line := range universe {
      column = append(column, line[j])
    }
    if !slices.Contains(column, "#") {
      expansionColumns = append(expansionColumns, j)
    }
  }
}

func mapGalaxies() {
  for i, line := range universe {
    for j, galaxy := range line {
      if galaxy == "#" {
        galaxyMap[len(galaxyMap) + 1] = Galaxy{len(galaxyMap) + 1, i, j}
      }
    }
  }
}

func (g Galaxy) findShortestPaths(expansion int) []int {
  closest := []int{}
  for _, galaxy := range galaxyMap {
    dist := math.Abs(float64(g.X - galaxy.X)) + math.Abs(float64(g.Y - galaxy.Y))
    if galaxy.Id != g.Id {
      for _, row := range expansionRows {
        if minInt(g.X, galaxy.X) < row && row < maxInt(g.X, galaxy.X) {
          dist += float64(expansion - 1)
        }
      }
      for _, column := range expansionColumns {
        if minInt(g.Y, galaxy.Y) < column && column < maxInt(g.Y, galaxy.Y) {
          dist += float64(expansion - 1)
        }
      }
      closest = append(closest, int(dist))
    }
  }
  delete(galaxyMap, g.Id)
  return closest
}

func minInt(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func maxInt(a, b int) int {
  if a > b {
    return a
  }
  return b
}
