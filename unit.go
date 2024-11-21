package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"

	"mygame/helpers"
)

type Unit struct {
	selected  bool
	size      rl.Vector2
	postition rl.Vector2
	target    rl.Vector2
}

func NewUnit(pos rl.Vector2, size rl.Vector2) *Unit {
	return &Unit{
		selected:  false,
		size:      size,
		postition: pos,
		target:    pos,
	}
}

func (u *Unit) GetTransform() (rl.Vector2, rl.Vector2) {
	return u.size, u.postition
}

func (u *Unit) MarkAsSelected(v bool) {
	u.selected = v
}

func (u *Unit) Update() {
	dist := rl.Vector2Distance(u.postition, u.target)
	dir := helpers.GetDirectionBetweenVectors(u.postition, u.target)

	if dist > 5 {
		u.postition.X += dir.X
		u.postition.Y += dir.Y
	}
}

func (u *Unit) Draw() {
	c := rl.Purple
	if u.selected {
		c = rl.Red
	}
	rect := helpers.NewRectangleVec2(u.postition, u.size)
	rl.DrawRectangleRec(rect, c)
}

func (u *Unit) SetTarget(tpos rl.Vector2) {
	randX := float32(rand.Intn(50))
	randY := float32(rand.Intn(50))

	if rand.Intn(1) == 1 {
		tpos.X += randX
	} else {
		tpos.X -= randX
	}

	if rand.Intn(1) == 1 {
		tpos.Y += randY
	} else {
		tpos.Y -= randY
	}

	u.target = tpos
}
