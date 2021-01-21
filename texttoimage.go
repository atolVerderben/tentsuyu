package tentsuyu

import (
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

//RenderTextToImage takes a given slice of strings and returns a pointer to an ebiten.Image
func RenderTextToImage(text []string, w, h int, fntSize float64, fnt *truetype.Font, textColor color.Color) *ebiten.Image {

	drawImage := ebiten.NewImage(w, h)

	//w, h := t.GetSize()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	//const size = 24
	const dpi = 72
	d := &font.Drawer{
		Dst: dst,
		Src: image.NewUniform(textColor), //image.White,
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size:    fntSize,
			DPI:     dpi,
			Hinting: font.HintingNone,
		}),
	}
	/*highlight := color.White
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
	}*/
	y := fntSize
	for _, s := range text {
		//d2.Dot = fixed.P(+1, int(y+1))
		//d2.DrawString(s)
		d.Dot = fixed.P(0, int(y))
		d.DrawString(s)
		y += fntSize
	}

	drawImage.ReplacePixels(dst.Pix)
	return drawImage

}
