package systems

import (
	"image/color"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/data"
	"github.com/PurityLake/thatsmyspot/entities"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type MainMenuSystem struct {
	ButtonEntities []*entities.ButtonEntity
}

func (mm *MainMenuSystem) New(world *ecs.World) {
	basic := ecs.NewBasic()
	transform := components.Transform{
		Pos:    maths.Vector2{X: 100, Y: 100},
		Scale:  maths.Vector2{X: 1, Y: 1},
		Rotate: 0,
		Anchor: maths.Vector2{X: 1.0, Y: 1.0},
	}
	rect := maths.Bounds{
		W: 150,
		H: 30,
	}
	label := "Start Game"
	file, err := data.GetFile("assets/fonts/joystix.otf")
	if err != nil {
		log.Fatal(err)
	}

	fontObj, err := opentype.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fontFace, err := opentype.NewFace(fontObj,
		&opentype.FaceOptions{
			Size:    12,
			DPI:     72,
			Hinting: font.HintingVertical,
		})
	if err != nil {
		log.Fatal(err)
	}

	image := ebiten.NewImage(rect.W, rect.H)
	image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	bounds, _ := font.BoundString(fontFace, label)
	text.Draw(image, label, fontFace, rect.W/2-bounds.Max.X.Ceil()/2, (rect.H/2 - bounds.Min.Y.Ceil()/2), color.RGBA{0, 0, 0, 255})

	renderable := components.Renderable{
		Image: image,
	}

	button := components.Button{
		Bounds:  rect,
		Hovered: false,
	}

	mm.Add(&basic, &transform, &renderable, &button)
}

func (mm *MainMenuSystem) Add(basic *ecs.BasicEntity, transform *components.Transform, renderable *components.Renderable, button *components.Button) {
	mm.ButtonEntities = append(mm.ButtonEntities,
		&entities.ButtonEntity{
			BasicEntity: *basic,
			Transform:   *transform,
			Renderable:  *renderable,
			Button:      *button,
		})
}

func (mm *MainMenuSystem) Update(dt float32) {
}

func (mm *MainMenuSystem) Remove(basic ecs.BasicEntity) {
}
