package main

import (
  _ "embed"
  "fmt"
  "strconv"
  "strings"
)

//go:embed input.txt
var input string

var maxCubeNums = map[string]int {
  "red": 12,
  "green": 13,
  "blue": 14,
}

func init() {
  input = strings.TrimRight(input, "\n")
}

func main() {
  part1()
  part2()
}

func part1() {
  sumValidGameIds := 0
  for _, s := range strings.Split(input, "\n") {
    gameId, game := parseGameNumber(s)
    validGame := validateGame(game)
    if validGame {
      sumValidGameIds += gameId
    }
  }
  
  fmt.Println("Part 1:", sumValidGameIds)
}

func part2() {
  sumPowers := 0
  for _, s := range strings.Split(input, "\n") {
    power := 1
    _, game := parseGameNumber(s)
    gameMap := powerGame(game)
    for _, number := range gameMap {
      power *= number
    }
    sumPowers += power
  }
  
  fmt.Println("Part 2:", sumPowers)
}

func powerGame(game string) map[string]int {
  var cubeMap = map[string]int {
    "red": 0,
    "green": 0,
    "blue": 0,
  }
  hands := strings.Split(game, ";")
  for _, hand := range hands {
    hand = strings.TrimSpace(hand)
    handMap := parseHand(hand)
    for color, number := range handMap {
      if number > cubeMap[color] {
        cubeMap[color] = number
      }
    }
  }
  return cubeMap
}

func parseGameNumber(game string) (int, string) {
  s := strings.Split(game, ":")
  gameId := strings.TrimSpace(strings.Split(s[0], " ")[1])
  gameIdNum, _ := strconv.Atoi(gameId)
  return gameIdNum, strings.TrimSpace(s[1])
}

func parseHand(hand string) map[string]int {
  var cubeMap = map[string]int {
    "red": 0,
    "green": 0,
    "blue": 0,
  }
  cubes := strings.Split(hand, ",")
  for _, cube := range cubes {
    cube = strings.TrimSpace(cube)
    number, color := parseCube(cube)
    cubeMap[color] += number
  }
  return cubeMap
}

func parseCube(cube string) (int, string) {
  s := strings.Split(cube, " ")
  color := s[1]
  number, _ := strconv.Atoi(s[0])
  return number, color
}

func validateGame(game string) bool {
  hands := strings.Split(game, ";")
  for _, hand := range hands {
    hand = strings.TrimSpace(hand)
    validHand := validateHand(parseHand(hand))
    if !validHand {
      return false
    }
  }
  return true
}

func validateHand(hand map[string]int) bool {
  switch {
  case hand["red"] > maxCubeNums["red"]:
    return false
  case hand["green"] > maxCubeNums["green"]:
    return false
  case hand["blue"] > maxCubeNums["blue"]:
    return false
  default:
    return true
  }
}
