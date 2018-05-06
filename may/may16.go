package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
)

// May16 generates a gif
func May16() {
	width := 400.0
	height := 400.0
	filename := "out/may16.gif"

	var circ func(*bitlibgo.BitSurface, float64, float64, float64, float64, float64, int)
	circ = func(surface *bitlibgo.BitSurface, rotation, x, y, r, lw float64, d int) {
		surface.Save()
		surface.Translate(x, y)
		surface.Rotate(rotation)
		surface.SetLineWidth(lw)
		surface.StrokeCircle(0, 0, r)
		if d > 0 {
			radius := r * 1.5
			for i := 0.0; i < 4.0; i += 1.0 {
				angle := math.Pi * 2.0 / 4.0 * i
				circ(surface, rotation, math.Cos(angle)*radius, math.Sin(angle)*radius, r*0.5, lw*0.5, d-1)
			}
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)

		circ(surface, percent*math.Pi*2.0/4.0, 0, 0, 80, 8, 6)

		surface.Restore()
	})

}
