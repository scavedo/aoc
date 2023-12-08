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
var handTypes = map[string]int{
  "5 of a kind": 7,
  "4 of a kind": 6,
  "full house": 5,
  "3 of a kind": 4,
  "2 pair": 3,
  "1 pair": 2,
  "high card": 1,
}

var cards = map[string]int{
  "A": 14,
  "K": 13,
  "Q": 12,
  "J": 11,
  "T": 10,
}

type hands struct {
  hands []hand
}

type hand struct {
  cards string
  bid int
  handType int
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
  total := 0
  hands := hands{}
  for _, line := range lines {
    l := strings.Fields(line)
    cards := l[0]
    bid, _ := strconv.Atoi(l[1])
    hand := hand{cards, bid, 0}
    hand.determineType()
    hands.hands = append(hands.hands, hand)
  }
  hands.sort()
  for i, hand := range hands.hands {
    total += hand.bid * (i + 1)
  }
  fmt.Println("Part 1:", total)
}

func part2() {
  cards["J"] = 1
  total := 0
  hands := hands{}
  for _, line := range lines {
    l := strings.Fields(line)
    cards := l[0]
    bid, _ := strconv.Atoi(l[1])
    hand := hand{cards, bid, 0}
    hand.determineTypePartTwo()
    hands.hands = append(hands.hands, hand)
  }
  hands.sort()
  for i, hand := range hands.hands {
    total += hand.bid * (i + 1)
  }
  fmt.Println("Part 2:", total)
}

func (h *hands) sort() {
  for i := 0; i < len(h.hands); i++ {
    for j := i + 1; j < len(h.hands); j++ {
      if h.hands[i].compare(h.hands[j]) > 0 {
        h.hands[i], h.hands[j] = h.hands[j], h.hands[i]
      }
    }
  }
}

func (h hand) compare(h2 hand) int {
  if h.handType > h2.handType {
    return 1
  } else if h.handType < h2.handType {
    return -1
  } else {
    v1 := 0
    v2 := 0
    for i := 0; i < len(h.cards); i++ {
      c1 := getCardValue(string(h.cards[i]))
      c2 := getCardValue(string(h2.cards[i]))
      v1 += c1
      v2 += c2
      if c1 != c2 {
        break
      }
    }
    return v1 - v2
  }
}

func getCardValue(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil {
    return cards[s]
  }
  return i
}

func (h *hand) determineType() {
  cards := make(map[string]int)
  for _, c := range h.cards {
    cards[string(c)]++
  }
  for _, v := range cards {
    switch v {
    case 5: h.handType = handTypes["5 of a kind"]
    case 4: h.handType = handTypes["4 of a kind"]
    case 3:
      if h.handType == handTypes["1 pair"] {
        h.handType = handTypes["full house"]
      } else {
        h.handType = handTypes["3 of a kind"]
      }
    case 2:
      if h.handType == handTypes["3 of a kind"] {
        h.handType = handTypes["full house"]
      } else if h.handType == handTypes["1 pair"] {
        h.handType = handTypes["2 pair"]
      } else {
        h.handType = handTypes["1 pair"]
      }
    case 1:
      if h.handType == 0 {
        h.handType = handTypes["high card"]
      }
    }
  }
}

func (h *hand) determineTypePartTwo() {
  cards := make(map[string]int)
  jCount := 0
  for _, c := range h.cards {
    if string(c) == "J" {
      jCount++
    } else {
      cards[string(c)]++
    }
  }
  mk, mv := "", 0
  for k, v := range cards {
    if v > mv {
      mk = k
      mv = v
    }
  }
  cards[mk] += jCount
  for _, v := range cards {
    switch v {
    case 5: h.handType = handTypes["5 of a kind"]
    case 4: h.handType = handTypes["4 of a kind"]
    case 3:
      if h.handType == handTypes["1 pair"] {
        h.handType = handTypes["full house"]
      } else {
        h.handType = handTypes["3 of a kind"]
      }
    case 2:
      if h.handType == handTypes["3 of a kind"] {
        h.handType = handTypes["full house"]
      } else if h.handType == handTypes["1 pair"] {
        h.handType = handTypes["2 pair"]
      } else {
        h.handType = handTypes["1 pair"]
      }
    case 1:
      if h.handType == 0 {
        h.handType = handTypes["high card"]
      }
    }
  }
}
