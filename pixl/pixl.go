package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/swatch"
	"zerotomastery.io/pixl/ui"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
