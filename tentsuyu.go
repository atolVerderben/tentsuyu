package tentsuyu

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//Pixel is a 1x1 white pixel that can be used for simple drawing
var Pixel *ebiten.Image

func init() {
	//Create Pixel
	//Simple white 1x1 pixel image for manipulation
	Pixel := ebiten.NewImage(1, 1)

	Pixel.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
}

//SetCustomCursor to allow for drawing your own cursor object
func SetCustomCursor(uiController *UIController, width, height, sx, sy int, spritesheet *ebiten.Image) {
	uiController.SetCustomCursor(width, height, sx, sy, spritesheet)
}

//CenterCursor sets whether the cursor is centered or not
func CenterCursor(uiController *UIController, center bool) {
	uiController.Cursor.NotCentered = !center
}

//Collision returns true if two given BasicObjects are overlapping
func Collision(obj1 *BasicObject, obj2 *BasicObject) bool {
	//Objects are to the left of each other
	if obj1.Left() > obj2.Right() || obj2.Left() > obj1.Right() {
		return false
	}
	if obj1.Bottom() < obj2.Top() || obj2.Bottom() < obj1.Top() {
		return false
	}
	return true
}

//Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// DrawLine draws a line segment on the given destination dst.
//
// DrawLine is intended to be used mainly for debugging or prototyping purpose.
func DrawLine(dst *ebiten.Image, x1, y1, x2, y2 float64, clr color.Color, camera *Camera) {
	ew, eh := Pixel.Size()
	length := math.Hypot(x2-x1, y2-y1)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(length/float64(ew), 1/float64(eh))
	op.GeoM.Rotate(math.Atan2(y2-y1, x2-x1))
	op.GeoM.Translate(x1, y1)
	op.ColorM.Scale(colorScale(clr))
	camera.ApplyCameraTransform(op, true)
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	dst.DrawImage(Pixel, op)
}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}
