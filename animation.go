package tentsuyu

//Animation is used to control animation frames
type Animation struct {
	currFrame, frameCount, frameSpeed int
	paused                            bool
	ImageParts                        *BasicImageParts
	SpriteSheet                       *SpriteSheet
	Frames                            []int
	LoopCompleted                     bool
	reverse                           bool
	Repeating                         bool
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
		Repeating:  true,
	}

	return a
}

//SetAnimationSpeed changes the speed of the animation to the passed value
func (a *Animation) SetAnimationSpeed(s int) {
	a.frameSpeed = s
}

//SetReverse tells the animation to play in reverse
func (a *Animation) SetReverse() {
	a.reverse = true
}

//SetForward tells the animation to play normally.
//This should only be necessary after calling SetReverse()
func (a *Animation) SetForward() {
	a.reverse = false
}

//Update the animation if not paused
func (a *Animation) Update() {
	if !a.paused {
		if !a.reverse {
			a.frameCount++
			if a.frameCount > a.frameSpeed {
				a.currFrame++
				if a.currFrame >= len(a.Frames) {
					if !a.Repeating {
						a.currFrame = len(a.Frames) - 1
					} else {
						a.currFrame = 0
					}
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
		} else {
			a.frameCount++
			if a.frameCount > a.frameSpeed {
				a.currFrame--
				if a.currFrame <= 0 {
					a.currFrame = len(a.Frames) - 1
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

	if !a.Repeating && !a.paused {
		if a.LoopCompleted {
			a.Stop()
			a.currFrame = len(a.Frames) - 1
		}
	}
}

//ReturnImageParts of the current animation
func (a *Animation) ReturnImageParts() *BasicImageParts {
	return a.ImageParts
}

//SetFrameSpeed of the current animation
func (a *Animation) SetFrameSpeed(speed int) {
	a.frameSpeed = speed
}

//CurrentFrame returns the current frame of the animation
func (a Animation) CurrentFrame() int {
	return a.currFrame
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

//Reset the animation to the starting point
func (a *Animation) Reset() {
	a.paused = false
	a.LoopCompleted = false
	a.frameCount = 0
	a.currFrame = 0
}

//Resume playing the animation
func (a *Animation) Resume() {
	a.paused = false
}

//Play the animation
func (a *Animation) Play() {
	a.paused = false
}

//IsPaused returns true if the animation is paused
func (a Animation) IsPaused() bool {
	return a.paused
}

//Stop the animation
func (a *Animation) Stop() {
	a.paused = true
	a.LoopCompleted = false
	a.frameCount = 0
	a.currFrame = 0
}
