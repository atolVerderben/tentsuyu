package tentsuyu

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

//Map represents an entire Tiled map
type Map struct {
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	Layers      []*Layer    `json:"layers"`
	Orientation string      `json:"orientation"`
	TileWidth   int         `json:"tilewidth"`
	TileHeight  int         `json:"tileheight"`
	TileSets    []*TileSet  `json:"tilesets"`
	Properties  []*Property `json:"properties"`
	Version     int         `json:"version"`
}

//Layer represents a Tiled Layer
type Layer struct {
	Name       string       `json:"name"`
	Type       string       `json:"type"`
	X          int          `json:"x"`
	Y          int          `json:"y"`
	Width      int          `json:"width"`
	Height     int          `json:"height"`
	Data       []int        `json:"data"`
	Opacity    int          `json:"opacity"`
	Visible    bool         `json:"visible"`
	Properties []*Property  `json:"properties"`
	DrawOrder  string       `json:"draworder"`
	Objects    []*MapObject `json:"objects"`
	//Encoding   string       `json:"encoding"`
	//	ImageName  string               `json:"image"`
}

//TileSet represents a Tiled TileSet
type TileSet struct {
	FirstGID      int         `json:"firstgid"`
	ImageName     string      `json:"image"`
	ImageWidth    int         `json:"imagewidth"`
	ImageHeight   int         `json:"imageheight"`
	Margin        int         `json:"margin"`
	Name          string      `json:"name"`
	Properties    []*Property `json:"properties"`
	Spacing       int         `json:"spacing"`
	TileWidth     int         `json:"tilewidth"`
	TileHeight    int         `json:"tileheight"`
	Columns       int         `json:"columns"`
	TileCount     int         `json:"tilecount"`
	Image         *ebiten.Image
	Rows, LastGID int
}

//MapObject is a representation of the Tiled Object
type MapObject struct {
	Height   int     `json:"height"`
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Rotation float64 `json:"rotation"`
	Type     string  `json:"type"`
	Visible  bool    `json:"visible"`
	Width    int     `json:"width"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

//Property of the layer/tile from Tiled
type Property struct {
	Name     string `json:"name"`
	PropType string `json:"type"`
	Value    string `json:"value"`
}

//Tile is a renderable tile used by the game
type Tile struct {
	Image     *ebiten.Image
	Collide   bool
	Gid       int
	ImageName string
	*BasicImageParts
	*BasicObject
}

//TileMap is the renderable map used by the game
type TileMap struct {
	Layers                               []*TileLayer
	Width, Height, TileWidth, TileHeight int
}

//TileLayer is the renderable layer used by the game
type TileLayer struct {
	Data          []*Tile
	DrawOrder     int
	Collide       bool
	Above         bool
	Name          string
	IsImageLayer  bool
	ImageName     string
	X, Y          float64
	Width, Height int
	Properties    []*Property
}

//ReadMap from JSON file and dump into Map
func ReadMap(filename string) *Map {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	m := &Map{}
	json.Unmarshal(raw, m)
	return m
}

//ReadMapfromString reads an entire Map struct from a given string (represented by json)
func ReadMapfromString(jString string) *Map {
	raw := []byte(jString)

	m := &Map{}
	json.Unmarshal(raw, m)
	return m
}

//ReadMapfromByte reads an entire Map struct from a byte slice
func ReadMapfromByte(byteMap []byte) *Map {

	m := &Map{}
	json.Unmarshal(byteMap, m)
	return m
}

//CreateTileMapFromFile creates a new TileMap based on the given filelocation
func CreateTileMapFromFile(fileLocation string) *TileMap {
	tilemap := ReadMap(fileLocation)
	return CreateTileMap(tilemap)

}

//CreateTileMap creates a renderable TileMap
func CreateTileMap(tilemap *Map) *TileMap {
	tm := &TileMap{
		Layers:     []*TileLayer{},
		Width:      tilemap.Width,
		Height:     tilemap.Height,
		TileHeight: tilemap.TileHeight,
		TileWidth:  tilemap.TileWidth,
	}
	PrepareTileSet(tilemap)
	for _, layer := range tilemap.Layers {
		if layer.Type == "imagelayer" {
			tl := &TileLayer{
				Name: layer.Name,
				//ImageName:    layer.ImageName,
				IsImageLayer: true,
				X:            float64(layer.X),
				Y:            float64(layer.Y),
				Width:        layer.Width * tilemap.TileWidth,
				Height:       layer.Height * tilemap.TileHeight,
				Properties:   layer.Properties,
			}
			tm.Layers = append(tm.Layers, tl)
		}
		if layer.Type == "tilelayer" {
			tl := &TileLayer{
				Name: layer.Name,
				Data: []*Tile{},
			}
			/*if layer.Properties["collision"] != nil {
				tl.Collide = true
			}
			if layer.Properties["above"] != nil {
				tl.Above = true
			}*/
			rowRange := 0
			//rowTiles := []*Tile{}

			x := 0.0
			y := 0.0
			for _, tileData := range layer.Data {

				rowRange++
				if tileData != 0 {
					t := &Tile{
						Gid:         tileData,
						BasicObject: NewBasicObject(x, y, tilemap.TileWidth, tilemap.TileHeight),
						ImageName:   tl.ImageName,
					}
					t.NotCentered = true
					t.DetermineTileSet(tilemap)
					tl.Data = append(tl.Data, t)

				}
				x += float64(tilemap.TileWidth)
				//Determine whether to start a new row
				if rowRange >= layer.Width {
					rowRange = 0
					x = 0.0 //Determine whether to start a new row
					if rowRange > layer.Width {
						rowRange = 0
						x = 0.0

					}
					y += float64(tilemap.TileHeight)

				}
			}
			tm.Layers = append(tm.Layers, tl)
			//tm.Layers[i] = tl
		}
	}

	return tm
}

//DetermineTileSet of the given tile based on the GID
func (t *Tile) DetermineTileSet(tilemap *Map) {
	for _, tileSet := range tilemap.TileSets {
		if tileSet.FirstGID <= t.Gid {
			if tileSet.LastGID >= t.Gid || tileSet.LastGID == tileSet.FirstGID {
				//t.Image = tileSet.Image

				t.ImageName = tileSet.Name
				sx, sy := tileSet.ReturnImagePosition(t.Gid)
				t.BasicImageParts = &BasicImageParts{
					Height:     tileSet.TileHeight,
					Width:      tileSet.TileWidth,
					Sx:         int(sx),
					Sy:         int(sy),
					DestWidth:  tileSet.TileWidth,
					DestHeight: tileSet.TileHeight,
				}
			}
		}
	}
}

//Draw a single renderable Tile
func (t *Tile) Draw(screen *ebiten.Image, imageManager *ImageManager) error {

	op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(float64(-t.BasicObject.Width/2), float64(-t.BasicObject.Height/2))
	//op.GeoM.Scale(float64(scalex), float64(scaley))
	op.GeoM.Translate(t.GetX(), t.GetY())

	//log.Printf("%v,%v\n", scalex, scaley)
	//ApplyCameraTransform(op, true)
	//if Components.Camera.OnScreen(t.X, t.Y) {
	if t.ImageName != "" {
		screen.DrawImage(t.BasicImageParts.SubImage(imageManager.ReturnImage(t.ImageName)), op)
	}
	//}

	return nil
}

//Draw the entire TileLayer
func (tl *TileLayer) Draw(screen *ebiten.Image, imageManager *ImageManager) error {
	if tl.IsImageLayer {
		op := &ebiten.DrawImageOptions{}
		//op.GeoM.Translate(float64(-t.BasicObject.Width/2), float64(-t.BasicObject.Height/2))
		//op.GeoM.Scale(float64(scalex), float64(scaley))
		op.GeoM.Translate(tl.X, tl.Y)

		//log.Printf("%v,%v\n", scalex, scaley)
		//ApplyCameraTransform(op, true)

		screen.DrawImage(BasicImageParts{Height: tl.Height, Width: tl.Width}.SubImage(imageManager.ReturnImage(tl.ImageName)), op)

		return nil
	}

	for x := range tl.Data {
		tl.Data[x].Draw(screen, imageManager)
	}
	return nil
}

//Draw the entire TileMap
func (tm *TileMap) Draw(screen *ebiten.Image, imageManager *ImageManager) error {
	for x := 0; x < len(tm.Layers); x++ {
		tm.Layers[x].Draw(screen, imageManager)
	}
	return nil
}

//PrepareTileSet imports the tileset image, and sets LastGID and Rows for easier calculation
func PrepareTileSet(tilemap *Map) { //, imageManager *ImageManager) {
	for x, tileSet := range tilemap.TileSets {
		//image := imageManager.ReturnImage(tilemap.TileSets[x].Name) //ebitenutil.NewImageFromFile(tilemap.TileSets[x].ImageName, ebiten.FilterNearest)

		//tilemap.TileSets[x].Image = image
		tilemap.TileSets[x].LastGID = tileSet.FirstGID + tileSet.TileCount
		tilemap.TileSets[x].Rows = tileSet.ImageHeight / tileSet.TileHeight
	}
}

//ReturnImagePosition returns the image location within the tileset
func (ts *TileSet) ReturnImagePosition(gid int) (float64, float64) {

	gid = gid - (ts.FirstGID - 1)
	row := gid / ts.Columns
	if math.Mod(float64(gid), float64(ts.Columns)) == 0 {
		row--
	}

	y := float64(row * ts.TileHeight)

	startRowGid := row*ts.Columns + 1
	x := float64((gid - startRowGid) * ts.TileWidth)

	//x *= float64(ts.TileWidth)
	//y *= float64(ts.TileHeight)
	return x, y
}
