package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

type Ball struct {
	X         float32
	Y         float32
	R         float32
	Color     color.Color
	FillColor color.Color
	xInc      float32
	yInc      float32
}

func (b *Ball) Updates() {
	if b.X >= screenWidth-b.R || b.X <= b.R {
		b.xInc = -b.xInc
	}
	if b.Y >= screenHeight-b.R || b.Y <= b.R {
		b.yInc = -b.yInc
	}
	b.X = b.X + b.xInc
	b.Y = b.Y + b.yInc
}

type App struct {
	Balls []*Ball
	// lastSync time.Time
}

func (a *App) Update() error {
	// if time.Since(a.lastSync) < 1000 {
	// 	return nil
	// }
	// a.lastSync = time.Now()

	for i := range a.Balls {
		a.Balls[i].Updates()
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	for i := range a.Balls {
		b := a.Balls[i]
		vector.DrawFilledCircle(screen, float32(b.X), float32(b.Y), float32(b.R-1), b.FillColor, true)
		vector.StrokeCircle(screen, float32(b.X), float32(b.Y), float32(b.R), 3, b.Color, true)
	}
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
