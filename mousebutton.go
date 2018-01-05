package tentsuyu

import "github.com/hajimehoshi/ebiten"

//MouseButton is a representation of one of the mouse buttons
type MouseButton struct {
	Name     string
	Triggers []ebiten.MouseButton
}

// JustPressed checks whether an input was pressed in the previous frame.
func (b MouseButton) JustPressed() bool {
	for _, trigger := range b.Triggers {
		v := Input.Mouse.Get(trigger).JustPressed()
		if v {
			return v
		}
	}

	return false
}

// JustReleased checks whether an input was released in the previous frame.
func (b MouseButton) JustReleased() bool {
	for _, trigger := range b.Triggers {
		v := Input.Mouse.Get(trigger).JustReleased()
		if v {
			return v
		}
	}

	return false
}

// Down checks whether the current input is being held down.
func (b MouseButton) Down() bool {
	for _, trigger := range b.Triggers {
		v := Input.Mouse.Get(trigger).Down()
		if v {
			return v
		}
	}

	return false
}
