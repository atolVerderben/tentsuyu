package tentsuyu

import (
	"image"
	"log"

	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

//TextElement contains the font, text, and position of that current text element
type TextElement struct {
	font                 *truetype.Font
	Name                 string
	drawImage            *ebiten.Image
	fntSize, fntDpi      float64
	text, prevText       []string
	visible              bool
	Stationary           bool
	textColor, origColor color.Color
	*BasicUIElement
}

//NewTextElement returns a new TextElement and creates the image
func NewTextElement(x, y float64, w, h int, font *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextElement {
	textImage, err := ebiten.NewImage(w, h, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	t := &TextElement{
		font:           font,
		fntSize:        fntSize,
		fntDpi:         72,
		drawImage:      textImage,
		text:           text,
		prevText:       []string{},
		BasicUIElement: NewBasicUIElement(x, y, w, h),
		textColor:      textColor,
		origColor:      textColor,
		visible:        true,
	}
	t.drawText(t.text)
	return t
}

//NewTextElementStationary returns a new TextElement and creates the image
func NewTextElementStationary(x, y float64, w, h int, font *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextElement {
	textImage, err := ebiten.NewImage(w, h, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	t := &TextElement{
		font:           font,
		fntSize:        fntSize,
		fntDpi:         72,
		drawImage:      textImage,
		text:           text,
		prevText:       []string{},
		BasicUIElement: NewBasicUIElement(x, y, w, h),
		textColor:      textColor,
		origColor:      textColor,
		visible:        true,
		Stationary:     true,
	}
	t.drawText(t.text)
	return t
}

func (t *TextElement) Hide() {
	t.visible = false
}

func (t *TextElement) Show() {
	t.visible = true
}

func (t *TextElement) Highlighted() bool {
	t.textColor = color.RGBA{153, 153, 0, 255}
	t.drawText(t.text)

	return true
}

func (t *TextElement) UnHighlighted() bool {
	t.textColor = t.origColor
	t.drawText(t.text)

	return true
}

func (t *TextElement) SetFontSize(fntSize float64) {
	t.fntSize = fntSize
}

func (t *TextElement) drawText(text []string) error {
	w, h := t.GetSize()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	const size = 24
	const dpi = 72
	d := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(t.textColor), //image.White,
		Face: truetype.NewFace(t.font, &truetype.Options{
			Size:    t.fntSize,
			DPI:     t.fntDpi,
			Hinting: font.HintingFull,
		}),
	}
	highlight := color.White
	if t.textColor != color.Black {
		highlight = color.Black
	}
	d2 := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(highlight),
		Face: truetype.NewFace(t.font, &truetype.Options{
			Size:    t.fntSize,
			DPI:     t.fntDpi,
			Hinting: font.HintingFull,
		}),
	}
	y := t.fntSize
	for _, s := range text {
		if t.fntSize > 16 {
			d2.Dot = fixed.P(+2, int(y+2))
			d2.DrawString(s)
		}
		d.Dot = fixed.P(0, int(y))
		d.DrawString(s)
		y += t.fntSize
	}

	return t.drawImage.ReplacePixels(dst.Pix)

}

//Update TextElement
func (t *TextElement) Update() {
	if !testEq(t.text, t.prevText) {
		t.drawText(t.text)
		t.prevText = t.text
	}
}

//SetText of the TextElement
func (t *TextElement) SetText(text []string) {
	t.prevText = t.text
	t.text = text
	t.drawText(t.text)
}

//SetColor of the TextElement
func (t *TextElement) SetColor(color color.Color) {
	t.textColor = color
	t.origColor = color
	t.drawText(t.text)
}

//ReturnText returns the text for debugging
func (t *TextElement) ReturnText() string {
	return t.text[0]
}

func (t *TextElement) SetPosition(x, y float64) {
	t.X = x
	t.Y = y
}

//Draw the TextElement
func (t *TextElement) Draw(screen *ebiten.Image) error {
	if t.visible == false {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.GetPosition())
	//GameCamera.DrawCameraTransform(op)
	if !t.Stationary {
		ApplyCameraTransform(op, false)
	}
	if err := screen.DrawImage(t.drawImage, op); err != nil {
		return err
	}
	return nil
}

func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
