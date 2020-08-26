package main

import (

	"fmt"
	"image"
	"log"
	//"reflect"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	//"github.com/hajimehoshi/ebiten/examples/keyboard/keyboard"
	//rkeyabord "github.com/hajimehoshi/ebiten/examples/resources/images/keyboard"
	"github.com/jumdtw/FC-emulator/FC_CPU"
)

type Game struct {
	NoiseImage *image.RGBA
	Ppuemu *Ppu
	Cpuemu *FC_CPU.CpuEmu
	pressed []ebiten.Key
	usekey []ebiten.Key
}


func cpuexecute(g *Game){

	/*
	if g.Cpuemu.RegPc >= 0xc5dc && g.Cpuemu.RegPc <= 0xc5f0 {
		fmt.Printf("PC : 0x%04x, A : 0x%02x, X : 0x%02x, Y : 0x%02x, P : 0x%02x, SP : 0x%02x\n",g.Cpuemu.RegPc, g.Cpuemu.Regi["A"], g.Cpuemu.Regi["X"], g.Cpuemu.Regi["Y"], g.Cpuemu.Regi["P"], g.Cpuemu.Regi["S"])
	}
	*/
	
	//fmt.Printf("PC : 0x%x\n",g.Cpuemu.RegPc)
	fmt.Printf("PC : 0x%04x, A : 0x%02x, X : 0x%02x, Y : 0x%02x, P : 0x%02x, SP : 0x%02x\n",g.Cpuemu.RegPc, g.Cpuemu.Regi["A"], g.Cpuemu.Regi["X"], g.Cpuemu.Regi["Y"], g.Cpuemu.Regi["P"], g.Cpuemu.Regi["S"])
	g.Cpuemu.Execute()

	
	if g.Cpuemu.VramWriteFlag {
		g.Ppuemu.Memory[g.Cpuemu.VramAddr] = g.Cpuemu.VramWriteValue
		g.Cpuemu.VramAddr++
		g.Cpuemu.VramWriteFlag = false
	}

	if g.Cpuemu.OamWriteFlag {
		g.Cpuemu.OamWriteFlag = false
		cc := g.Cpuemu.OamWritecount
		if cc == 1 {
			g.Ppuemu.Oam[g.Cpuemu.Oamnum].Y = g.Cpuemu.OamWriteValue
		}
		if cc == 2 {
			g.Ppuemu.Oam[g.Cpuemu.Oamnum].Spritenum = g.Cpuemu.OamWriteValue
		}
		if cc == 3 {
			g.Ppuemu.Oam[g.Cpuemu.Oamnum].Sflag = uint8(g.Cpuemu.OamWriteValue)
		}
		if cc == 4 {
			g.Ppuemu.Oam[g.Cpuemu.Oamnum].X = g.Cpuemu.OamWriteValue
		}

	}

	
	if g.Cpuemu.DAMflag {
		g.Cpuemu.DAMflag = false
		var addr uint16 = uint16(g.Cpuemu.DAMvalue) * 16 * 16
		for v := 0 ; v < 0x100 ; v++ {
			g.Ppuemu.Oam[v].Y = int(g.Cpuemu.Memory[addr + uint16(v)])
			g.Ppuemu.Oam[v].Spritenum = int(g.Cpuemu.Memory[addr + uint16(v) + 1])
			g.Ppuemu.Oam[v].Sflag = g.Cpuemu.Memory[addr + uint16(v) + 2]
			g.Ppuemu.Oam[v].X = int(g.Cpuemu.Memory[addr + uint16(v) + 3])
		}
	}

}

func cp_mirrorvram(g *Game){
	// vertical
	if g.Cpuemu.Mirror == 1{
		for i:=0; i<1024 ; i++{
			g.Ppuemu.Memory[0x2800+i] = g.Ppuemu.Memory[0x2000+i]
			g.Ppuemu.Memory[0x2c00+i] = g.Ppuemu.Memory[0x2400+i]
		}
	// horizontal
	}else {
		for i:=0; i<1024 ; i++{
			g.Ppuemu.Memory[0x2400+i] = g.Ppuemu.Memory[0x2000+i]
			g.Ppuemu.Memory[0x2c00+i] = g.Ppuemu.Memory[0x2800+i]
		}
	}
	// ram mirror
	for i:=0; i<(0x1fff-0x800) ; i++{
		g.Cpuemu.Memory[0x800+i] = g.Cpuemu.Memory[0x200+i]
	}
}

func drawTile(g *Game,tileheadaddr int,numblock int,numtile int){


	var patternNum uint8
	var palletnum uint8
	
	//var palletnum uint8
	// chrnumにはname tableから、bgまたはspriteの番号を取り出す。
	nametableaddr := 0x2000
	chrnum := g.Ppuemu.Memory[nametableaddr + numtile]

	// 番号からほしいタイルのメモリ番号の頭アドレス.ここから128bitとりだす。
	Raddr := Selectbgcartridge(g)
	var romaddr uint64 = Raddr + uint64(chrnum)*16
	
	highbit, lowbit := Romdatareturn(g,romaddr)

	// patternNumがどのpalletnumであるか。入りえる値は0~3.
	palletnum = Palletnumreturn(g,numblock,numtile)	

	var pp int
	for i :=0; i<8; i++ {

		for k :=0; k < 8; k++ {

			// pp : 書き込もうとしているピクセルのrcの場所の配列数宇
			pp = tileheadaddr+k*4*Pixlsize+i*256*Pixlsize*4*Pixlsize
			// N tile目のpattern number
			patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
			var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f00)
			/*
			g.NoiseImage.Pix[pp] = rc
			g.NoiseImage.Pix[pp+1] = gc
			g.NoiseImage.Pix[pp+2] = bc
			g.NoiseImage.Pix[pp+3] = 0xff
			*/
			g.Cpuemu.Bufvram[pp] = rc
			g.Cpuemu.Bufvram[pp+1] = gc
			g.Cpuemu.Bufvram[pp+2] = bc
			g.Cpuemu.Bufvram[pp+3] = 0xff
			

			/*
			// ピクセルの大きさを上がるにはこの処理が必要。下記ソースはピクセルを2x2の大きさで一つのドットにするときに必要になる。
			g.NoiseImage.Pix[pp+4] = rc
			g.NoiseImage.Pix[pp+4+1] = gc
			g.NoiseImage.Pix[pp+4+2] = bc
			g.NoiseImage.Pix[pp+4+3] = 0xff
			
			g.NoiseImage.Pix[pp+256*Pixlsize*4] = rc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+1] = gc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+2] = bc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+3] = 0xff

			g.NoiseImage.Pix[pp+256*Pixlsize*4+4] = rc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+4+1] = gc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+4+2] = bc
			g.NoiseImage.Pix[pp+256*Pixlsize*4+4+3] = 0xff
			*/
			
		}
	}
}

func (g *Game) padread(){
	g.pressed = nil
	for _, k := range g.usekey {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}

	for _, p := range g.pressed {
		// A
		if p == g.usekey[0] {
			g.Cpuemu.Abotmflag = true
		}
		// B
		if p == g.usekey[1] {
			g.Cpuemu.Bbotmflag = true
		}
		// Select q key
		if p == g.usekey[2] {
			g.Cpuemu.Selectbotmflag = true
		}
		// Start w key
		if p == g.usekey[3] {
			g.Cpuemu.Startbotmflag = true
		} 
		// Down
		if p == g.usekey[4] {
			g.Cpuemu.Downbotmflag = true
		} 
		// Left
		if p == g.usekey[5] {
			g.Cpuemu.Leftbotmflag = true
		} 
		// Right
		if p == g.usekey[6] {
			g.Cpuemu.Rightbotmflag = true
		} 
		// Up
		if p == g.usekey[7] {
			g.Cpuemu.Upbotmflag = true
		} 
		// reset
		if p == g.usekey[8] {
			g.Cpuemu.RegPc = g.Cpuemu.Resetaddr
			g.Cpuemu.Regi["A"] = 0
			g.Cpuemu.Regi["X"] = 0
			g.Cpuemu.Regi["Y"] = 0
			g.Cpuemu.Regi["S"] = 0xff
			g.Cpuemu.Regi["P"] = 0b00000000
			for i, _ := range g.Ppuemu.Oam{
				g.Ppuemu.Oam[i].X = 0
				g.Ppuemu.Oam[i].Y = 0
				g.Ppuemu.Oam[i].Spritenum = 0
				g.Ppuemu.Oam[i].Sflag = 0
			}
		}
	}
}


func (g *Game) Update(screen *ebiten.Image) error {

	// Generate the noise with random RGB values.
	var numblock int
	var numtile int

	var vv int	
	
	for i :=0; i<30; i++ {

		for k :=0; k < 32; k++ {

			// vv は n枚目のタイルの一番左上のピクセルの最初の配列番号
			// k は一増えるごとに8pixl×ピクセル倍率分増える
			// i は一増えるごとに256×ピクセル倍率に8pixl×ピクセル倍率を掛ける
			// 最後に4を掛けることにより配列数を出す。
			vv = (k*8*Pixlsize+i*256*Pixlsize*8*Pixlsize)*4
			// patternNum がどのブロックにあるか
			// tileは32x30の半分なのでブロックは16x15.ナンバリングは0スタート
			numblock = Numblockreturn(i,k)
			numtile = k+i*32
			drawTile(g,vv,numblock,numtile)	
		}
	}

	
	return nil
}

func (g *Game) vblanckInterrupt() {
	// NMI割り込み
	// 多分本来のハードウェア動作だと次にこの動作処理が来るまでにCPU側で処理が終わっている。
	// でもこのエミュレータだと実質同期処理しているような状態なので処理が終わる前に再び
	// nmiaddrで割り込みが入ってしまう。そのため割り込み処理が終わるまでフラッグで判断して処理する。
	nmiflag := g.Cpuemu.Memory[0x2000] & 0b10000000
	vblankflag := g.Cpuemu.Memory[0x2002] & 0b10000000
	if nmiflag == 0b10000000 && g.Cpuemu.InterruptFlag == false && vblankflag == 0b10000000 {
		// ステータスレジスタの変化
		g.Cpuemu.Regi["P"] = g.Cpuemu.Regi["P"] & 0b11101011
		g.Cpuemu.Regi["P"] = g.Cpuemu.Regi["P"] + 0b00010100

		// 割り込みフラッグ　割り込み中はtrue
		// 割り込み処理に入ったのでtrueにする
		g.Cpuemu.InterruptFlag = true

		// スタックに現在のPCを積む
		var sppos uint16 = 0x0100 + uint16(g.Cpuemu.Regi["S"])
		// 0x4050 だったら　50 40 の順でスタックに積む。個々の動作は等で確認が取れていないためこのエミュレータだけの可能性あり。
		lowaddr := g.Cpuemu.RegPc & 0x00ff
		highaddr := g.Cpuemu.RegPc & 0xff00
		highaddr = highaddr >> 8
		g.Cpuemu.Memory[sppos-2] = g.Cpuemu.Regi["P"]
		g.Cpuemu.Memory[sppos-1] = uint8(lowaddr)
		g.Cpuemu.Memory[sppos] = uint8(highaddr)
		g.Cpuemu.Regi["S"] = g.Cpuemu.Regi["S"] - 3

		// 割り込みアドレスの代入
		g.Cpuemu.RegPc = g.Cpuemu.Nmiaddr
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	
	for v :=0 ; v < 260 ; v++ {
		
		// vblanck
		if v > 240 {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b10000000
		} else {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
		}
		if v==241 {
			g.vblanckInterrupt()
		}
		
		
		// 0 bom
		if (g.Ppuemu.Oam[0].Y+3) == v {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b01000000
		}else{
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
		}
		

		for i := 0 ; i < 256 ; i++ {
			// 0startだから239まで
			if v <= 239 {
				vv := i + v*256
				g.NoiseImage.Pix[vv*4] = g.Cpuemu.Bufvram[vv*4]
				g.NoiseImage.Pix[vv*4+1] = g.Cpuemu.Bufvram[vv*4+1]
				g.NoiseImage.Pix[vv*4+2] = g.Cpuemu.Bufvram[vv*4+2]
				g.NoiseImage.Pix[vv*4+3] = g.Cpuemu.Bufvram[vv*4+3]
			}
		}
		for q := 0 ; q < 114 ; q++ {
			if q%20 == 0 {
				g.padread()
			}
			cpuexecute(g)
		}
	}
	cp_mirrorvram(g)
	for i := 0 ; i < 256 ; i++ {
		DrawSprite(g,i)
	}
	screen.ReplacePixels(g.NoiseImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}


func main(){

	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("FC-emulator")

	// a b select(q) start(w) down left right up reset
	rusekey := []ebiten.Key{
		ebiten.Key(10),  // a 
		ebiten.Key(11),  // b
		ebiten.Key(26),  // q
		ebiten.Key(32),  // w
		ebiten.Key(42),  // down
		ebiten.Key(79),  // left 
		ebiten.Key(92),  // right 
		ebiten.Key(102), // up
		ebiten.Key(27), // reset
	}

	cpuEmu := FC_CPU.CpuEmu{}
	ppuEmu := Ppu{}
	chrrombuf := FC_CPU.InitEmu(&cpuEmu)
	Initppuemu(&ppuEmu, chrrombuf)


	g := &Game{
		NoiseImage: image.NewRGBA(image.Rect(0, 0, ScreenWidth, ScreenHeight)),
		Ppuemu: &ppuEmu,
		Cpuemu: &cpuEmu,
		pressed: nil,
		usekey: rusekey,
	}


	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("0x2000 : 0b%08b\n",g.Cpuemu.Memory[0x2000])
	
	/*
	for _, e := range g.Cpuemu.Exeopcdlist{
		fmt.Printf("opcd : 0x%x\n",e)
	}
	*/
}