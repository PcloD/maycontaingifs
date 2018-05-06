package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
)

// May14 generates a gif
func May14() {
	width := 400.0
	height := 400.0
	numRows := 10.0
	numCols := 10.0
	rowHeight := height / numRows
	colWidth := width / numCols

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		// rowEven := true
		for r := 0.0; r < numRows; r++ {
			even := true
			surface.SetSourceRGB(0, 0, 0)
			surface.MoveTo(0, rowHeight*r)
			surface.LineTo(width, rowHeight*r)
			surface.Stroke()
			for c := -1.0; c < numCols+1.0; c++ {
				surface.Save()
				surface.Translate(math.Sin(r/numRows*math.Pi*4)*30*blmath.SinRange(percent*math.Pi*2, -1, 1), rowHeight*r)
				if even {
					surface.SetSourceRGB(1, 1, 1)
				} else {
					surface.SetSourceRGB(0, 0, 0)
				}
				surface.FillRectangle(c*colWidth, 0, colWidth, rowHeight)
				even = !even
				surface.Restore()
			}
			// rowEven = !rowEven
		}

	})
}
