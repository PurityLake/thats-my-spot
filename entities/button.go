package entities

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/PurityLake/thatsmyspot/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type ButtonEntity struct {
	ecs.BasicEntity
	components.Transform
	components.Renderable
	components.Button
}

func (be *ButtonEntity) Update(fontFace font.Face) {
	x, y := ebiten.CursorPosition()
	oldHovered := be.Hovered
	be.Hovered = be.Bounds.IsPointInBounds(x, y)
	if oldHovered != be.Hovered {
		be.Redraw(fontFace)
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		be.Pressed = true
	}
}

func (be *ButtonEntity) Redraw(fontFace font.Face) {
	if be.Hovered {
		be.Image.Fill(be.HoverColor)
	} else {
		be.Image.Fill(be.Color)
	}
	bounds, _ := font.BoundString(fontFace, be.Text)
	text.Draw(be.Image, be.Text, fontFace,
		be.Bounds.W/2-bounds.Max.X.Ceil()/2,
		be.Bounds.H/2-bounds.Min.Y.Ceil()/2,
		color.RGBA{0, 0, 0, 255})
}
