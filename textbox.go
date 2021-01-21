package tentsuyu

import (
	"image/color"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

//TextBox represents a large text area
type TextBox struct {
	*BasicUIElement
	Text      *TextElement
	Selected  bool
	delayTick int
}

//NewTextBox returns a pointer to a TextBox struct using the given parameters
func NewTextBox(x, y float64, w, h int, font *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextBox {
	tb := &TextBox{
		Text:           NewTextElement(x, y, w, h, font, text, textColor, fntSize),
		BasicUIElement: NewBasicUIElement(x, y, w, h),
	}
	return tb
}

//Update the TextBox checkts whether it is selecter and what keystrokes have happened
func (tb *TextBox) Update(input *InputController) {

	tb.Text.Update()

	if tb.Text.Contains(input.Mouse.X, input.Mouse.Y) {
		if input.LeftClick().JustPressed() {
			if !tb.Selected {
				tb.Text.Highlighted()
				tb.Selected = true
			}
		}
	} else {
		if input.LeftClick().JustPressed() {
			tb.Text.UnHighlighted()
			tb.Selected = false
		}
	}
	if tb.Selected {
		tb.Text.text[0] += string(ebiten.InputChars())
		tb.Text.SetText(tb.Text.text)
		if input.keyManager.Get(ebiten.KeyEnter).JustPressed() || input.keyManager.Get(ebiten.KeyKPEnter).JustPressed() {
			tb.Text.UnHighlighted()
			tb.Selected = false
		}
		if input.keyManager.Get(ebiten.KeyBackspace).JustPressed() {
			if utf8.RuneCountInString(tb.Text.text[0]) > 0 {
				tb.Text.text[0] = tb.Text.text[0][:len(tb.Text.text[0])-1]
				tb.Text.SetText(tb.Text.text)
			}
		}

	}

}

//Draw the textbox
func (tb *TextBox) Draw(screen *ebiten.Image, camera *Camera) error {
	return tb.Text.Draw(screen, camera)
}
