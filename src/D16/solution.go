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
var grid [][]Tile

type Direction struct {
  X int
  Y int
}

type Beam struct {
  X int
  Y int
  Direction string
  Split bool
}

type Tile struct {
  Symbol string
  Energized bool
}

type Mirror struct {
  Direction map[string][]string
}

const (
  UP = "UP"
  DOWN = "DOWN"
  LEFT = "LEFT"
  RIGHT = "RIGHT"
)

var mirrors = map[string]Mirror{
  "/": Mirror{Direction: map[string][]string{
    RIGHT: []string{UP},
    DOWN: []string{LEFT},
    UP: []string{RIGHT},
    LEFT: []string{DOWN},
  }},
  "\\": Mirror{Direction: map[string][]string{
    LEFT: []string{UP},
    UP: []string{LEFT},
    RIGHT: []string{DOWN},
    DOWN: []string{RIGHT},
  }},
  "-": Mirror{Direction: map[string][]string{
    LEFT: []string{LEFT},
    RIGHT: []string{RIGHT},
    UP: []string{LEFT, RIGHT},
    DOWN: []string{LEFT, RIGHT},
  }},
  "|": Mirror{Direction: map[string][]string{
    LEFT: []string{UP, DOWN},
    RIGHT: []string{UP, DOWN},
    UP: []string{UP},
    DOWN: []string{DOWN},
  }},
  ".": Mirror{Direction: map[string][]string{
    LEFT: []string{LEFT},
    RIGHT: []string{RIGHT},
    UP: []string{UP},
    DOWN: []string{DOWN},
  }},
}

var directions = map[string]Direction{
  UP: Direction{X: -1, Y: 0},
  DOWN: Direction{X: 1, Y: 0},
  LEFT: Direction{X: 0, Y: -1},
  RIGHT: Direction{X: 0, Y: 1},
}

var startings = []Beam{}
var beams []Beam

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  for _, line := range lines {
    grid = append(grid, []Tile{})
    for _, char := range line {
      grid[len(grid)-1] = append(grid[len(grid)-1], Tile{Symbol: string(char), Energized: false})
    }
  }
}

func main() {
  part1()
  part2()
}

func part1() {
  output := 0
  beams = []Beam{Beam{X: 0, Y: -1, Direction: RIGHT, Split: false}}
  startings = append(startings, beams[0])

  run()
  output = numEnergyzed()
  reset()
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0

  // From Left Side
  for x := 0; x < len(grid); x++ {
    beams = []Beam{Beam{X: x, Y: -1, Direction: RIGHT, Split: false}}
    startings = append(startings, beams[0])
    run()
    num := numEnergyzed()
    if num > output {
      output = num
    }
    reset()
  }

  // From Top Side
  for y := 0; y < len(grid[0]); y++ {
    beams = []Beam{Beam{X: -1, Y: y, Direction: DOWN, Split: false}}
    startings = append(startings, beams[0])
    run()
    num := numEnergyzed()
    if num > output {
      output = num
    }
    reset()
  }

  // From Right Side
  for x := 0; x < len(grid); x++ {
    beams = []Beam{Beam{X: x, Y: len(grid[0]), Direction: LEFT, Split: false}}
    startings = append(startings, beams[0])
    run()
    num := numEnergyzed()
    if num > output {
      output = num
    }
    reset()
  }

  // From Bottom Side
  for y := 0; y < len(grid[0]); y++ {
    beams = []Beam{Beam{X: len(grid), Y: y, Direction: UP, Split: false}}
    startings = append(startings, beams[0])
    run()
    num := numEnergyzed()
    if num > output {
      output = num
    }
    reset()
  }
  
  fmt.Println("Part 2:", output)
}

func run() {
  for i := 0; i < len(beams); i++ {
    beam := beams[i]
    for !beam.willBeOutOfBounds() && !beam.Split {
      beam.move()
    }
  }
}

func numEnergyzed() int {
  output := 0
  for _, row := range grid {
    for _, tile := range row {
      if tile.Energized {
        output++
      }
    }
  }
  return output
}

func reset() {
  beams = []Beam{}
  startings = []Beam{}
  for x := 0; x < len(grid); x++ {
    for y := 0; y < len(grid[0]); y++ {
      grid[x][y].Energized = false
    }
  }
}

func (b Beam) willBeOutOfBounds() bool {
  return b.X + directions[b.Direction].X < 0 ||
    b.X + directions[b.Direction].X >= len(grid) ||
    b.Y + directions[b.Direction].Y < 0 ||
    b.Y + directions[b.Direction].Y >= len(grid[0])
}

func (b *Beam) move() {
  b.X += directions[b.Direction].X
  b.Y += directions[b.Direction].Y
  currentTile := &grid[b.X][b.Y]
  currentTile.Energized = true

  newDirection := mirrors[currentTile.Symbol].Direction[b.Direction]
  if len(newDirection) == 2 {
    b.Split = true
    for _, dir := range newDirection {
      beam := Beam{X: b.X, Y: b.Y, Direction: dir}
      if slices.Contains(startings, beam) {
        continue
      } else {
        startings = append(startings, beam)
        beams = append(beams, beam)
      }
    }
  } else {
    b.Direction = newDirection[0]
  }
}
