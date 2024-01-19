package data

import (
	"errors"
	"image/color"
)

const (
	EmptyTile = 0
	WallTile  = 1
	CarTile   = 2
	MaxTile   = 3
)

var (
	EmptyColor color.RGBA = color.RGBA{0, 0, 0, 255}
	WallColor  color.RGBA = color.RGBA{125, 125, 125, 255}
)

type Tile struct {
	Id int
}

func NewTile(id int) (*Tile, error) {
	if id < 0 || id >= MaxTile {
		return nil, errors.New("invalid tile id")
	}
	return &Tile{
		Id: id,
	}, nil
}

func (t Tile) IsEmpty() bool {
	return t.Id == EmptyTile
}
