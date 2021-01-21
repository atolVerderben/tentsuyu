package tentsuyu

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// MouseStateUp is a state for when the key is not currently being pressed
	MouseStateUp = iota
	// MouseStateDown is a state for when the key is currently being pressed
	MouseStateDown
	// MouseStateJustDown is a state for when a key was just pressed
	MouseStateJustDown
	// MouseStateJustUp is a state for when a key was just released
	MouseStateJustUp
	// MouseStateWheel is a state for when the mouse wheel is moved
	MouseStateWheel
)

//Mouse represents the cursor
type Mouse struct {
	X, Y             float64
	wX, wY           float64
	buttonMap        map[ebiten.MouseButton]MouseState
	mutex            sync.RWMutex
	mouseWheelMoving bool
	scroll           int
}

//Set tells the mouse to press the selected mouse button
func (m *Mouse) Set(key ebiten.MouseButton, state bool) {
	m.mutex.Lock()

	ms := m.buttonMap[key]
	ms.set(state)
	m.buttonMap[key] = ms

	m.mutex.Unlock()
}

// Get retrieves a keys state.
func (m *Mouse) Get(k ebiten.MouseButton) MouseState {
	m.mutex.RLock()
	ms := m.buttonMap[k]
	m.mutex.RUnlock()

	return ms
}

//IsScrollDown returns true if the user is scrolling down on the mouse wheel
func (m *Mouse) IsScrollDown() bool {
	if m.scroll < 0 {
		return true
	}
	return false
}

//IsScrollUp returns true if the user is scrolling up on the mouse wheel
func (m *Mouse) IsScrollUp() bool {
	if m.scroll > 0 {
		return true
	}
	return false
}

func (m *Mouse) update(input *InputController) {
	m.X, m.Y = input.GetMouseCoords() //m.GetGameMouseCoordsNoZoom()
	m.mouseWheelMoving = false
	wX, wY := ebiten.Wheel()
	if wX != 0 || wY != 0 {
		//moving the mouse wheel
		m.mouseWheelMoving = true
		if wX > m.wX {
			//click left
		} else if wX < m.wX {
			//click right
		} else {

		}

		if wY > m.wY {
			//scrollUp
			m.scroll = 1
		} else if wY < m.wY {
			//scrollDown
			m.scroll = -1
		} else {
			m.scroll = 0
		}
	} else {
		m.scroll = 0
	}
	m.wX = wX
	m.wY = wY
	for key := range m.buttonMap {
		if ebiten.IsMouseButtonPressed(key) {
			m.Set(key, true)
		} else {
			m.Set(key, false)
		}
	}

}

//GetGameMouseCoordsNoZoom is the same as GetGameMouseCoords but ignores the camera's zoom level (useful for drawing the cursor)
func (m *Mouse) GetGameMouseCoordsNoZoom(camera *Camera) (x, y float64) {
	mx, my := ebiten.CursorPosition()
	x, y = (float64(mx) + (camera.GetX())), (float64(my) + (camera.GetY()))
	return x, y
}

//NewMouse returns a new pointer to a Mouse struct
func NewMouse() *Mouse {
	m := &Mouse{
		buttonMap: make(map[ebiten.MouseButton]MouseState),
	}
	m.AddKey(ebiten.MouseButtonLeft)
	m.AddKey(ebiten.MouseButtonRight)
	m.AddKey(ebiten.MouseButtonMiddle)
	return m
}

//AddKey adds the ebiten.Key to the list of keys being managed
func (m *Mouse) AddKey(k ebiten.MouseButton) {
	m.buttonMap[k] = MouseState{}
}

// MouseState is used for detecting the state of a key press.
type MouseState struct {
	lastState    bool
	currentState bool
}

func (key *MouseState) set(state bool) {
	key.lastState = key.currentState
	key.currentState = state
}

// State returns the raw state of a key.
func (key *MouseState) State() int {
	if key.lastState {
		if key.currentState {
			return MouseStateDown
		}
		return MouseStateJustUp

	}
	if key.currentState {
		return MouseStateJustDown
	}
	return MouseStateUp

}

// JustPressed returns whether a key was just pressed
func (key MouseState) JustPressed() bool {
	return (!key.lastState && key.currentState)
}

// JustReleased returns whether a key was just released
func (key MouseState) JustReleased() bool {
	return (key.lastState && !key.currentState)
}

// Up returns wheter a key is not being pressed
func (key MouseState) Up() bool {
	return (!key.lastState && !key.currentState)
}

// Down returns wether a key is being pressed
func (key MouseState) Down() bool {
	return (key.lastState && key.currentState)
}
