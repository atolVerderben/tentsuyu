package tentsuyu

import (
	"image/color"
	"unicode/utf8"

	"github.com/golang/freetype/truetype"
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
func (ta *TextArea) Update() {
	if ta.Selected {
		if Input.Button("TextBox0").JustPressed() {
			ta.text[ta.Lines-1] += "0"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox1").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "!"
			} else {
				ta.text[ta.Lines-1] += "1"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox2").JustPressed() {
			ta.text[ta.Lines-1] += "2"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox3").JustPressed() {
			ta.text[ta.Lines-1] += "3"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox4").JustPressed() {
			ta.text[ta.Lines-1] += "4"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox5").JustPressed() {
			ta.text[ta.Lines-1] += "5"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox6").JustPressed() {
			ta.text[ta.Lines-1] += "6"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox7").JustPressed() {
			ta.text[ta.Lines-1] += "7"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox8").JustPressed() {
			ta.text[ta.Lines-1] += "8"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBox9").JustPressed() {
			ta.text[ta.Lines-1] += "9"
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxPeriod").JustPressed() {
			ta.text[ta.Lines-1] += "."
			ta.SetText(ta.text)
		}

		//Letters
		if Input.Button("TextBoxA").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "A"
			} else {
				ta.text[ta.Lines-1] += "a"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxB").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "B"
			} else {
				ta.text[ta.Lines-1] += "b"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxC").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "C"
			} else {
				ta.text[ta.Lines-1] += "c"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxD").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "D"
			} else {
				ta.text[ta.Lines-1] += "d"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxE").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "E"
			} else {
				ta.text[ta.Lines-1] += "e"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxF").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "F"
			} else {
				ta.text[ta.Lines-1] += "f"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxG").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "G"
			} else {
				ta.text[ta.Lines-1] += "g"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxH").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "H"
			} else {
				ta.text[ta.Lines-1] += "h"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxI").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "I"
			} else {
				ta.text[ta.Lines-1] += "i"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxJ").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "J"
			} else {
				ta.text[ta.Lines-1] += "j"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxK").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "K"
			} else {
				ta.text[ta.Lines-1] += "k"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxL").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "L"
			} else {
				ta.text[ta.Lines-1] += "l"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxM").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "M"
			} else {
				ta.text[ta.Lines-1] += "m"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxN").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "N"
			} else {
				ta.text[ta.Lines-1] += "n"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxO").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "O"
			} else {
				ta.text[ta.Lines-1] += "o"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxP").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "P"
			} else {
				ta.text[ta.Lines-1] += "p"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxQ").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "Q"
			} else {
				ta.text[ta.Lines-1] += "q"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxR").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "R"
			} else {
				ta.text[ta.Lines-1] += "r"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxS").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "S"
			} else {
				ta.text[ta.Lines-1] += "s"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxT").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "T"
			} else {
				ta.text[ta.Lines-1] += "t"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxU").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "U"
			} else {
				ta.text[ta.Lines-1] += "u"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxV").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "V"
			} else {
				ta.text[ta.Lines-1] += "v"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxW").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "W"
			} else {
				ta.text[ta.Lines-1] += "w"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxX").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "X"
			} else {
				ta.text[ta.Lines-1] += "x"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxY").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "Y"
			} else {
				ta.text[ta.Lines-1] += "y"
			}
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxZ").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += "Z"
			} else {
				ta.text[ta.Lines-1] += "z"
			}
			ta.SetText(ta.text)
		}

		if Input.Button("TextBoxSemiColon").JustPressed() {
			if Input.Button("TextBoxShift").Down() {
				ta.text[ta.Lines-1] += ":"
			} else {
				ta.text[ta.Lines-1] += ";"
			}
			ta.SetText(ta.text)
		}

		if Input.Button("TextBoxBackspace").JustPressed() {
			if utf8.RuneCountInString(ta.text[ta.Lines-1]) > 0 {
				ta.text[ta.Lines-1] = ta.text[ta.Lines-1][:len(ta.text[ta.Lines-1])-1]
				ta.SetText(ta.text)
			}
		}
		if Input.Button("TextBoxSpace").JustPressed() {
			ta.text[ta.Lines-1] += " "
			ta.SetText(ta.text)
		}
		if Input.Button("TextBoxEnter").JustPressed() {
			ta.NewLine()
		}
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
