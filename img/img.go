package img

import (
	"go-perlin-generator/noise"
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

// Img data structure to handle generated images
type Img struct {
	width      int
	height     int
	resolution float64
}

// NewImg constructor for Img struct type
func NewImg(w, h int, r float64) Img {
	return Img{w, h, r}
}

// Generate return a new perlin-noised generated image
func (img *Img) Generate() image.Image {
	ctx := gg.NewContext(img.width, img.height)

	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			nx := float64(x) / img.resolution
			ny := float64(y) / img.resolution
			n := noise.Perlin(nx, ny)
			i := uint8(n * 255.0)

			col := color.RGBA{i, i, i, 255}
			ctx.SetColor(col)
			ctx.SetPixel(x, y)
		}
	}

	return ctx.Image()
}
