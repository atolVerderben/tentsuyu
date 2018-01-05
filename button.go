package tentsuyu

import "github.com/hajimehoshi/ebiten"

// A Button is an input which can be either JustPressed, JustReleased or Down. Common uses would be for, a jump key or an action key.
type Button struct {
	Triggers []ebiten.Key
	Name     string
}

//NewButton creates a new button
func NewButton(name string, triggerKeys []ebiten.Key) Button {
	b := Button{
		Triggers: triggerKeys,
		Name:     name,
	}
	return b
}

// JustPressed checks whether an input was pressed in the previous frame.
func (b Button) JustPressed() bool {
	for _, trigger := range b.Triggers {
		v := Input.keyManager.Get(trigger).JustPressed()
		if v {
			return v
		}
	}

	return false
}

// JustReleased checks whether an input was released in the previous frame.
func (b Button) JustReleased() bool {
	for _, trigger := range b.Triggers {
		v := Input.keyManager.Get(trigger).JustReleased()
		if v {
			return v
		}
	}

	return false
}

// Down checks whether the current input is being held down.
func (b Button) Down() bool {
	for _, trigger := range b.Triggers {
		v := Input.keyManager.Get(trigger).Down()
		if v {
			return v
		}
	}

	return false
}
