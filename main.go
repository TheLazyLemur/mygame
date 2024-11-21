package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 800
)

type Updatable interface {
	Update()
}

type Drawable interface {
	Draw()
}

type WorldSelectable interface {
	GetTransform() (rl.Vector2, rl.Vector2)
	MarkAsSelected(bool)
}

var g = NewGame()

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "RTS Game")
	defer rl.CloseWindow()

	player := NewPlayer(rl.NewCamera2D(rl.Vector2{
		X: float32(ScreenWidth) / 2,
		Y: float32(ScreenHeight) / 2,
	}, rl.Vector2{
		X: 0,
		Y: 0,
	}, 0, 1))

	g.AddUpdatable(player)
	g.AddDrawable(player)

	for i := 0; i < 30; i++ {
		angle := 2 * math.Pi * float64(i) / float64(30)

		x := 200 + 100*float32(math.Cos(angle))
		y := 200 + 100*float32(math.Sin(angle))

		u := NewUnit(rl.Vector2{
			X: float32(x),
			Y: float32(y),
		}, rl.Vector2{
			X: 10,
			Y: 10,
		})

		g.AddUpdatable(u)
		g.AddDrawable(u)
		g.AddWorldSelectable(u)
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode2D(player.camera)
		g.Draw()
		rl.EndMode2D()

		g.DrawUI()
		rl.EndDrawing()
	}
}
