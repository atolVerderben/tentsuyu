package tentsuyu

import (
	"fmt"
	"math"
)

type Vector2d struct {
	X, Y float64
}

func (v *Vector2d) Add(other Vector2d) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vector2d) Sub(other Vector2d) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vector2d) Mul(other float64) {
	v.X *= other
	v.Y *= other
}

func (v *Vector2d) Div(other float64) {
	v.X /= other
	v.Y /= other
}

//Limit the vector to the given float
func (v *Vector2d) Limit(limit float64) {

	if v.Length() > limit {
		v.Normalize()
		v.Mul(limit)
	}

	/*if math.Abs(v.X) > limit {
		if v.X < 0 {
			v.X = -limit
		}
		v.X = limit
	}

	if math.Abs(v.Y) > limit {
		if v.Y < 0 {
			v.Y = -limit
		}
		v.Y = limit
	}*/
}

func (v Vector2d) Dot(other Vector2d) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2d) Cross(other Vector2d) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vector2d) Crossf(other float64) Vector2d {
	return Vector2d{-v.Y * other, v.X * other}
}

func (v Vector2d) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector2d) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vector2d) Normalize() {
	v.Mul(1 / v.Length())
}

func (v Vector2d) Normalized() *Vector2d {
	return VectorMul(v, 1/v.Length())
}

func (v Vector2d) Plus(v2 Vector2d) *Vector2d {
	return VectorAdd(v, v2)
}

func (v Vector2d) Minus(v2 Vector2d) *Vector2d {
	return VectorSub(v, v2)
}

func (v Vector2d) Times(r float64) *Vector2d {
	return VectorMul(v, r)
}

func VectorAdd(v, u Vector2d) *Vector2d {
	return &Vector2d{v.X + u.X, v.Y + u.Y}
}

func VectorSub(v, u Vector2d) *Vector2d {
	return &Vector2d{v.X - u.X, v.Y - u.Y}
}

func VectorMul(v Vector2d, r float64) *Vector2d {
	return &Vector2d{v.X * r, v.Y * r}
}

//ToString nicely formats the coords in written form
func (v Vector2d) ToString() string {
	return fmt.Sprintf("{X:%f, Y:%f}", v.X, v.Y)
}
