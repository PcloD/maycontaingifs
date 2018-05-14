package may

import (
	"fmt"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/noise"
	"github.com/bit101/blg/random"
)

// May24 generates a gif
func May24() {
	width := 400.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		random.Seed(0)
		for i := 0; i < 20000; i++ {
			x := random.FloatRange(-1, 1)
			y := random.FloatRange(-1, 1)
			z := blmath.LerpSin(percent, -1, 1)
			xx := blmath.Map(noise.Perlin(x, y, z), -1, 1, 0, 400)
			yy := blmath.Map(noise.Perlin(y, x, z), -1, 1, 0, 400)
			surface.FillRectangle(xx, yy, 1, 1)
		}
	})
}
