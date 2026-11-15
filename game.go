package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	img         *ebiten.Image // спрайт
	dragging    bool          // флаг смещения в текущий момент
	dragOffsetX int           // смещение мыши внутри спрайта при начале перетаскивания
	dragOffsetY int
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !g.dragging {
		// check if cursor is inside the sprite bounds
		w, h := g.img.Size()
		if 0 <= mx && mx < w && 0 <= my && my < h {
			g.dragging = true
			g.dragOffsetX = mx
			g.dragOffsetY = my
		}
	}

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.dragging = false
	}

	if g.dragging {
		wx, wy := ebiten.WindowPosition()

		newX := wx + mx - g.dragOffsetX
		newY := wy + my - g.dragOffsetY

		ebiten.SetWindowPosition(newX, newY)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.25, 0.25)

	screen.DrawImage(g.img, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 300
}
