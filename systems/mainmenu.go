package systems

import (
	"image/color"
	"log"
	"runtime"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/PurityLake/thatsmyspot/constants"
	"github.com/PurityLake/thatsmyspot/data"
	"github.com/PurityLake/thatsmyspot/entities"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type ButtonEntry struct {
	X, Y       int
	W, H       int
	Text       string
	Name       string
	Color      color.RGBA
	HoverColor color.RGBA
}

var ButtonList = []ButtonEntry{
	{
		X:          constants.HalfWindowWidth,
		Y:          constants.HalfWindowHeight + 100,
		W:          150,
		H:          30,
		Text:       "Start Game",
		Name:       "start",
		Color:      color.RGBA{0xff, 0xff, 0xff, 0xff},
		HoverColor: color.RGBA{0xff, 0x00, 0x00, 0xff},
	},
	{
		X:          constants.HalfWindowWidth,
		Y:          constants.HalfWindowHeight + 150,
		W:          150,
		H:          30,
		Text:       "Exit Game",
		Name:       "exit",
		Color:      color.RGBA{0xff, 0xff, 0xff, 0xff},
		HoverColor: color.RGBA{0xff, 0x00, 0x00, 0xff},
	},
}

type MainMenuSystem struct {
	ButtonEntities []*entities.Button
	FontFace       font.Face
}

func (mm *MainMenuSystem) New(world *ecs.World) {
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
	mm.FontFace = fontFace

	for _, button := range ButtonList {
		if button.Name == "exit" && runtime.GOOS == "js" {
			continue
		}
		basic := ecs.NewBasic()
		transform := components.Transform{
			Pos: maths.Vector2{
				X: float32(button.X),
				Y: float32(button.Y),
			},
			Rotate: 0,
			Scale:  maths.Vector2{X: 1, Y: 1},
		}
		button := components.Button{
			Bounds: maths.Bounds{
				X: button.X,
				Y: button.Y,
				W: button.W,
				H: button.H,
			},
			Hovered:    false,
			Text:       button.Text,
			Pressed:    false,
			Name:       button.Name,
			Color:      button.Color,
			HoverColor: button.HoverColor,
		}

		image := ebiten.NewImage(button.Bounds.W, button.Bounds.H)
		image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		bounds, _ := font.BoundString(fontFace, button.Text)
		text.Draw(image, button.Text, fontFace,
			button.Bounds.W/2-bounds.Max.X.Ceil()/2,
			button.Bounds.H/2-bounds.Min.Y.Ceil()/2,
			color.RGBA{0, 0, 0, 255})

		renderable := components.Renderable{
			Image: image,
		}
		mm.Add(&basic, &transform, &renderable, &button)
	}
}

func (mm *MainMenuSystem) Add(basic *ecs.BasicEntity, transform *components.Transform, renderable *components.Renderable, button *components.Button) {
	mm.ButtonEntities = append(mm.ButtonEntities,
		&entities.Button{
			BasicEntity: *basic,
			Transform:   *transform,
			Renderable:  *renderable,
			Button:      *button,
		})
}

func (mm *MainMenuSystem) Update(dt float32) {
	for _, entity := range mm.ButtonEntities {
		entity.Update(mm.FontFace)
	}
}

func (mm *MainMenuSystem) Remove(basic ecs.BasicEntity) {
}
