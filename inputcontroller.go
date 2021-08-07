package tentsuyu

import "github.com/hajimehoshi/ebiten/v2"

//InputController controls all input for the game
type InputController struct {
	drawCursor            bool
	keyManager            *KeyManager
	buttons               map[string]Button
	mouseButtons          map[string]MouseButton
	Mouse                 *Mouse
	keyDelay, keyInterval float64
}

//NewInputController returns a new InputController
func NewInputController() *InputController {
	ic := &InputController{
		buttons:      make(map[string]Button),
		mouseButtons: make(map[string]MouseButton),
		keyManager:   NewKeyManager(),
		drawCursor:   true,
		Mouse:        NewMouse(),
		keyDelay:     30,
		keyInterval:  3,
	}
	ic.RegisterMouseButton("LeftMouse", ebiten.MouseButtonLeft)
	ic.RegisterMouseButton("RightMouse", ebiten.MouseButtonRight)
	ic.RegisterMouseButton("MiddleMouse", ebiten.MouseButtonMiddle)

	return ic
}

//LeftClick returns the Mouse left click button
func (ic *InputController) LeftClick() MouseState {
	return ic.Mouse.buttonMap[ebiten.MouseButtonLeft]
}

//RightClick returns the Mouse left click button
func (ic *InputController) RightClick() MouseState {
	return ic.Mouse.buttonMap[ebiten.MouseButtonRight]
}

//MouseWheelUp returns true if the user is scrolling up
func (ic *InputController) MouseWheelUp() bool {
	return ic.Mouse.IsScrollUp()
}

//MouseWheelDown returns true if the user is scrolling down
func (ic *InputController) MouseWheelDown() bool {
	return ic.Mouse.IsScrollDown()
}

//GetMouseCoords returns the ebiten mouse coords
func (ic *InputController) GetMouseCoords() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x), float64(y)
}

//GetGameMouseCoords returns the game coords based on the camera position and zoom
func (ic *InputController) GetGameMouseCoords(camera *Camera) (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx)+(camera.GetX()))/(camera.Zoom), (float64(my)+(camera.GetY()))/(camera.Zoom)
	return x, y
}

//GetGameMouseCoordsOffset returns the game coords based on the camera position and zoom with added offsets
func (ic *InputController) GetGameMouseCoordsOffset(camera *Camera, xOffset, yOffset float64) (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx)+(camera.GetX()+xOffset))/(camera.Zoom), (float64(my)+(camera.GetY()+yOffset))/(camera.Zoom)
	return x, y
}

//GetGameMouseCoordsNoZoom is the same as GetGameMouseCoords but ignores the camera's zoom level (useful for drawing the cursor)
func (ic *InputController) GetGameMouseCoordsNoZoom(camera *Camera) (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx) + (camera.GetX())), (float64(my) + (camera.GetY()))
	return x, y
}

//Update InputController
func (ic *InputController) Update() {
	ic.keyManager.update()
	ic.Mouse.update(ic)

}

// RegisterButton registers a new button input.
func (ic *InputController) RegisterButton(name string, triggerKeys ...ebiten.Key) {
	ic.buttons[name] = NewButton(name, triggerKeys, ic)
	for i := range triggerKeys {
		ic.keyManager.AddKey(triggerKeys[i])
	}
}

//RegisterMouseButton adds the mouse button to the game with the given name
func (ic *InputController) RegisterMouseButton(name string, buttons ...ebiten.MouseButton) {
	ic.mouseButtons[name] = MouseButton{
		Name:     name,
		Triggers: buttons,
		input:    ic,
	}
}

//Button retrieves a Button with a specified name.
func (ic *InputController) Button(name string) Button {
	return ic.buttons[name]
}

//MouseButton retrieves a Button with a specified name.
func (ic *InputController) MouseButton(name string) MouseButton {
	return ic.mouseButtons[name]
}
