package systems

import (
	_ "image/png"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/entities"
	"github.com/PurityLake/thatsmyspot/maths"
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
	renderable := components.Renderable{
		Image:  img,
		Pos:    maths.Vector2{X: 100, Y: 100},
		Scale:  maths.Vector2{X: 0.5, Y: 0.5},
		Rotate: 0.0,
	}
	gs.Add(&ecs.BasicEntity{}, &renderable)
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
