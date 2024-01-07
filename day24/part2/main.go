package main

import (
	"fmt"
	"math/big"
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
			point:    aoc.NewVector3(point[0], point[1], point[2]),
			velocity: aoc.NewVector3(velocity[0], velocity[1], velocity[2]),
		}
	}

	return rays
}

func process(rays []ray) *big.Int {
	/*
		Let:
			p0 - the vector of the start position of the stone
			v0 - the vector of the velocity of the stone
			pi - the vector of the start position of the i-1 hailstone
			vi - the vector of the velocity of the i-1 hailstone
			ti - the time until the intersection of the i-1 hailstone

		p0+ti*v0 = pi+ti*vi
		p0-pi = ti*(vi-v0)

		In order to intersect, p0-pi and vi-v0 must be parallel. So, their cross product must be
		equal to 0. ti may be excluded as a scalar.

		(p0-pi)x(vi-v0) = 0

		Comparing two different hailstones (i and j):

		(p0-pi) x (vi-v0) = (p0-pj) x (vj-v0)
		(p0-pi) x vi + (p0-pi) x (-v0) = (p0-pj) x vj + (p0-pj) x (-v0)
		p0 x vi + (-pi) x vi + p0 x (-v0) + (-pi) x (-v0) = p0 x vj + (-pj) x vj + p0 x (-v0) + (-pj) x (-v0)
		p0 x vi + (-pi) x vi + + (-pi) x (-v0) = p0 x vj + (-pj) x vj + + (-pj) x (-v0)

		Comparing the stone with hailstones 0-1 and 0-2 and solving the system for x, y, and z coordinates,
		we get 6 linear equations with 6 unknown values - coordinates of p0 and v0.
	*/
	ray1, ray2, ray3 := rays[0], rays[1], rays[2]
	equations := []equation{
		newXEquation(ray1).sub(newXEquation(ray2)),
		newXEquation(ray1).sub(newXEquation(ray3)),
		newYEquation(ray1).sub(newYEquation(ray2)),
		newYEquation(ray1).sub(newYEquation(ray3)),
		newZEquation(ray1).sub(newZEquation(ray2)),
		newZEquation(ray1).sub(newZEquation(ray3)),
	}
	solution := solveSystem(equations)

	sum := bigInt(0)
	for i := 0; i < 3; i++ {
		sum = sum.Add(sum, solution[i])
	}

	return sum
}

const (
	p0x int = iota
	p0y
	p0z
	v0x
	v0y
	v0z
	value
)

func bigInt(v int) *big.Int {
	return big.NewInt(int64(v))
}

func bigIntPiVi(a, b, c, d int) *big.Int {
	p0 := bigInt(0)
	p1 := bigInt(0)

	p0 = p0.Mul(bigInt(a), bigInt(b))
	p1 = p1.Mul(bigInt(c), bigInt(d))

	return p0.Sub(p0, p1)
}

// p0x, p0y, p0z, v0x, v0y, v0z, value
type equation [7]*big.Int

func newEquation() equation {
	var e equation
	for i := range e {
		e[i] = bigInt(0)
	}

	return e
}

func newXEquation(r ray) equation {
	pi, vi := r.point, r.velocity

	eq := newEquation()
	eq[p0y] = bigInt(-vi.Z)
	eq[p0z] = bigInt(vi.Y)
	eq[v0y] = bigInt(pi.Z)
	eq[v0z] = bigInt(-pi.Y)
	eq[value] = bigIntPiVi(pi.Z, vi.Y, pi.Y, vi.Z)

	return eq
}

func newYEquation(r ray) equation {
	pi, vi := r.point, r.velocity

	eq := newEquation()
	eq[p0x] = bigInt(vi.Z)
	eq[p0z] = bigInt(-vi.X)
	eq[v0x] = bigInt(-pi.Z)
	eq[v0z] = bigInt(pi.X)
	eq[value] = bigIntPiVi(pi.X, vi.Z, pi.Z, vi.X)

	return eq
}

func newZEquation(r ray) equation {
	pi, vi := r.point, r.velocity

	eq := newEquation()
	eq[p0x] = bigInt(-vi.Y)
	eq[p0y] = bigInt(vi.X)
	eq[v0x] = bigInt(pi.Y)
	eq[v0y] = bigInt(-pi.X)
	eq[value] = bigIntPiVi(pi.Y, vi.X, pi.X, vi.Y)

	return eq
}

func (e equation) sub(e2 equation) equation {
	var res equation
	for i, v := range e {
		res[i] = bigInt(0).Sub(v, e2[i])
	}

	return res
}

func (e equation) mul(scalar *big.Int) equation {
	var res equation
	for i := range e {
		res[i] = bigInt(0).Mul(e[i], scalar)
	}

	return res
}

func (e equation) div(scalar *big.Int) equation {
	var res equation
	for i := range e {
		res[i] = bigInt(0).Div(e[i], scalar)
	}

	return res
}

func (e equation) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t=\t%v", e[0], e[1], e[2], e[3], e[4], e[5], e[6])
}

func solveSystem(system []equation) []*big.Int {
	gauss(system)
	solution := make([]*big.Int, len(system))
	for i := range system {
		solution[i] = system[i][value]
	}

	return solution
}

func gauss(system []equation) {
	for i := 0; i < len(system)-1; i++ {
		rearrangeEquations(system, i)
		eq := system[i]
		currentValue := eq[i]
		for j := i + 1; j < len(system); j++ {
			v := system[j][i]
			if v.Cmp(bigInt(0)) == 0 {
				continue
			}
			system[j] = system[j].mul(currentValue).sub(eq.mul(v))
		}
	}

	for i := len(system) - 1; i >= 0; i-- {
		system[i] = system[i].div(system[i][i])
		for j := 0; j < i; j++ {
			system[j] = system[j].sub(system[i].mul(system[j][i]))
		}
	}
}

func rearrangeEquations(system []equation, idx int) {
	if system[idx][idx].Cmp(bigInt(0)) != 0 {
		return
	}

	for i := idx + 1; i < len(system); i++ {
		if system[i][idx].Cmp(bigInt(0)) != 0 {
			system[idx], system[i] = system[i], system[idx]
			return
		}
	}

	panic("unexpected system state")
}

type ray struct {
	point    aoc.Vector3
	velocity aoc.Vector3
}
