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

func read() map[string]module {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	modules := make(map[string]module, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		targets := strings.Split(parts[1], ", ")
		name := parts[0]
		switch {
		case name == broadcasterID:
			modules[broadcasterID] = broadcastModule{targets: targets}
		case name[0] == '%':
			modules[name[1:]] = &flipFlopModule{name: name[1:], targets: targets}
		case name[0] == '&':
			modules[name[1:]] = &conjuctionModule{name: name[1:], targets: targets}
		default:
			panic("unknown module type")
		}
	}

	if len(modules) != modulesCount {
		panic("unexpected number of modules")
	}

	return modules
}

// &lv -> rx

// all must become high
// &st -> lv
// &tn -> lv
// &hh -> lv
// &dt -> lv

func process(modules map[string]module) int {
	// All incoming to lv must be high. There are 4 sources:
	// &st -> lv
	// &tn -> lv
	// &hh -> lv
	// &dt -> lv
	// Every device sends high signal with the same rate. The rate is different for every device.
	initConjunctionInputs(modules)
	highSignals := map[string]*int{}

	for i := 0; len(highSignals) < 4; i++ {
		pulses := []pulse{{level: 0, source: "button", target: broadcasterID}}
		for len(pulses) > 0 {
			p := pulses[0]
			pulses = pulses[1:]
			if p.target == "lv" && p.level == 1 && highSignals[p.source] == nil {
				rate := i + 1
				highSignals[p.source] = &rate
			}
			if m, ok := modules[p.target]; ok {
				nextPulses := m.handleIncomingPulse(p.source, p.level)
				pulses = append(pulses, nextPulses...)
			} else if p.level == 0 {
				return i + 1
			}
		}
	}

	count := 1
	for _, v := range highSignals {
		count *= *v
	}

	return count
}

func initConjunctionInputs(modules map[string]module) {
	conjunctions := map[string]*conjuctionModule{}
	for name, m := range modules {
		if cm, ok := m.(*conjuctionModule); ok {
			conjunctions[name] = cm
		}
	}

	for name, m := range modules {
		for _, targetName := range m.getTargets() {
			if cm, ok := conjunctions[targetName]; ok {
				cm.addSource(name)
				break
			}
		}
	}
}

type module interface {
	getTargets() []string
	handleIncomingPulse(source string, level int) []pulse
}

type broadcastModule struct {
	targets []string
}

func (bm broadcastModule) getTargets() []string {
	return bm.targets
}

func (bm broadcastModule) handleIncomingPulse(_ string, _ int) []pulse {
	output := make([]pulse, len(bm.targets))
	for i, name := range bm.targets {
		output[i] = pulse{level: 0, source: broadcasterID, target: name}
	}

	return output
}

type flipFlopModule struct {
	name    string
	on      bool
	targets []string
}

func (ffm *flipFlopModule) getTargets() []string {
	return ffm.targets
}

func (ffm *flipFlopModule) handleIncomingPulse(_ string, level int) []pulse {
	if level > 0 {
		return nil
	}

	ffm.on = !ffm.on

	outputLevel := 0
	if ffm.on {
		outputLevel = 1
	}

	output := make([]pulse, len(ffm.targets))
	for i, name := range ffm.targets {
		output[i] = pulse{level: outputLevel, source: ffm.name, target: name}
	}

	return output
}

type conjuctionModule struct {
	name      string
	countHigh int
	sources   map[string]int
	targets   []string
}

func (cm *conjuctionModule) getTargets() []string {
	return cm.targets
}

func (cm *conjuctionModule) addSource(name string) {
	if cm.sources == nil {
		cm.sources = map[string]int{}
	}
	cm.sources[name] = 0
}

func (cm *conjuctionModule) handleIncomingPulse(source string, level int) []pulse {
	if level == 0 {
		if cm.sources[source] > 0 {
			cm.sources[source] = 0
			cm.countHigh--
		}
	} else if cm.sources[source] == 0 {
		cm.sources[source] = 1
		cm.countHigh++
	}

	outputValue := 0
	if cm.countHigh < len(cm.sources) {
		outputValue = 1
	}

	output := make([]pulse, len(cm.targets))
	for i, name := range cm.targets {
		output[i] = pulse{level: outputValue, source: cm.name, target: name}
	}

	return output
}

type pulse struct {
	level  int
	source string
	target string
}

const (
	broadcasterID = "broadcaster"
	modulesCount  = 58
)
