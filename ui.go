package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
	maxAngle     = 256
)

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
}

var ebImg *ebiten.Image
var op = &ebiten.DrawImageOptions{}

func init() {
	img, _, _ := ebitenutil.NewImageFromFile("images/floor_2.png", ebiten.FilterDefault)

	w, h := img.Size()
	ebImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	ebImg.DrawImage(img, op)

}

func update(screen *ebiten.Image) error {
	w, _ := ebImg.Size()
	offset := 0
	for i := 0; i < 20; i++ {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(offset), float64(0))
		offset += w
		screen.DrawImage(ebImg, op)
		//fmt.Println(w)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "FRyan Demo"); err != nil {
		log.Fatal(err)
	}
}
