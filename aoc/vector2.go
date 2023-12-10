package aoc

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) Add(av Vector2) Vector2 {
	return NewVector2(v.X+av.X, v.Y+av.Y)
}

func (v Vector2) Sub(av Vector2) Vector2 {
	return NewVector2(v.X-av.X, v.Y-av.Y)
}

func (v Vector2) Mul(scalar int) Vector2 {
	return NewVector2(v.X*scalar, v.Y*scalar)
}

// RotateLeft rotates the vector to the left (left-handed system)
func (v Vector2) RotateLeft() Vector2 {
	return NewVector2(v.Y, -v.X)
}

// RotateRight rotates the vector to the right (left-handed system)
func (v Vector2) RotateRight() Vector2 {
	return NewVector2(-v.Y, v.X)
}
