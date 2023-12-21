package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

//go:embed input.txt
var input string
var lines []string
var modules map[string]Module = make(map[string]Module)
var queue Queue = Queue{[]Signal{}}

const (
  LOW = 0
  HIGH = 1
)

type Queue struct {
  Items []Signal
}

type Signal struct {
  Value int
  Source string
  Destination string
}

type Module interface {
  getName() string
  addDestination(string)
  getDestinations() []string
  process(Signal)
}

type Basic struct {
  Name string
  Destinations []string
}

type FlipFlop struct {
  Name string
  On bool
  Destinations []string
  
}

type Conjunction struct {
  Name string
  Memory map[string]int
  Destinations []string
}

func (s Signal) process() {
  m, ok := modules[s.Destination]
  if ok {
    m.process(s)
  }
}

func (s Signal) String() string {
  return fmt.Sprintf("%s -> %d -> %s", s.Source, s.Value, s.Destination)
}

func (q *Queue) push(s Signal) {
  q.Items = append(q.Items, s)
}

func (q *Queue) pop() Signal {
  if q.isEmpty() {
    return Signal{}
  }
  item := q.Items[0]
  q.Items = q.Items[1:]
  return item
}

func (q Queue) peek() Signal {
  if q.isEmpty() {
    return Signal{}
  }
  return q.Items[0]
}

func (q Queue) isEmpty() bool {
  return len(q.Items) == 0
}

func newModule(in string) Module {
  var m Module
  parts := strings.Split(in, " -> ")
  namePart := parts[0]
  if strings.HasPrefix(namePart, "%") {
    m = newFlipFlop(namePart[1:])
  } else if strings.HasPrefix(namePart, "&") {
    m = newConjunction(namePart[1:])
  } else {
    m = newBasic(namePart)
  }
  destinations := strings.Split(parts[1], ", ")
  for _, d := range destinations {
    m.addDestination(d)
  }
  return m
}

func newBasic(name string) *Basic {
  return &Basic{name, []string{}}
}

func (b *Basic) getName() string {
  return b.Name
}

func (b *Basic) addDestination(d string) {
  b.Destinations = append(b.Destinations, d)
}

func (b *Basic) getDestinations() []string {
  return b.Destinations
}

func (b *Basic) process(signal Signal) {
  for _, d := range b.Destinations {
    signal := Signal{LOW, b.Name, d}
    queue.push(signal)
  }
}

func newFlipFlop(name string) *FlipFlop {
  return &FlipFlop{name, false, []string{}}
}

func (f *FlipFlop) getName() string {
  return f.Name
}

func (f *FlipFlop) addDestination(d string) {
  f.Destinations = append(f.Destinations, d)
}

func (f *FlipFlop) getDestinations() []string {
  return f.Destinations
}

func (f *FlipFlop) process(signal Signal) {
  pulse := signal.Value
  if pulse == LOW {
    outPulse := HIGH
    if f.On {
      f.On = false
      outPulse = LOW
    } else {
      f.On = true
    }
    for _, d := range f.Destinations {
      signal := Signal{outPulse, f.Name, d}
      queue.push(signal)
    }
  }

  modules[f.Name] = f
}

func newConjunction(name string) *Conjunction {
  return &Conjunction{name, map[string]int{}, []string{}}
}

func (c *Conjunction) getName() string {
  return c.Name
}

func (c *Conjunction) addDestination(d string) {
  c.Destinations = append(c.Destinations, d)
}

func (c *Conjunction) getDestinations() []string {
  return c.Destinations
}

func (c *Conjunction) process(signal Signal) {
  c.Memory[signal.Source] = signal.Value
  allHigh := true
  for _, m := range c.Memory {
    if m == LOW {
      allHigh = false
    }
  }
  for _, d := range c.Destinations {
    signal := Signal{HIGH, c.Name, d}
    if allHigh {
      signal.Value = LOW
    }
    queue.push(signal)
  }
  
  modules[c.Name] = c
}

func init() {
  input = strings.TrimRight(input, "\n")
  lines = strings.Split(input, "\n")
}

func reset() {
  for _, line := range lines {
    m := newModule(line)
    modules[m.getName()] = m
  }
  for _, m := range modules {
    if reflect.TypeOf(m) == reflect.TypeOf(&Conjunction{}) {
      for _, n := range modules {
        if slices.Contains(n.getDestinations(), m.getName()) {
          (m.(*Conjunction)).Memory[n.getName()] = LOW
        }
      }
    }
  }
}

func main() {
  part1()
  part2()
}

func part1() {
  output := 0
  low := 0
  high := 0
  reset()
  
  for i := 0; i < 1000; i++ {
    pushButton()
    for !queue.isEmpty() {
      m := queue.pop()
      if m.Value == HIGH {
        high++
      } else {
        low++
      }
      m.process()
    }
  }

  output = low * high
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  reset()

  var feed Module = nil
  for _, m := range modules {
    for _, d := range m.getDestinations() {
      if d == "rx" {
        feed = m
      }
    }
  }

  var cycleLengths map[string]int = make(map[string]int)
  var seen map[string]int = make(map[string]int)
  for _, m := range modules {
    for _, d := range m.getDestinations() {
      if d == feed.getName() {
        seen[m.getName()] = 0
      }
    }
  }

  presses := 0
  looking := true
  for looking {
    pushButton()
    presses++
    for !queue.isEmpty() {
      m := queue.pop()
      if m.Destination == feed.getName() && m.Value == HIGH {
        seen[m.Source] = seen[m.Source] + 1
        _, ok := cycleLengths[m.Source]
        if !ok {
          cycleLengths[m.Source] = presses
        }
      }

      allSeen := true
      for _, s := range seen {
        if s == 0 {
          allSeen = false
        }
      }
      if allSeen {
        x := 1
        for _, c := range cycleLengths {
          x = LCM(x, c)
        }
        output = x
        looking = false
        break
      }
      m.process()
    }
  }
  
  fmt.Println("Part 2:", output)
}

func pushButton() {
  signal := Signal{LOW, "button", "broadcaster"}
  queue.push(signal)
}

func GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

func LCM(a, b int, integers ...int) int {
  result := a * b / GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}
