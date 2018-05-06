package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
)

// May17 generates a gif
func May17() {
	width := 500.0
	height := 300.0
	xres := 2.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.SetLineWidth(0.5)

		for x := -100.0; x < width+100.0; x += xres {
			w := blmath.SinRange(x*0.03, 20, 100)
			h := 120 - w
			surface.Save()
			surface.Translate(x, height/2+blmath.SinRange(x*0.04, -50, 50))
			// surface.Rotate(blmath.LerpSin(percent, -2, 2))
			surface.Rotate(percent * math.Pi * 2)
			surface.Ellipse(0, 0, w, h)
			surface.SetSourceRGB(1, 1, 1)
			surface.FillPreserve()
			surface.SetSourceRGB(0, 0, 0)
			surface.Stroke()
			surface.Restore()
		}
	})

}
