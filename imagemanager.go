package tentsuyu

import (
	"bytes"
	"image"

	//I want to accept png and jpg files by default
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type imageManager struct {
	Images map[string]*ebiten.Image
}

func (im *imageManager) LoadImageFromFile(name string, path string) error {

	fImg1, _ := os.Open(path)
	defer fImg1.Close()
	img, _, err := image.Decode(fImg1)
	if err != nil {
		return err
	}

	img2, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if err != nil {
		return err
	}

	im.Images[name] = img2
	return nil
}

//AddImage adds an ebiten.Image to the map with a given name
func (im *imageManager) AddImage(name string, image *ebiten.Image) {
	im.Images[name] = image
}

//ReturnImage retrieves the specified image name
func (im *imageManager) ReturnImage(name string) *ebiten.Image {
	return im.Images[name]
}

//AddImageFromBytes adds in the image based on a byte slice
//Very helpful with using file2byteslice by HajimeHoshi
func (im *imageManager) AddImageFromBytes(name string, b []byte) error {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}
	img2, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if err != nil {
		return err
	}

	im.Images[name] = img2

	return nil
}
