package tentsuyu

// Action corresponds to a control action such as move, press, release
type Action int

// Modifier represents a special key pressed along with another key
type Modifier int

var (
	// Move is an action representing mouse movement
	Move = Action(0)
	// Press is an action representing a mouse press/click
	Press = Action(1)
	// Release is an action representing a mouse a release
	Release = Action(2)
	// Neutral represents a neutral action
	Neutral = Action(99)
	// Shift represents the shift modifier.
	// It is triggered when the shift key is pressed simultaneously with another key
	Shift = Modifier(0x0001)
	// Control represents the control modifier
	// It is triggered when the ctrl key is pressed simultaneously with another key
	Control = Modifier(0x0002)
	// Alt represents the alt modifier
	// It is triggered when the alt key is pressed simultaneously with another key
	Alt = Modifier(0x0004)
	// Super represents the super modifier
	// (Windows key on Microsoft Windows, Command key on Apple OSX, and varies on Linux)
	// It is triggered when the super key is pressed simultaneously with another key
	Super = Modifier(0x0008)
)
