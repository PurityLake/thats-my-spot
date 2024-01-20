package components

import (
	"math"

	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	Pos    maths.Vector2
	Scale  maths.Vector2
	Rotate float64
	Anchor maths.Vector2
}

func (t Transform) GetDrawOptions(w, h float64) *ebiten.DrawImageOptions {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Reset()
	options.GeoM.Translate(-w/2, -h/2)
	options.GeoM.Scale(float64(t.Scale.X), float64(t.Scale.Y))
	options.GeoM.Rotate((math.Pi / 180) * t.Rotate)
	options.GeoM.Translate(w/2*float64(t.Anchor.X), h/2*float64(t.Anchor.Y))
	options.GeoM.Translate(float64(t.Pos.X), float64(t.Pos.Y))
	return options
}
