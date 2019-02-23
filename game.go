package tentsuyu

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//GameHelperFunction is a function that takes no parameters and returns an error
type GameHelperFunction func() error

//GameLoadImagesFunction returns an ImageManager which is used to load new images into the game
type GameLoadImagesFunction func() *ImageManager

//GameLoadAudioFunction returns an AudioPlayer and is used to load new audio into the game
type GameLoadAudioFunction func() *AudioPlayer

//Game represents, well... the game
type Game struct {
	imageLoadedCh                  chan *ImageManager
	audioLoadedCh                  chan *AudioPlayer
	gameState                      GameState
	pausedState                    GameState
	GameData                       *GameData
	Screen                         *ebiten.Image
	DefaultCamera                  *Camera
	UIController                   *UIController
	Random                         *rand.Rand
	highScoreDisplay, scoreDisplay *MenuElement
	Input                          *InputController
	ImageManager                   *ImageManager
	GameStateLoop                  GameHelperFunction
	AudioPlayer                    *AudioPlayer
}

//NewGame returns a new Game while setting the width and height of the screen
func NewGame(screenWidth, screenHeight float64) (game *Game, err error) {
	game = &Game{
		imageLoadedCh: make(chan *ImageManager),
		audioLoadedCh: make(chan *AudioPlayer),
		GameData:      NewGameData(),
		Random:        rand.New(rand.NewSource(time.Now().UnixNano())),
		Input:         NewInputController(),
		DefaultCamera: CreateCamera(screenWidth, screenHeight),
		UIController:  NewUIController(),
		ImageManager: &ImageManager{
			Images: map[string]*ebiten.Image{},
		},
	}
	game.AudioPlayer, err = NewAudioPlayer()
	if err != nil {
		return nil, err
	}

	//Basic Default Inputs
	game.Input.RegisterButton("Up", ebiten.KeyW, ebiten.KeyUp)
	game.Input.RegisterButton("Down", ebiten.KeyS, ebiten.KeyDown)
	game.Input.RegisterButton("Left", ebiten.KeyA, ebiten.KeyLeft)
	game.Input.RegisterButton("Right", ebiten.KeyD, ebiten.KeyRight)
	game.Input.RegisterButton("Escape", ebiten.KeyEscape)
	game.Input.RegisterButton("Enter", ebiten.KeyEnter)
	game.Input.RegisterButton("Space", ebiten.KeySpace)

	//Default Numbers:
	game.Input.RegisterButton("1", ebiten.Key1)
	game.Input.RegisterButton("2", ebiten.Key2)
	game.Input.RegisterButton("3", ebiten.Key3)
	game.Input.RegisterButton("4", ebiten.Key4)
	game.Input.RegisterButton("5", ebiten.Key5)
	game.Input.RegisterButton("6", ebiten.Key6)
	game.Input.RegisterButton("7", ebiten.Key7)
	game.Input.RegisterButton("8", ebiten.Key8)
	game.Input.RegisterButton("9", ebiten.Key9)
	game.Input.RegisterButton("0", ebiten.Key0)

	//ToggleFullscreen default button is F11
	game.Input.RegisterButton("ToggleFullscreen", ebiten.KeyF11)

	return
}

//ToggleFullscreen toggles the game in or out of full screen
func (g *Game) ToggleFullscreen() {
	if ebiten.IsFullscreen() {
		ebiten.SetFullscreen(false)
	} else {
		ebiten.SetFullscreen(true)
	}
}

//Loop is the main game loop
func (g *Game) Loop(screen *ebiten.Image) error {
	if g.imageLoadedCh != nil || g.audioLoadedCh != nil {
		select {
		case g.ImageManager = <-g.imageLoadedCh:

			g.imageLoadedCh = nil
		case g.AudioPlayer = <-g.audioLoadedCh:

			g.audioLoadedCh = nil
		default:
		}
	}
	if g.imageLoadedCh != nil || g.audioLoadedCh != nil {
		return ebitenutil.DebugPrint(screen, "Now Loading...")
	}

	g.Input.Update()
	g.Screen = screen
	//g.Screen.Fill(color.RGBA{R: 192, G: 192, B: 192, A: 255})

	if g.gameState == nil {
		g.gameState = defaultGameState()

	} else {
		if err := g.GameStateLoop(); err != nil {
			return err
		}
	}

	g.gameState.Update(g)
	g.GameData.Update()
	g.UIController.Update()
	if g.Input.Button("ToggleFullscreen").JustPressed() {
		g.ToggleFullscreen()
	}
	if !ebiten.IsRunningSlowly() {
		//g.background.Draw(g.screen, true)
		if err := g.gameState.Draw(g); err != nil {
			return err
		}
		g.scoreDisplay.Update()
		g.highScoreDisplay.Update()
		g.scoreDisplay.Draw(g.Screen)
		g.highScoreDisplay.Draw(g.Screen)
		g.UIController.Draw(g.Screen)
	}

	return nil
	//return ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %.2f", ebiten.CurrentFPS()))
}

//SetGameState of the game
func (g *Game) SetGameState(gs GameState) {
	g.gameState = gs
}

//SetPauseState of the game
//This changes the pausedState to the current GameState then switches to the passed GameState.
//Used to preserve the current game state
func (g *Game) SetPauseState(gs GameState) {
	g.pausedState = g.gameState
	g.gameState = gs
	g.pausedState.SetMsg(GameStateMsgNone)
}

//UnPause switches back the the puasedState GameState of the Game
func (g *Game) UnPause() {
	g.gameState = g.pausedState
}

//SetGameStateLoop should be a switch statement telling the game when to switch to what gamestate
//This is where your gamestate logic will exist
func (g *Game) SetGameStateLoop(gFunction GameHelperFunction) {
	g.GameStateLoop = gFunction
}

//LoadImages will set the imageLoadedCh to the passed GameHlperFunction
//This is used to load images before a gamestate is set
func (g *Game) LoadImages(gFunction GameLoadImagesFunction) {
	go func() {
		imageManager := &ImageManager{
			Images: map[string]*ebiten.Image{},
		}
		if imageManager = gFunction(); imageManager != nil {
			g.imageLoadedCh <- imageManager
		}
		close(g.imageLoadedCh)
	}()

}

//LoadAudio will set the audioLoadedCh to the passed GameHelperFunction
//This is used to load audio before a gamestate is set
func (g *Game) LoadAudio(gFunction GameLoadAudioFunction) {
	go func() {
		audioPlayer, _ := NewAudioPlayer()
		if audioPlayer = gFunction(); audioPlayer != nil {
			g.audioLoadedCh <- audioPlayer
		}
		close(g.audioLoadedCh)
	}()
}
