package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/bitmath"
	"github.com/bit101/bitlibgo/geom"
	"github.com/bit101/bitlibgo/random"
)

// May02 generates a gif
func May02() {
	width := 400.0
	height := 400.0
	filename := "out/may02.gif"
	numPoints := 500
	var points []*geom.Point
	for i := 0; i < numPoints; i++ {
		points = append(points, geom.NewPoint(random.FloatRange(0, width), random.FloatRange(0, height)))
	}

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 120
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		maxDist := bitmath.SinRange(percent*math.Pi*2, 10, 125)
		for i, p0 := range points {
			for _, p1 := range points[i:] {
				dist := p0.Distance(p1)
				if dist < maxDist {
					surface.SetLineWidth((1 - dist/maxDist) * 0.5)
					surface.MoveTo(p0.X, p0.Y)
					surface.LineTo(p1.X, p1.Y)
					surface.Stroke()
				}
			}
		}
	})
}
