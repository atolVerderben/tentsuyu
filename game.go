package tentsuyu

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//GameHelperFunction is a function that takes no parameters and returns an error
type GameHelperFunction func() error

//GameDrawHelperFunction is meant to draw something on the passed ebiten.Image
type GameDrawHelperFunction func(*ebiten.Image) error

//GameLoadImagesFunction returns an ImageManager which is used to load new images into the game
type GameLoadImagesFunction func() *ImageManager

//GameLoadAudioFunction returns an AudioPlayer and is used to load new audio into the game
type GameLoadAudioFunction func() *AudioPlayer

//Game represents, well... the game
type Game struct {
	imageLoadedCh             chan *ImageManager
	audioLoadedCh             chan *AudioPlayer
	gameState                 GameState
	PausedState               GameState
	GameData                  *GameData
	Screen                    *ebiten.Image
	DefaultCamera             *Camera
	UIController              *UIController
	Random                    *rand.Rand
	Input                     *InputController
	ImageManager              *ImageManager
	GameStateLoop             GameHelperFunction
	GameDrawLoop              GameDrawHelperFunction
	AudioPlayer               *AudioPlayer
	AdditionalCameras         map[string]*Camera
	IsMobile                  bool
	screenWidth, screenHeight int
}

//NewGame returns a new Game while setting the width and height of the screen
func NewGame(screenWidth, screenHeight float64) (game *Game, err error) {
	game = &Game{
		imageLoadedCh: make(chan *ImageManager),
		audioLoadedCh: make(chan *AudioPlayer),
		GameData:      NewGameData(),
		//Random:            rand.New(rand.NewSource(time.Now().UnixNano())),
		Input:             NewInputController(),
		DefaultCamera:     CreateCamera(screenWidth, screenHeight),
		ImageManager:      NewImageManager(),
		AdditionalCameras: map[string]*Camera{},
	}
	game.screenWidth = int(screenWidth)
	game.screenHeight = int(screenHeight)
	game.UIController = NewUIController(game.Input)
	game.AudioPlayer, err = NewAudioPlayer()
	game.gameState = NewBaseGameState()
	game.gameState.SetMsg(GameStateMsgNotStarted)

	game.SetGameDrawLoop(func(screen *ebiten.Image) error {

		return nil
	})
	if err != nil {
		return nil, err
	}

	//=====================================
	//Create Default Inputs
	//All inputs can be overriden
	//=====================================

	//Basic Default Inputs - Can be overwritten
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

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
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
		return nil
	}

	g.Input.Update()

	if err := g.GameStateLoop(); err != nil {
		return err
	}

	if err := g.gameState.Update(g); err != nil {
		return err
	}
	g.GameData.Update()
	g.UIController.Update()
	if g.Input.Button("ToggleFullscreen").JustPressed() {
		g.ToggleFullscreen()
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Screen = screen

	if g.imageLoadedCh != nil || g.audioLoadedCh != nil {
		ebitenutil.DebugPrint(screen, "Now Loading...")
		return
	}

	if err := g.gameState.Draw(g); err != nil {
		log.Fatal(err)
	}
	/*if err := g.UIController.Draw(g.Screen); err != nil {
		log.Fatal(err)
	}*/
	if err := g.GameDrawLoop(screen); err != nil {
		log.Fatal(err)
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

//SetMobile tells the game if it's on mobile or not
//This is useful to know whether to check for touches or keys
func (g *Game) SetMobile(m bool) {
	g.IsMobile = m
}

//SetGameState of the game
func (g *Game) SetGameState(gs GameState) {
	g.gameState = gs
}

//GetGameState of the game
func (g Game) GetGameState() GameState {
	return g.gameState
}

//SetPauseState of the game
//This changes the PausedState to the current GameState then switches to the passed GameState.
//Used to preserve the current game state
func (g *Game) SetPauseState(gs GameState) {
	g.PausedState = g.gameState
	g.gameState = gs
	g.PausedState.SetMsg(GameStateMsgNone)
}

//UnPause switches back the the puasedState GameState of the Game
func (g *Game) UnPause() {
	g.gameState = g.PausedState
}

//SetGameStateLoop should be a switch statement telling the game when to switch to what gamestate
//This is where your gamestate logic will exist
func (g *Game) SetGameStateLoop(gFunction GameHelperFunction) {
	g.GameStateLoop = gFunction
}

//SetGameDrawLoop allows the user to add a final draw over the game screen no matter what state the game is in.
func (g *Game) SetGameDrawLoop(gFunction GameDrawHelperFunction) {
	g.GameDrawLoop = gFunction
}

//LoadImages will set the imageLoadedCh to the passed GameHlperFunction
//This is used to load images before a gamestate is set
func (g *Game) LoadImages(gFunction GameLoadImagesFunction) {
	go func() {
		/*var imageManager *ImageManager
		if imageManager = gFunction(); imageManager != nil {
			g.imageLoadedCh <- imageManager
			//close(g.imageLoadedCh)
		}*/
		imageManager := gFunction()
		g.imageLoadedCh <- imageManager

	}()

}

//LoadAudio will set the audioLoadedCh to the passed GameHelperFunction
//This is used to load audio before a gamestate is set
func (g *Game) LoadAudio(gFunction GameLoadAudioFunction) {
	go func() {
		/*var audioPlayer *AudioPlayer
		if audioPlayer = gFunction(); audioPlayer != nil {
			g.audioLoadedCh <- audioPlayer
			//close(g.audioLoadedCh)
		}*/
		audioPlayer := gFunction()
		g.audioLoadedCh <- audioPlayer
	}()
}

//ScreenSize returns the width and height of the game screen
func (g Game) ScreenSize() (w, h int) {
	w, h = g.Screen.Size()
	return w, h
}

//ScreenWidth returns the width of the game screen
func (g Game) ScreenWidth() int {
	w, _ := g.Screen.Size()
	return w
}

//ScreenHeight returns the height of the game screen
func (g Game) ScreenHeight() int {
	_, h := g.Screen.Size()
	return h
}
