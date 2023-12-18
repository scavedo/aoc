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
var lagoon map[Coord]string = make(map[Coord]string)

const (
  UP = "U"
  DOWN = "D"
  LEFT = "L"
  RIGHT = "R"
)

type Coord struct {
  X int
  Y int
}

var directions = map[string]Coord{
  UP: Coord{-1, 0},
  DOWN: Coord{1, 0},
  LEFT: Coord{0, -1},
  RIGHT: Coord{0, 1},
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
  output := 0
  edges, positions := parse(parseSteps)
  output = calculateArea(edges, positions)
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  edges, positions := parse(parseColors)
  output = calculateArea(edges, positions)
  fmt.Println("Part 2:", output)
}

func parse(fun func(string) (Coord, int)) (int, []Coord) {
  coords := []Coord{}
  coord := Coord{}
  edges := 0
  for _, line := range lines {
    direction, steps := fun(line)
    coord.move(direction, steps)
    edges += steps
    coords = append(coords, coord)
  }

  return edges, coords
}

func parseSteps(line string) (Coord, int) {
  fields := strings.Fields(line)
  direction := fields[0]
  steps, _ := strconv.Atoi(fields[1])
  return directions[direction], steps
}

func parseColors(line string) (Coord, int) {
  fields := strings.Fields(line)
  color := fields[2]
  distance, _ := strconv.ParseInt(color[2:7], 16, 64)
  dir := color[7:8]
  direction := ""
  switch dir {
  case "0": direction = RIGHT
  case "1": direction = DOWN
  case "2": direction = LEFT
  case "3": direction = UP
  }
  
  return directions[direction], int(distance)
}

func calculateArea(perimeter int, positions []Coord) int {
	n := len(positions)
	positions = append(positions, Coord{X: positions[0].X, Y: positions[0].Y})
	positions = append(positions, Coord{X: positions[1].X, Y: positions[1].Y})

  area := 0
  for i := 1; i <= n; i++ {
		area += positions[i].Y * (positions[i+1].X - positions[i-1].X)
	}
  
  return area / 2 + perimeter / 2 + 1
}

func (p *Coord) move(direction Coord, count int) {
	x := p.X + (direction.X * count)
  y := p.Y + (direction.Y * count)

	p.X = x
	p.Y = y
}
