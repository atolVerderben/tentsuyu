package tentsuyu

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
	"golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/gofont/gomonobolditalic"
	"golang.org/x/image/font/gofont/gomonoitalic"
	"golang.org/x/image/font/gofont/goregular"
)

//Available Go Fonts
const (
	FntRegular        string = "FontGoRegular"
	FntMono                  = "FontGoMono"
	FntBold                  = "FontGoBold"
	FntItalic                = "FontGoItalic"
	FntBoldItalic            = "FontGoBoldItalic"
	FntMonoBold              = "FontMonoBold"
	FntMonoItalic            = "FontMonoItalic"
	FntMonoBoldItalic        = "FontMonoBoldItalic"
)

//LoadDefaultFonts is a method that loads the Go fonts into the UIController.
func (ui *UIController) LoadDefaultFonts() {
	font, _ := truetype.Parse(goregular.TTF)
	ui.AddFont(FntRegular, font)

	font1, _ := truetype.Parse(gomono.TTF)
	ui.AddFont(FntMono, font1)

	font2, _ := truetype.Parse(gobold.TTF)
	ui.AddFont(FntBold, font2)

	font3, _ := truetype.Parse(goitalic.TTF)
	ui.AddFont(FntItalic, font3)

	font4, _ := truetype.Parse(gobolditalic.TTF)
	ui.AddFont(FntBoldItalic, font4)

	font5, _ := truetype.Parse(gomonobold.TTF)
	ui.AddFont(FntMonoBold, font5)

	font6, _ := truetype.Parse(gomonoitalic.TTF)
	ui.AddFont(FntMonoItalic, font6)

	font7, _ := truetype.Parse(gomonobolditalic.TTF)
	ui.AddFont(FntMonoBoldItalic, font7)
}

//LoadDefaultFonts adds the Go supplied fonts to the tentsuyu UIController
func LoadDefaultFonts(uiController *UIController) {
	font, _ := truetype.Parse(goregular.TTF)
	uiController.AddFont(FntRegular, font)

	font1, _ := truetype.Parse(gomono.TTF)
	uiController.AddFont(FntMono, font1)

	font2, _ := truetype.Parse(gobold.TTF)
	uiController.AddFont(FntBold, font2)

	font3, _ := truetype.Parse(goitalic.TTF)
	uiController.AddFont(FntItalic, font3)

	font4, _ := truetype.Parse(gobolditalic.TTF)
	uiController.AddFont(FntBoldItalic, font4)

	font5, _ := truetype.Parse(gomonobold.TTF)
	uiController.AddFont(FntMonoBold, font5)

	font6, _ := truetype.Parse(gomonoitalic.TTF)
	uiController.AddFont(FntMonoItalic, font6)

	font7, _ := truetype.Parse(gomonobolditalic.TTF)
	uiController.AddFont(FntMonoBoldItalic, font7)
}
