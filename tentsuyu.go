package tentsuyu

import (
	"github.com/hajimehoshi/ebiten"
)

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
