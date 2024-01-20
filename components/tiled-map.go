package components

import (
	"bytes"
	"fmt"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type TiledMap struct {
	Width, Height int
	TileW, TileH  int
	TempImage     *ebiten.Image
}

func NewTiledMap(filename string) (*TiledMap, error) {
	tiledMap := &TiledMap{}
	gameMap, err := tiled.LoadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tiledMap.TileW = 40
	tiledMap.TileH = 40
	tiledMap.Width = gameMap.Width * tiledMap.TileW
	tiledMap.Height = gameMap.Height * tiledMap.TileH
	fmt.Printf("width x height: %d x %d\n", tiledMap.Width, tiledMap.Height)

	mapRenderer, err := render.NewRenderer(gameMap)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = mapRenderer.RenderVisibleLayers()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var buf []byte
	buffer := bytes.NewBuffer(buf)

	mapRenderer.SaveAsPng(buffer)

	im, err := png.Decode(buffer)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tiledMap.TempImage = ebiten.NewImageFromImage(im)

	return tiledMap, nil
}
