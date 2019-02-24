package tentsuyu

//GameState represents what all game states must conform to
type GameState interface {
	Update(*Game) error
	Draw(*Game) error
	Msg() GameStateMsg
	SetMsg(GameStateMsg)
}

//GameStateMsg is what it is
type GameStateMsg string

const (
	//GameStateMsgNone is the default empty message
	GameStateMsgNone GameStateMsg = ""
	//GameStateMsgPause pauses the game
	GameStateMsgPause = "Paused"
	//GameStateMsgUnPause resumes the game from a paused state
	GameStateMsgUnPause = "UnPaused"
)

func defaultGameState() GameState {
	return &defaultGS{}
}

type defaultGS struct {
}

func (d *defaultGS) Update(g *Game) error {
	return nil
}

func (d *defaultGS) Draw(g *Game) error {
	return nil
}

func (d *defaultGS) Msg() GameStateMsg {
	return GameStateMsgNone
}

func (d *defaultGS) SetMsg(g GameStateMsg) {

}
