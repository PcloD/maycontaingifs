package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
)

// May25 generates a gif
func May25() {
	width := 200.0
	height := 300.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		// surface.SetSourceColor(color.HSV(percent*360.0, 1, 1))
		surface.Save()
		surface.Translate(width/2, height/2)
		surface.Rotate(math.Pi / 4.0)
		random.Seed(1)
		for i := 0.0; i < 5000.0; i++ {
			angle := i / 5000.0 * math.Pi * 2.0
			x := math.Cos(angle)
			y := math.Sin(angle)
			z := blmath.LerpSin(percent, 0, 1)
			// z := percent * 40
			xx := blmath.Map(noise.Perlin(x, y, z), -1, 1, -200, 200)
			yy := blmath.Map(noise.Perlin(y, x, z), -1, 1, -200, 200)
			surface.SetSourceRGB(1, 0, 0)
			surface.FillRectangle(xx, yy, 1, 1)

			xx = blmath.Map(noise.Perlin(x*1.1, y*1.1, z), -1, 1, -200, 200)
			yy = blmath.Map(noise.Perlin(y*1.1, x*1.1, z), -1, 1, -200, 200)
			surface.SetSourceRGB(0, 1, 0)
			surface.FillRectangle(xx, yy, 1, 1)

			xx = blmath.Map(noise.Perlin(x*1.2, y*1.2, z), -1, 1, -200, 200)
			yy = blmath.Map(noise.Perlin(y*1.2, x*1.2, z), -1, 1, -200, 200)
			surface.SetSourceRGB(0, 0, 1)
			surface.FillRectangle(xx, yy, 1, 1)
			surface.FillRectangle(xx, yy, 1, 1)
		}
		surface.Restore()
	})
}
