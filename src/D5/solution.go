package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var lines []string

var almanac map[string]Mapper

type Range struct {
  sourceStart int
  destinationStart int
  length int
}

type Mapper struct {
  source string
  destination string
  ranges []Range
  correspondance map[int]int
}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
}

func main() {
  almanac = parseAlmanac(lines[2:])
  part1()
  part2()
}


func part1() {
  lowest := math.MaxInt64
  seeds := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
  for _, x := range seeds {
    seedInt, _ := strconv.Atoi(x)
    result := almanac["seed"]
    destNum := getDestNum(seedInt, result)
    n := drillDown(destNum, result.destination)
    if n < lowest {
      lowest = n
    }
  }
  fmt.Println("Part 1:", lowest)
}

func part2() {
  lowest := math.MaxInt64
  seeds := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
  seedsRanges := [][]int{}
  i := 0
  for j := 0; j < len(seeds); j++ {
    seedInt, _ := strconv.Atoi(seeds[j])
    if j % 2 != 0 {
      seedsRanges = append(seedsRanges, []int{i, seedInt})
    } else {
      i = seedInt
    }
  }

  for _, x := range seedsRanges {
    for i := x[0]; i <= x[0] + x[1]; i++ {
      result := almanac["seed"]
      destNum := getDestNum(i, result)
      n := drillDown(destNum, result.destination)
      if n < lowest {
        lowest = n
      }
    }
  }
    
  fmt.Println("Part 2:", lowest)
}

func drillDown(n int, source string) int {
  result, ok := almanac[source]
  if !ok {
    return n
  }
  destNum := getDestNum(n, result)
  return drillDown(destNum, result.destination)
}

func parseSeeds(in string) []string {
  return strings.Split(strings.TrimSpace(strings.Split(in, ":")[1]), " ")
}

func parseAlmanac(in []string) map[string]Mapper {
  almanac := map[string]Mapper{}
  mapper := Mapper{}
  for _, line := range in {
    if line == "" {
      almanac[mapper.source] = mapper
      mapper = Mapper{}
      continue
    } else {
      splits := strings.Split(line, "map:")
      if len(splits) == 2 {
        key := strings.TrimSpace(splits[0])
        mapper.addSourceAndDestination(key)
      } else {
        mapper.addRange(line)
      }
    }

    if !mapper.empty() {
      almanac[mapper.source] = mapper
    }
      
  }
  return almanac
}

func getDestNum(n int, m Mapper) int {
  for _, r := range m.ranges {
    if (n >= r.sourceStart && n <= r.sourceStart + r.length) {
      return r.destinationStart + (n - r.sourceStart)
    }
  }
  return n
}

func (m Mapper) empty() bool {
  return m.source == "" && m.destination == "" && len(m.ranges) == 0
}

func (m *Mapper) addSourceAndDestination(in string) {
  splits := strings.Split(in, "-to-")
  source := strings.TrimSpace(splits[0])
  destination := strings.TrimSpace(splits[1])
  m.source = source
  m.destination = destination
}

func (m *Mapper) addRange(in string) {
  splits := strings.Split(in, " ")
  sourceStart, _ := strconv.Atoi(splits[1])
  destinationStart, _ := strconv.Atoi(splits[0])
  length, _ := strconv.Atoi(splits[2])
  m.ranges = append(m.ranges, Range{sourceStart, destinationStart, length})
}

func (m *Mapper) setCorrespondance(source, destination int) {
  m.correspondance[source] = destination
}
