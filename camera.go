package tentsuyu

import "github.com/hajimehoshi/ebiten"

//Camera is the entity that follows our player so he doesn't walk off screen
type Camera struct {
	x, y, rotation, Width, Height, Zoom, ScreenWidth, ScreenHeight float64
	zoomCount, zoomCountMax                                        int
	FreeFloating                                                   bool
	MaxZoomOut, MaxZoomIn                                          float64
}

//CreateCamera intializes a camera struct
func CreateCamera(width, height float64) *Camera {
	c := &Camera{
		Height:       height,
		Width:        width,
		Zoom:         1,
		zoomCountMax: 1,
		ScreenHeight: height,
		ScreenWidth:  width,
		FreeFloating: false,
		MaxZoomOut:   0.1,
		MaxZoomIn:    2.0,
	}
	return c
}

//SetDimensions sets the width and height of the camera
func (c *Camera) SetDimensions(width, height float64) {
	c.Height = height
	c.Width = width
}

//SetZoom of the camera
func (c *Camera) SetZoom(zoom float64) {
	c.Zoom = zoom
}

//GetX returns the camera X position
func (c *Camera) GetX() float64 {
	return c.x
}

//GetY returns the camera Y position
func (c *Camera) GetY() float64 {
	return c.y
}

//Center camera on point
func (c *Camera) Center(x, y float64) {
	c.x = x - c.Width/2
	c.y = y - c.Height/2
}

//CenterX centers camera X position
func (c *Camera) CenterX(x float64) {
	c.x = x - c.Width/2
}

//CenterY centers camera Y position
func (c *Camera) CenterY(y float64) {
	c.y = y - c.Height/2
}

//ChangeZoom increments or decrements the camera zoom level
func (c *Camera) ChangeZoom() {
	if c.zoomCount > 0 {
		c.zoomCount++
		if c.zoomCount > c.zoomCountMax {
			c.zoomCount = 0
		}
	} else {
		increment := 0.01
		if ebiten.IsKeyPressed(ebiten.KeyQ) && c.Zoom < 2.0 {
			c.Zoom += increment
			c.zoomCount++
		}
		if ebiten.IsKeyPressed(ebiten.KeyE) && c.Zoom > 0.1 {
			c.Zoom -= increment
			c.zoomCount++
		}

	}
}

//ZoomIn move the camera closer towards the player
func (c *Camera) ZoomIn() {
	if c.zoomCount > 0 {
		c.zoomCount++
		if c.zoomCount > c.zoomCountMax {
			c.zoomCount = 0
		}
	} else {
		increment := 0.01
		if c.Zoom < c.MaxZoomIn {
			c.Zoom += increment
			c.zoomCount++
		}

	}
}

//ZoomOut moves the camera further away from the player
func (c *Camera) ZoomOut() {
	if c.zoomCount > 0 {
		c.zoomCount++
		if c.zoomCount > c.zoomCountMax {
			c.zoomCount = 0
		}
	} else {
		increment := 0.01
		if ebiten.IsKeyPressed(ebiten.KeyE) && c.Zoom > c.MaxZoomOut {
			c.Zoom -= increment
			c.zoomCount++
		}

	}
}

//OnScreen determines if the given position is within the camera viewport
func (c Camera) OnScreen(x, y float64, w, h int) bool {
	containsX, containsY := false, false
	x, y = x*c.Zoom, y*c.Zoom
	width, height := float64(w)*c.Zoom, float64(h)*c.Zoom
	if x-width < c.x+c.Width && x+width > c.x {
		containsX = true
	}
	if y-height < c.y+c.Height && y+height > c.y {
		containsY = true
	}

	return containsX && containsY
}

//Position of the camera
func (c Camera) Position() (x, y float64) {
	return c.x, c.y
}

//SetX position of the camera
func (c *Camera) SetX(x float64) {
	c.x = x
}

//SetY position of the camera
func (c *Camera) SetY(y float64) {
	c.y = y
}

//SetPosition by passing both x and y coordinates of the camera
func (c *Camera) SetPosition(x, y float64) {
	c.x = x
	c.y = y
}

//TransformMatrix of the camera (currently for concept purposes only... but is correct)
func (c *Camera) TransformMatrix() *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(c.rotation)
	op.GeoM.Scale(c.Zoom, c.Zoom)
	op.GeoM.Translate(c.x, c.y)
	return op
}

//DrawCameraTransform appls the TransformMatrix of the camera to the specified image options
//This translates the opposite direction of the TransformMatrix
func (c *Camera) DrawCameraTransform(op *ebiten.DrawImageOptions) {
	op.GeoM.Rotate(c.rotation)
	op.GeoM.Scale(c.Zoom, c.Zoom)
	op.GeoM.Translate(-c.x, -c.y)
}

//DrawCameraTransformIgnoreZoom is same as DrawCameraTransform minus the zoom
func (c *Camera) DrawCameraTransformIgnoreZoom(op *ebiten.DrawImageOptions) {
	op.GeoM.Rotate(c.rotation)
	op.GeoM.Translate(-c.x, -c.y)
}

//FollowPlayer follows the specified character (in this case the player)
func (c *Camera) FollowPlayer(player GameObject, worldWidth, worldHeight float64) {

	//c.ChangeZoom()

	worldHeight *= c.Zoom
	worldWidth *= c.Zoom
	cameraOverWidth, cameraOverHeight := false, false
	if worldWidth < c.ScreenWidth {
		//c.x = 0
		//c.y = 0
		//c.Center(worldWidth/2, worldHeight/2)
		c.CenterX(worldWidth / 2)
		cameraOverWidth = true

	}
	if worldHeight < c.ScreenHeight {
		c.CenterY(worldHeight / 2)
		cameraOverHeight = true

	}
	if cameraOverHeight && cameraOverWidth {
		return
	}
	x, y := player.GetPosition()
	x, y = x*c.Zoom, y*c.Zoom

	// X-Axis
	if !cameraOverWidth {
		// Follow Player Freely
		if x-c.Width/2 > 0 && x+c.Width/2 < worldWidth {
			c.x = (x - c.Width/2)
		} else if x+c.Width/2 >= worldWidth { // Stop at right edge
			c.x = worldWidth - c.Width
		} else { // Stop at left edge
			c.x = 0
		}
	}

	// Y-Axis
	if !cameraOverHeight {
		// Follow Player Freely
		if y-c.Height/2 > 0 && y+c.Height/2 < worldHeight {
			c.y = y - c.Height/2
		} else if y+c.Height/2 >= worldHeight { // Stop at bottom
			c.y = worldHeight - c.Height
		} else { // Stop at top
			c.y = 0
		}
	}
}
