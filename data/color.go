package data

import (
	"image/color"
	"strconv"
)

func HexToRGBA(hex string) color.RGBA {
	hex = hex[1:]
	colorInt, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return color.RGBA{0, 0, 0, 255}
	}
	r := uint8(colorInt >> 16)
	g := uint8(colorInt >> 8)
	b := uint8(colorInt)
	return color.RGBA{r, g, b, 255}
}
