package tentsuyu

//Animation is used to control animation frames
type Animation struct {
	currFrame, frameCount, frameSpeed int
	paused                            bool
	ImageParts                        *BasicImageParts
	SpriteSheet                       *SpriteSheet
	Frames                            []int
	LoopCompleted                     bool
}

//NewAnimation takes a spritesheet, []int of frames to play, and speed to return an Animation
func NewAnimation(spriteSheet *SpriteSheet, frames []int, speed int) *Animation {
	a := &Animation{
		SpriteSheet: spriteSheet,
		Frames:      frames,
		ImageParts: &BasicImageParts{
			Sx:     spriteSheet.Frames[frames[0]].Frame["x"],
			Sy:     spriteSheet.Frames[frames[0]].Frame["y"],
			Width:  spriteSheet.Frames[frames[0]].Frame["w"],
			Height: spriteSheet.Frames[frames[0]].Frame["h"],
		},
		frameSpeed: speed,
		paused:     false,
	}

	return a
}

//Update the animation if not paused
func (a *Animation) Update() {
	if !a.paused {
		a.frameCount++
		if a.frameCount > a.frameSpeed {
			a.currFrame++
			if a.currFrame >= len(a.Frames) {
				a.currFrame = 0
				a.LoopCompleted = true
			} else {
				if a.LoopCompleted == true {
					a.LoopCompleted = false
				}
			}
			a.frameCount = 0
			a.ImageParts.Sx = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["x"]
			a.ImageParts.Sy = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["y"]
			a.ImageParts.Width = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["w"]
			a.ImageParts.Height = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["h"]
		}
	}
}

//ReturnImageParts of the current animation
func (a *Animation) ReturnImageParts() *BasicImageParts {
	return a.ImageParts
}

//SetAnimationSpeed of the current animation
func (a *Animation) SetAnimationSpeed(speed int) {
	a.frameSpeed = speed
}

//SetCurrentFrame of the current animation
func (a *Animation) SetCurrentFrame(frame int) {
	a.currFrame = frame
	a.frameCount = 0
	a.ImageParts.Sx = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["x"]
	a.ImageParts.Sy = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["y"]
	a.ImageParts.Width = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["w"]
	a.ImageParts.Height = a.SpriteSheet.Frames[a.Frames[a.currFrame]].Frame["h"]
	a.LoopCompleted = false
}

//Pause the animation
func (a *Animation) Pause() {
	a.paused = true
}

//Resume playing the animation
func (a *Animation) Resume() {
	a.paused = false
}
