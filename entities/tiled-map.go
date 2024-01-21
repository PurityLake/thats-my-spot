package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
)

type TiledMap struct {
	ecs.BasicEntity
	components.Transform
	components.Renderable
	components.TiledMap
}

func (t *TiledMap) Update() {
	if t.Image == nil {
		t.Image = t.TempImage
		t.TempImage = nil
	}
}
