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
	world    ecs.World
	mainMenu bool
	won      bool
}

func (g *Game) Update() error {
	if !g.won {
		g.world.Update(0.016)
	}
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			g.won = sys.TiledMap.Won
		case *systems.MainMenuSystem:
			if g.mainMenu {
				for _, entity := range sys.ButtonEntities {
					if entity.Name == "start" && entity.Pressed {
						g.mainMenu = false
						g.world = ecs.World{}
						g.world.AddSystem(&systems.GameSystem{})
						break
					}
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			if sys.TiledMap.Image == nil {
				continue
			}
			screen.Fill(sys.TiledMap.BackgroundColor)
			bounds := sys.TiledMap.Image.Bounds()
			w, h := bounds.Dx(), bounds.Dy()
			options := sys.TiledMap.GetDrawOptions(float64(w), float64(h))
			screen.DrawImage(sys.TiledMap.Image, options)
			for _, entity := range sys.Entities {
				bounds := entity.Image.Bounds()
				w, h := bounds.Dx(), bounds.Dy()
				options := entity.GetDrawOptions(float64(w), float64(h))
				screen.DrawImage(entity.Image, options)
			}
		case *systems.MainMenuSystem:
			if g.mainMenu {
				screen.Fill(color.RGBA{0x55, 0xff, 0x55, 0xff})
				for _, entity := range sys.ButtonEntities {
					bounds := entity.Image.Bounds()
					w, h := bounds.Dx(), bounds.Dy()
					options := entity.GetDrawOptions(float64(w), float64(h))
					screen.DrawImage(entity.Image, options)
				}
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
	game := &Game{mainMenu: true, world: world}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
