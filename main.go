package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var sourceImage image.Image

func init() {
	var err error
	img, sourceImage, err = ebitenutil.NewImageFromFile("assets/hamster/idle.png")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(400, 300)
	ebiten.SetWindowTitle("Yoink!")

	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetScreenClearedEveryFrame(false)

	game := &Game{
		img:         img,
		sourceImage: sourceImage,
	}
	options := &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}

	if err := ebiten.RunGameWithOptions(game, options); err != nil {
		log.Fatal(err)
	}
}
