package aoc

import "fmt"

type Fraction struct {
	Numerator   int
	Denominator int
}

func NewFraction(numerator, denominator int) Fraction {
	return Fraction{
		Numerator:   numerator,
		Denominator: denominator,
	}
}

func (f1 Fraction) Equals(f2 Fraction) bool {
	if f1.Denominator == 0 || f2.Denominator == 0 {
		panic("a denominator is 0")
	}

	if f1.Numerator == 0 && f2.Numerator == 0 {
		return true
	}

	if f1.Numerator == f2.Numerator && f1.Denominator == f2.Denominator {
		return true
	}

	f1, f2 = f1.Simplify(), f2.Simplify()

	return f1.Numerator == f2.Numerator && f1.Denominator == f2.Denominator
}

func (f Fraction) Simplify() Fraction {
	gcd := GCD(f.Numerator, f.Denominator)

	return Fraction{
		Numerator:   f.Numerator / gcd,
		Denominator: f.Denominator / gcd,
	}
}

func (f Fraction) Mul(s int) Fraction {
	f.Numerator *= s

	return f.Simplify()
}

func (f Fraction) Add(s int) Fraction {
	f.Numerator += s * f.Denominator

	return f.Simplify()
}

func (f Fraction) Sub(s int) Fraction {
	return f.Add(-s)
}

func (f Fraction) Div(f2 Fraction) Fraction {
	f.Numerator *= f2.Denominator
	f.Denominator *= f2.Numerator

	return f.Simplify()
}

func (f Fraction) Float() float64 {
	return float64(f.Numerator) / float64(f.Denominator)
}

func (f Fraction) String() string {
	return fmt.Sprintf("%v/%v", f.Numerator, f.Denominator)
}
