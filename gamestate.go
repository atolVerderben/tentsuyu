package tentsuyu

//GameState represents what all game states must conform to
type GameState interface {
	Update(*Game) error
	Draw(*Game) error
	Msg() GameStateMsg
	SetMsg(GameStateMsg)
}

//GameStateMsg represents
type GameStateMsg string

const (
	//GameStateMsgNone is the default empty message
	GameStateMsgNone GameStateMsg = ""
	//GameStateMsgPause pauses the game
	GameStateMsgPause = "Paused"
	//GameStateMsgUnPause resumes the game from a paused state
	GameStateMsgUnPause = "UnPaused"
	//GameStateMsgNotStarted is used to call the game init
	GameStateMsgNotStarted = "Game Not Init"
)

//NewBaseGameState returns a BaseGameState with default GameStateMessage
func NewBaseGameState() *BaseGameState {
	return &BaseGameState{
		GameStateMessage: GameStateMsgNone,
	}
}

//BaseGameState implements the basic methods to satisfy the GameState interface
type BaseGameState struct {
	GameStateMessage GameStateMsg
}

//Update is a dummy method for BaseGameState
func (b *BaseGameState) Update(g *Game) error {
	return nil
}

//Draw is a dummy method for BaseGameState
func (b *BaseGameState) Draw(g *Game) error {
	return nil
}

//Msg returns the GameStateMessage of the BaseGameState
func (b *BaseGameState) Msg() GameStateMsg {
	return b.GameStateMessage
}

//SetMsg sets the GameStateMessage of BaseGameState to the passed parameter
func (b *BaseGameState) SetMsg(g GameStateMsg) {
	b.GameStateMessage = g
}
