package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/bitmath"
	"github.com/bit101/bitlibgo/random"
)

// May08 generates a gif
func May08() {
	width := 400.0
	height := 400.0
	yres := 5.0
	xres := 5.0
	filename := "out/may08.gif"

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.5)
		for y := 0.0; y < height+yres; y += yres {
			surface.MoveTo(0, y)
			for x := xres; x < width+xres; x += xres {
				dx := x - width/2
				dy := y - height/2
				dist := math.Hypot(dx, dy)
				mult := bitmath.SinRange(percent*math.Pi*2, -0.15, 0.15)
				h := bitmath.Clamp(-50, 50, math.Sin(dist*mult)*1000/dist)
				surface.LineTo(x, y+h+random.FloatRange(-2, 2))
			}
		}
		surface.Stroke()
	})
}
