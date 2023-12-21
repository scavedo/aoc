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
var garden [][]string
var sx int
var sy int

type Tile struct {
  x int
  y int
}

var directions = []Tile{
  Tile{0, 1},
  Tile{0, -1},
  Tile{1, 0},
  Tile{-1, 0},
}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  for _, line := range lines {
    garden = append(garden, strings.Split(line, ""))
  }
  for r, row := range garden {
    for c, col := range row {
      if col == "S" {
        sx = r
        sy = c
      }
    }
  }
}

func main() {
  part1()
  part2()
}

func part1() {
  output := fill(sx, sy, 64)
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0

  steps := 26501365
  size := len(garden)
  grid_width := steps / size - 1
  odd := int(math.Pow(float64(grid_width / 2 * 2 + 1), 2))
  even := int(math.Pow(float64((grid_width + 1) / 2 * 2), 2))
  
  odd_points := fill(sx, sy, size * 2 + 1)
  even_points := fill(sx, sy, size * 2)
  corner_t := fill(size - 1, sy, size - 1)
  corner_r := fill(sx, 0, size - 1)
  corner_b := fill(0, sy, size - 1)
  corner_l := fill(sx, size - 1, size - 1)
  small_tr := fill(size - 1, 0, size / 2 - 1)
  small_br := fill(0, 0, size / 2 - 1)
  small_tl := fill(size - 1, size - 1, size / 2 - 1)
  small_bl := fill(0, size - 1, size / 2 - 1)
  large_tr := fill(size - 1, 0, size * 3 / 2 - 1)
  large_br := fill(0, 0, size * 3 / 2 - 1)
  large_tl := fill(size - 1, size - 1, size * 3 / 2 - 1)
  large_bl := fill(0, size - 1, size * 3 / 2 - 1)

  output = odd * odd_points +
    even * even_points +
    corner_t + corner_r + corner_b + corner_l +
    (grid_width + 1) * (small_tr + small_br + small_tl + small_bl) +
    (grid_width) * (large_tr + large_br + large_tl + large_bl)
  
  fmt.Println("Part 2:", output)
}

func (t Tile) Walk() []Tile {
  tiles := []Tile{}
  for _, dir := range directions {
    newSpot := Tile{t.x + dir.x, t.y + dir.y}
    if newSpot.x >= 0 &&
      newSpot.x < len(garden) && 
      newSpot.y >= 0 &&
      newSpot.y < len(garden[0]) {
      if garden[newSpot.x][newSpot.y] != "#" {
        tiles = append(tiles, newSpot)
      }
    }
  }
  return tiles
}

func fill(sx int, sy int, steps int) int {
  outer := []Tile{Tile{sx, sy}}
  for x := 0; x < steps; x++ {
    spots := outer
    outer = []Tile{}
    for _, spot := range spots {
      newSpots := spot.Walk()
      for _, newSpot := range newSpots {
        if !slices.Contains(outer, newSpot) {
          outer = append(outer, newSpot)
        }
      }
    }
  }
  return len(outer)
}
