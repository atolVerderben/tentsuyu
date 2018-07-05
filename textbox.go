package tentsuyu

import (
	"image/color"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
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
func (tb *TextBox) Update() {

	tb.Text.Update()

	if tb.Text.Contains(Input.Mouse.X, Input.Mouse.Y) {
		if Input.LeftClick().JustPressed() {
			if !tb.Selected {
				tb.Text.Highlighted()
				tb.Selected = true
			}
		}
	} else {
		if Input.LeftClick().JustPressed() {
			tb.Text.UnHighlighted()
			tb.Selected = false
		}
	}
	if tb.Selected {
		tb.Text.text[0] += string(ebiten.InputChars())
		tb.Text.SetText(tb.Text.text)
		if Input.keyManager.Get(ebiten.KeyEnter).JustPressed() || Input.keyManager.Get(ebiten.KeyKPEnter).JustPressed() {
			tb.Text.UnHighlighted()
			tb.Selected = false
		}
		if Input.keyManager.Get(ebiten.KeyBackspace).JustPressed() {
			if utf8.RuneCountInString(tb.Text.text[0]) > 0 {
				tb.Text.text[0] = tb.Text.text[0][:len(tb.Text.text[0])-1]
				tb.Text.SetText(tb.Text.text)
			}
		}

	}

}

//Draw the textbox
func (tb *TextBox) Draw(screen *ebiten.Image) error {
	return tb.Text.Draw(screen)
}

/*

	if Input.Button("TextBox0").JustPressed() {
		tb.Text.text[0] += "0"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox1").JustPressed() {
		tb.Text.text[0] += "1"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox2").JustPressed() {
		tb.Text.text[0] += "2"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox3").JustPressed() {
		tb.Text.text[0] += "3"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox4").JustPressed() {
		tb.Text.text[0] += "4"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox5").JustPressed() {
		tb.Text.text[0] += "5"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox6").JustPressed() {
		tb.Text.text[0] += "6"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox7").JustPressed() {
		tb.Text.text[0] += "7"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox8").JustPressed() {
		tb.Text.text[0] += "8"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBox9").JustPressed() {
		tb.Text.text[0] += "9"
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxPeriod").JustPressed() {
		tb.Text.text[0] += "."
		tb.Text.SetText(tb.Text.text)
	}

	//Letters
	if Input.Button("TextBoxA").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "A"
		} else {
			tb.Text.text[0] += "a"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxB").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "B"
		} else {
			tb.Text.text[0] += "b"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxC").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "C"
		} else {
			tb.Text.text[0] += "c"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxD").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "D"
		} else {
			tb.Text.text[0] += "d"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxE").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "E"
		} else {
			tb.Text.text[0] += "e"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxF").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "F"
		} else {
			tb.Text.text[0] += "f"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxG").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "G"
		} else {
			tb.Text.text[0] += "g"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxH").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "H"
		} else {
			tb.Text.text[0] += "h"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxI").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "I"
		} else {
			tb.Text.text[0] += "i"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxJ").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "J"
		} else {
			tb.Text.text[0] += "j"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxK").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "K"
		} else {
			tb.Text.text[0] += "k"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxL").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "L"
		} else {
			tb.Text.text[0] += "l"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxM").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "M"
		} else {
			tb.Text.text[0] += "m"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxN").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "N"
		} else {
			tb.Text.text[0] += "n"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxO").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "O"
		} else {
			tb.Text.text[0] += "o"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxP").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "P"
		} else {
			tb.Text.text[0] += "p"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxQ").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "Q"
		} else {
			tb.Text.text[0] += "q"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxR").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "R"
		} else {
			tb.Text.text[0] += "r"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxS").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "S"
		} else {
			tb.Text.text[0] += "s"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxT").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "T"
		} else {
			tb.Text.text[0] += "t"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxU").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "U"
		} else {
			tb.Text.text[0] += "u"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxV").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "V"
		} else {
			tb.Text.text[0] += "v"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxW").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "W"
		} else {
			tb.Text.text[0] += "w"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxX").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "X"
		} else {
			tb.Text.text[0] += "x"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxY").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "Y"
		} else {
			tb.Text.text[0] += "y"
		}
		tb.Text.SetText(tb.Text.text)
	}
	if Input.Button("TextBoxZ").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += "Z"
		} else {
			tb.Text.text[0] += "z"
		}
		tb.Text.SetText(tb.Text.text)
	}

	if Input.Button("TextBoxSemiColon").JustPressed() {
		if Input.Button("TextBoxShift").Down() {
			tb.Text.text[0] += ":"
		} else {
			tb.Text.text[0] += ";"
		}
		tb.Text.SetText(tb.Text.text)
	}

	if Input.Button("TextBoxBackspace").JustPressed() {
		if utf8.RuneCountInString(tb.Text.text[0]) > 0 {
			tb.Text.text[0] = tb.Text.text[0][:len(tb.Text.text[0])-1]
			tb.Text.SetText(tb.Text.text)
		}
	}
	if Input.Button("TextBoxEnter").JustPressed() {
		tb.Text.UnHighlighted()
		tb.Selected = false
	}
*/
