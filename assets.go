package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Assets struct {
	Hamster       *ebiten.Image
	HamsterSource image.Image
	CloseButton   *ebiten.Image
}

func loadAssets() (*Assets, error) {
	var assets Assets
	var err error
	assets.Hamster, assets.HamsterSource, err = ebitenutil.NewImageFromFile("assets/hamster/idle.png")

	if err != nil {
		return nil, fmt.Errorf("загрузка хомяка: %w", err)
	}

	assets.CloseButton, _, err = ebitenutil.NewImageFromFile("assets/ui/close-btn.png")

	if err != nil {
		return nil, fmt.Errorf("загрузка кнопки закрытия: %w", err)
	}

	return &assets, nil
}
