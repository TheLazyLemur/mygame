package helpers

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewRectangleVec2(pos rl.Vector2, size rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(pos.X, pos.Y, size.X, size.Y)
}

func GetDirectionBetweenVectors(A, B rl.Vector2) rl.Vector2 {
	dir := rl.Vector2{
		X: B.X - A.X,
		Y: B.Y - A.Y,
	}

	magnitude := float32(math.Sqrt(float64(dir.X*dir.X + dir.Y*dir.Y)))

	if magnitude != 0 {
		dir.X /= magnitude
		dir.Y /= magnitude
	}

	return dir
}
