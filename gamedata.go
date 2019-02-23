package tentsuyu

import "time"

//GameValuePair is used to add settings and various values to the GameData structure
type GameValuePair struct {
	Name       string
	ValueType  GameValueType
	ValueText  string
	ValueInt   int
	ValueFloat float64
}

//GameValueType is the available types of values use in the GameValuePair struct
type GameValueType int

const (
	//GameValueText means the value is a string
	GameValueText GameValueType = iota
	//GameValueInt means the value is an integer
	GameValueInt
	//GameValueFloat means the value is a float64
	GameValueFloat
)

//GameMode represents the different game modes possible
type GameMode int

//List of possible game modes
const (
	GameModeNormal GameMode = iota
)

//GameData holds the various data to be passed between game states
type GameData struct {
	time      int
	startTime time.Time
	gameMode  GameMode
	highScore float64
	currScore float64
	Settings  map[string]GameValuePair
}

//NewGameData creates a new GameData
func NewGameData() *GameData {
	g := &GameData{
		startTime: time.Now(),
	}

	return g
}

//Update GameData
//Ticks time ahead
func (g *GameData) Update() {
	g.time++
}

//TimeInSecond returns the current time in seconds
func (g *GameData) TimeInSecond() int {

	return int(time.Now().Sub(g.startTime).Seconds())

	//return g.time / 60
}

//TimeInMilliseconds returns the current time in seconds
func (g *GameData) TimeInMilliseconds() int {

	return int(time.Now().Sub(g.startTime).Nanoseconds() / int64(time.Millisecond))

	//return g.time / 60
}

//SetHighScore compares the given score to current highscore and updates if it's greater
func (g *GameData) SetHighScore(score float64) {
	if g.highScore < score {
		g.highScore = score
	}
}

//SetCurrentScore sets the current score of the game
func (g *GameData) SetCurrentScore(score float64) {
	g.currScore = score
}
