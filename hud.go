package tentsuyu

import "github.com/hajimehoshi/ebiten/v2"

//HUD represents the player's Heads Up Display
type HUD struct {
	x, y, prevX, prevY                         float64
	w, h                                       float64
	uiElements                                 map[string]*hudElement
	topLeft, topRight, bottomLeft, bottomRight *hudSection
}

//hudElement is just for use in by the HUD
type hudElement struct {
	UIElement
	hudX, hudY float64
}

//NewHUD returns a new HUD object
func NewHUD(screenWidth, screenHeight float64) *HUD {
	hud := &HUD{
		w:          screenWidth,
		h:          screenHeight,
		uiElements: make(map[string]*hudElement),
	}
	hud.createSections()
	return hud
}

//Update updates the HUD....prevHealth
func (hud *HUD) Update() {
	//hud.x = Components.Camera.GetX()
	//hud.y = Components.Camera.GetY()
	//if hud.x != hud.prevX || hud.y != hud.prevY {
	for _, element := range hud.topLeft.elements {
		element.UIElement.SetPosition(hud.x+element.hudX, hud.y+element.hudY)
		element.UIElement.Update()
	}
	for _, element := range hud.topRight.elements {
		element.UIElement.SetPosition(hud.x+element.hudX, hud.y+element.hudY)
		element.UIElement.Update()
	}
	for _, element := range hud.bottomRight.elements {
		element.UIElement.SetPosition(hud.x+element.hudX, hud.y+element.hudY)
		element.UIElement.Update()
	}
	for _, element := range hud.bottomLeft.elements {
		element.UIElement.SetPosition(hud.x+element.hudX, hud.y+element.hudY)
		element.UIElement.Update()
	}
	//}
	hud.prevX = hud.x
	hud.prevY = hud.y
}

//Draw All elements
func (hud *HUD) Draw(screen *ebiten.Image, camera *Camera) {
	for _, element := range hud.topLeft.elements {
		element.Draw(screen, camera)
	}
	for _, element := range hud.topRight.elements {
		element.Draw(screen, camera)
	}
	for _, element := range hud.bottomRight.elements {
		element.Draw(screen, camera)
	}
	for _, element := range hud.bottomLeft.elements {
		element.Draw(screen, camera)
	}
}

//AddTopLeft of the HUD
func (hud *HUD) AddTopLeft(element UIElement) {
	section := hud.topLeft
	h := &hudElement{
		hudX:      section.startX,
		UIElement: element,
		hudY:      section.startY,
	}
	if section.count > 0 {
		addition := 0
		for _, element := range section.elements {
			_, h := element.Size()
			addition += h
		}
		h.hudY += float64(addition)
	}
	section.count++
	section.addElement(h)
}

//AddTopRight of the hud
func (hud *HUD) AddTopRight(element UIElement) {
	section := hud.topRight
	width, _ := element.Size()
	h := &hudElement{
		hudX:      section.startX - float64(width),
		UIElement: element,
		hudY:      section.startY,
	}
	if section.count > 0 {
		addition := 0
		for _, element := range section.elements {
			_, h := element.Size()
			addition += h
		}
		h.hudY += float64(addition)
	}
	section.count++
	section.addElement(h)
}

//AddBottomRight of the HUD
func (hud *HUD) AddBottomRight(element UIElement) {
	section := hud.bottomRight
	width, height := element.Size()
	h := &hudElement{
		hudX:      section.startX - float64(width),
		UIElement: element,
		hudY:      section.startY - float64(height),
	}
	if section.count > 0 {
		addition := 0
		for _, element := range section.elements {
			_, h := element.Size()
			addition += h
		}
		h.hudY -= float64(addition)
	}

	section.count++
	section.addElement(h)
}

//AddBottomLeft of the HUD
func (hud *HUD) AddBottomLeft(element UIElement) {
	section := hud.bottomLeft
	_, height := element.Size()
	h := &hudElement{
		hudX:      section.startX,
		UIElement: element,
		hudY:      section.startY - float64(height),
	}
	if section.count > 0 {
		addition := 0
		for _, element := range section.elements {
			_, h := element.Size()
			addition += h
		}
		h.hudY -= float64(addition)
	}

	section.count++
	section.addElement(h)
}

type hudSection struct {
	startX, startY float64
	invertY        bool
	count          int
	elements       []*hudElement
}

func (hs *hudSection) addElement(element *hudElement) {
	hs.elements = append(hs.elements, element)
}

func (hud *HUD) createSections() {
	hud.topLeft = &hudSection{
		startX:   10,
		startY:   5,
		elements: []*hudElement{},
	}
	hud.topRight = &hudSection{
		startX:   hud.w - 10,
		startY:   5,
		elements: []*hudElement{},
	}
	hud.bottomLeft = &hudSection{
		elements: []*hudElement{},
		startX:   10,
		startY:   hud.h - 10,
		invertY:  true,
	}
	hud.bottomRight = &hudSection{
		elements: []*hudElement{},
		startX:   hud.w - 10,
		startY:   hud.h - 10,
		invertY:  true,
	}
}
