package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
)

type RenderableEntity struct {
	ecs.BasicEntity
	components.Renderable
}
