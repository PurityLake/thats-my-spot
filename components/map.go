package components

import (
	"errors"
	"image"

	"github.com/PurityLake/thatsmyspot/data"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	Width, Height int
	TileW, TileH  int
	Tiles         []data.Tile
}

func (m *Map) GetTile(x, y int) (*data.Tile, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return nil, errors.New("out of bounds")
	}
	return &m.Tiles[x+y*m.Width], nil
}

func (m Map) CanGo(x, y int) bool {
	tile, err := m.GetTile(x, y)
	if err != nil {
		return false
	}
	return tile.IsEmpty()
}

func (m Map) GetLastTileInDir(x, y int, dir maths.Vector2) (*data.Tile, error) {
	lastTile, err := m.GetTile(x, y)
	if err != nil {
		return nil, err
	}
	for {
		x += int(dir.X)
		y += int(dir.Y)
		tile, err := m.GetTile(x, y)
		if err != nil {
			return nil, err
		}
		if !tile.IsEmpty() {
			return lastTile, nil
		}
		lastTile = tile
	}
}

func (m *Map) CreateImage() *ebiten.Image {
	img := ebiten.NewImage(m.Width*m.TileW, m.Height*m.TileH)

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			tile, err := m.GetTile(x, y)
			if err != nil {
				continue
			}
			rect := image.Rect(x*m.TileW, y*m.TileH, (x+1)*m.TileW, (y+1)*m.TileH)
			switch tile.Id {
			case data.EmptyTile:
				img.SubImage(rect).(*ebiten.Image).Fill(data.EmptyColor)
			case data.WallTile:
				img.SubImage(rect).(*ebiten.Image).Fill(data.WallColor)
			}
		}
	}
	return img
}
