package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type RenderableEntity struct {
	ecs.BasicEntity
	components.Renderable
	components.Transform
	IsPlayer bool
}

func (r *RenderableEntity) Update() {
	if !r.IsPlayer {
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if r.Rotate == 270 {
			r.Pos.X -= 32
		} else {
			r.Rotate = 270
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if r.Rotate == 90 {
			r.Pos.X += 32
		} else {
			r.Rotate = 90
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if r.Rotate == 0 {
			r.Pos.Y -= 32
		} else {
			r.Rotate = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if r.Rotate == 180 {
			r.Pos.Y += 32
		} else {
			r.Rotate = 180
		}
	}
}
