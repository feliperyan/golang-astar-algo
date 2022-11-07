package main

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	mapWidth  = 40
	mapHeight = 25

	tunnels               = 130
	tunnelLength          = 15
	mapProportionModifier = 1
	screenFactor          = 2
)

var (
	tileSize int
	op       = &ebiten.DrawImageOptions{}
	dungeon  Map2d
	bob      *MapElement
	gold     *MapElement

	floorSprite  *Sprite
	knightSprite *Sprite
	chestSprite  *Sprite
	coinSprite   *Sprite
)

type Sprite struct {
	imageWidth  int
	imageHeight int
	imagePath   string
	image       *ebiten.Image
}

type Game struct{}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		dungeon = generateDungeon(mapWidth, mapHeight, tunnels, tunnelLength)
		bob = getRandomPosition(&dungeon, "b", false)
		gold = getRandomPosition(&dungeon, "g", true)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		bob.path = aStarAlgorithm(&dungeon, bob, gold)
		fmt.Println(len(bob.path))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw map
	w, h := tileSize, tileSize
	for colNum, row := range dungeon.two_d {
		for elNum, e := range row {
			op.GeoM.Reset()
			if e.name != "#" {
				op.GeoM.Translate(float64(colNum*h), float64(elNum*w))
				screen.DrawImage(floorSprite.image, op)
			}
		}
	}

	if bob != nil {
		msg := fmt.Sprintf("bob %v %v", bob.pos_x, bob.pos_y)
		ebitenutil.DebugPrint(screen, msg)
		op.GeoM.Reset()
		op.GeoM.Translate(float64((bob.pos_x)*tileSize), float64((bob.pos_y-1)*tileSize))
		screen.DrawImage(knightSprite.image, op)
	}

	if gold != nil {
		msg := fmt.Sprintf("         | gold %v %v", gold.pos_x, gold.pos_y)
		ebitenutil.DebugPrint(screen, msg)
		op.GeoM.Reset()
		op.GeoM.Translate(float64((gold.pos_x)*tileSize), float64((gold.pos_y)*tileSize))
		screen.DrawImage(chestSprite.image, op)
	}

	if bob.path != nil {
		for _, p := range bob.path {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(p.pos_x*tileSize)+2, float64(p.pos_y*tileSize)+2)
			screen.DrawImage(coinSprite.image, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (s *Sprite) initSpriteImage() {

	op := &ebiten.DrawImageOptions{}

	img, _, _ := ebitenutil.NewImageFromFile(s.imagePath)
	s.imageWidth, s.imageHeight = img.Size()
	s.imageWidth = int(float64(s.imageWidth) * mapProportionModifier)
	s.imageHeight = int(float64(s.imageHeight) * mapProportionModifier)

	s.image = ebiten.NewImage(s.imageWidth, s.imageHeight)

	op.ColorM.Scale(1, 1, 1, 1.0)
	s.image.DrawImage(img, op)
}

func createSprite(imagePath string) *Sprite {
	s := Sprite{0, 0, imagePath, nil}
	s.initSpriteImage()
	return &s
}

func init() {

	tileSize = 16
	tileSize = int(float64(tileSize) * mapProportionModifier)

	// Init the sprite images
	floorSprite = createSprite("images/floor_2.png")
	knightSprite = createSprite("images/knight_f_idle_anim_f0.png")
	chestSprite = createSprite("images/chest_empty_open_anim_f0.png")
	coinSprite = createSprite("images/coin_anim_f0.png")

	// Generate a random Dungeon.
	dungeon = generateDungeon(mapWidth, mapHeight, tunnels, tunnelLength)

	// Place bob the knight and a chest of gold at a random floor tile
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
