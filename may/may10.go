package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
)

// May10 generates a gif
func May10() {
	width := 400.0
	height := 400.0
	filename := "out/may10.gif"

	type Point3D struct {
		x float64
		y float64
		z float64
	}

	radius := 100.0
	var points []*Point3D
	for y := -height * 0.8; y < height*0.8; y += 20.0 {
		angle := y / height * 6.0

		p0 := &Point3D{
			math.Cos(angle) * radius,
			y,
			math.Sin(angle) * radius,
		}
		points = append(points, p0)
		p1 := &Point3D{
			math.Cos(angle) * -radius,
			y,
			math.Sin(angle) * -radius,
		}
		points = append(points, p1)
	}

	fl := 300.0

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		angle := percent * math.Pi * 2

		for _, p := range points {
			xx := math.Cos(angle)*p.x - math.Sin(angle)*p.z
			zz := math.Cos(angle)*p.z + math.Sin(angle)*p.x
			scale := fl / (fl + zz + 200)
			x := scale * xx
			y := scale * p.y
			r := scale * 20

			surface.FillCircle(x, y, r)
		}
		surface.Restore()
	})
}
