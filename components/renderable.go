package components

import (
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderable struct {
	Image  *ebiten.Image
	Pos    maths.Vector2
	Scale  maths.Vector2
	Rotate float64
}

func (r Renderable) GetDrawOptions(options *ebiten.DrawImageOptions) {
	options.GeoM.Reset()
	options.GeoM.Scale(float64(r.Scale.X), float64(r.Scale.Y))
	options.GeoM.Rotate(r.Rotate)
	options.GeoM.Translate(float64(r.Pos.X), float64(r.Pos.Y))
}
