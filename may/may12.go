package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
)

// May12 generates a gif
func May12() {
	width := 400.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.Save()
		angle := percent * math.Pi * 4
		surface.Translate(width/2+math.Sin(angle)*100.0, height/2+math.Cos(angle)*100.0)
		even := true
		res := 10.0
		for i := width; i >= res; i -= res {
			r := (width - i) * 0.008
			surface.Save()
			surface.Rotate(blmath.SinRange(percent*math.Pi*2, -r, r))
			if even {
				surface.SetSourceRGB(1, 1, 1)
			} else {
				surface.SetSourceRGB(0, 0, 0)
			}
			surface.FillRectangle(-i, -i, i*2, i*2)
			even = !even
			surface.Restore()
		}
		surface.Restore()
	})

}
