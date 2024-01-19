package systems

import (
	_ "image/png"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/entities"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameSystem struct {
	Entities []entities.RenderableEntity
}

func (gs *GameSystem) New(world *ecs.World) {
	img, _, err := ebitenutil.NewImageFromFile("assets/car0.png")
	if err != nil {
		log.Fatal(err)
	}
	gs.Add(&ecs.BasicEntity{}, &components.Renderable{Image: img})
}

func (gs *GameSystem) Add(basic *ecs.BasicEntity, rect *components.Renderable) {
	gs.Entities = append(gs.Entities, entities.RenderableEntity{BasicEntity: *basic, Renderable: *rect})
}

func (gs *GameSystem) Update(dt float32) {
}

func (gs *GameSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, entity := range gs.Entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		gs.Entities = append(gs.Entities[:delete], gs.Entities[:delete+1]...)
	}
}
