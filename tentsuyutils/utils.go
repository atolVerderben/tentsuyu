package tentsuyutils

//NearCoords tests if (x,y) are "near" (x1, y1) based on radius
func NearCoords(x, y, x1, y1, radius float64) bool {
	//radius := 20.0
	if x <= (x1+radius) && x >= (x1-radius) {
		if y <= (y1+radius) && y >= (y1-radius) {
			return true
		}
	}

	return false
}
