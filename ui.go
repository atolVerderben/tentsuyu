package tentsuyu

import (
	"io/ioutil"
	"os"
	"strconv"

	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
)

//UIElement is the interface for all UI elements
type UIElement interface {
	Update()
	Draw(*ebiten.Image) error
	Highlighted() bool
	UnHighlighted() bool
	AddPosition(float64, float64)
	SetPosition(float64, float64)
	Size() (int, int)
	SetSize(int, int)
	Contains(float64, float64) bool
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

//AddFont adds a new truetype font to the map
func (ui *UIController) AddFont(fntName, fntFileLoc string) error {

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

func (ui *UIController) AddFontFile(fntName string, tt *truetype.Font) {
	ui.fonts[fntName] = tt
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

func (ui *UIController) ActiveMenu() string {
	for name, value := range ui.menus {
		if value.Active {
			return name
		}
	}
	return "None Active"
}

func (ui *UIController) TextElementExists(name string) bool {
	if ui.textElements[name] != nil {
		return true
	}
	return false
}

func (ui *UIController) HideTextElement(name string) {
	ui.textElements[name].Hide()
}
func (ui *UIController) ShowTextElement(name string) {
	ui.textElements[name].Show()
}

//WriteText creates a new TextElement
func (ui *UIController) WriteText(text []string, name, font string, x, y float64, w, h int, textColor color.Color, fntSize float64) {
	t := NewTextElement(x, y, w, h, ui.fonts[font], text, textColor, fntSize)

	ui.textElements[name] = t
}

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
			X:      ui.screenWidth / 2,
			Y:      ui.screenHeight / 2,
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
	ebiten.SetCursorVisibility(false)
}

//HideMouse both custom and default
func (ui *UIController) HideMouse() {
	ebiten.SetCursorVisibility(false)
	ui.DrawCursor = false
}

//ShowMouse shows the mouse, custom or default
func (ui *UIController) ShowMouse() {
	ebiten.SetCursorVisibility(true)
	ui.DrawCursor = true
}

//Draw all ui elements
func (ui UIController) Draw(screen *ebiten.Image) error {

	for _, t := range ui.textElements {
		t.Draw(screen)
	}
	for _, m := range ui.menus {
		if m.Active {
			m.Draw(screen)
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
			ui.menus[i].Update(ui.input)
		}
	}
	if ui.customCursor == true {
		ui.Cursor.Update(ui.input.GetMouseCoords()) //Input.Mouse.X, Input.Mouse.Y)
	}
}

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

//UINumberDisplay allows a pointer to a float64 that updates and draws a TextElement on update
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

//NewUINumberDisplay creates a new UINumberDisplay
func NewUINumberDisplayInt(number *int, x, y float64, w, h int, font *truetype.Font, fntSize float64, textColor color.Color) *UINumberDisplayInt {
	nd := &UINumberDisplayInt{
		currNumber:  number,
		prevNumber:  *number,
		TextElement: NewTextElement(x, y, w, h, font, []string{strconv.Itoa(int(*number))}, textColor, fntSize),
	}
	return nd
}

//UINumberDisplay allows a pointer to a float64 that updates and draws a TextElement on update
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

//NewUINumberDisplay creates a new UINumberDisplay
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

//NewUINumberDisplayStationary creates a new UINumberDisplay
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
	obj := &BasicObject{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
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
	u.X += x
	u.Y += y
}

//Size returns the elements size
func (u BasicUIElement) Size() (int, int) {
	return u.GetSize()
}

//Contains the specified x and y position
func (u *BasicUIElement) Contains(x, y float64) bool {
	return u.ContainsNoCenter(x, y)
}
