package tentsuyu

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//BasicImageParts is easy to set up basic sprite image
type BasicImageParts struct {
	Width, Height, Sx, Sy, DestWidth, DestHeight int
	Reverse                                      bool
	SourceRect                                   *image.Rectangle
}

//NewBasicImageParts returns a pointer to new BasicImageParts
func NewBasicImageParts(sx, sy, width, height int) *BasicImageParts {
	b := &BasicImageParts{
		Sx:         sx,
		Sy:         sy,
		Width:      width,
		Height:     height,
		DestHeight: height,
		DestWidth:  width,
	}
	return b
}

//ReturnSourceRect returns the image.Rectangle for the subImage in ebtien
//This replaces the overall ImageParts struct
func (b *BasicImageParts) ReturnSourceRect() image.Rectangle {
	if b.Reverse {
		return image.Rect((b.Sx + b.Width), (b.Sy), (b.Sx), (b.Sy + b.Height))
	}
	return image.Rect((b.Sx), (b.Sy), (b.Sx + b.Width), (b.Sy + b.Height))
}

//SetDestinationDimensions can be used to set the size the image should be drawn to the screen
func (b *BasicImageParts) SetDestinationDimensions(width, height int) {
	b.DestWidth = width
	b.DestHeight = height
}

//ReverseX flips the image
func (b *BasicImageParts) ReverseX(reverse bool) {
	b.Reverse = reverse
}

//Len returns 1
func (b *BasicImageParts) Len() int {
	return 1
}

//Dst we just make it 1:1
func (b *BasicImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	if b.DestHeight == 0 && b.DestWidth == 0 {
		return 0, 0, b.Width, b.Height
	}
	return 0, 0, b.DestWidth, b.DestHeight
}

//Src cuts out the specified rectangle from the source image to display the sprite
func (b *BasicImageParts) Src(i int) (x0, y0, x1, y1 int) {
	x := b.Sx
	y := b.Sy
	if b.Reverse {
		return x + b.Width, y, x, y + b.Height
	}
	return x, y, x + b.Width, y + b.Height
}

//SubImage returns the sub image of the passed ebiten.Image based on the BasicImageParts properties
//Reduces the amount of coding needed in the actual game to get to drawing the image
func (b BasicImageParts) SubImage(img *ebiten.Image) *ebiten.Image {
	if b.Reverse {
		return img.SubImage(image.Rect(b.Sx+b.Width, b.Sy, b.Sx, b.Sy+b.Height)).(*ebiten.Image)
	}
	return img.SubImage(image.Rect(b.Sx, b.Sy, b.Sx+b.Width, b.Sy+b.Height)).(*ebiten.Image)
}

//SetScale sets the scale of the DrawImageOptions based on the given DestHeight and DestWidth of the BasicImageParts
func (b *BasicImageParts) SetScale(op *ebiten.DrawImageOptions) {
	if b.DestWidth == 0 {
		b.DestWidth = b.Width
	}
	if b.DestHeight == 0 {
		b.DestHeight = b.Height
	}
	op.GeoM.Scale(float64(b.DestWidth)/float64(b.Width), float64(b.DestHeight)/float64(b.Height))
}

//BasicImagePartsFromSpriteSheet creates a BasicImageParts from a passed spritesheet on the passed frame.
//This is helpful to easily get the correct sx,sy,w,h without manually typing it in every time.
func BasicImagePartsFromSpriteSheet(spriteSheet *SpriteSheet, frame int) *BasicImageParts {
	return &BasicImageParts{
		Sx:     spriteSheet.Frames[frame].Frame["x"],
		Sy:     spriteSheet.Frames[frame].Frame["y"],
		Width:  spriteSheet.Frames[frame].Frame["w"],
		Height: spriteSheet.Frames[frame].Frame["h"],
	}
}
