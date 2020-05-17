package main

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var rc, gc, bc uint8 = 1,0,0

const (
	screenWidth  = 320
	screenHeight = 240
)

type rand struct {
	x, y, z, w uint32
}

func (r *rand) next() uint32 {
	// math/rand is too slow to keep 60 FPS on web browsers.
	// Use Xorshift instead: http://en.wikipedia.org/wiki/Xorshift
	t := r.x ^ (r.x << 11)
	r.x, r.y, r.z = r.y, r.z, r.w
	r.w = (r.w ^ (r.w >> 19)) ^ (t ^ (t >> 8))
	return r.w
}

var theRand = &rand{12345678, 4185243, 776511, 45411}

type Game struct {
	noiseImage *image.RGBA
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Generate the noise with random RGB values.
	const l = screenWidth * screenHeight
	for i := 0; i < l; i++ {
		//x := theRand.next()
		g.noiseImage.Pix[4*i] = rc
		g.noiseImage.Pix[4*i+1] = gc
		g.noiseImage.Pix[4*i+2] = bc
		g.noiseImage.Pix[4*i+3] = 0xff

		if rc==0xff{
			rc = 0
			gc = 0xff
		}else if gc ==0xff{
			gc = 0
			bc = 0xff
		}else{
			bc =0 
			rc = 0xff
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.noiseImage.Pix)
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebiten Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
