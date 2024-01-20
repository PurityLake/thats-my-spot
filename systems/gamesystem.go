package systems

import (
	_ "image/png"
	"log"
	"math/rand"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/data"
	"github.com/PurityLake/thatsmyspot/entities"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameSystem struct {
	Entities       []*entities.RenderableEntity
	MapEntity      *entities.MapEntity
	TiledMapEntity *entities.TiledMapEntity
}

func (gs *GameSystem) New(world *ecs.World) {
	var mapTiles []data.Tile
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			i := rand.Intn(data.MaxTile)
			tile, err := data.NewTile(i)
			if err != nil {
				log.Fatal(err)
			}
			mapTiles = append(mapTiles, *tile)
		}
	}
	tileMap, err := components.NewTiledMap("assets/maps/tiled/map0.png")
	if err != nil {
		log.Fatal(err)
	}
	gs.TiledMapEntity = &entities.TiledMapEntity{
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
	gs.MapEntity = &entities.MapEntity{
		BasicEntity: ecs.NewBasic(),
		Transform: components.Transform{
			Pos:    maths.Vector2{X: 0, Y: 0},
			Scale:  maths.Vector2{X: 1, Y: 1},
			Rotate: 0,
		},
		Renderable: components.Renderable{
			Image: nil,
		},
		Map: components.Map{
			Width:  20,
			Height: 20,
			TileW:  40,
			TileH:  40,
			Tiles:  mapTiles,
		},
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/sprites/car0.png")
	if err != nil {
		log.Fatal(err)
	}
	basic := ecs.NewBasic()
	renderable := components.Renderable{Image: img}
	transform := components.Transform{
		Pos:    maths.Vector2{X: 20, Y: 20},
		Scale:  maths.Vector2{X: 0.25, Y: 0.25},
		Rotate: 0,
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
	gs.MapEntity.Update()
	gs.TiledMapEntity.Update()
	for _, entity := range gs.Entities {
		entity.Update()
	}
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
