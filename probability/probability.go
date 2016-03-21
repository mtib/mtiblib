package probability

import (
	"math/rand"
	"time"
)

// Initialize sets the Seed
func Initialize() {
	rand.Seed(time.Now().UnixNano())
}

// RandomFloat returns a float between 0.0 and 1.0
func RandomFloat() float64 {
	return rand.Float64()
}

// RandomInt returns an int between 0 and max
func RandomInt(max int) int {
	return rand.Intn(max)
}

// RandomF returns a float between min and max
func RandomF(min, max float64) float64 {
	return RandomFloat()*(max-min) + min
}

// RandomI returns and int between min and max
func RandomI(min, max int) int {
	return RandomInt(max-min) + min
}

// Probability returns true prob*100% of the time
func Probability(prob float64) bool {
	if prob <= 0 {
		return false
	}
	if prob >= 1 {
		return true
	}
	if prob > RandomFloat() {
		return true
	}
	return false
}

// Half is true 50% of the time
func Half() bool {
	return Probability(0.50)
}
