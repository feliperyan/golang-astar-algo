package main

import (
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	maxAngle     = 256
)

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
}

func init() {
	ebImg, img, _ := ebitenutil.NewImageFromFile("images/floor_2.png")
}
