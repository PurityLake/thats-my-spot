package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
)

type ButtonEntity struct {
	ecs.BasicEntity
	components.Transform
	components.Renderable
	components.Button
}
