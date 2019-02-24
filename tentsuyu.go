package tentsuyu

import (
	"github.com/hajimehoshi/ebiten"
)

//Components stores all the different components used together in the engine
var Components *Component

//Input is the InputController for the entire game
var (
//Input *InputController
//ImageManager *ImageManager
)

//BootUp initializes the package
func BootUp(screenWidth, screenHeight float64) {
	Components = &Component{
		Camera:          CreateCamera(screenWidth, screenHeight),
		InputController: NewInputController(),
		//UIController:    NewUIController(),
		ScreenHeight: screenHeight,
		ScreenWidth:  screenWidth,
	}
	//Input = Components.InputController
	/*ImageManager = &ImageManager{
		Images: map[string]*ebiten.Image{},
	}*/
}

//Component holds different game elements and controllers
type Component struct {
	*Camera
	*InputController
	*UIController
	ScreenWidth, ScreenHeight float64
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

//Point represents a point in 2D space
type Point struct {
	X, Y float64
}
