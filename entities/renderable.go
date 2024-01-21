package entities

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type RenderableEntity struct {
	ecs.BasicEntity
	components.Renderable
	components.Transform
	IsPlayer bool
}

func (r *RenderableEntity) Update(mapEntity *TiledMapEntity) {
	if !r.IsPlayer {
		return
	}
	playerMove := maths.Vector2{X: 0, Y: 0}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		r.Rotate -= 90
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		r.Rotate += 90
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		switch r.Rotate {
		case 0:
			playerMove.Y = -1
		case 90:
			playerMove.X = 1
		case 180:
			playerMove.Y = 1
		case 270:
			playerMove.X = -1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		switch r.Rotate {
		case 0:
			playerMove.Y = 1
		case 90:
			playerMove.X = -1
		case 180:
			playerMove.Y = -1
		case 270:
			playerMove.X = 1
		}
	}

	if r.Rotate < 0 {
		r.Rotate += 360
	}
	r.Rotate = float64(int(r.Rotate) % 360)

	if playerMove.X != 0 || playerMove.Y != 0 {
		x, y := mapEntity.MapPosFromScreenPos(int(r.Pos.X), int(r.Pos.Y))
		newX, newY, err := mapEntity.GetLastTileInDir(x, y, playerMove)
		if err != nil {
			log.Fatal(err)
		}
		r.Pos.X, r.Pos.Y = mapEntity.ScreenPosFromMapPos(newX, newY)
		mapEntity.HasWon(newX, newY)
	}
}
