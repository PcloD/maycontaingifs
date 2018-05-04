package may

import (
	"fmt"
	"math"
	"os/exec"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/bitmath"
	"github.com/bit101/bitlibgo/geom"
	"github.com/bit101/bitlibgo/random"
)

// May04 generates a gif
func May04() {
	width := 380.0
	height := 380.0
	center := geom.NewPoint(width/2, height/2)
	filename := "out/may04.gif"
	numPoints := 400
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
		for i, p0 := range points {
			dActual := p0.Distance(center)
			dIdeal := bitmath.SinRange(percent*math.Pi*2, 0, width)
			dx := math.Abs(dActual - dIdeal)
			maxDist := (1.0 - dx/width) * 100.0
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
	cmd := exec.Command("cp", filename, "out/latest.gif")
	cmd.Run()
}
