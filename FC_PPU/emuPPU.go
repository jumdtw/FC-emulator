package main

import (
	"fmt"
	"image"
	"log"
	//"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//var rc, gc, bc uint8 = 1,0,0

const (
	pixlsize = 1
	screenWidth  = 256 * pixlsize
	screenHeight = 240 * pixlsize
	memCap = 65535	
	
)

ppuemu = ppu{}
	

type rand struct {
	x, y, z, w uint32
}

type ppu struct {
	// memory
	// 0x0000 ~ 0x0fff : bg rom #0　　　 4096
	// 0x1000 ~ 0x1fff : sprite rom  #1 4096
	// 0x2000 ~ 0x23bf : name table #0  960 = 32 * 30
	// 0x23c0 ~ 0x23ff : attribute #0	64 = 0xff - 0xc0
	// 0x2400 ~ 0x27bf : name table #1  960 = 32 * 30	.
	// 0x27c0 ~ 0x27ff : attribute #1  	64 = 0xff - 0xc0
	//          .
	//	ミラーやら拡張機能やら
	//			.
	// 0x3f00 ~ 0x3f0f : BG palette  4 color * 4 palette 
	// 0x3f10 ~ 0x3f1f : sprite palette 4 color * 4 palette
	memory [memCap]uint8
}

type Game struct {
	noiseImage *image.RGBA
}


func rgbreturn(patternnum uint8,palletnum uint8)(uint8,uint8,uint8){
	var rc, gc, bc uint8 

	return rc, gc, bc
}

func romdatareturn(romaddr uint64)(uint64,uint64){
	var highbit, lowbit uint64 = 0, 0
	for i:=0; i<8; i++{
		lowbit += ppuemu.memory[romaddr+i]
		lowbit = lowbit << 8 
	}
	romaddr += 8
	for i:=0; i<8; i++{
		highbit += ppuemu.memory[romaddr+i]
		highbit = highbit << 8 
	}
	return highbit, lowbit
}

func drawTile(g *Game,tilenum int,numblock uint8,numtile uint8){
	// chrnumにはname tableから、bgまたはspriteの番号を取り出す。
	chrnum = ppu.memory[0x2400 + numtile]
	// 番号からほしいタイルのメモリ番号の頭アドレス.ここから128bitとりだす。
	var romaddr uint64 = 0x1000 + chrnum*16

	highbit, lowbit = romdatareturn(romaddr)

	for i :=0; i<8; i++ {
		for k :=0; k < 8; k++ {
			// pp : 書き込もうとしているピクセルのrcの場所の配列数宇
			pp := tilenum+k*4*pixlsize+i*256*pixlsize*4*pixlsize
			// N tile目のpattern number
			var patternNum uint8 = 0
			// patternNumがどのpalletnumであるか。入りえる値は0~3.
			var palletnum uint8 = 0
			var rc, gc, bc := rgbreturn(patternNum,palletnum)

			g.noiseImage.Pix[pp] = rc
			g.noiseImage.Pix[pp+1] = gc
			g.noiseImage.Pix[pp+2] = bc
			g.noiseImage.Pix[pp+3] = 0xff

			/*
			// ピクセルの大きさを上がるにはこの処理が必要。下記ソースはピクセルを2x2の大きさで一つのドットにするときに必要になる。
			g.noiseImage.Pix[pp+4] = rc
			g.noiseImage.Pix[pp+4+1] = gc
			g.noiseImage.Pix[pp+4+2] = bc
			g.noiseImage.Pix[pp+4+3] = 0xff
			
			g.noiseImage.Pix[pp+256*pixlsize*4] = rc
			g.noiseImage.Pix[pp+256*pixlsize*4+1] = gc
			g.noiseImage.Pix[pp+256*pixlsize*4+2] = bc
			g.noiseImage.Pix[pp+256*pixlsize*4+3] = 0xff

			g.noiseImage.Pix[pp+256*pixlsize*4+4] = rc
			g.noiseImage.Pix[pp+256*pixlsize*4+4+1] = gc
			g.noiseImage.Pix[pp+256*pixlsize*4+4+2] = bc
			g.noiseImage.Pix[pp+256*pixlsize*4+4+3] = 0xff
			*/
			
		}
	}
	
	
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Generate the noise with random RGB values.

	for i :=0; i<30; i++ {
		for k :=0; k < 32; k++ {
			// vv は n枚目のタイルの一番左上のピクセルの最初の配列番号
			// k は一増えるごとに8pixl×ピクセル倍率分増える
			// i は一増えるごとに256×ピクセル倍率に8pixl×ピクセル倍率を掛ける
			// 最後に4を掛けることにより配列数を出す。
			vv := (k*8*pixlsize+i*256*pixlsize*8*pixlsize)*4
			// patternNum がどのブロックにあるか
			// tileは32x30の半分なのでブロックは16x15.ナンバリングは0スタート
			var numblock uint8 = ((k/2)-k%2)+i*16
			var numtile uint8 = k+i*16 
			drawTile(g,vv,numblock,numtile)
			if rc==0xff{
				rc = 0
				gc = 0xff
			}else if gc == 0xff{
				gc = 0
				bc = 0xff
			}else{
				bc =0 
				rc = 0xff
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.noiseImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (ppuemu *ppu) initppumem(){
	// カートリッジ0x1を初期化
	for i:=0 ;i<0xfff; i++{
		ppu.memory[0x1000+i] = 0x01
	}
	// name table #1 を初期化
	for i:=0 ;i<0x3bf;i++{
		ppu.memory[0x2400+i] = 0x01
	}
	// BG pallet を初期化
	for i:=0 ;i<0xf;i++{
		ppu.memory[0x3f00+i] = 0x01
	}
}

func main() {

	ppuemu.initppumem()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebiten Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
