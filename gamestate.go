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
	GameStateMsgNone   GameStateMsg = ""
	GameStateMsgPaused              = "Paused"
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
