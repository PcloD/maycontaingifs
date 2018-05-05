package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/bitmath"
)

// May13 generates a gif
func May13() {
	width := 400.0
	height := 400.0
	filename := "out/may13.gif"

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.Save()

		angle := percent * math.Pi * 4
		surface.Translate(width/2+math.Sin(angle)*100.0, height/2+math.Cos(angle)*100.0)
		even := true
		res := 5.0

		for i := width * 2; i >= res; i -= res {
			r := (width - i) * 0.012
			surface.Save()
			surface.Rotate(bitmath.LerpSin(percent, -r, r))
			if even {
				surface.SetSourceRGB(bitmath.LerpSin(percent, 1, 0), 0.5, bitmath.LerpSin(percent, 0, 1))
			} else {
				surface.SetSourceRGB(0, 0, 0)
			}

			surface.FillEllipse(-i, -i, i*2.5, i*1.5)
			even = !even
			surface.Restore()
		}
		surface.Restore()
	})
}
