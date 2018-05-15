package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/color"
	"github.com/bit101/blg/noise"
	"github.com/bit101/blg/random"
)

// May24 generates a gif
func May24() {
	width := 200.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceColor(color.HSV(percent*360.0, 1, 1))
		surface.Save()
		surface.Translate(width/2, height*0.6)
		surface.Rotate(math.Pi / 4.0)
		random.Seed(0)
		for i := 0.0; i < 10000.0; i++ {
			angle := i / 10000.0 * math.Pi * 2.0
			// r := random.FloatRange(0.95, 1.05)
			r := blmath.LerpSin(i/10000.0*100, 0.98, 1.02)
			x := math.Cos(angle) * r
			y := math.Sin(angle) * r
			z := blmath.LerpSin(percent, -1, 1)
			// z := percent * 40
			xx := blmath.Map(noise.Perlin(x, y, z), -1, 1, -200, 200)
			yy := blmath.Map(noise.Perlin(y, x, z), -1, 1, -200, 200)
			surface.FillRectangle(xx, yy, 1, 1)
		}
		surface.Restore()
	})
}
