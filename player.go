package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type playerSelectionBox struct {
	start       rl.Vector2
	end         rl.Vector2
	isSelecting bool
}

type Player struct {
	camera        rl.Camera2D
	isDragging    bool
	previousMouse rl.Vector2

	sb       playerSelectionBox
	selected []WorldSelectable
	save1    []WorldSelectable
	save2    []WorldSelectable
}

func NewPlayer(c rl.Camera2D) *Player {
	return &Player{
		camera:   c,
		selected: []WorldSelectable{},
		save1:    []WorldSelectable{},
		save2:    []WorldSelectable{},
	}
}

func (p *Player) dragMove() {
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		currentMouse := rl.GetMousePosition()
		if !p.isDragging {
			p.previousMouse = currentMouse
			p.isDragging = true
		}

		delta := rl.Vector2Subtract(p.previousMouse, currentMouse)

		p.camera.Target.X += delta.X * CameraPanSpeed * rl.GetFrameTime()
		p.camera.Target.Y += delta.Y * CameraPanSpeed * rl.GetFrameTime()

		p.previousMouse = currentMouse
	} else {
		p.isDragging = false
	}
}

func (p *Player) wasdMove() {
	if rl.IsKeyDown(rl.KeyRight) {
		p.camera.Target.X += 1
	} else if rl.IsKeyDown(rl.KeyLeft) {
		p.camera.Target.X -= 1
	}

	if rl.IsKeyDown(rl.KeyUp) {
		p.camera.Target.Y -= 1
	} else if rl.IsKeyDown(rl.KeyDown) {
		p.camera.Target.Y += 1
	}
}

func (p *Player) selectionBox() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		p.sb.start = rl.GetScreenToWorld2D(rl.GetMousePosition(), p.camera)
		p.sb.isSelecting = true
		for _, s := range p.selected {
			s.MarkAsSelected(false)
		}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		p.sb.end = rl.GetScreenToWorld2D(rl.GetMousePosition(), p.camera)
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		p.sb.isSelecting = false
		rectX := min(p.sb.start.X, p.sb.end.X)
		rectY := min(p.sb.start.Y, p.sb.end.Y)
		rectWidth := float32(math.Abs(float64(p.sb.end.X - p.sb.start.X)))
		rectHeight := float32(math.Abs(float64(p.sb.end.Y - p.sb.start.Y)))
		rect := rl.NewRectangle(rectX, rectY, rectWidth, rectHeight)

		p.selected = g.GetWithinSelection(rect)
		for _, s := range p.selected {
			s.MarkAsSelected(true)
		}
	}
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyLeftShift) {
		p.sb.isSelecting = false
		p.wasdMove()
		p.dragMove()
	}

	if !rl.IsKeyDown(rl.KeyLeftShift) {
		p.selectionBox()
	}

	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		mPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), p.camera)
		if len(p.selected) > 0 {
			for _, u := range p.selected {
				x, ok := u.(*Unit)
				if !ok {
					continue
				}

				x.SetTarget(mPos)
			}
		}
	}

	if rl.IsKeyDown(rl.KeyLeftControl) && len(p.selected) > 0 {
		if rl.IsKeyPressed(rl.KeyOne) {
			p.save1 = make([]WorldSelectable, 0)
			for _, u := range p.selected {
				p.save1 = append(p.save1, u)
			}
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			p.save2 = make([]WorldSelectable, 0)
			for _, u := range p.selected {
				p.save2 = append(p.save2, u)
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyOne) {
		for _, u := range p.selected {
			u.MarkAsSelected(false)
		}

		p.selected = make([]WorldSelectable, 0)

		for _, u := range p.save1 {
			p.selected = append(p.selected, u)
			u.MarkAsSelected(true)
		}
	}

	if rl.IsKeyPressed(rl.KeyTwo) {
		for _, u := range p.selected {
			u.MarkAsSelected(false)
		}

		p.selected = make([]WorldSelectable, 0)

		for _, u := range p.save2 {
			p.selected = append(p.selected, u)
			u.MarkAsSelected(true)
		}
	}
}

func (p *Player) Draw() {
	rl.DrawRectangleRec(
		rl.NewRectangle(p.camera.Target.X, p.camera.Target.Y, 10, 10),
		rl.Blue,
	)

	if p.sb.isSelecting {
		rectX := min(p.sb.start.X, p.sb.end.X)
		rectY := min(p.sb.start.Y, p.sb.end.Y)
		rectWidth := float32(math.Abs(float64(p.sb.end.X - p.sb.start.X)))
		rectHeight := float32(math.Abs(float64(p.sb.end.Y - p.sb.start.Y)))

		rl.DrawRectangleLinesEx(
			rl.NewRectangle(rectX, rectY, rectWidth, rectHeight),
			2,
			rl.Green,
		)
	}
}
