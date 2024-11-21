package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"mygame/helpers"
)

type Game struct {
	updatableEntities []Updatable
	drawableEntities  []Drawable
	worldSelectable   []WorldSelectable
}

func NewGame() *Game {
	return &Game{
		updatableEntities: []Updatable{},
		drawableEntities:  []Drawable{},
	}
}

func (g *Game) GetWithinSelection(sb rl.Rectangle) []WorldSelectable {
	withinSelection := []WorldSelectable{}

	for _, s := range g.worldSelectable {
		size, pos := s.GetTransform()
		rect := helpers.NewRectangleVec2(pos, size)

		if rl.CheckCollisionRecs(rect, sb) {
			withinSelection = append(withinSelection, s)
		}
	}

	return withinSelection
}

func (g *Game) AddUpdatable(ent Updatable) {
	g.updatableEntities = append(g.updatableEntities, ent)
}

func (g *Game) AddDrawable(ent Drawable) {
	g.drawableEntities = append(g.drawableEntities, ent)
}

func (g *Game) AddWorldSelectable(ent WorldSelectable) {
	g.worldSelectable = append(g.worldSelectable, ent)
}

func (g *Game) Update() {
	for _, ent := range g.updatableEntities {
		ent.Update()
	}
}

func (g *Game) Draw() {
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	for _, ent := range g.drawableEntities {
		ent.Draw()
	}
}

func (g *Game) DrawUI() {
	rl.DrawFPS(10, 10)
}
