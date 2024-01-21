package data

import (
	"errors"
	"image/color"
)

const (
	EmptyTile = iota
	WallTile
	CarTile
	PretendTile
	FinishTile
	MaxTile
)

var (
	EmptyColor color.RGBA = color.RGBA{0, 0, 0, 255}
	WallColor  color.RGBA = color.RGBA{125, 125, 125, 255}
)

type Tile struct {
	Id int
}

func NewTile(id int) (*Tile, error) {
	if id >= MaxTile {
		return nil, errors.New("invalid tile id")
	}
	return &Tile{
		Id: id,
	}, nil
}

func (t Tile) IsNil() bool {
	return t.Id == -1
}

func (t Tile) IsWall() bool {
	return t.Id == WallTile
}

func (t Tile) IsEmpty() bool {
	return t.Id == EmptyTile || t.Id == FinishTile
}
