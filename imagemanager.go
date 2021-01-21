package tentsuyu

import (
	"bytes"
	"image"

	//I want to accept png and jpg files by default
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//NewImageManager creates a new pointer to ImageManager
func NewImageManager() *ImageManager {
	return &ImageManager{
		Images: map[string]*ebiten.Image{},
	}
}

//ImageManager is a struct containing a map of named ebiten.Images
type ImageManager struct {
	Images map[string]*ebiten.Image
}

//LoadImageFromFile [Deprecated - use AddImageFromFile]
func (im *ImageManager) LoadImageFromFile(name string, path string) error {
	return im.AddImageFromFile(name, path)
}

//AddImageFromFile loads the given image at "path" with "name"
func (im *ImageManager) AddImageFromFile(name string, path string) error {

	img2, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return err
	}

	im.Images[name] = img2
	return nil
}

//AddImage adds an ebiten.Image to the map with a given name
func (im *ImageManager) AddImage(name string, image *ebiten.Image) {
	im.Images[name] = image
}

//ReturnImage retrieves the specified image name
func (im *ImageManager) ReturnImage(name string) *ebiten.Image {
	return im.Images[name]
}

//AddImageFromBytes adds in the image based on a byte slice
//Very helpful with using file2byteslice by HajimeHoshi
func (im *ImageManager) AddImageFromBytes(name string, b []byte) error {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}
	img2 := ebiten.NewImageFromImage(img)

	im.Images[name] = img2

	return nil
}
