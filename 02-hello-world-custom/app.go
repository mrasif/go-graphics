package main

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mrasif/go-graphics/02-hello-world-custom/fonts"
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
	academyEngravedLETFonts, _ := text.NewGoTextFaceSource(
		bytes.NewReader(fonts.AtmaRegularTTF),
	)
	face := &text.GoTextFace{
		Source: academyEngravedLETFonts,
		Size:   48,
	}
	op := &text.DrawOptions{}
	t := "আমার পছন্দের বাংলা!"
	w, h := text.Measure(t, face, 48)
	x := (screenWidth - int(w)) / 2
	y := (screenHeight - int(h)) / 2
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.Scale(0, 200, 0, 255)
	text.Draw(screen, t, face, op)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}
