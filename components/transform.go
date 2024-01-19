package components

import (
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	Pos    maths.Vector2
	Scale  maths.Vector2
	Rotate float64
}

func (t Transform) GetDrawOptions(options *ebiten.DrawImageOptions) {
	options.GeoM.Reset()
	options.GeoM.Scale(float64(t.Scale.X), float64(t.Scale.Y))
	options.GeoM.Rotate(t.Rotate)
	options.GeoM.Translate(float64(t.Pos.X), float64(t.Pos.Y))
}
