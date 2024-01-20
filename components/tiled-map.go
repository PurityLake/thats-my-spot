package components

import (
	"errors"
	_ "image/png"
	"log"

	"github.com/PurityLake/thatsmyspot/data"
	"github.com/PurityLake/thatsmyspot/mapreader"
	"github.com/PurityLake/thatsmyspot/maths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TiledMap struct {
	Width, Height int
	TileW, TileH  int
	Tiles         []data.Tile
	TempImage     *ebiten.Image
}

func NewTiledMap(imgFilename, mapFilename, tilesetFilename string) (*TiledMap, map[string]data.Property, error) {
	tiledMap, _, err := ebitenutil.NewImageFromFile(imgFilename)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	mapObj, err := mapreader.ReadJson(mapFilename)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	tilesetObj, err := mapreader.ReadJson(tilesetFilename)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	mapData := mapreader.ParseMapData(mapObj)
	tilesetData := mapreader.ParseTilesetData(tilesetObj)

	propertyList := tilesetData["properties"].Value.([]data.Property)
	tiles := make([][]data.Property, 0)
	for _, prop := range propertyList {
		props := prop.Value.(map[string]data.Property)
		tileProps := make([]data.Property, 0)
		for _, p := range props {
			tileProps = append(tileProps, p)
		}
		tiles = append(tiles, tileProps)
	}
	var tilesTypes []data.Tile
	for _, layer := range mapData["layers"].Value.([]data.Property) {
		for _, tileIndex := range layer.Value.([]int) {
			if tileIndex == 0 {
				t, _ := data.NewTile(-1)
				tilesTypes = append(tilesTypes, *t)
				continue
			}
			tile := tiles[tileIndex-1]
			t, _ := data.NewTile(int(tile[0].Value.(float64)))
			tilesTypes = append(tilesTypes, *t)
		}
	}
	bounds := tiledMap.Bounds()
	return &TiledMap{
		Width:     bounds.Dx(),
		Height:    bounds.Dy(),
		TileW:     40,
		TileH:     40,
		TempImage: tiledMap,
		Tiles:     tilesTypes,
	}, mapData, nil
}

func (m *TiledMap) GetTile(x, y int) (*data.Tile, error) {
	width := m.Width / m.TileW
	height := m.Height / m.TileH
	if x < 0 || x >= width || y < 0 || y >= height {
		return nil, errors.New("out of bounds")
	}
	return &m.Tiles[x+y*width], nil
}

func (m TiledMap) CanGo(x, y int) bool {
	tile, err := m.GetTile(x, y)
	if err != nil {
		return false
	}
	return tile.IsEmpty()
}

func (m TiledMap) GetLastTileInDir(x, y int, dir maths.Vector2) (int, int, error) {
	lastTileX, lastTileY := x, y
	for {
		x += int(dir.X)
		y += int(dir.Y)

		if !m.CanGo(x, y) {
			return lastTileX, lastTileY, nil
		}

		lastTileX, lastTileY = x, y
	}
}

func (m TiledMap) MapPosFromScreenPos(x, y int) (int, int) {
	return (x - 20) / m.TileW, (y - 20) / m.TileH
}

func (m TiledMap) ScreenPosFromMapPos(x, y int) (float32, float32) {
	return float32(20 + (x * m.TileW)), float32(20 + (y * m.TileH))
}
