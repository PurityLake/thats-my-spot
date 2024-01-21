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
	Entities []*entities.RenderableEntity
	TiledMap *entities.TiledMap
}

func (gs *GameSystem) New(world *ecs.World) {
	tileMap, properties, err := components.NewTiledMap(
		"assets/maps/tiled/map0.png",
		"assets/maps/tiled/map0.json",
		"assets/maps/tiled/jamegam.json")
	if err != nil {
		log.Fatal(err)
	}
	gs.TiledMap = &entities.TiledMap{
		BasicEntity: ecs.NewBasic(),
		Transform: components.Transform{
			Pos:    maths.Vector2{X: 0, Y: 0},
			Scale:  maths.Vector2{X: 1, Y: 1},
			Rotate: 0,
			Anchor: maths.Vector2{X: 1.0, Y: 1.0},
		},
		Renderable: components.Renderable{
			Image: nil,
		},
		TiledMap: *tileMap,
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/sprites/car0.png")
	if err != nil {
		log.Fatal(err)
	}
	playerX := int(properties["PlayerX"].Value.(float64))
	playerY := int(properties["PlayerY"].Value.(float64))
	playerRotation := properties["PlayerRotation"].Value.(float64)
	basic := ecs.NewBasic()
	renderable := components.Renderable{Image: img}
	transform := components.Transform{
		Pos: maths.Vector2{
			X: float32(20 + (playerX * tileMap.TileW)),
			Y: float32(20 + (playerY * tileMap.TileH)),
		},
		Scale:  maths.Vector2{X: 0.25, Y: 0.25},
		Rotate: playerRotation,
	}
	gs.Add(&basic, &renderable, &transform, true)
}

func (gs *GameSystem) Add(basic *ecs.BasicEntity, rect *components.Renderable, transform *components.Transform, isPlayer bool) {
	gs.Entities = append(gs.Entities,
		&entities.RenderableEntity{
			BasicEntity: *basic,
			Renderable:  *rect,
			Transform:   *transform,
			IsPlayer:    isPlayer,
		})
}

func (gs *GameSystem) Update(dt float32) {
	for _, entity := range gs.Entities {
		entity.Update(gs.TiledMap)
	}
	gs.TiledMap.Update()
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
