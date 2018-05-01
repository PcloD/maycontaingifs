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

// May05 generates a gif
func May05() {
	width := 400.0
	height := 400.0
	center := geom.NewPoint(width/2, height/2)
	filename := "out/may05.gif"
	numPoints := 500
	maxDist := 50.0
	var points []*geom.Point
	for i := 0; i < numPoints; i++ {
		angle := random.FloatRange(0, math.Pi * 2)
		radius := random.FloatRange(0, width)
		points = append(points, geom.NewPoint(width / 2 + math.Cos(angle) * radius, height / 2 + math.Sin(angle) * radius))
	}

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		var lpoints []*geom.Point
		for _, p := range points {
			lp := geom.LerpPoint(bitmath.SinRange(percent * math.Pi * 2, 0, 1), center, p)
			lpoints = append(lpoints, &lp)
		}
		surface.ClearRGB(1, 1, 1)
		for i, p0 := range lpoints {
			for _, p1 := range lpoints[i:] {
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
