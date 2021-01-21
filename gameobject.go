package tentsuyu

import (
	"math"

	"github.com/rs/xid"

	"github.com/hajimehoshi/ebiten/v2"
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
	Angle, Speed, PrevX, PrevY float64
	VX, VY                     float64
	NotCentered                bool
	Width, Height              int
	WidthF, HeightF            float64
	ID                         xid.ID
	isCentered                 bool
	Velocity                   *Vector2d
	Position                   *Vector2d
}

//SetCollision2D will allow the object to use more advanced collision functions
func (obj *BasicObject) SetCollision2D(isCircle bool) {
	x, y := obj.GetPosition()
	if obj.NotCentered {
		x += float64(obj.Width) / 2
		y += float64(obj.Height) / 2
	}
}

//========================================================
//Implement GameObject

//GetSpeed returns the BasicObject speed
func (obj *BasicObject) GetSpeed() float64 {
	return obj.Speed
}

//GetVelocity returns Vx and Vy of the BasicObject
func (obj *BasicObject) GetVelocity() (float64, float64) {
	//vecX, vecY := math.Cos(obj.Angle)*obj.Speed, math.Sin(obj.Angle)*obj.Speed
	return obj.VX, obj.VY //vecX, vecY
}

//GetPosition returns the x,y coords
func (obj BasicObject) GetPosition() (float64, float64) {
	return obj.Position.X, obj.Position.Y
}

//GetY returns the y coords
func (obj BasicObject) GetY() float64 {
	return obj.Position.Y
}

//GetX returns the x coords
func (obj BasicObject) GetX() float64 {
	return obj.Position.X
}

//GetWidth returns width
func (obj BasicObject) GetWidth() int {
	return obj.Width
}

//GetHeight returns height
func (obj BasicObject) GetHeight() int {
	return obj.Height
}

//GetWidthF returns width as a float64
func (obj BasicObject) GetWidthF() float64 {
	return float64(obj.Width)
}

//GetHeightF returns height as a float64
func (obj BasicObject) GetHeightF() float64 {
	return float64(obj.Height)
}

//AddPosition increases X and Y position by vX and vY respectively
func (obj *BasicObject) AddPosition(vX, vY float64) {
	obj.Position.X += vX
	obj.Position.Y += vY
}

//SetPosition of x,y
func (obj *BasicObject) SetPosition(x, y float64) {
	obj.Position.X = x
	obj.Position.Y = y
}

//SetSize with width , height
func (obj *BasicObject) SetSize(width, height int) {
	obj.Width = width
	obj.Height = height
	obj.WidthF = float64(width)
	obj.HeightF = float64(height)
}

//=========================================================

//SetX of the position
func (obj *BasicObject) SetX(x float64) {
	obj.Position.X = x
}

//SetY of the position
func (obj *BasicObject) SetY(y float64) {
	obj.Position.Y = y
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
	obj.Position.X += vX
}

//AddY adds to the y value
func (obj *BasicObject) AddY(vY float64) {
	obj.Position.Y += vY
}

//NewBasicObject returns a new oject
func NewBasicObject(x, y float64, w, h int) *BasicObject {
	obj := &BasicObject{
		ID:         xid.New(),
		Width:      w,
		Height:     h,
		WidthF:     float64(w),
		HeightF:    float64(h),
		isCentered: true,
		Velocity:   &Vector2d{},
		Position:   &Vector2d{X: x, Y: y},
	}

	return obj
}

//GetID returns the guid of the Basic Object
func (obj BasicObject) GetID() xid.ID {
	return obj.ID
}

//GetIDasString returns the string representation of the guid
func (obj BasicObject) GetIDasString() string {
	return obj.ID.String()
}

//SetID sets the BasicObject's ID to a new guid
func (obj *BasicObject) SetID() {
	obj.ID = xid.New()
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
	if srcX <= obj.Right() && srcX >= obj.Left() {
		xIn = true
	}
	if srcY <= obj.Bottom() && srcY >= obj.Top() {
		yIn = true
	}
	return xIn && yIn
}

//LeftNoCenter edge of the rectangle
func (obj *BasicObject) LeftNoCenter() float64 {
	x, _ := obj.GetPosition()
	return x

}

//RightNoCenter edge of the rectangle
func (obj *BasicObject) RightNoCenter() float64 {
	x, _ := obj.GetPosition()
	return x + float64(obj.GetWidth())

}

//TopNoCenter edge of the rectangle
func (obj *BasicObject) TopNoCenter() float64 {
	_, y := obj.GetPosition()
	return y

}

//BottomNoCenter edge of the rectangle
func (obj *BasicObject) BottomNoCenter() float64 {
	_, y := obj.GetPosition()
	return y + float64(obj.GetHeight())

}

//ContainsNoCenter returns true if the given point is withing the rectangle of the object
func (obj *BasicObject) ContainsNoCenter(srcX, srcY float64) bool {

	xIn, yIn := false, false
	if srcX <= obj.RightNoCenter() && srcX >= obj.LeftNoCenter() {
		xIn = true
	}
	if srcY <= obj.BottomNoCenter() && srcY >= obj.TopNoCenter() {
		yIn = true
	}
	return xIn && yIn
}

//GetHealth should return BasicObject Health
func (obj *BasicObject) GetHealth() float64 {
	return 1.0
}

//Update for GameObject
//This should be overriden by user
func (obj *BasicObject) Update() {

}

//Draw for GameObject
//This should be overriden by user
func (obj *BasicObject) Draw(screen *ebiten.Image) error {
	return nil
}

//ReturnVectorPosition returns the X,Y position as a vector2d
//This is useful for vector math
func (obj BasicObject) ReturnVectorPosition() Vector2d {
	return *obj.Position //Vector2d{X: obj.X, Y: obj.Y}
}

//SetCentered of the object. If true the object coords refer to the center of the object.
//If false the object coords refer to the top left of the object.
func (obj *BasicObject) SetCentered(isCentered bool) {
	obj.NotCentered = !isCentered
	obj.isCentered = isCentered
}
