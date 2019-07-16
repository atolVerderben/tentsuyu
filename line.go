package tentsuyu

import "math"

//Line represents a line from (X1,Y1) to (X2,Y2)
type Line struct {
	X1, Y1, X2, Y2, Angle, Length float64
}

//CalculateAngle returns the line's angle
func (l Line) CalculateAngle() float64 {
	return math.Atan2(l.Y2-l.Y1, l.X2-l.X1)
}

//CalculateLength returns the line's length
func (l Line) CalculateLength() float64 {
	return math.Sqrt(math.Pow(l.X2-l.X1, 2) + math.Pow(l.Y2-l.Y1, 2))
}

//NewLineFromTo returns a Line beginning at (fromX,fromY) to (toX,toY)
func NewLineFromTo(fromX, fromY, toX, toY float64) *Line {
	l := &Line{
		X1: fromX,
		X2: toX,
		Y1: fromY,
		Y2: toY,
	}
	l.Length = l.CalculateLength()
	l.Angle = l.CalculateAngle()

	return l
}
