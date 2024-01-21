package constants

import "image/color"

const (
	WindowWidth      = 800
	WindowHeight     = 600
	HalfWindowWidth  = WindowWidth / 2
	HalfWindowHeight = WindowHeight / 2
	ButtonWidth      = 150
	ButtonHeight     = 30
)

var (
	ButtonColor      color.RGBA = color.RGBA{0xff, 0xff, 0xff, 0xff}
	ButtonHoverColor color.RGBA = color.RGBA{0xff, 0x00, 0x00, 0xff}
)
