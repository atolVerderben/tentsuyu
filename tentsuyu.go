package tentsuyu

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten"
)

//Components stores all the different components used together in the engine
var Components *Component

//Input is the InputController for the entire game
var (
	Input        *InputController
	ImageManager *imageManager
)

//BootUp initializes the package
func BootUp(screenWidth, screenHeight float64) {
	Components = &Component{
		Camera:          CreateCamera(screenWidth, screenHeight),
		InputController: NewInputController(),
		UIController:    NewUIController(),
		ScreenHeight:    screenHeight,
		ScreenWidth:     screenWidth,
	}
	Input = Components.InputController
	ImageManager = &imageManager{
		Images: map[string]*ebiten.Image{},
	}
}

//Component holds different game elements and controllers
type Component struct {
	*Camera
	*InputController
	*UIController
	ScreenWidth, ScreenHeight float64
}

//ApplyCameraTransform applies the camera's position to the DrawImageOptions, bool toggles whether zoom is applied or not
func ApplyCameraTransform(op *ebiten.DrawImageOptions, applyZoom bool) {
	if applyZoom {
		Components.Camera.DrawCameraTransform(op)
	} else {
		Components.Camera.DrawCameraTransformIgnoreZoom(op)
	}
}

//SetCustomCursor to allow for drawing your own cursor object
func SetCustomCursor(width, height, sx, sy int, spritesheet *ebiten.Image) {
	Components.UIController.SetCustomCursor(width, height, sx, sy, spritesheet)
}

//CenterCursor sets whether the cursor is centered or not
func CenterCursor(center bool) {
	Components.UIController.Cursor.NotCentered = !center
}

//OnScreen returns true if the specified coords are currently on the screenWidth
//i.e. the position is within the camera's view
func OnScreen(x, y float64, width, height int) bool {
	return Components.Camera.OnScreen(x, y, width, height)
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

//Collision2D uses the collision2d library to get a more accurate collision detail
func Collision2D(obj1 *BasicObject, obj2 *BasicObject) (bool, collision2d.Response) {
	if Collision(obj1, obj2) {

		if obj1.IsCircle {
			obj1.Circle.Pos = collision2d.NewVector(obj1.X, obj1.Y)
			if obj2.IsCircle {
				obj2.Circle.Pos = collision2d.NewVector(obj2.X, obj2.Y)
				return collision2d.TestCircleCircle(obj1.Circle, obj2.Circle)
			}
			obj2.Box.Pos = collision2d.NewVector(obj2.X, obj2.Y)
			return collision2d.TestCirclePolygon(obj1.Circle, obj2.Box.ToPolygon())

		}
		if obj2.IsCircle {
			obj2.Circle.Pos = collision2d.NewVector(obj2.X, obj2.Y)
			obj1.Box.Pos = collision2d.NewVector(obj1.X, obj1.Y)
			return collision2d.TestPolygonCircle(obj1.Box.ToPolygon(), obj2.Circle)
		}
		obj1.Box.Pos = collision2d.NewVector(obj1.X, obj1.Y)
		obj2.Box.Pos = collision2d.NewVector(obj2.X, obj2.Y)
		return collision2d.TestPolygonPolygon(obj1.Box.ToPolygon(), obj2.Box.ToPolygon())
	}
	return false, collision2d.NewResponse()
}
