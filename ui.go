package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"
	"time"

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
var knightImg *ebiten.Image
var chestImg *ebiten.Image
var op = &ebiten.DrawImageOptions{}
var dungeon Map2d
var bob *MapElement
var gold *MapElement

func init() {
	op := &ebiten.DrawImageOptions{}

	img, _, _ := ebitenutil.NewImageFromFile("images/floor_2.png", ebiten.FilterDefault)
	w, h := img.Size()
	ebImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	op.ColorM.Scale(1, 1, 1, 1.0)
	ebImg.DrawImage(img, op)

	img2, _, _ := ebitenutil.NewImageFromFile("images/knight_f_idle_anim_f0.png", ebiten.FilterDefault)
	w, h = img2.Size()
	knightImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	op.ColorM.Scale(1, 1, 1, 1.0)
	knightImg.DrawImage(img2, op)

	img3, _, _ := ebitenutil.NewImageFromFile("images/chest_empty_open_anim_f0.png", ebiten.FilterDefault)
	w, h = img3.Size()
	chestImg, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	op.ColorM.Scale(1, 1, 1, 1.0)
	chestImg.DrawImage(img3, op)

	dungeon = generateDungeon(15, 10, 60, 10)

	bob = getRandomPosition(&dungeon, "b", false)
	gold = getRandomPosition(&dungeon, "g", true)
}

func getRandomPosition(aMap *Map2d, name string, pass bool) *MapElement {
	rand.Seed(time.Now().UnixNano())
	var e *MapElement
	for {
		e = aMap.two_d[rand.Intn(aMap.x-1)][rand.Intn(aMap.y-1)]
		if e.passable {
			e, _ = putElementinMap2d(aMap, name, pass, e.pos_x, e.pos_y)
			break
		}
	}
	return e
}

func update(screen *ebiten.Image) error {

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		dungeon = generateDungeon(15, 10, 60, 10)
		bob = getRandomPosition(&dungeon, "b", false)
		gold = getRandomPosition(&dungeon, "g", true)

		// fmt.Printf(print_map(&dungeon))
	}

	// Draw map
	w, h := ebImg.Size()
	for colNum, row := range dungeon.two_d {
		for elNum, e := range row {
			op.GeoM.Reset()
			if e.name != "#" {
				op.GeoM.Translate(float64(colNum*h), float64(elNum*w))
				screen.DrawImage(ebImg, op)
			}
		}
	}

	if bob != nil {
		msg := fmt.Sprintf("bob %v %v", bob.pos_x, bob.pos_y)
		ebitenutil.DebugPrint(screen, msg)
		op.GeoM.Reset()
		op.GeoM.Translate(float64((bob.pos_x)*w), float64((bob.pos_y-1)*w))
		screen.DrawImage(knightImg, op)
	}
	if gold != nil {
		msg := fmt.Sprintf("         | gold %v %v", gold.pos_x, gold.pos_y)
		ebitenutil.DebugPrint(screen, msg)
		op.GeoM.Reset()
		op.GeoM.Translate(float64((gold.pos_x)*w), float64((gold.pos_y)*w))
		screen.DrawImage(chestImg, op)
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
