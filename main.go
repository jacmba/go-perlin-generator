package main

import (
	"go-perlin-generator/img"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

const (
	defaultWidth  = 512
	defaultHeight = 512
)

func main() {
	mApp := app.New()
	w := mApp.NewWindow("Perlin noise generator")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(defaultWidth, defaultHeight))

	mImg := img.NewImg(defaultWidth, defaultHeight)
	i := mImg.Generate()

	raster := canvas.NewImageFromImage(i)
	w.SetContent(raster)

	w.ShowAndRun()
}
