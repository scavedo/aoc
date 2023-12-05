package main

import (
  _ "embed"
  "fmt"
  "sort"
  "strings"
  "strconv"
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
    numbers := strings.Split(strings.Split(line, ":")[1], "|")
    winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
    current := strings.Split(strings.TrimSpace(numbers[1]), " ")
    lineTotal := 0
    for _, curr := range current {
      if curr == "" {
        continue
      }
      hasWon := false
      for _, winner := range winning {
        if winner == "" {
          continue
        }
        if winner == curr {
          if lineTotal == 0 {
            lineTotal = 1
          } else {
            lineTotal *= 2
          }
          hasWon = true
          break
        }
        if hasWon {
          break
        }
      }
    }
    total += lineTotal
  }
  fmt.Println("Part 1:", total)
}

func part2() {
  total := 0
  cardMap := make(map[int][]string)
  for _, line := range lines {
    cardData := strings.Split(line, ":")
    card := cardData[0]
    x := strings.Split(card, " ")
    cardNumber := x[len(x)-1]
    cn, _ := strconv.Atoi(cardNumber)
    cardMap[cn] = []string{cardData[1]}
  }

  keys := make([]int, 0, len(cardMap))
  for k := range cardMap {
    keys = append(keys, k)
  }
  sort.Ints(keys)
  for _, k := range keys {
    v := cardMap[k]
    for _, line := range v {
      total += 1
      numbers := strings.Split(line, "|")
      winning := strings.Split(strings.TrimSpace(numbers[0]), " ")
      current := strings.Split(strings.TrimSpace(numbers[1]), " ")
      lineTotal := 0
      for _, curr := range current {
        if curr == "" {
          continue
        }
        hasWon := false
        for _, winner := range winning {
          if winner == "" {
            continue
          }
          if winner == curr {
            lineTotal += 1
            hasWon = true
            break
          }
          if hasWon {
            break
          }
        }
      }

      for i := 1; i <= lineTotal; i++ {
        x := k + i
        if x > len(keys) {
          continue
        }
          
        cardMap[x] = append(cardMap[x], cardMap[x][0])
      }
    }
  }
  
  fmt.Println("Part 2:", total)
}
