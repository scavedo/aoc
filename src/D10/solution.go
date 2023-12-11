package main

import (
  _ "embed"
  "fmt"
  "strings"
)

//go:embed input.txt
var input string
var lines []string

type Coord struct {
  X int
  Y int
}

var pipes = map[rune][]rune{
  '|': []rune{'U', 'D'},
  '-': []rune{'L', 'R'},
  'L': []rune{'U', 'R'},
  'J': []rune{'U', 'L'},
  '7': []rune{'D', 'L'},
  'F': []rune{'D', 'R'},
}

var directions = map[rune]Coord{
  'U': Coord{0, -1},
  'D': Coord{0, 1},
  'L': Coord{-1, 0},
  'R': Coord{1, 0},
}

var inputMap [][]rune

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
  parseMap()
}

func main() {
  part1()
  part2()
}


func part1() {
  var starting Coord
  for r, row := range lines {
    for c, char := range row {
      if char == 'S' {
        starting = Coord{c, r}
      }
    }
  }

  var d []Coord
  for k, dir := range directions {
    x := dir.X
    y := dir.Y
    if starting.X + x < 0 || starting.X + x >= len(inputMap) {
      continue
    }
    if starting.Y + y < 0 || starting.Y + y >= len(inputMap[0]) {
      continue
    }
    symbol := inputMap[starting.Y + y][starting.X + x]
    switch k {
      case 'U':
        if symbol == '|' || symbol == 'F' || symbol == '7' {
          d = append(d, Coord{x, y})
        }
      case 'D':
        if symbol == '|' || symbol == 'L' || symbol == 'J' {
          d = append(d, Coord{x, y})
        }
      case 'L':
        if symbol == '-' || symbol == 'L' || symbol == 'F' {
          d = append(d, Coord{x, y})
        }
      case 'R':
        if symbol == '-' || symbol == 'J' || symbol == '7' {
          d = append(d, Coord{x, y})
        }
    }
  }

  left := starting
  right := starting
  nLeft := Coord{starting.X + d[0].X, starting.Y + d[0].Y}
  nRight := Coord{starting.X + d[1].X, starting.Y + d[1].Y}
  i := 0
  for ok := true; ok; ok = left != right {
    l := nLeft
    r := nRight
    nLeft = travelPath(left, l)
    nRight = travelPath(right, r)
    left = l
    right = r
    i++
  }
  fmt.Println("Part 1:", i)
}

func part2() {
  fmt.Println("Part 2:", 0)
}

func parseMap() {
  inputMap = make([][]rune, len(lines))
  for r, row := range lines {
    inputMap[r] = make([]rune, len(row))
    for c, char := range row {
      inputMap[r][c] = char
    }
  }
}

func travelPath(o Coord, n Coord) Coord {
  symbol := inputMap[n.Y][n.X]
  pipe := pipes[symbol]
  next := Coord{}
  for _, dir := range pipe {
    // fmt.Println(string(dir))
    next = Coord{n.X + directions[dir].X, n.Y + directions[dir].Y}
    if next != o {
      // fmt.Println(o, n, next)
      return next
    }
  }
  return next
}
