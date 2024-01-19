package maths

type Vector2 struct {
	X, Y float32
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vector2) Sub(other *Vector2) *Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}

func (v *Vector2) Mul(other *Vector2) *Vector2 {
	v.X *= other.X
	v.Y *= other.Y
	return v
}

func (v *Vector2) MulScalar(scalar float32) *Vector2 {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector2) DivScalar(scalar float32) *Vector2 {
	v.X /= scalar
	v.Y /= scalar
	return v
}
