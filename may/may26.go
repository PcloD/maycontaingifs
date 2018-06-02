package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
)

// May26 generates a gif
func May26() {
	width := 400.0
	height := 300.0

	f := func(x, num, percent float64) float64 {
		total := 4.0 * math.Sin(x) / math.Pi
		for i := 3.0; i <= num; i += 2.0 {
			total += 4.0 * percent * math.Sin(i*x) / (i * math.Pi)
		}
		return total * 0.5
	}

	fourier := func(surface *blgo.Surface, num, percent float64) {
		for i := 0.0; i < surface.Width; i++ {
			x := i / surface.Width * math.Pi * 5.0
			y := f(x, num, percent) * -150.0
			// surface.FillRectangle(i, y, 1, 1)
			surface.LineTo(i, y)
		}
		surface.Stroke()
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		// surface.SetSourceColor(color.HSV(percent*360.0, 1, 1))
		surface.Save()
		surface.Translate(0, height/2)

		fourier(surface, 21, blmath.LerpSin(percent, 0, 1))
		surface.Restore()
	})
}
