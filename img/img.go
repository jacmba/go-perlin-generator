package img

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

// Img data structure to handle generated images
type Img struct {
	width  int
	height int
}

// NewImg constructor for Img struct type
func NewImg(w, h int) Img {
	return Img{w, h}
}

// Generate return a new perlin-noised generated image
func (img *Img) Generate() image.Image {
	ctx := gg.NewContext(img.width, img.height)

	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			c := uint8((x * y) % 255)
			col := color.RGBA{c, c, c, 255}
			ctx.SetColor(col)
			ctx.SetPixel(x, y)
		}
	}

	return ctx.Image()
}
