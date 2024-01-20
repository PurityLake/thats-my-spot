package components

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TiledMap struct {
	Width, Height int
	TileW, TileH  int
	TempImage     *ebiten.Image
}

func NewTiledMap(filename string) (*TiledMap, error) {
	tiledMap, _, err := ebitenutil.NewImageFromFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	bounds := tiledMap.Bounds()
	return &TiledMap{
		Width:     bounds.Dx(),
		Height:    bounds.Dy(),
		TileW:     40,
		TileH:     40,
		TempImage: tiledMap,
	}, nil
}
