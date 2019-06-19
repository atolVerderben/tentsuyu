package tentsuyu

import "math"

//Line represents a line from (X1,Y1) to (X2,Y2)
type Line struct {
	X1, Y1, X2, Y2, Angle, Length float64
}

//CalculateAngle returns the line's angle
func (l *Line) CalculateAngle() float64 {
	return math.Atan2(l.Y2-l.Y1, l.X2-l.X1)
}
