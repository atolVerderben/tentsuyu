package tentsuyu

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	txt "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

//TextElement contains the font, text, and position of that current text element
type TextElement struct {
	dropShadow                                      bool
	font                                            *truetype.Font
	Name                                            string
	drawImage                                       *ebiten.Image
	fntSize, fntDpi                                 float64
	text, prevText                                  []string
	visible                                         bool
	Stationary                                      bool
	textColor, origColor, highlightColor, dropColor color.Color
	*BasicUIElement
	fntFace font.Face
}

//NewTextElement returns a new TextElement and creates the image
func NewTextElement(x, y float64, w, h int, fnt *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextElement {
	textImage := ebiten.NewImage(w, h)

	t := &TextElement{
		font:           fnt,
		fntSize:        fntSize,
		fntDpi:         72,
		drawImage:      textImage,
		text:           text,
		prevText:       []string{},
		BasicUIElement: NewBasicUIElement(x, y, w, h),
		textColor:      textColor,
		origColor:      textColor,
		visible:        true,
		highlightColor: color.RGBA{153, 153, 0, 255},
		dropShadow:     true,
		dropColor:      color.Black,
	}
	if t.textColor == color.Black {
		t.dropColor = color.White
	}
	t.fntFace = truetype.NewFace(t.font, &truetype.Options{
		Size:    t.fntSize,
		DPI:     t.fntDpi,
		Hinting: font.HintingNone,
	})
	t.drawText(t.text)
	t.SetCentered(false)
	return t
}

func NewTextElementCentered(x, y float64, w, h int, fnt *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextElement {
	t := NewTextElement(x, y, w, h, fnt, text, textColor, fntSize)
	t.SetCentered(true)
	return t
}

//NewTextElementStationary returns a new TextElement and creates the image
func NewTextElementStationary(x, y float64, w, h int, fnt *truetype.Font, text []string, textColor color.Color, fntSize float64) *TextElement {
	textImage := ebiten.NewImage(w, h)

	t := &TextElement{
		font:           fnt,
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
		highlightColor: color.RGBA{153, 153, 0, 255},
		dropColor:      color.Black,
	}
	if t.textColor == color.Black {
		t.dropColor = color.White
	}
	t.fntFace = truetype.NewFace(t.font, &truetype.Options{
		Size:    t.fntSize,
		DPI:     t.fntDpi,
		Hinting: font.HintingNone,
	})
	t.drawText(t.text)
	return t
}

//SetHighlightColor sets the color of the text element when it's highlighted
func (t *TextElement) SetHighlightColor(c color.Color) {
	t.highlightColor = c
}

//SetTextColor sets the color of the text element
func (t *TextElement) SetTextColor(c color.Color) {
	t.textColor = c
	t.origColor = c
}

//Hide the TextElement
func (t *TextElement) Hide() {
	t.visible = false
}

//Show the TextElement
func (t *TextElement) Show() {
	t.visible = true
}

func (t *TextElement) GetDrawImage() *ebiten.Image {
	return t.drawImage
}

//Highlighted sets the TextElement to its highlighted color
func (t *TextElement) Highlighted() bool {
	t.textColor = t.highlightColor
	t.drawText(t.text)

	return true
}

//UnHighlighted returns the TextElement to its original color
func (t *TextElement) UnHighlighted() bool {
	t.textColor = t.origColor
	t.drawText(t.text)

	return true
}

//SetFontSize of the TextElement
func (t *TextElement) SetFontSize(fntSize float64) {
	t.fntSize = fntSize
}

func (t *TextElement) drawText(text []string) error {
	t.drawImage.Clear()

	tx := ""
	for i := range text {
		tx += text[i]
		//Add a newline if multiple lines
		if i != len(text)-1 {
			tx += "\n"
		}
	}
	txt.Draw(t.drawImage, tx, t.fntFace, 0+3, int(t.fntSize)+3, t.dropColor)
	txt.Draw(t.drawImage, tx, t.fntFace, 0, int(t.fntSize), t.textColor)

	return nil
	/*w, h := t.GetSize()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	d := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(t.textColor), //image.White,
		Face: truetype.NewFace(t.font, &truetype.Options{
			Size: t.fntSize,
			DPI:  t.fntDpi,
			//Hinting: font.HintingFull,
		}),
	}
	highlight := color.Black
	/*if t.textColor != color.Black {
		highlight = color.Black
	}*/
	/*d2 := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(highlight),
		Face: truetype.NewFace(t.font, &truetype.Options{
			Size: t.fntSize,
			DPI:  t.fntDpi,
			//Hinting: font.HintingFull,
		}),
	}
	y := t.fntSize
	for _, s := range text {
		if t.dropShadow {
			d2.Dot = fixed.P(+2, int(y+2))
			d2.DrawString(s)
		}
		d.Dot = fixed.P(0, int(y))
		d.DrawString(s)
		y += t.fntSize
	}

	return t.drawImage.ReplacePixels(dst.Pix)
	*/
}

//SetDropShadow of the TextElement. If true then a second outline will be drawn.
func (t *TextElement) SetDropShadow(drop bool) {
	t.dropShadow = drop
}

//SetDropShadowColor sets the color to be used for the dropshadow
func (t *TextElement) SetDropShadowColor(color color.Color) {
	t.dropColor = color
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

//SetPosition of TextElement to given x,y coords
func (t *TextElement) SetPosition(x, y float64) {
	t.Position.X = x
	t.Position.Y = y
}

//Draw the TextElement
func (t *TextElement) Draw(screen *ebiten.Image, camera *Camera) error {
	if t.visible == false {
		return nil
	}

	/*txt.Draw(screen, t.text[0], truetype.NewFace(t.font, &truetype.Options{
		Size:    t.fntSize,
		DPI:     t.fntDpi,
		Hinting: font.HintingNone,
	}), int(t.GetX()), int(t.GetY()), t.textColor)
	*/
	op := &ebiten.DrawImageOptions{}
	//move to center to rotate
	op.GeoM.Translate(-t.GetWidthF()/2, -t.GetHeightF()/2)
	op.GeoM.Rotate(t.GetAngle())
	if t.NotCentered {
		//move back to corner for placement
		op.GeoM.Translate(t.GetWidthF()/2, t.GetHeightF()/2)
	}
	op.GeoM.Translate(t.GetPosition())
	//GameCamera.DrawCameraTransform(op)
	if !t.Stationary {
		camera.ApplyCameraTransform(op, true)
	}
	screen.DrawImage(t.drawImage, op)
	return nil
}

//DrawPosition applies the camera transform to the text element
func (t *TextElement) DrawPosition(screen *ebiten.Image, camera *Camera) error {
	if t.visible == false {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.GetPosition())
	//GameCamera.DrawCameraTransform(op)
	if !t.Stationary {
		camera.ApplyCameraTransform(op, true)
	}
	screen.DrawImage(t.drawImage, op)
	return nil
}

//DrawApplyZoom the TextElement
func (t *TextElement) DrawApplyZoom(screen *ebiten.Image) error {
	if t.visible == false {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.GetPosition())
	//GameCamera.DrawCameraTransform(op)
	if !t.Stationary {
		//ApplyCameraTransform(op, true)
	}
	screen.DrawImage(t.drawImage, op)
	return nil
}

//SetCentered sets whether the TextElement is centered with bool value c
func (t *TextElement) SetCentered(c bool) {
	t.NotCentered = !c
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
