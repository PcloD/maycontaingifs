package util

import (
	"os/exec"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
)

// Project is a animated gif maker.
type Project struct {
	Width          float64
	Height         float64
	Frames         int
	FPS            int
	Filename       string
	RenderCallback anim.RenderCallback
}

// NewProject creates a new animated gif project.
func NewProject(filename string) *Project {
	return &Project{
		Width:    200,
		Height:   200,
		Frames:   60,
		FPS:      30,
		Filename: filename,
	}
}

// Render renders the gif
func (p *Project) Render(renderCallback anim.RenderCallback) {
	surface := bitlibgo.NewBitSurface(p.Width, p.Height)
	anim.MakeGif(surface, p.Frames, p.FPS, p.Filename, renderCallback)
	cmd := exec.Command("cp", p.Filename, "out/latest.gif")
	cmd.Run()
}

// SetSize sets the size
func (p *Project) SetSize(w float64, h float64) {
	p.Width = w
	p.Height = h
}
