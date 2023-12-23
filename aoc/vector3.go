package aoc

type Vector3 struct {
	X int
	Y int
	Z int
}

func NewVector3(x, y, z int) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}
