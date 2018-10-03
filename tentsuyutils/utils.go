package tentsuyutils

import (
	"math"
	"math/rand"
)

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

//WithinDistance is a simple function to check if x is withing a certain distance of x1
func WithinDistance(x, x1, radius float64) bool {
	if x <= (x1+radius) && x >= (x1-radius) {
		return true
	}
	return false
}

//RandomBetween returns a random int between min and max respectively
//This is just a useful function for many reasons
func RandomBetween(min, max int) int {
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//Distance returns the scalar distance between two points
func Distance(x, y, x1, y1 float64) float64 {

	return math.Sqrt(math.Pow(x-x1, 2) + math.Pow(y-y1, 2))

}
