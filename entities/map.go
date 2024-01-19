package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
)

type MapEntity struct {
	ecs.BasicEntity
	components.Map
	components.Transform
	components.Renderable
}

func (m *MapEntity) Update() {
	if m.Image == nil {
		m.Image = m.CreateImage()
	}
}
