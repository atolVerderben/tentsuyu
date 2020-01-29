package tentsuyu

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//SpriteSheet holds all frames of a spritesheet from a json
type SpriteSheet struct {
	Frames                                []Frame `json:"frames"`
	framePaddingWidth, framePaddingHeight int
}

//Frame represents a single frame of a spritesheet
type Frame struct {
	Filename         string             `json:"filename"`
	Frame            map[string]int     `json:"frame"` // x, y, w, h
	Rotated          bool               `json:"rotated"`
	Trimmed          bool               `json:"trimmed"`
	SpriteSourceSize map[string]int     `json:"spriteSourceSize"`
	SourceSize       map[string]int     `json:"sourceSize"`
	Pivot            map[string]float64 `json:"pivot"`
}

//ReadSpriteSheet reads a json file and returns a SpriteSheet struct
func ReadSpriteSheet(filename string) *SpriteSheet {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	m := &SpriteSheet{}
	json.Unmarshal(raw, m)
	return m
}

//ReadSpriteSheetJSON from a byte slice
func ReadSpriteSheetJSON(jsonByte []byte) *SpriteSheet {
	m := &SpriteSheet{}
	json.Unmarshal(jsonByte, m)
	return m
}

//NewSpriteSheet returns a SpriteSheet struct with just a basic
func NewSpriteSheet(imageWidth, imageHeight, frameWidth, frameHeight, paddingWidth, paddingHeight int) *SpriteSheet {
	s := &SpriteSheet{}

	for height := 0 + paddingHeight; height <= imageHeight; height += paddingHeight + frameHeight {
		for width := 0 + paddingWidth; width <= imageWidth; width += paddingWidth + frameWidth {
			if width >= imageWidth {
				continue
			}
			if height >= imageHeight {
				continue
			}
			frame := map[string]int{}
			frame["x"] = width
			frame["y"] = height
			frame["w"] = frameWidth
			frame["h"] = frameHeight
			s.Frames = append(s.Frames, Frame{
				Filename: "Frame" + strconv.Itoa(height+width),
				Frame:    frame,
			})
		}
	}

	return s
}

//ReturnImageParts returns the BasicImageParts at the given frame of the animation.
func (s SpriteSheet) ReturnImageParts(f int) *BasicImageParts {
	i := &BasicImageParts{
		Sx:     s.Frames[f].Frame["x"],
		Sy:     s.Frames[f].Frame["y"],
		Width:  s.Frames[f].Frame["w"],
		Height: s.Frames[f].Frame["h"],
	}

	return i

}
