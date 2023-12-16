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

type Contents struct {
  Key string
  Value int
  Slot int
}

type Box struct {
  Key int
  Contents []Contents
}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, ",")
}

func main() {
  part1()
  part2()
}

func part1() {
  output := 0

  for _, line := range lines {
    output += hash(line)
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  boxes := []Box{}
  for i := 0; i < 256; i++ {
    boxes = append(boxes, Box{Key: i})
  }

  for _, line := range lines {
    if strings.Contains(line, "=") {
      split := strings.Split(line, "=")
      box := hash(split[0])
      value, _ := strconv.Atoi(split[1])
      contents := Contents{Key: split[0], Value: value}
      boxes[box].AddContents(contents)
    } else if strings.Contains(line, "-") {
      split := strings.Split(line, "-")
      box := hash(split[0])
      boxes[box].RemoveContents(split[0])
    }
  }

  for i := len(boxes) - 1; i >= 0; i-- {
    if (len(boxes[i].Contents) == 0) {
      boxes = append(boxes[:i], boxes[i+1:]...)
    }
  }

  for _, box := range boxes {
    for j, contents := range box.Contents {
      x := 1 + box.Key
      y := j + 1
      z := contents.Value
      output += x * y * z
    }
  }
  
  fmt.Println("Part 2:", output)
}

func hash(input string) int {
  x := 0
  for _, char := range input {
    x += int(char)
    x *= 17
    x %= 256
  }
  return x
}

func (b *Box) AddContents(contents Contents) {
  if b.Contains(contents.Key) {
    b.ReplaceContents(contents.Key, contents)
  } else {
    contents.Slot = len(b.Contents) + 1
    b.Contents = append(b.Contents, contents)
  }
}

func (b Box) Contains(key string) bool {
  for _, contents := range b.Contents {
    if contents.Key == key {
      return true
    }
  }
  return false
}

func (b *Box) ReplaceContents(key string, c Contents) {
  for i, contents := range b.Contents {
    if contents.Key == key {
      b.Contents[i].Value = c.Value
    }
  }
}

func (b *Box) RemoveContents(key string) {
  for i, contents := range b.Contents {
    if contents.Key == key {
      b.Contents = append(b.Contents[:i], b.Contents[i+1:]...)
    } else if contents.Slot > i {
      contents.Slot--
    }
  }
}
