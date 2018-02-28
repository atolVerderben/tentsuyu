package tentsuyu

import "github.com/hajimehoshi/ebiten"

//InputController controls all input for the game
type InputController struct {
	drawCursor   bool
	keyManager   *KeyManager
	buttons      map[string]Button
	mouseButtons map[string]MouseButton
	*Mouse
}

//NewInputController returns a new InputController
func NewInputController() *InputController {
	ic := &InputController{
		buttons:      make(map[string]Button),
		mouseButtons: make(map[string]MouseButton),
		keyManager:   NewKeyManager(),
		drawCursor:   true,
		Mouse:        NewMouse(),
	}
	ic.RegisterMouseButton("LeftMouse", ebiten.MouseButtonLeft)
	ic.RegisterMouseButton("RightMouse", ebiten.MouseButtonRight)
	ic.RegisterMouseButton("MiddleMouse", ebiten.MouseButtonMiddle)
	ic.RegisterButton("TextBox0", ebiten.Key0)
	ic.RegisterButton("TextBox1", ebiten.Key1)
	ic.RegisterButton("TextBox2", ebiten.Key2)
	ic.RegisterButton("TextBox3", ebiten.Key3)
	ic.RegisterButton("TextBox4", ebiten.Key4)
	ic.RegisterButton("TextBox5", ebiten.Key5)
	ic.RegisterButton("TextBox6", ebiten.Key6)
	ic.RegisterButton("TextBox7", ebiten.Key7)
	ic.RegisterButton("TextBox8", ebiten.Key8)
	ic.RegisterButton("TextBox9", ebiten.Key9)
	ic.RegisterButton("TextBoxPeriod", ebiten.KeyPeriod)
	ic.RegisterButton("TextBoxSpace", ebiten.KeySpace)
	ic.RegisterButton("TextBoxBackspace", ebiten.KeyBackspace)
	ic.RegisterButton("TextBoxEnter", ebiten.KeyEnter)
	ic.RegisterButton("TextBoxSemiColon", ebiten.KeySemicolon)
	ic.RegisterButton("TextBoxShift", ebiten.KeyShift)

	ic.RegisterButton("TextBoxSlash", ebiten.KeySlash)
	ic.RegisterButton("TextBoxBackSlash", ebiten.KeyBackslash)

	//Letters
	ic.RegisterButton("TextBoxA", ebiten.KeyA)
	ic.RegisterButton("TextBoxB", ebiten.KeyB)
	ic.RegisterButton("TextBoxC", ebiten.KeyC)
	ic.RegisterButton("TextBoxD", ebiten.KeyD)
	ic.RegisterButton("TextBoxE", ebiten.KeyE)
	ic.RegisterButton("TextBoxF", ebiten.KeyF)
	ic.RegisterButton("TextBoxG", ebiten.KeyG)
	ic.RegisterButton("TextBoxH", ebiten.KeyH)
	ic.RegisterButton("TextBoxI", ebiten.KeyI)
	ic.RegisterButton("TextBoxJ", ebiten.KeyJ)
	ic.RegisterButton("TextBoxK", ebiten.KeyK)
	ic.RegisterButton("TextBoxL", ebiten.KeyL)
	ic.RegisterButton("TextBoxM", ebiten.KeyM)
	ic.RegisterButton("TextBoxN", ebiten.KeyN)
	ic.RegisterButton("TextBoxO", ebiten.KeyO)
	ic.RegisterButton("TextBoxP", ebiten.KeyP)
	ic.RegisterButton("TextBoxQ", ebiten.KeyQ)
	ic.RegisterButton("TextBoxR", ebiten.KeyR)
	ic.RegisterButton("TextBoxS", ebiten.KeyS)
	ic.RegisterButton("TextBoxT", ebiten.KeyT)
	ic.RegisterButton("TextBoxU", ebiten.KeyU)
	ic.RegisterButton("TextBoxV", ebiten.KeyV)
	ic.RegisterButton("TextBoxW", ebiten.KeyW)
	ic.RegisterButton("TextBoxX", ebiten.KeyX)
	ic.RegisterButton("TextBoxY", ebiten.KeyY)
	ic.RegisterButton("TextBoxZ", ebiten.KeyZ)
	return ic
}

//LeftClick returns the Mouse left click button
func (ic *InputController) LeftClick() MouseState {
	return ic.buttonMap[ebiten.MouseButtonLeft]
}

//RightClick returns the Mouse left click button
func (ic *InputController) RightClick() MouseState {
	return ic.buttonMap[ebiten.MouseButtonRight]
}

//GetMouseCoords returns the ebiten mouse coords
func (ic *InputController) GetMouseCoords() (float64, float64) {
	x, y := ebiten.CursorPosition()
	return float64(x), float64(y)
}

//GetGameMouseCoords returns the game coords based on the camera position and zoom
func (ic *InputController) GetGameMouseCoords() (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx)+(Components.Camera.GetX()))/(Components.Camera.Zoom), (float64(my)+(Components.Camera.GetY()))/(Components.Camera.Zoom)
	return x, y
}

//GetGameMouseCoordsNoZoom is the same as GetGameMouseCoords but ignores the camera's zoom level (useful for drawing the cursor)
func (ic *InputController) GetGameMouseCoordsNoZoom() (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx) + (Components.Camera.GetX())), (float64(my) + (Components.Camera.GetY()))
	return x, y
}

//Update InputController
func (ic *InputController) Update() {
	ic.keyManager.update()
	ic.Mouse.update()

}

// RegisterButton registers a new button input.
func (ic *InputController) RegisterButton(name string, triggerKeys ...ebiten.Key) {
	ic.buttons[name] = NewButton(name, triggerKeys)
	for i := range triggerKeys {
		ic.keyManager.AddKey(triggerKeys[i])
	}
}

//RegisterMouseButton adds the mouse button to the game with the given name
func (ic *InputController) RegisterMouseButton(name string, buttons ...ebiten.MouseButton) {
	ic.mouseButtons[name] = MouseButton{
		Name:     name,
		Triggers: buttons,
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
