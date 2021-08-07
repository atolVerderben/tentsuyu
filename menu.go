package tentsuyu

import "github.com/hajimehoshi/ebiten/v2"

//Menu is a collection of MenuElements
type Menu struct {
	Elements                               [][]*MenuElement
	Active, background                     bool
	Name                                   string
	x, y, midStartX, midStartY, minx, miny float64
	paddingX, paddingY                     int
	maxWidth, maxHeight                    int
	backgroundImage                        *ebiten.Image
	backgroundImgParts                     *BasicImageParts
	selectedRow, selectedCol               int
}

//NewMenu creates a new menu
func NewMenu(screenWidth, screenHeight float64) *Menu {
	m := &Menu{
		midStartX:   screenWidth / 2,
		midStartY:   screenHeight / 8,
		paddingX:    2,
		paddingY:    5,
		selectedRow: -1,
		selectedCol: -1,
	}
	return m
}

//SetPadding of the menus x and y values
func (m *Menu) SetPadding(x, y int) {
	m.paddingX = x
	m.paddingY = y
}

//SetBackground image for menu
func (m *Menu) SetBackground(src *ebiten.Image, imgParts *BasicImageParts) {
	m.background = true

	m.backgroundImage = src
	m.backgroundImgParts = imgParts
}

//SelectHorizontal moves the selection alrong the x-axis
func (m *Menu) SelectHorizontal(id int) {
	m.selectedCol = 0
	m.selectedRow = id
}

//SelectVertical move the selection along the y-axis
func (m *Menu) SelectVertical(id int) {
	m.selectedCol = id
}

//ReturnSelected returns the Row and Col int of the current selected option
func (m *Menu) ReturnSelected() (int, int) {
	return m.selectedRow, m.selectedCol
}

//PressSelected exectures the Action of the currently selected option
func (m *Menu) PressSelected() {
	if m.selectedCol == -1 || m.selectedRow == -1 {
		return
	}
	m.Elements[m.selectedRow][m.selectedCol].Action()
}

//Update the menu
func (m *Menu) Update(input *InputController, offsetX, offsetY float64) {
	//log.Printf("%v,%v\n", m.maxWidth, m.maxHeight)
	//m.x = Components.Camera.GetX()
	//m.y = Components.Camera.GetY()

	for x := range m.Elements {
		for y := range m.Elements[x] {
			switch e := m.Elements[x][y].UIElement.(type) {
			case *TextElement:
				e.NotCentered = true
			}
			mx := m.x + m.Elements[x][y].menuX
			my := m.y + m.Elements[x][y].menuY

			if x == 0 {
				if y == 0 {
					m.miny = my
				}
				m.minx = mx
			} else {
				if mx < m.minx {
					m.minx = mx
				}
			}

			//m.elements[x][y].UIElement.SetPosition(mx, my)
			if m.Elements[x][y].Update(input, offsetX, offsetY) {
				m.selectedCol = -1
				m.selectedRow = -1
			}
			if x == m.selectedRow && y == m.selectedCol {
				m.Elements[x][y].highlighted = true
				m.Elements[x][y].Highlighted()
			}
		}
	}
}

//Draw window
func (m *Menu) Draw(screen *ebiten.Image, camera *Camera) {
	if m.background {
		//w := 50.0 //float64(m.maxWidth) / 2
		//h := 50.0 //float64(m.maxHeight) / 2
		scalex := float64(m.maxWidth)/100 + .85
		scaley := float64(m.maxHeight)/100 + .2
		op := &ebiten.DrawImageOptions{}
		//op.ImageParts = m.backgroundImgParts
		//op.GeoM.Translate(-w, -h)
		op.GeoM.Scale(float64(scalex), float64(scaley))
		//op.GeoM.Translate(w*float64(scalex), h*float64(scaley))
		op.GeoM.Translate(m.minx-50, m.miny-10)
		//log.Printf("%v,%v\n", scalex, scaley)
		//ApplyCameraTransform(op, false)

		screen.DrawImage(m.backgroundImgParts.SubImage(m.backgroundImage), op)

	}
	for x := range m.Elements {
		for y := range m.Elements[x] {
			if !m.Elements[x][y].hidden {
				m.Elements[x][y].UIElement.Draw(screen, camera)
			}
		}
	}
}

//AddElement adds a new Line of UIElements
func (m *Menu) AddElement(element []UIElement, action []func()) {

	menuY := m.midStartY + float64(m.maxHeight)

	/*if len(m.elements) > 0 {
		addition := 0
		for x := range m.elements {
			maxHeight := 0
			for y := range m.elements[x] {
				_, h := m.elements[x][y].Size()
				if h > maxHeight {
					maxHeight = h
				}

			}
			addition += m.paddingY + maxHeight
		}
		menuY += float64(addition)
		if addition > m.maxHeight {
			m.maxHeight += addition
		}
	}*/

	MenuElements := []*MenuElement{}
	maxWidth := 0
	maxHeight := 0
	for i := range element {
		width, height := element[i].Size()
		mE := &MenuElement{
			menuX:     m.midStartX - float64(width/2),
			UIElement: element[i],
			menuY:     menuY,
		}
		switch u := element[i].(type) {
		case *TextElement:
			u.Stationary = true
		}
		mE.SetAction(action[i])
		mE.SetPosition(mE.menuX, mE.menuY)
		MenuElements = append(MenuElements, mE)
		maxWidth += width + m.paddingX
		if height > maxHeight {
			maxHeight = height
		}
	}

	if maxWidth > m.maxWidth {
		m.maxWidth = maxWidth
	}
	maxWidth -= m.paddingX
	m.maxHeight += maxHeight + m.paddingY
	if len(element) > 1 {
		lineStartX := m.midStartX - float64(maxWidth/2)
		for i := range MenuElements {
			MenuElements[i].menuX = lineStartX
			MenuElements[i].SetPosition(MenuElements[i].menuX, MenuElements[i].menuY)
			width, _ := element[i].Size()
			lineStartX += float64(width + m.paddingX)
		}
		lineStartX -= float64(m.paddingX)
	}

	m.Elements = append(m.Elements, MenuElements)
}

//MenuElement represents a single element/option within a menu
type MenuElement struct {
	UIElement
	Action                  func()
	highlighted, Selectable bool
	menuX, menuY            float64
	hidden                  bool
}

//SetAction of the MenuElement
func (m *MenuElement) SetAction(function func()) {
	m.Action = function
	if function == nil {
		m.Selectable = false
	} else {
		m.Selectable = true
	}
}

//Update the MenuElement
func (m *MenuElement) Update(input *InputController, offsetX, offsetY float64) bool {
	if m.hidden {
		return false
	}
	mx, my := input.GetMouseCoords()
	mouseHighlight := false
	if m.UIElement.Contains(mx+offsetX, my+offsetY) {
		if m.Selectable {
			mouseHighlight = true
			m.Highlighted()
			m.highlighted = true
			if input.LeftClick().JustPressed() {

				if m.Action != nil {
					m.Action()
				}
			}
		}
	} else {
		if m.highlighted == true {
			m.highlighted = false
			m.UnHighlighted()
		}
	}

	m.UIElement.Update()
	return mouseHighlight
}

//UpdateWithCamera the MenuElement in context of the game camera
func (m *MenuElement) UpdateWithCamera(input *InputController, camera *Camera, offsetX, offsetY float64) bool {
	if m.hidden {
		return false
	}
	mx, my := input.GetGameMouseCoords(camera)
	mouseHighlight := false
	if m.UIElement.Contains(mx+offsetX, my+offsetY) {
		if m.Selectable {
			mouseHighlight = true
			m.Highlighted()
			m.highlighted = true
			if input.LeftClick().JustPressed() {

				if m.Action != nil {
					m.Action()
				}
			}
		}
	} else {
		if m.highlighted == true {
			m.highlighted = false
			m.UnHighlighted()
		}
	}

	m.UIElement.Update()
	return mouseHighlight
}

//Hide takes a bool and sets the hidden variable
func (m *MenuElement) Hide(h bool) {
	m.hidden = h
}

//SetCentered centers the MenuElement based on bool value c
func (m *MenuElement) SetCentered(c bool) {
	switch e := m.UIElement.(type) { //Check the type of element
	case *TextElement:
		e.SetCentered(c)
	}
}
