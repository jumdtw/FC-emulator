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

	// NMI割り込み
	// 多分本来のハードウェア動作だと次にこの動作処理が来るまでにCPU側で処理が終わっている。
	// でもこのエミュレータだと実質同期処理しているような状態なので処理が終わる前に再び
	// nmiaddrで割り込みが入ってしまう。そのため割り込み処理が終わるまでフラッグで判断して処理する。
	nmiflag := g.Cpuemu.Memory[0x2000] & 0b10000000
	if nmiflag == 0b10000000 && g.Cpuemu.InterruptFlag == false{
		// ステータスレジスタの変化
		g.Cpuemu.Regi["P"] = g.Cpuemu.Regi["P"] & 0b11101011
		g.Cpuemu.Regi["P"] = g.Cpuemu.Regi["P"] + 0b00010100

		// 割り込みフラッグ　割り込み中はtrue
		// 割り込み処理に入ったのでtrueにする
		g.Cpuemu.InterruptFlag = true

		// スタックに現在のPCを積む
		var sppos uint16 = 0x1000 + uint16(g.Cpuemu.Regi["S"])
		// 0x4050 だったら　50 40 の順でスタックに積む。個々の動作は等で確認が取れていないためこのエミュレータだけの可能性あり。
		lowaddr := g.Cpuemu.RegPc & 0x00ff
		highaddr := g.Cpuemu.RegPc & 0xff00
		highaddr = highaddr >> 8
		g.Cpuemu.Memory[sppos] = uint8(lowaddr)
		g.Cpuemu.Memory[sppos-1] = uint8(highaddr)
		g.Cpuemu.Regi["S"] = g.Cpuemu.Regi["S"] - 2

		// 割り込みアドレスの代入
		g.Cpuemu.RegPc = g.Cpuemu.Nmiaddr
	}
	fmt.Printf("PC : 0x%x\n",g.Cpuemu.RegPc)
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

}

func cp_mirrorvram(g *Game){
	// vertical
	if g.Cpuemu.Mirror == 1{
		for i:=0; i<1024 ; i++{
			g.Cpuemu.Memory[0x2000+i] = g.Cpuemu.Memory[0x2800+i]
			g.Cpuemu.Memory[0x2400+i] = g.Cpuemu.Memory[0x2c00+i]
		}
	// horizontal
	}else {
		for i:=0; i<1024 ; i++{
			g.Cpuemu.Memory[0x2000+i] = g.Cpuemu.Memory[0x2400+i]
			g.Cpuemu.Memory[0x2800+i] = g.Cpuemu.Memory[0x2c00+i]
		}
	}
}

func drawTile(g *Game,tileheadaddr int,numblock int,numtile int){
	var patternNum uint8
	var palletnum uint8
	var vblankflag bool = false
	if numtile >= 960 {
		vblankflag = true
	}
	//var palletnum uint8
	// chrnumにはname tableから、bgまたはspriteの番号を取り出す。
	nametableaddr := 0x2000
	chrnum := g.Ppuemu.Memory[nametableaddr + numtile]
	// 番号からほしいタイルのメモリ番号の頭アドレス.ここから128bitとりだす。
	Raddr := Selectbgcartridge(g)
	var romaddr uint64 = Raddr + uint64(chrnum)*16
	var pp int

	highbit, lowbit := Romdatareturn(g,romaddr)

	// patternNumがどのpalletnumであるか。入りえる値は0~3.
	palletnum = Palletnumreturn(g,numblock,numtile)	


	for i :=0; i<8; i++ {
		for k :=0; k < 8; k++ {
			if vblankflag != true {
				// pp : 書き込もうとしているピクセルのrcの場所の配列数宇
				pp = tileheadaddr+k*4*Pixlsize+i*256*Pixlsize*4*Pixlsize
				// N tile目のpattern number
				patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
				var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f00)

				g.NoiseImage.Pix[pp] = rc
				g.NoiseImage.Pix[pp+1] = gc
				g.NoiseImage.Pix[pp+2] = bc
				g.NoiseImage.Pix[pp+3] = 0xff
			}

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


func (g *Game) Update(screen *ebiten.Image) error {

	// Generate the noise with random RGB values.
	var numblock int
	var numtile int
	var vv int

	g.pressed = nil
	for _, k := range g.usekey {
		if ebiten.IsKeyPressed(k) {
			g.pressed = append(g.pressed, k)
		}
	}
	g.Cpuemu.Abotmflag = false
	g.Cpuemu.Bbotmflag = false
	g.Cpuemu.Selectbotmflag = false
	g.Cpuemu.Startbotmflag = false
	g.Cpuemu.Leftbotmflag = false
	g.Cpuemu.Upbotmflag = false
	g.Cpuemu.Downbotmflag = false
	g.Cpuemu.Rightbotmflag = false
	for _, p := range g.pressed {
		switch p {
		case g.usekey[0]:
			g.Cpuemu.Abotmflag = true
		case g.usekey[1]:
			g.Cpuemu.Bbotmflag = true
		case g.usekey[2]:
			g.Cpuemu.Selectbotmflag = true
		case g.usekey[3]:
			g.Cpuemu.Startbotmflag = true
		case g.usekey[4]:
			g.Cpuemu.Downbotmflag = true
		case g.usekey[5]:
			g.Cpuemu.Leftbotmflag = true
		case g.usekey[6]:
			g.Cpuemu.Rightbotmflag = true
		case g.usekey[7]:
			g.Cpuemu.Upbotmflag = true
		case g.usekey[8]:
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

	
	
	for i :=0; i<35; i++ {

		if i > 30 {
			//fmt.Println("Vblank")
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b10000000
		} else {
			g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b01111111
		}

		for k :=0; k < 32; k++ {
			if i >= 27 {
				pp := (i-27)*32+k
				if pp == 0 {
					g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
					g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] + 0b01000000
				}else{
					g.Cpuemu.Memory[0x2002] = g.Cpuemu.Memory[0x2002] & 0b10111111
				}
				DrawSprite(g,pp)
			}
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
			if k % 5 == 0 {
				cpuexecute(g)
				cp_mirrorvram(g)
			}
		}
	}

	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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

}