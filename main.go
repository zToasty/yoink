package main

import (
	"log"
	"time"

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

	memText, err := getMemoryUsage()

	if err != nil {
		log.Fatal(err)
	}

	timestmp := time.Now()

	game := &Game{
		assets:           assets,
		memoryText:       memText,
		lastMemoryUpdate: timestmp,
	}

	options := &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}

	if err := ebiten.RunGameWithOptions(game, options); err != nil {
		log.Fatal(err)
	}
}
