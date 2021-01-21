package tentsuyu

import (
	"math"
	"math/rand"

	"github.com/atolVerderben/tentsuyu/tentsuyutils"

	"github.com/hajimehoshi/ebiten/v2"
)

//Camera is the entity that follows our player so he doesn't walk off screen
type Camera struct {
	x, y, rotation, Width, Height, Zoom, ScreenWidth, ScreenHeight float64
	zoomCount, zoomCountMax                                        int
	FreeFloating                                                   bool
	MaxZoomOut, MaxZoomIn                                          float64
	startShaking, isShaking                                        bool
	shakeAngle, shakeRadius                                        float64
	preShakeX, preShakeY                                           float64
	destX, destY                                                   float64
	freeFloatSpeed                                                 float64
	moving                                                         bool
	zoomingIn                                                      bool
	offSetX, offSetY                                               float64
	lowerBoundX, upperBoundX                                       float64
	lowerBoundY, upperBoundY                                       float64
	clamp                                                          bool
}

//CreateCamera intializes a camera struct
func CreateCamera(width, height float64) *Camera {
	c := &Camera{
		Height:         height,
		Width:          width,
		Zoom:           1,
		zoomCountMax:   1,
		ScreenHeight:   height,
		ScreenWidth:    width,
		FreeFloating:   false,
		MaxZoomOut:     0.1,
		MaxZoomIn:      2.0,
		shakeRadius:    60.0,
		freeFloatSpeed: 4.5,
		clamp:          true,
	}
	return c
}

//SetClamp sets whether the camera takes the floor of its coordinates
func (c *Camera) SetClamp(clamp bool) {
	c.clamp = clamp
}

//SetBounds that the camera operates in
func (c *Camera) SetBounds(lowerX, upperX, lowerY, upperY float64) {
	c.lowerBoundX = lowerX
	c.upperBoundX = upperX
	c.lowerBoundY = lowerY
	c.upperBoundY = upperY
}

//SetDimensions sets the width and height of the camera
func (c *Camera) SetDimensions(width, height float64) {
	c.Height = height
	c.Width = width
}

//SetZoom of the camera
func (c *Camera) SetZoom(zoom float64) {
	c.Zoom = zoom
	c.x, c.y = c.destX*c.Zoom, c.destY*c.Zoom

}

//SetZoomGradual zooms into the passed zoom value and a given speed
//TODO: Implement this
func (c *Camera) SetZoomGradual(zoom, speed float64) {
	c.SetZoom(zoom) //temporarily just call SetZoom
}

//GetOffsetX returns the camera X offset position
func (c Camera) GetOffsetX() float64 {
	return c.offSetX
}

//GetOffsetY returns the camera Y offset position
func (c Camera) GetOffsetY() float64 {
	return c.offSetY
}

//SetOffsetX sets the offset X
func (c *Camera) SetOffsetX(offset float64) {
	c.offSetX = offset
}

//SetOffsetY sets the offset Y
func (c *Camera) SetOffsetY(offset float64) {
	c.offSetY = offset
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

//GetDestination of camera (x,y)
func (c *Camera) GetDestination() (float64, float64) {
	return c.destX, c.destY
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

//SetSpeed that the camera moves at
func (c *Camera) SetSpeed(s float64) {
	c.freeFloatSpeed = s
}

func (c *Camera) moveToDestination() {
	//moveX, moveY := false, false
	if c.moving {
		angle := math.Atan2(c.destY-c.y, c.destX-c.x)

		c.x += c.freeFloatSpeed * math.Cos(angle)
		c.y += c.freeFloatSpeed * math.Sin(angle)

		if tentsuyutils.Distance(c.x, c.y, c.destX, c.destY) <= 5 {
			c.moving = false
			c.x = c.destX
			c.y = c.destY
		}

		/*if moveX && moveY {
			c.x += c.freeFloatSpeed / 2
			c.y += c.freeFloatSpeed / 2
		} else {
			if moveX {
				c.x += c.freeFloatSpeed
			}
			if moveY {
				c.y += c.freeFloatSpeed
			}
		}*/
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

//ApplyCameraTransform applies the camera's position to the DrawImageOptions, bool toggles whether zoom is applied or not
func (c Camera) ApplyCameraTransform(op *ebiten.DrawImageOptions, applyZoom bool) {
	if applyZoom {
		c.DrawCameraTransform(op)
	} else {
		c.DrawCameraTransformIgnoreZoom(op)
	}
}

//Update the camera
func (c *Camera) Update() {
	if c.startShaking {
		c.moving = false
		if !c.isShaking { //retain the preshake coords
			c.preShakeX = c.x
			c.preShakeY = c.y
		}
		c.shakeAngle = rand.Float64() * math.Pi * 2
		/*c.shakeAngle = rand.Float64() * math.Pi * 2
		offsetX := math.Sin(c.shakeAngle) * c.shakeRadius
		offsetY := math.Cos(c.shakeAngle) * c.shakeRadius
		c.y += offsetY
		c.x += offsetX*/
		c.startShaking = false
		c.isShaking = true
	}

	if c.isShaking {
		c.Shake()
	}

	if c.moving {
		c.moveToDestination()
		if c.clamp {
			c.x = math.Round(c.x)
			c.y = math.Round(c.y)
		}
	}

}

//GetScreenCoords returns where the passed coords would be on the screen space
func (c Camera) GetScreenCoords(x, y float64) (float64, float64) {
	return x*c.Zoom - c.x, y*c.Zoom - c.y
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
	x, y = (x+c.offSetX)*c.Zoom, (y+c.offSetY)*c.Zoom

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
	if c.isShaking {
		c.Shake()
	}
}

//FollowObject follows the given GameObject either within or without bounds
func (c *Camera) FollowObject(player GameObject, bounds bool) {
	if bounds {
		c.FollowObjectInBounds(player)
	} else {
		c.FollowObjectNoBounds(player)
	}
}

//FollowObjectInBounds follows the given GameObject within the bounds of the camera
func (c *Camera) FollowObjectInBounds(player GameObject) {

	worldHeight := c.upperBoundY * c.Zoom
	worldWidth := c.upperBoundX * c.Zoom
	lowerHeight := c.lowerBoundY * c.Zoom
	lowerWidth := c.lowerBoundX * c.Zoom
	cameraOverWidth, cameraOverHeight := false, false
	if worldWidth < c.ScreenWidth {
		//c.x = 0
		//c.y = 0
		//c.Center(worldWidth/2, worldHeight/2)
		c.CenterX((worldWidth + lowerWidth) / 2)
		cameraOverWidth = true

	}
	if worldHeight < c.ScreenHeight {
		c.CenterY((worldHeight + lowerHeight) / 2)
		cameraOverHeight = true

	}
	if cameraOverHeight && cameraOverWidth {
		return
	}
	x, y := player.GetPosition()
	x, y = (x+c.offSetX)*c.Zoom, (y+c.offSetY)*c.Zoom

	// X-Axis
	if !cameraOverWidth {
		// Follow Player Freely
		if x-c.Width/2 > lowerWidth && x+c.Width/2 < worldWidth {
			c.destX = (x - c.Width/2)
		} else if x+c.Width/2 >= worldWidth { // Stop at right edge
			c.destX = worldWidth - c.Width
		} else { // Stop at left edge
			c.destX = lowerWidth
		}
	}

	// Y-Axis
	if !cameraOverHeight {
		// Follow Player Freely
		if y-c.Height/2 > lowerHeight && y+c.Height/2 < worldHeight {
			c.destY = y - c.Height/2
		} else if y+c.Height/2 >= worldHeight { // Stop at bottom
			c.destY = worldHeight - c.Height
		} else { // Stop at top
			c.destY = lowerHeight
		}
	}
	c.x = c.destX
	c.y = c.destY
	if c.clamp {
		c.x = math.Round(c.x)
		c.y = math.Round(c.y)
	}
	c.moving = true
}

//FollowObjectNoBounds follows the given GameObject without boundaries
func (c *Camera) FollowObjectNoBounds(player GameObject) {

	x, y := player.GetPosition()
	x, y = (x+c.offSetX)*c.Zoom, (y+c.offSetY)*c.Zoom

	// X-Axis

	// Follow Player Freely

	c.x = (x - c.Width/2)

	// Y-Axis

	// Follow Player Freely

	c.y = y - c.Height/2

	if c.clamp {
		c.x = math.Round(c.x)
		c.y = math.Round(c.y)
	}
}

//StartShaking begins the camera shake with the passed intensity
func (c *Camera) StartShaking(r float64) {
	c.shakeRadius = r
	c.startShaking = true

}

//SetShakeRadius for the camera shake
func (c *Camera) SetShakeRadius(radius float64) {
	c.shakeRadius = radius
}

//Shake shakes the camera
func (c *Camera) Shake() {

	/*if c.startShaking {
		if !c.isShaking { //retain the preshake coords
			c.preShakeX = c.x
			c.preShakeY = c.y
		}
		c.shakeAngle = rand.Float64() * math.Pi * 2
		c.shakeAngle = rand.Float64() * math.Pi * 2
		offsetX := math.Sin(c.shakeAngle) * c.shakeRadius
		offsetY := math.Cos(c.shakeAngle) * c.shakeRadius
		c.y += offsetY
		c.x += offsetX
		c.startShaking = false
		c.isShaking = true
	}*/

	if c.shakeRadius >= 0.2 {
		c.shakeRadius *= 0.9
		c.shakeAngle += (150 + rand.Float64()*1.0472)
		offsetX := math.Sin(c.shakeAngle) * c.shakeRadius
		offsetY := math.Cos(c.shakeAngle) * c.shakeRadius
		c.y += offsetY
		c.x += offsetX
	} else {
		c.isShaking = false
		c.destX = c.preShakeX
		c.destY = c.preShakeY
		c.moving = true
		//c.x = c.preShakeX
		//c.y = c.preShakeY
	}
}
