package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	world ecs.World
}

func (g *Game) Update() error {
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			for _, entity := range sys.Entities {
				entity.Update()
			}
		}
	}
	g.world.Update(0.016)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, system := range g.world.Systems() {
		switch sys := system.(type) {
		case *systems.GameSystem:
			for _, entity := range sys.Entities {
				bounds := entity.Image.Bounds()
				w, h := bounds.Dx(), bounds.Dy()
				options := entity.GetDrawOptions(float64(w), float64(h))
				screen.DrawImage(entity.Image, options)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	world := ecs.World{}
	world.AddSystem(&systems.GameSystem{})
	if err := ebiten.RunGame(&Game{world: world}); err != nil {
		log.Fatal(err)
	}
}
