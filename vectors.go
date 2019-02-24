package tentsuyu

import (
	"fmt"
	"math"
)

//Vector2d represents a 2 dimensional vector
type Vector2d struct {
	X, Y float64
}

//Add other vector to the vector
func (v *Vector2d) Add(other Vector2d) {
	v.X += other.X
	v.Y += other.Y
}

//Sub (tract) other vector from the vector
func (v *Vector2d) Sub(other Vector2d) {
	v.X -= other.X
	v.Y -= other.Y
}

//Subf subtract float from the vector
func (v *Vector2d) Subf(number float64) {
	v.X -= number
	v.Y -= number
}

//Mul (tiply) the other vector with the vector
func (v *Vector2d) Mul(other float64) {
	v.X *= other
	v.Y *= other
}

//Div (ide) the vector by the other vector
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
}

//Dot product of vector with other vector
func (v Vector2d) Dot(other Vector2d) float64 {
	return v.X*other.X + v.Y*other.Y
}

//Cross Product of two 2D Vectors
func (v Vector2d) Cross(other Vector2d) float64 {
	return v.X*other.Y - v.Y*other.X
}

//Crossf is cross of Vector and float64
func (v Vector2d) Crossf(other float64) Vector2d {
	return Vector2d{-v.Y * other, v.X * other}
}

//LengthSquared returns the magnitude before square root function
func (v Vector2d) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

//Length returns the vector magnitude
func (v Vector2d) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

//Normalize the vector by making it equal to magnitude of 1
func (v *Vector2d) Normalize() {
	if v.Length() == 0 {
		return
	}
	v.Mul(1 / v.Length())
}

//Normalized returns a new normalized Vector2d
func (v Vector2d) Normalized() *Vector2d {
	if v.Length() == 0 {
		return &Vector2d{}
	}
	return VectorMul(v, 1/v.Length())
}

//Plus returns a new vector of the vector plus v2
func (v Vector2d) Plus(v2 Vector2d) *Vector2d {
	return VectorAdd(v, v2)
}

//Minus returns a new vector of the vector minus v2
func (v Vector2d) Minus(v2 Vector2d) *Vector2d {
	return VectorSub(v, v2)
}

//Times returns a new vector of the vectors multiplied together
func (v Vector2d) Times(r float64) *Vector2d {
	return VectorMul(v, r)
}

//VectorAdd adds vectors v and u
func VectorAdd(v, u Vector2d) *Vector2d {
	return &Vector2d{v.X + u.X, v.Y + u.Y}
}

//VectorSub subtracts vectors v and u
func VectorSub(v, u Vector2d) *Vector2d {
	return &Vector2d{v.X - u.X, v.Y - u.Y}
}

//VectorMul multiples vector v by float r
func VectorMul(v Vector2d, r float64) *Vector2d {
	return &Vector2d{v.X * r, v.Y * r}
}

//ToString nicely formats the coords in written form
func (v Vector2d) ToString() string {
	return fmt.Sprintf("{X:%f, Y:%f}", v.X, v.Y)
}
