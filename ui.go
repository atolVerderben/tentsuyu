package tentsuyu

import (
	"io/ioutil"
	"os"
	"strconv"

	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

//UIElement is the interface for all UI elements
type UIElement interface {
	Update()
	Draw(*ebiten.Image, *Camera) error
	Highlighted() bool
	UnHighlighted() bool
	AddPosition(float64, float64)
	SetPosition(float64, float64)
	Size() (int, int)
	SetSize(int, int)
	Contains(float64, float64) bool
	SetHighlightColor(color.Color)
}

//UIController controls all UI elements
type UIController struct {
	Cursor                    *Cursor
	DrawCursor, customCursor  bool
	screenWidth, screenHeight float64
	fonts                     map[string]*truetype.Font
	textElements              map[string]*TextElement
	menus                     map[string]*Menu
	input                     *InputController
}

//AddFontFile adds a new truetype font from the given file location with the given name
func (ui *UIController) AddFontFile(fntName, fntFileLoc string) error {

	f, err := os.Open(fntFileLoc) //ebitenutil.OpenFile(fntFileLoc)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	tt, err := truetype.Parse(b)
	if err != nil {
		return err
	}
	//log.Print(tt)
	ui.fonts[fntName] = tt

	return nil
}

//AddFont adds the given font with the given name to the ui fonts list
func (ui *UIController) AddFont(fntName string, tt *truetype.Font) {
	ui.fonts[fntName] = tt
}

//AddFontFromBytes adds the font through a byte slice
func (ui *UIController) AddFontFromBytes(fntName string, bytes []byte) error {
	fnt, err := truetype.Parse(bytes)
	if err != nil {
		return err
	}
	ui.AddFont(fntName, fnt)
	return nil
}

//AddMenu to the UIController
func (ui *UIController) AddMenu(name string, menu *Menu) {
	ui.menus[name] = menu
}

//ToggleMenu if the menu is on it turns off else it turns on
func (ui *UIController) ToggleMenu(name string) {
	ui.menus[name].Active = !ui.menus[name].Active

}

//ActivateMenu activates specified menu
func (ui *UIController) ActivateMenu(name string) {
	ui.menus[name].Active = true

}

//DeActivateMenu deactivates specified menu
func (ui *UIController) DeActivateMenu(name string) {
	ui.menus[name].Active = false

}

//ActiveMenu returns the name of the active menu.
//Returns "None Active" if there is no active menu.
func (ui *UIController) ActiveMenu() string {
	for name, value := range ui.menus {
		if value.Active {
			return name
		}
	}
	return "None Active"
}

//TextElementExists returns true if the given name is found in the UIController's text elements
func (ui *UIController) TextElementExists(name string) bool {
	if ui.textElements[name] != nil {
		return true
	}
	return false
}

//HideTextElement hides the named element
func (ui *UIController) HideTextElement(name string) {
	ui.textElements[name].Hide()
}

//ShowTextElement shows the named element
func (ui *UIController) ShowTextElement(name string) {
	ui.textElements[name].Show()
}

//WriteText creates a new TextElement
func (ui *UIController) WriteText(text []string, name, font string, x, y float64, w, h int, textColor color.Color, fntSize float64) {
	t := NewTextElement(x, y, w, h, ui.fonts[font], text, textColor, fntSize)

	ui.textElements[name] = t
}

//UpdateTextPosition sets the position of the named text element
func (ui *UIController) UpdateTextPosition(name string, x, y float64) {
	ui.textElements[name].SetPosition(x, y)
}

//ReturnFont returns the truetype font with the specified name
func (ui *UIController) ReturnFont(name string) *truetype.Font {
	return ui.fonts[name]
}

//NewUIController creates a default UI controller
func NewUIController(input *InputController) *UIController {
	ui := &UIController{
		DrawCursor:   true,
		fonts:        make(map[string]*truetype.Font),
		textElements: make(map[string]*TextElement),
		menus:        make(map[string]*Menu),
		input:        input,
	}
	return ui
}

//SetCustomCursor allows the addition of a cursor image
func (ui *UIController) SetCustomCursor(width, height, sx, sy int, spritesheet *ebiten.Image) {
	c := &Cursor{
		BasicObject: &BasicObject{
			Position: &Vector2d{
				X: ui.screenWidth / 2,
				Y: ui.screenHeight / 2,
			},

			Width:  width,
			Height: height,
		},
		BasicImageParts: &BasicImageParts{
			Sx:     sx,
			Sy:     sy,
			Width:  width,
			Height: height,
		},
		style:       CursorCrosshair,
		spritesheet: spritesheet,
	}
	ui.DrawCursor = true
	ui.customCursor = true
	ui.Cursor = c
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

//HideMouse both custom and default
func (ui *UIController) HideMouse() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ui.DrawCursor = false
}

//ShowMouse shows the mouse, custom or default
func (ui *UIController) ShowMouse() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ui.DrawCursor = true
}

//Draw all ui elements
func (ui UIController) Draw(screen *ebiten.Image, camera *Camera) error {

	for _, t := range ui.textElements {
		t.Draw(screen, camera)
	}
	for _, m := range ui.menus {
		if m.Active {
			m.Draw(screen, camera)
		}
	}
	if ui.customCursor == true && ui.DrawCursor == true {
		ui.Cursor.Draw(screen)
	}
	return nil
}

//Update all ui elements
func (ui *UIController) Update() {

	for _, t := range ui.textElements {
		t.Update()
	}
	for i := range ui.menus {
		if ui.menus[i].Active {
			ui.menus[i].Update(ui.input, 0, 0) //TODO: I need to update this to work without offsets
		}
	}
	if ui.customCursor == true {
		ui.Cursor.Update(ui.input.GetMouseCoords()) //Input.Mouse.X, Input.Mouse.Y)
	}
}

//AddTextDisplay adds the passed TextElement with the given name to the UIController
func (ui *UIController) AddTextDisplay(name string, textElement *TextElement) {
	ui.textElements[name] = textElement
}

//UINumberDisplay allows a pointer to a float64 that updates and draws a TextElement on update
type UINumberDisplay struct {
	*TextElement
	currNumber *float64
	prevNumber float64
}

//Update UINumberDisplay
func (nd *UINumberDisplay) Update() {
	if *nd.currNumber != nd.prevNumber {
		nd.prevNumber = *nd.currNumber
		nd.TextElement.SetText([]string{strconv.Itoa(int(nd.prevNumber))})
	}

	nd.TextElement.Update()
}

//NewUINumberDisplay creates a new UINumberDisplay
func NewUINumberDisplay(number *float64, x, y float64, w, h int, font *truetype.Font, textColor color.Color, fntSize float64) *UINumberDisplay {
	nd := &UINumberDisplay{
		currNumber:  number,
		prevNumber:  *number,
		TextElement: NewTextElement(x, y, w, h, font, []string{strconv.Itoa(int(*number))}, textColor, fntSize),
	}
	return nd
}

//UINumberDisplayInt allows a pointer to an integer that updates and draws a TextElement on update
type UINumberDisplayInt struct {
	*TextElement
	currNumber *int
	prevNumber int
}

//Update UINumberDisplay
func (nd *UINumberDisplayInt) Update() {
	if *nd.currNumber != nd.prevNumber {
		nd.prevNumber = *nd.currNumber
		nd.TextElement.SetText([]string{strconv.Itoa(int(nd.prevNumber))})
	}

	nd.TextElement.Update()
}

//NewUINumberDisplayInt creates a new UINumberDisplayInt
func NewUINumberDisplayInt(number *int, x, y float64, w, h int, font *truetype.Font, fntSize float64, textColor color.Color) *UINumberDisplayInt {
	nd := &UINumberDisplayInt{
		currNumber:  number,
		prevNumber:  *number,
		TextElement: NewTextElement(x, y, w, h, font, []string{strconv.Itoa(int(*number))}, textColor, fntSize),
	}
	return nd
}

//UITextDisplay shows a changable text element on the display
type UITextDisplay struct {
	*TextElement
	currText *string
	prevText string
}

//Update UINumberDisplay
func (td *UITextDisplay) Update() {
	if *td.currText != td.prevText {
		td.prevText = *td.currText
		td.TextElement.SetText([]string{td.prevText})
	}

	td.TextElement.Update()
}

//NewUITextDisplay creates a new UITExtDisplay
func NewUITextDisplay(text *string, x, y float64, w, h int, font *truetype.Font, textColor color.Color, fntSize float64) *UITextDisplay {
	nd := &UITextDisplay{
		currText:    text,
		prevText:    *text,
		TextElement: NewTextElement(x, y, w, h, font, []string{*text}, textColor, fntSize),
	}
	nd.TextElement.fntSize = fntSize
	return nd
}

//NewUINumberDisplayStationary creates a new UINumberDisplay
func NewUINumberDisplayStationary(number *float64, x, y float64, w, h int, font *truetype.Font, textColor color.Color) *UINumberDisplay {
	nd := &UINumberDisplay{
		currNumber:  number,
		prevNumber:  *number,
		TextElement: NewTextElementStationary(x, y, w, h, font, []string{strconv.Itoa(int(*number))}, textColor, 24),
	}
	return nd
}

//NewUINumberDisplayIntStationary creates a new UINumberDisplayInt
func NewUINumberDisplayIntStationary(number *int, x, y float64, w, h int, font *truetype.Font, textColor color.Color, fntSize float64) *UINumberDisplayInt {
	nd := &UINumberDisplayInt{
		currNumber:  number,
		prevNumber:  *number,
		TextElement: NewTextElementStationary(x, y, w, h, font, []string{strconv.Itoa(int(*number))}, textColor, fntSize),
	}
	return nd
}

//BasicUIElement ================================================================================

//BasicUIElement is so I don't have to rewrite the same functions
type BasicUIElement struct {
	isHighlited bool
	*BasicObject
}

//NewBasicUIElement  creates a BasicUIElement
func NewBasicUIElement(x, y float64, w, h int) *BasicUIElement {
	obj := NewBasicObject(x, y, w, h)
	u := &BasicUIElement{
		BasicObject: obj,
	}
	return u
}

//Highlighted returns if the element is Highlighted
func (u *BasicUIElement) Highlighted() bool {
	return u.isHighlited
}

//UnHighlighted returns if the element is Highlighted
func (u *BasicUIElement) UnHighlighted() bool {
	return u.isHighlited
}

//AddPosition adds the specified values to the x and y position
func (u *BasicUIElement) AddPosition(x, y float64) {
	u.Position.X += x
	u.Position.Y += y
}

//Size returns the elements size
func (u BasicUIElement) Size() (int, int) {
	return u.GetSize()
}

//Contains the specified x and y position
func (u *BasicUIElement) Contains(x, y float64) bool {
	return u.BasicObject.Contains(x, y)
}
