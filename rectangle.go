package tentsuyu

//Rectangle used for collision checks
type Rectangle struct {
	W, H int
	X, Y float64
}

//Left edge of the rectangle
func (rect *Rectangle) Left() float64 {

	return rect.X //- float64(rect.w/2)

}

//Right edge of the rectangle
func (rect *Rectangle) Right() float64 {

	return rect.X + float64(rect.W)

}

//Top edge of the rectangle
func (rect *Rectangle) Top() float64 {

	return rect.Y //- float64(rect.h/2)

}

//Bottom edge of the rectangle
func (rect *Rectangle) Bottom() float64 {

	return rect.Y + float64(rect.H)

}

//Contains returns true if the given point is withing the bounding box
func (rect *Rectangle) Contains(srcX, srcY float64) bool {

	xIn, yIn := false, false
	if srcX < rect.Right() && srcX > rect.Left() {
		xIn = true
	}
	if srcY < rect.Bottom() && srcY > rect.Top() {
		yIn = true
	}
	return xIn && yIn
}
