package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	w, h := ebiten.Monitor().Size()
	h = h - 35
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Bounce Ball")

	balls := []*Ball{
		&Ball{
			X:         41,
			Y:         41,
			R:         40,
			Color:     color.RGBA{255, 0, 0, 255},
			FillColor: color.RGBA{255, 0, 0, 10},
			xInc:      3,
			yInc:      3,
		},
		&Ball{
			X:         201,
			Y:         101,
			R:         55,
			Color:     color.RGBA{255, 255, 0, 255},
			FillColor: color.RGBA{255, 255, 0, 10},
			xInc:      2,
			yInc:      2,
		},
		&Ball{
			X:         71,
			Y:         201,
			R:         40,
			Color:     color.RGBA{255, 0, 255, 255},
			FillColor: color.RGBA{255, 0, 255, 10},
			xInc:      3,
			yInc:      3,
		},
		&Ball{
			X:         400,
			Y:         301,
			R:         35,
			Color:     color.RGBA{255, 255, 255, 255},
			FillColor: color.RGBA{255, 255, 255, 10},
			xInc:      3,
			yInc:      -3,
		},
		&Ball{
			X:         301,
			Y:         401,
			R:         45,
			Color:     color.RGBA{0, 255, 255, 255},
			FillColor: color.RGBA{0, 255, 255, 10},
			xInc:      -2,
			yInc:      2,
		},
		&Ball{
			X:         671,
			Y:         201,
			R:         40,
			Color:     color.RGBA{0, 255, 255, 255},
			FillColor: color.RGBA{0, 255, 255, 10},
			xInc:      -3,
			yInc:      -3,
		},
		&Ball{
			X:         671,
			Y:         401,
			R:         45,
			Color:     color.RGBA{255, 255, 0, 255},
			FillColor: color.RGBA{255, 255, 0, 10},
			xInc:      2,
			yInc:      -2,
		},
		&Ball{
			X:         1000,
			Y:         600,
			R:         55,
			Color:     color.RGBA{255, 0, 0, 255},
			FillColor: color.RGBA{255, 0, 0, 10},
			xInc:      2,
			yInc:      2,
		},
	}

	app := &App{
		Balls:        balls,
		ScreenWidth:  w,
		ScreenHeight: h,
	}
	ebiten.SetTPS(30)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	if err := ebiten.RunGameWithOptions(app, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}

}
