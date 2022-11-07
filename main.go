package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game := &Game{}

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("AStart Algorithm")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
