package may

import (
	"fmt"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
)

// May21 generates a gif
func May21() {
	width := 400.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)

		scale := 0.01
		res := 1.0

		for x := 0.0; x < width; x += res {
			for y := 0.0; y < width; y += res {
				yy := y + blmath.LerpSin(percent+0.25, 0, 200)
				v := blmath.Map(noise.Perlin(x*scale, yy*scale, blmath.LerpSin(percent, 0, 200)*scale), -1, 1, 0, 1)
				surface.SetSourceRGB(v, v, v)
				surface.FillRectangle(x, y, res, res)
			}
		}

	})
}
