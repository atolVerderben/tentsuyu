package tentsuyu

import "github.com/hajimehoshi/ebiten/v2"

//Cursor represents the player's mouse position
type Cursor struct {
	style int
	*BasicObject
	*BasicImageParts
	spritesheet *ebiten.Image
}

//Types of cursors
const (
	CursorCrosshair = iota
	CursorPointer
)

//NewCursor creates the cursor... should be called during game set up
func NewCursor(screenWidth, screenHeight float64, spritesheet *ebiten.Image) *Cursor {
	c := &Cursor{
		BasicObject: &BasicObject{
			Position: &Vector2d{
				X: screenWidth / 2,
				Y: screenHeight / 2,
			},
			Width:  32,
			Height: 32,
		},
		BasicImageParts: &BasicImageParts{
			Sx:     332,
			Sy:     468,
			Width:  32,
			Height: 32,
		},
		style:       CursorCrosshair,
		spritesheet: spritesheet,
	}
	return c
}

//Update sets the curosr position
func (c *Cursor) Update(mx, my float64) {
	c.SetPosition(mx, my)
}

//SetStyle of the cursor: center or top corner
func (c *Cursor) SetStyle(cursorstyle int) {
	c.style = cursorstyle
}

//Draw the cursor
func (c *Cursor) Draw(screen *ebiten.Image) error {
	w, h := c.GetSize()
	op := &ebiten.DrawImageOptions{}
	//op.ImageParts = c.BasicImageParts
	if !c.NotCentered {
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	}
	op.GeoM.Translate(c.GetPosition())
	//ApplyCameraTransform(op, false)

	screen.DrawImage(c.BasicImageParts.SubImage(c.spritesheet), op)

	return nil
}
