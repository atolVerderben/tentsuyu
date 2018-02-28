package tentsuyu

import (
	"math"

	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten"
)

//GameObject is any renderable object
type GameObject interface {
	GetPosition() (float64, float64)
	SetPosition(float64, float64)
	GetWidth() int
	GetHeight() int
	//Update()
	Draw(*ebiten.Image) error
	Contains(float64, float64) bool
	GetHealth() float64
	GetAngle() float64
	GetVelocity() (float64, float64)
}

//BasicObject is the bare implementation of a GameObject
type BasicObject struct {
	X, Y, Angle, Speed float64
	VX, VY             float64
	NotCentered        bool
	Width, Height      int
	IsCircle           bool
	Box                collision2d.Box
	Circle             collision2d.Circle
}

//SetCollision2D will allow the object to use more advanced collision functions
func (obj *BasicObject) SetCollision2D(isCircle bool) {
	x, y := obj.GetPosition()
	if obj.NotCentered {
		x += float64(obj.Width) / 2
		y += float64(obj.Height) / 2
	}

	if isCircle {
		obj.IsCircle = isCircle
		obj.Circle = collision2d.NewCircle(collision2d.NewVector(x, y), float64(obj.Width))
	} else {
		obj.Box = collision2d.NewBox(collision2d.NewVector(x, y), float64(obj.Width), float64(obj.Height))
	}
}

//========================================================
//Implement GameObject

func (obj *BasicObject) GetSpeed() float64 {
	return obj.Speed
}

func (obj *BasicObject) GetVelocity() (float64, float64) {
	//vecX, vecY := math.Cos(obj.Angle)*obj.Speed, math.Sin(obj.Angle)*obj.Speed
	return obj.VX, obj.VY //vecX, vecY
}

//GetPosition returns the x,y coords
func (obj BasicObject) GetPosition() (float64, float64) {
	return obj.X, obj.Y
}

//GetWidth returns width
func (obj BasicObject) GetWidth() int {
	return obj.Width
}

//GetHeight returns height
func (obj BasicObject) GetHeight() int {
	return obj.Height
}

//AddPosition increases X and Y position by vX and vY respectively
func (obj *BasicObject) AddPosition(vX, vY float64) {
	obj.X += vX
	obj.Y += vY
}

//SetPosition of x,y
func (obj *BasicObject) SetPosition(x, y float64) {
	obj.X = x
	obj.Y = y
	obj.Box.Pos.X = obj.X
	obj.Box.Pos.Y = obj.Y
}

//SetSize with width , height
func (obj *BasicObject) SetSize(width, height int) {
	obj.Width = width
	obj.Height = height
}

//=========================================================

//SetX of the position
func (obj *BasicObject) SetX(x float64) {
	obj.X = x
}

//SetY of the position
func (obj *BasicObject) SetY(y float64) {
	obj.Y = y
}

//GetAngle  of BasicObject
func (obj BasicObject) GetAngle() float64 {
	return obj.Angle
}

//SetAngle of BasicObject
func (obj *BasicObject) SetAngle(angle float64) {
	obj.Angle = angle
}

//AddAngle adds the provided increment to the angle RADIAN
func (obj *BasicObject) AddAngle(inc float64) {
	obj.Angle += inc
	if obj.Angle > 2*math.Pi {
		obj.Angle = 0
	}
	if obj.Angle < 0 {
		obj.Angle = (2 * math.Pi) + obj.Angle
	}
}

//GetSize returns width and height
func (obj BasicObject) GetSize() (int, int) {
	return obj.GetWidth(), obj.GetHeight()
}

//AddX adds to the x value
func (obj *BasicObject) AddX(vX float64) {
	obj.X += vX
}

//AddY adds to the y value
func (obj *BasicObject) AddY(vY float64) {
	obj.X += vY
}

//NewBasicObject returns a new oject
func NewBasicObject() *BasicObject {
	obj := &BasicObject{}

	return obj
}

//Left edge of the rectangle
func (obj *BasicObject) Left() float64 {
	x, _ := obj.GetPosition()
	if obj.NotCentered {
		return x
	}
	return x - float64(obj.GetWidth()/2)

}

//Right edge of the rectangle
func (obj *BasicObject) Right() float64 {
	x, _ := obj.GetPosition()
	if obj.NotCentered {
		return x + float64(obj.GetWidth())
	}
	return x + float64(obj.GetWidth()/2)

}

//Top edge of the rectangle
func (obj *BasicObject) Top() float64 {
	_, y := obj.GetPosition()
	if obj.NotCentered {
		return y
	}
	return y - float64(obj.GetHeight()/2)

}

//Bottom edge of the rectangle
func (obj *BasicObject) Bottom() float64 {
	_, y := obj.GetPosition()
	if obj.NotCentered {
		return y + float64(obj.GetHeight())
	}
	return y + float64(obj.GetHeight()/2)

}

//Contains returns true if the given point is withing the rectangle of the object
func (obj *BasicObject) Contains(srcX, srcY float64) bool {
	if obj.NotCentered == true {
		return obj.ContainsNoCenter(srcX, srcY)
	}

	xIn, yIn := false, false
	if srcX < obj.Right() && srcX > obj.Left() {
		xIn = true
	}
	if srcY < obj.Bottom() && srcY > obj.Top() {
		yIn = true
	}
	return xIn && yIn
}

//Left edge of the rectangle
func (obj *BasicObject) LeftNoCenter() float64 {
	x, _ := obj.GetPosition()
	return x

}

//Right edge of the rectangle
func (obj *BasicObject) RightNoCenter() float64 {
	x, _ := obj.GetPosition()
	return x + float64(obj.GetWidth())

}

//Top edge of the rectangle
func (obj *BasicObject) TopNoCenter() float64 {
	_, y := obj.GetPosition()
	return y

}

//Bottom edge of the rectangle
func (obj *BasicObject) BottomNoCenter() float64 {
	_, y := obj.GetPosition()
	return y + float64(obj.GetHeight())

}

//Contains returns true if the given point is withing the rectangle of the object
func (obj *BasicObject) ContainsNoCenter(srcX, srcY float64) bool {

	xIn, yIn := false, false
	if srcX < obj.RightNoCenter() && srcX > obj.LeftNoCenter() {
		xIn = true
	}
	if srcY < obj.BottomNoCenter() && srcY > obj.TopNoCenter() {
		yIn = true
	}
	return xIn && yIn
}

func (obj *BasicObject) GetHealth() float64 {
	return 1.0
}

//Update for GameObject
func (obj *BasicObject) Update() {

}

//Draw for GameObject
func (obj *BasicObject) Draw(screen *ebiten.Image) error {
	return nil
}

//======================
//Image Parts
//======================

//BasicImageParts is easy to set up basic sprite image
type BasicImageParts struct {
	name                                         string
	Width, Height, Sx, Sy, DestWidth, DestHeight int
	Reverse                                      bool
}

func NewBasicImageParts(sx, sy, width, height int) *BasicImageParts {
	b := &BasicImageParts{
		Sx:         sx,
		Sy:         sy,
		Width:      width,
		Height:     height,
		DestHeight: height,
		DestWidth:  width,
	}
	return b
}

func (b *BasicImageParts) SetDestinationDimensions(width, height int) {
	b.DestWidth = width
	b.DestHeight = height
}

func (b *BasicImageParts) ReverseX(reverse bool) {
	b.Reverse = reverse
}

//Len returns 1
func (b *BasicImageParts) Len() int {
	return 1
}

//Dst we just make it 1:1
func (b *BasicImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	if b.DestHeight == 0 && b.DestWidth == 0 {
		return 0, 0, b.Width, b.Height
	}
	return 0, 0, b.DestWidth, b.DestHeight
}

//Src cuts out the specified rectangle from the source image to display the sprite
func (b *BasicImageParts) Src(i int) (x0, y0, x1, y1 int) {
	x := b.Sx
	y := b.Sy
	if b.Reverse {
		return x + b.Width, y, x, y + b.Height
	}
	return x, y, x + b.Width, y + b.Height
}
