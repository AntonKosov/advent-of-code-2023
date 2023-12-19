package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2023/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (map[string]workflow, []part) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	workflows := map[string]workflow{}
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		}
		id, w := parseWorkflow(line)
		workflows[id] = w
	}

	parts := make([]part, len(lines))
	for i, line := range lines {
		parts[i] = parsePart(line)
	}

	return workflows, parts
}

func process(workflows map[string]workflow, parts []part) int {
	sum := 0
	for _, p := range parts {
		if acceptedPart(workflows, p) {
			sum += p.rating()
		}
	}

	return sum
}

func acceptedPart(workflows map[string]workflow, p part) bool {
	currentWorkflow := initialWorkflow
nextWorkflow:
	for {
		if currentWorkflow == rejected {
			return false
		}
		if currentWorkflow == accepted {
			return true
		}

		w := workflows[currentWorkflow]
		for _, c := range w.conditions {
			v := p.categories[c.category]
			if c.less {
				if v < c.value {
					currentWorkflow = c.then
					continue nextWorkflow
				}
			} else {
				if v > c.value {
					currentWorkflow = c.then
					continue nextWorkflow
				}
			}
		}

		v := p.categories[w.lastCondition.category]
		if w.lastCondition.less {
			if v < w.lastCondition.value {
				currentWorkflow = w.lastCondition.then
				continue
			}
			currentWorkflow = w.lastCondition.otherwise
		} else {
			if v > w.lastCondition.value {
				currentWorkflow = w.lastCondition.then
				continue
			}
			currentWorkflow = w.lastCondition.otherwise
		}
	}
}

type condition struct {
	category byte
	less     bool
	value    int
	then     string
}

type lastCondition struct {
	condition
	otherwise string
}

type workflow struct {
	conditions    []condition
	lastCondition lastCondition
}

func parseWorkflow(value string) (string, workflow) {
	parts := strings.Split(value, "{")
	id := parts[0]
	conditionsString := parts[1][:len(parts[1])-1]
	conditionsParts := strings.Split(conditionsString, ",")
	var w workflow
	for i := 0; i < len(conditionsParts)-2; i++ {
		w.conditions = append(w.conditions, parseCondition(conditionsParts[i]))
	}

	w.lastCondition = lastCondition{
		condition: parseCondition(conditionsParts[len(conditionsParts)-2]),
		otherwise: conditionsParts[len(conditionsParts)-1],
	}

	return id, w
}

func parseCondition(value string) condition {
	parts := strings.Split(value, ":")
	return condition{
		category: value[0],
		less:     value[1] == '<',
		value:    aoc.StrToInts(parts[0])[0],
		then:     parts[1],
	}
}

func parsePart(value string) part {
	values := aoc.StrToInts(value)

	return part{
		categories: map[byte]int{
			'x': values[0],
			'm': values[1],
			'a': values[2],
			's': values[3],
		},
	}
}

type part struct {
	categories map[byte]int
}

func (p part) rating() int {
	sum := 0
	for _, v := range p.categories {
		sum += v
	}

	return sum
}

const (
	rejected        = "R"
	accepted        = "A"
	initialWorkflow = "in"
)
