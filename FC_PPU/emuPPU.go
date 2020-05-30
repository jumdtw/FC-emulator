package main

import (
	"fmt"
	"image"
	"log"
	//"time"
	"./FileOp"
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
	ppuemu *ppu
}

func romdatareturn(g *Game,romaddr uint64)(uint64,uint64){
	var highbit, lowbit uint64 = 0, 0
	for i:=0; i<8; i++{
		lowbit += uint64(g.ppuemu.memory[romaddr+uint64(i)])
		lowbit = lowbit << 8 
	}
	romaddr += 8
	for i:=0; i<8; i++{
		highbit += uint64(g.ppuemu.memory[romaddr+uint64(i)])
		highbit = highbit << 8 
	}
	return highbit, lowbit
}

func patternNumreturn(highbit uint64,lowbit uint64)(uint8,uint64,uint64){
	var patternNum uint64
	//一番最初のビットのみ必要なので取り出す。
	patternNum = highbit & 0x80000000
	//lowbitが入るため0b0~62bit~x0としたい
	patternNum = patternNum >> 62
	// 一番最初のビットはいらないので捨てる。
	highbit = highbit << 1
	if lowbit&0x80000000 >= 0{
		patternNum++
	}
	lowbit = lowbit << 1
	return uint8(patternNum),highbit,lowbit
}

func palletnumreturn(g *Game,numblock int,numtile int)(uint8){
	var counter int
	var highflag bool = false 
	// まずnumblockを利用してメモリからattribute情報を取り出す。
	var attribute uint8 = g.ppuemu.memory[0x27c0 + int(numblock)]
	// numtileの判定。ダルい。
	// 1, まず、16を順に引いていき奇数回なら下側。偶数なら上側(つまり数の少ない方)
	counter = 0
	bufnumtile := numtile
	for bufnumtile > 0{
		bufnumtile -= 16
		counter++
	}
	if counter%2==1{
		highflag = true
	}
	if highflag {
		attribute = attribute >> 4
	}
	// 2, 次にその数自体が偶数なら小さい方。奇数なら大きいほう
	if numtile%2 == 0 {
		attribute = 0x03 & attribute
	}else{
		attribute = attribute >> 2
	}
	return attribute
}

func rgbreturn(g *Game,patternnum uint8,palletnum uint8)(uint8,uint8,uint8){
	var rc, gc, bc uint8
	palletaddrhead := 0x3f00 + int(palletnum * 4)
	colornum := g.ppuemu.memory[palletaddrhead + int(patternnum)]

	switch colornum{
	case 0x01:
		rc = 3
		gc = 35
		bc = 144
	default:
		rc = 0xff
		gc = 0
		bc = 0
	}

	return rc, gc, bc
}

func drawTile(g *Game,tilenum int,numblock int,numtile int){
	var patternNum uint8
	//var palletnum uint8
	// chrnumにはname tableから、bgまたはspriteの番号を取り出す。
	chrnum := g.ppuemu.memory[0x2400 + numtile]
	// 番号からほしいタイルのメモリ番号の頭アドレス.ここから128bitとりだす。
	var romaddr uint64 = uint64(0x1000) + uint64(chrnum*16)


	highbit, lowbit := romdatareturn(g,romaddr)

	for i :=0; i<8; i++ {
		for k :=0; k < 8; k++ {
			// pp : 書き込もうとしているピクセルのrcの場所の配列数宇
			pp := tilenum+k*4*pixlsize+i*256*pixlsize*4*pixlsize
			// N tile目のpattern number
			patternNum, highbit, lowbit = patternNumreturn(highbit,lowbit)
			// patternNumがどのpalletnumであるか。入りえる値は0~3.
			palletnum := palletnumreturn(g,numblock,numtile)
			var rc, gc, bc uint8 = rgbreturn(g,patternNum,palletnum)

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

func numblockreturn(i int, k int)(int){
	var numblock int
	numblock += i*16
	numblock += k/2
	return numblock
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Generate the noise with random RGB values.
	var numblock int
	var numtile int
	for i :=0; i<30; i++ {
		for k :=0; k < 32; k++ {
			// vv は n枚目のタイルの一番左上のピクセルの最初の配列番号
			// k は一増えるごとに8pixl×ピクセル倍率分増える
			// i は一増えるごとに256×ピクセル倍率に8pixl×ピクセル倍率を掛ける
			// 最後に4を掛けることにより配列数を出す。
			vv := (k*8*pixlsize+i*256*pixlsize*8*pixlsize)*4
			// patternNum がどのブロックにあるか
			// tileは32x30の半分なのでブロックは16x15.ナンバリングは0スタート
			numblock = numblockreturn(i,k)
			numtile = k+i*32
			drawTile(g,vv,numblock,numtile)
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

func initppuemu(ppuemu *ppu){
	// カートリッジ0x1を初期化
	//ReadBinary(path string, readsize int,memory []uint8,baseaddr uint32)([]uint8) {
	ppuemu.memory = ReadBinary("C:/Users/ttnmr/HOME/emulator/FC/kotoha.nes",0xfff,ppuemu.memory,0x1000)
	for i:=0 ;i<0xfff; i++{
		ppuemu.memory[0x1000+i] = 0x01
	}
	// name table #1 を初期化
	for i:=0 ;i<0x3bf;i++{
		ppuemu.memory[0x2400+i] = 0x01
	}
	// BG pallet を初期化
	for i:=0 ;i<0xf;i++{
		ppuemu.memory[0x3f00+i] = 0x01
	}
}

func newppuemu()(*ppu){
	ppuemu := ppu{}
	initppuemu(&ppuemu)
	return &ppuemu
}

func main() {

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Noise (Ebiten Demo)")
	g := &Game{
		noiseImage: image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight)),
		ppuemu: newppuemu(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
