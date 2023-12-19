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

func read() map[string]workflow {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	workflows := map[string]workflow{}
	for _, line := range lines {
		if line == "" {
			break
		}
		id, w := parseWorkflow(line)
		workflows[id] = w
	}

	return workflows
}

func process(workflows map[string]workflow) int {
	var comb combination
	for i := range comb {
		comb[i] = valuesRange{minimum: minValue, maximum: maxValue}
	}

	return countCombinations(workflows, initialWorkflow, comb)
}

func countCombinations(workflows map[string]workflow, id string, comb combination) int {
	if id == rejected || !comb.valid() {
		return 0
	}

	if id == accepted {
		return comb.total()
	}

	count := 0
	w := workflows[id]
	for _, c := range w.conditions {
		cat := comb.categoryRange(c.category)
		if c.less {
			prevCat := *cat
			cat.maximum = c.value - 1
			count += countCombinations(workflows, c.then, comb)
			*cat = prevCat
			cat.minimum = c.value
			continue
		}

		prevCat := *cat
		cat.minimum = c.value + 1
		count += countCombinations(workflows, c.then, comb)
		*cat = prevCat
		cat.maximum = c.value
	}

	lc := w.lastCondition
	cat := comb.categoryRange(lc.category)
	if lc.less {
		prevCat := *cat
		cat.maximum = lc.value - 1
		count += countCombinations(workflows, lc.then, comb)
		*cat = prevCat

		cat.minimum = lc.value
		count += countCombinations(workflows, lc.otherwise, comb)

		return count
	}

	prevCat := *cat
	cat.minimum = lc.value + 1
	count += countCombinations(workflows, lc.then, comb)
	*cat = prevCat

	cat.maximum = lc.value
	count += countCombinations(workflows, lc.otherwise, comb)

	return count
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

type valuesRange struct {
	minimum int
	maximum int
}

func (r valuesRange) total() int {
	return r.maximum - r.minimum + 1
}

type combination [4]valuesRange

func (c *combination) categoryRange(category byte) *valuesRange {
	switch category {
	case 'x':
		return &c[0]
	case 'm':
		return &c[1]
	case 'a':
		return &c[2]
	case 's':
		return &c[3]
	}

	panic("unexpected category")
}

func (c combination) total() int {
	totalComb := 1
	for _, vr := range c {
		totalComb *= vr.total()
	}

	return totalComb
}

func (c combination) valid() bool {
	for _, r := range c {
		if r.maximum < r.minimum {
			return false
		}
	}

	return true
}

const (
	rejected        = "R"
	accepted        = "A"
	initialWorkflow = "in"
	minValue        = 1
	maxValue        = 4_000
)
