package main

import (
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/constants"
	"github.com/PurityLake/thatsmyspot/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	world ecs.World
	won   bool
}

func (g *Game) Update() error {
	if !g.won {
		g.world.Update(0.016)
	}
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			g.won = sys.TiledMapEntity.Won
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			if sys.TiledMapEntity.Image == nil {
				continue
			}
			screen.Fill(sys.TiledMapEntity.BackgroundColor)
			bounds := sys.TiledMapEntity.Image.Bounds()
			w, h := bounds.Dx(), bounds.Dy()
			options := sys.TiledMapEntity.GetDrawOptions(float64(w), float64(h))
			screen.DrawImage(sys.TiledMapEntity.Image, options)
			for _, entity := range sys.Entities {
				bounds := entity.Image.Bounds()
				w, h := bounds.Dx(), bounds.Dy()
				options := entity.GetDrawOptions(float64(w), float64(h))
				screen.DrawImage(entity.Image, options)
			}
		case *systems.MainMenuSystem:
			screen.Fill(color.RGBA{0x55, 0xff, 0x55, 0xff})
			for _, entity := range sys.ButtonEntities {
				bounds := entity.Image.Bounds()
				w, h := bounds.Dx(), bounds.Dy()
				options := entity.GetDrawOptions(float64(w), float64(h))
				screen.DrawImage(entity.Image, options)
			}
		}
	}
	if g.won {
		ebitenutil.DebugPrint(screen, "You won!")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.WindowWidth, constants.WindowHeight
}

func main() {
	ebiten.SetWindowSize(constants.WindowWidth, constants.WindowHeight)
	ebiten.SetWindowTitle("That's My Spot!")
	world := ecs.World{}
	// world.AddSystem(&systems.GameSystem{})
	world.AddSystem(&systems.MainMenuSystem{})
	if err := ebiten.RunGame(&Game{world: world}); err != nil {
		log.Fatal(err)
	}
}
