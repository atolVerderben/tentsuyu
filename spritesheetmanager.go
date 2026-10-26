package tentsuyu

//SpriteSheetManager handles the spritesheets for the game
type SpriteSheetManager struct {
	SpriteSheets map[string]*SpriteSheet
}

//NewSpriteSheetManager returns a new SpriteSheetManager
func NewSpriteSheetManager() *SpriteSheetManager {
	return &SpriteSheetManager{
		SpriteSheets: map[string]*SpriteSheet{},
	}
}

//AddSpriteSheet to the manager class based on the given name
func (s *SpriteSheetManager) AddSpriteSheet(name string, ss *SpriteSheet) {
	s.SpriteSheets[name] = ss
}

//ReturnSpriteSheet based on the name give
func (s *SpriteSheetManager) ReturnSpriteSheet(name string) *SpriteSheet {
	return s.SpriteSheets[name]
}
