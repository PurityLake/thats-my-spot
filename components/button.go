package components

import (
	"image/color"

	"github.com/PurityLake/thatsmyspot/maths"
)

type Button struct {
	Hovered    bool
	Text       string
	Bounds     maths.Bounds
	Pressed    bool
	Name       string
	Color      color.Color
	HoverColor color.Color
}
