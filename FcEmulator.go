package main

import (

	"fmt"
	"image"
	"log"
	//"os"
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

	var incppu uint8
	
	//PC : 0x800d, A : 0x10, X : 0xff, Y : 0x00, P : 0x24, SP : 0xff
	//fmt.Printf("PC : $%x\n",g.Cpuemu.RegPc)
	g.Cpuemu.Execute()

	incppu = g.Cpuemu.Memory[0x2000] & 0x4

	if g.Cpuemu.VramReadFlag {
		if g.Cpuemu.VramAddr >=0 && g.Cpuemu.VramAddr <= 0x3eff {
			g.Cpuemu.VramnonereadFlag = true
		} else {
			g.Cpuemu.VramnonereadFlag = false 
		}
		g.Cpuemu.Regi[g.Cpuemu.VramReadReg] = g.Ppuemu.Memory[g.Cpuemu.VramAddr]
		//fmt.Printf("read PC : $%x, ppumemory : $%x, ppuvalue : $%x\n",g.Cpuemu.RegPc, g.Cpuemu.VramAddr, g.Ppuemu.Memory[g.Cpuemu.VramAddr])
		if !(g.Cpuemu.VramnonereadFlag && !g.Cpuemu.VramonecFlag) {
			if incppu == 0x4 {
				g.Cpuemu.VramAddr+=32
			} else {
				g.Cpuemu.VramAddr++
			}
		} else {
			g.Cpuemu.VramonecFlag = true
		}
		g.Cpuemu.VramReadFlag = false
	}
	
	if g.Cpuemu.VramWriteFlag {
		g.Ppuemu.Memory[g.Cpuemu.VramAddr] = g.Cpuemu.VramWriteValue
		//fmt.Printf("write PC : $%x, ppumemory : $%x, value : $%x\n",g.Cpuemu.RegPc, g.Cpuemu.VramAddr, g.Cpuemu.VramWriteValue)
		if incppu == 0x4 {
			g.Cpuemu.VramAddr+=32
		} else {
			g.Cpuemu.VramAddr++
		}
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
		for v := 0 ; v < 256 ; v+=4 {
			g.Ppuemu.Oam[v/4].Y = int(g.Cpuemu.Memory[addr + uint16(v)])
			g.Ppuemu.Oam[v/4].Spritenum = int(g.Cpuemu.Memory[addr + uint16(v) + 1])
			g.Ppuemu.Oam[v/4].Sflag = g.Cpuemu.Memory[addr + uint16(v) + 2]
			g.Ppuemu.Oam[v/4].X = int(g.Cpuemu.Memory[addr + uint16(v) + 3])
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
			vv = (k*8+i*256*8)*4
			// patternNum がどのブロックにあるか
			// tileは32x30の半分なのでブロックは16x15.ナンバリングは0スタート
			numblock = Numblockreturn(i,k)
			numtile = k+i*32
			drawTile(g,vv,numblock,numtile,0x2000)
			drawTile(g,vv,numblock,numtile,0x2400)
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
	
	Screenupdate(g)
	Veiwerupdate(g)

	for v :=0 ; v < 260 ; v++ {
		
		// vblanck
		if v > 240 && v < 250 {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b10000000
		} else {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
		}
		if v==241 {
			g.vblanckInterrupt()
		}
		
		
		// 0 bom
		if (g.Ppuemu.Oam[0].Y) == v {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b01000000
		}else{
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
		}

		for q := 0 ; q < 200 ; q++ {
			if q%20 == 0 {
				//g.padread()
			}
			cpuexecute(g)
		}

		if g.Cpuemu.Displayupdateflag{
			g.Cpuemu.ZerobomY = v
		}
		
		if g.Cpuemu.Displayupdateflag {
			Screenupdate(g)
			Veiwerupdate(g)
			g.Cpuemu.Displayupdateflag = false
		}
		
		for i := 0 ; i < 256 ; i++ {
			// 0startだから239まで
			if v <= 239 {
				vv := i + v*256
				g.NoiseImage.Pix[vv*4] = g.Cpuemu.Veiwbuf[vv*4]
				g.NoiseImage.Pix[vv*4+1] = g.Cpuemu.Veiwbuf[vv*4+1]
				g.NoiseImage.Pix[vv*4+2] = g.Cpuemu.Veiwbuf[vv*4+2]
				g.NoiseImage.Pix[vv*4+3] = g.Cpuemu.Veiwbuf[vv*4+3]
			}
		}

		
	}
	cp_mirrorvram(g)
	for i := 0 ; i < 64 ; i++ {
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