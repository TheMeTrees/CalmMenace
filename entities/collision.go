package entities

import rl "github.com/gen2brain/raylib-go/raylib"

// Minimum Translation Vector
func ResolveCollision(a, b rl.Rectangle) (float32, float32) {
	if !rl.CheckCollisionRecs(a, b) {
		return 0, 0
	}

	dx1 := (b.X + b.Width) - a.X  // Push right
	dx2 := (a.X + a.Width) - b.X  // Push left
	dy1 := (b.Y + b.Height) - a.Y // Push down
	dy2 := (a.Y + a.Width) - b.Y  // Push up

	minX := dx1
	if dx2 < dx1 {
		minX = -dx2
	}

	minY := dy1
	if dy2 < dy1 {
		minY = -dy2
	}

	// Choose the smallest absolute correction
	if max(minX, -minX) < max(minY, -minY) {
		return minX, 0
	}

	return 0, minY
}
