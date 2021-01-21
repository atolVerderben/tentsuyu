package tentsuyu

import (
	"image/color"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
)

//TextArea represents a typeable area
type TextArea struct {
	*BasicObject
	*TextElement
	Selected bool
	Lines    int
}

//NewTextArea returns a TextArea
func NewTextArea(x, y float64, w, h, lines int, font *truetype.Font, textColor color.Color, fntSize float64) *TextArea {
	ta := &TextArea{
		TextElement: NewTextElement(x, y, w, h, font, make([]string, lines), textColor, fntSize),
		Lines:       lines,
	}

	return ta
}

//Update the TextArea. If it's selected you can type in the area
func (ta *TextArea) Update(input *InputController) {
	if ta.Selected {

		if input.keyManager.Get(ebiten.KeyEnter).JustPressed() || input.keyManager.Get(ebiten.KeyKPEnter).JustPressed() {
			ta.NewLine()
		}
		if input.keyManager.Get(ebiten.KeyBackspace).JustPressed() {
			if utf8.RuneCountInString(ta.text[ta.Lines-1]) > 0 {
				ta.text[ta.Lines-1] = ta.text[ta.Lines-1][:len(ta.text[ta.Lines-1])-1]
				ta.SetText(ta.text)
			}
		}
		ta.text[ta.Lines-1] += string(ebiten.InputChars())
		ta.SetText(ta.text)

	}
	ta.TextElement.Update()
}

//NewLine adds a new line, removes the first element and shifts all the elements up
func (ta *TextArea) NewLine() {
	ta.text = append(ta.text[:0], ta.text[0+1:]...)
	ta.text = append(ta.text, "")
	ta.SetText(ta.text)
}

//AddLine adds a line of text to the TextArea as if it were entered
func (ta *TextArea) AddLine(text string) {
	ta.text[ta.Lines-1] += text
	ta.NewLine()
}

//ReturnLastEntered Returns the line that was just entered
func (ta *TextArea) ReturnLastEntered() string {
	return ta.text[ta.Lines-2]
}
