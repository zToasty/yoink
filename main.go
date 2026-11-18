package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(400, 300)
	ebiten.SetWindowTitle("Yoink!")

	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetScreenClearedEveryFrame(false)

	// Загрузка ассетов
	assets, err := loadAssets()

	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		assets: assets,
	}

	options := &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}

	if err := ebiten.RunGameWithOptions(game, options); err != nil {
		log.Fatal(err)
	}
}
