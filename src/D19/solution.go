package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var parts []Part

var workflows map[string]Workflow = make(map[string]Workflow)

type Part struct {
  X int
  M int
  A int
  S int
  Workflow string
}

type Workflow struct {
  Rules map[int]Rule
}

type Rule struct {
  Attribute string
  Comparison string
  Value int
  NextWorkflow string
}

var accepted = []Part{}

func init() {
  input = strings.TrimRight(input, "\n")
  sections := strings.Split(input, "\n\n")
  for _, workflow := range strings.Split(sections[0], "\n") {
    n, w := parseWorkflow(workflow)
    workflows[n] = w
  }
  w := workflows["in"]
  for _, part := range strings.Split(sections[1], "\n") {
    p := parsePart(part)
    p.Workflow = "in"
    parts = append(parts, p)
  }
  workflows["in"] = w
}

func main() {
  part1()
  part2()
}

func part1() {
  output := 0

  for _, part := range parts {
    part.Process()
  }

  for _, part := range accepted {
    output += part.X
    output += part.M
    output += part.A
    output += part.S
  }
  
  fmt.Println("Part 1:", output)
}

func part2() {
  output := 0
  fmt.Println("Part 2:", output)
}

func parseWorkflow(in string) (string, Workflow) {
  workflow := Workflow{}
  workflow.Rules = make(map[int]Rule)
  in = strings.TrimRight(in, "}")
  sections := strings.Split(in, "{")
  name := sections[0]
  rules := strings.Split(sections[1], ",")
  for i, rule := range rules {
    r := parseRule(rule)
    workflow.Rules[i] = r
  }
  return name, workflow
}

func parseRule(in string) Rule {
  r := Rule{}
  if strings.Contains(in, ":") {
    idx := strings.Index(in, ":")
    r.Attribute = in[:1]
    r.Comparison = in[1:2]
    r.Value, _ = strconv.Atoi(in[2:idx])
    r.NextWorkflow = in[idx+1:]
  } else {
    r.NextWorkflow = in
  }
  return r
}

func parsePart(in string) Part {
  part := Part{}
  in = strings.TrimRight(in, "}")
  in = strings.TrimLeft(in, "{")
  for _, attribute := range strings.Split(in, ",") {
    if strings.Contains(attribute, "x") {
      part.X,_ = strconv.Atoi(attribute[2:])
    } else if strings.Contains(attribute, "m") {
      part.M,_ = strconv.Atoi(attribute[2:])
    } else if strings.Contains(attribute, "a") {
      part.A,_ = strconv.Atoi(attribute[2:])
    } else if strings.Contains(attribute, "s") {
      part.S,_ = strconv.Atoi(attribute[2:])
    }
  }
  return part
}

func (p *Part) Process() {
  for p.Workflow != "A" && p.Workflow != "R" {
    workflow := workflows[p.Workflow]
    nextWorkflow := ""
    for i := 0; i < len(workflow.Rules); i++ {
      rule := workflow.Rules[i]
      if rule.Attribute == "" {
        nextWorkflow = rule.NextWorkflow
        break
      } else {
        v := reflect.ValueOf(p).Elem()
        val := v.FieldByName(strings.ToUpper(rule.Attribute))
        if rule.Comparison == "<" {
          if val.Int() < int64(rule.Value) {
            nextWorkflow = rule.NextWorkflow
            break
          }
        } else if rule.Comparison == ">" {
          if val.Int() > int64(rule.Value) {
            nextWorkflow = rule.NextWorkflow
            break
          }
        }
      }
    }
    if nextWorkflow == "A" {
      accepted = append(accepted, *p)
    }

    p.Workflow = nextWorkflow
  }
}
