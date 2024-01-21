package maths

type Bounds struct {
	X, Y int
	W, H int
}

func (b Bounds) IsPointInBounds(x, y int) bool {
	return x >= b.X && x <= b.X+b.W && y >= b.Y && y <= b.Y+b.H
}
