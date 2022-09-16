package ui

import (
	"fyne.io/fyne/v2"

	"zerotomastery/pixl/swatch"

	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/swatch"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
