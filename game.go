package main

import (
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	assets           *Assets   // Ассеты (очевидно же)
	dragging         bool      // флаг смещения в текущий момент
	dragOffsetX      int       // смещение мыши внутри спрайта при начале перетаскивания
	dragOffsetY      int       //
	drawn            bool      // статус отрисовки
	menuOpen         bool      // статус меню
	quitRequested    bool      // нажали чтобы выйти
	memoryText       string    // данные о памяти
	lastMemoryUpdate time.Time // последний timestamp получения памяти
}

const hamsterScale = 0.25

var closeButtonRect = image.Rect(340, 88, 372, 120)
var redSquare = image.Rect(230, 100, 305, 120)

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()

	g.updateDragging(mx, my)
	g.updateMenu(mx, my)

	if time.Since(g.lastMemoryUpdate) >= time.Second {
		g.lastMemoryUpdate = time.Now()

		memoryText, err := getMemoryUsage()
		if err != nil {
			log.Printf("обновление метрик памяти: %v", err)
		} else {
			g.memoryText = memoryText
			g.drawn = false
		}
	}

	if g.quitRequested {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.drawn {
		return
	}
	screen.Clear()

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(hamsterScale, hamsterScale)

	screen.DrawImage(g.assets.Hamster, options)

	if g.menuOpen {
		g.drawMenu(screen)
	}

	//пишем текст
	ebitenutil.DebugPrintAt(screen,
		g.memoryText, 10, 270)

	g.drawn = true

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 300
}

func (g *Game) isHamsterAt(mx, my int) bool {
	px := int(float64(mx) / hamsterScale)
	py := int(float64(my) / hamsterScale)
	w, h := g.assets.Hamster.Bounds().Dx(), g.assets.Hamster.Bounds().Dy()

	if 0 <= px && px < w && 0 <= py && py < h {
		_, _, _, alpha := g.assets.HamsterSource.At(px, py).RGBA()
		return alpha != 0
	}
	return false
}

func (g *Game) updateDragging(mx, my int) {
	// перетаскивание при Зажатии ЛКМ на хомяке
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		!g.menuOpen && !g.dragging && g.isHamsterAt(mx, my) {
		g.dragging = true
		g.dragOffsetX = mx
		g.dragOffsetY = my

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
}

func (g *Game) updateMenu(mx, my int) {
	// вызов меню по нажатию ПКМ на хомяка
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) &&
		g.isHamsterAt(mx, my) {
		g.menuOpen = !g.menuOpen
		g.drawn = false
		g.dragging = false
	}

	// кнопка закрытия
	if g.menuOpen &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		image.Pt(mx, my).In(closeButtonRect) {

		g.menuOpen = false
		g.drawn = false
	}

	// кнопка закрыть хомяка
	if g.menuOpen &&
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) &&
		image.Pt(mx, my).In(redSquare) {
		g.quitRequested = true
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	vector.FillRect(screen, 220, 80, 160, 100,
		color.RGBA{R: 45, G: 40, B: 50, A: 255}, false)

	options := ebiten.DrawImageOptions{}
	options.GeoM.Scale(2, 2)
	options.GeoM.Translate(340, 88)
	screen.DrawImage(g.assets.CloseButton, &options)

	vector.FillRect(screen,
		float32(redSquare.Min.X), float32(redSquare.Min.Y),
		float32(redSquare.Dx()), float32(redSquare.Dy()),
		color.RGBA{R: 255, A: 255}, false)
}
