package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type App struct {
}

func (a *App) Update() error {

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	w := 40.0
	h := 40.0
	x := (screenHeight - w) / 2
	y := (screenWidth - h) / 2

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), color.White, false)
	vector.DrawFilledCircle(screen, 100, 100, 50, color.RGBA{255, 255, 0, 255}, true)
	vector.StrokeCircle(screen, 200, 200, 10, 3, color.RGBA{255, 0, 0, 255}, true)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
