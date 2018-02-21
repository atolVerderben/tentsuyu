package tentsuyu

//Animation is used to control animation frames
type Animation struct {
	currFrame, frameCount, frameSpeed int
	ImageParts                        *BasicImageParts
	SpriteSheet                       *SpriteSheet
}

//Update the animation
func (a *Animation) Update() {
	a.frameCount++
	if a.frameCount > a.frameSpeed {
		switch a.currFrame {
		case 1:
			a.currFrame = 2
		case 2:
			a.currFrame = 3
		case 3:
			a.currFrame = 1
		}
		a.frameCount = 0
		a.ImageParts.Sx = a.SpriteSheet.Frames[a.currFrame].Frame["x"]
		a.ImageParts.Sy = a.SpriteSheet.Frames[a.currFrame].Frame["y"]
	}
}
