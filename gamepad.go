package tentsuyu

//GamePad holds the relevant gamepad logic
type GamePad struct {
}

//GamePadManager contains all the available gamepads
type GamePadManager struct {
	GamePads []*GamePad
}

//NewGamePadManager returns an empty GamePadManager
func NewGamePadManager() *GamePadManager {
	g := &GamePadManager{
		GamePads: []*GamePad{},
	}

	return g
}
