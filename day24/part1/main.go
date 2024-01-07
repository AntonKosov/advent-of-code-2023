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

func read() []ray {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	rays := make([]ray, len(lines))
	for i, row := range lines {
		parts := strings.Split(row, " @ ")
		point := aoc.StrToInts(parts[0])
		velocity := aoc.StrToInts(parts[1])
		rays[i] = ray{
			point:    aoc.NewVector2(point[0], point[1]),
			velocity: aoc.NewVector2(velocity[0], velocity[1]),
		}
	}

	return rays
}

func process(rays []ray) int {
	count := 0
	for i := 0; i < len(rays)-1; i++ {
		ray1 := rays[i]
		for j := i + 1; j < len(rays); j++ {
			if point, ok := ray1.intersects(rays[j]); ok && testingArea(point) {
				count++
			}
		}
	}

	return count
}

type ray struct {
	point    aoc.Vector2[int]
	velocity aoc.Vector2[int]
}

func (r1 ray) intersects(r2 ray) (point aoc.Vector2[float64], ok bool) {
	a1, b1 := r1.coefficientA(), r1.coefficientB()
	a2, b2 := r2.coefficientA(), r2.coefficientB()
	if a1.Equals(a2) {
		return point, false
	}

	point.X = (b2.Float() - b1.Float()) / (a1.Float() - a2.Float())
	point.Y = a1.Float()*point.X + b1.Float()

	return point, r1.pointOnRay(point) && r2.pointOnRay(point)
}

func (r ray) pointOnRay(p aoc.Vector2[float64]) bool {
	// line = start_point + a*direction (any a)
	// ray = start_point + a*direction (a >= 0)
	// if both a are non-negative, the rays intersect
	sp := aoc.NewVector2[float64](float64(r.point.X), float64(r.point.Y))
	offset := p.Sub(sp)

	return aoc.Sign(r.velocity.X) == aoc.Sign(offset.X) && aoc.Sign(r.velocity.Y) == aoc.Sign(offset.Y)
}

func (r ray) coefficientA() aoc.Fraction {
	p1 := r.point
	p2 := p1.Add(r.velocity)

	return aoc.Fraction{
		Numerator:   p1.Y - p2.Y,
		Denominator: p1.X - p2.X,
	}
}

func (r ray) coefficientB() aoc.Fraction {
	p := r.point
	m := r.coefficientA()

	return m.Mul(-p.X).Add(p.Y)
}

func testingArea(p aoc.Vector2[float64]) bool {
	return p.X >= minPosition && p.X <= maxPosition && p.Y >= minPosition && p.Y <= maxPosition
}

const (
	minPosition = 200000000000000
	maxPosition = 400000000000000
)
