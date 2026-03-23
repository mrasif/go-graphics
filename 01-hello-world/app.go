package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type App struct {
	x, y float64
}

func (g *App) Update() error {
	return nil
}

func (g *App) Draw(screen *ebiten.Image) {
	msg := "Hello World!"
	text.Draw(screen, msg, basicfont.Face7x13, 10, 30, color.White)
}

func (g *App) Layout(_, _ int) (int, int) {
	return 640, 480
}
