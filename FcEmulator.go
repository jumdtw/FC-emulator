package main

import (

	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jumdtw/FC-emulator/FC_CPU"

)

type Game struct {
	NoiseImage *image.RGBA
	Ppuemu *Ppu
	Cpuemu *FC_CPU.CpuEmu
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
		g.Cpuemu.Memory[sppos] = uint8(highaddr)
		g.Cpuemu.Regi["S"] = g.Cpuemu.Regi["S"] - 2

		// 割り込みアドレスの代入
		g.Cpuemu.RegPc = g.Cpuemu.Nmiaddr
	}

	fmt.Printf("execute code : 0x%x\n",g.Cpuemu.Memory[g.Cpuemu.RegPc])
	g.Cpuemu.Execute()
	
	if g.Cpuemu.VramWriteFlag {
		g.Ppuemu.Memory[g.Cpuemu.VramAddr] = g.Cpuemu.VramWriteValue
		//fmt.Printf("vramaddr : 0x%x\n",g.Cpuemu.VramAddr)
		g.Cpuemu.VramAddr++
		g.Cpuemu.VramWriteFlag = false
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
			DrawTile(g,vv,numblock,numtile)

			// ppu はcpuの3倍のクロックを持っているのでppuの方が動作的には早いが
			// さすがにここで実行すると遅すぎるので将来的には修正が必要
			cpuexecute(g)

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

	cpuEmu := FC_CPU.CpuEmu{}
	ppuEmu := Ppu{}
	chrrombuf := FC_CPU.InitEmu(&cpuEmu)
	Initppuemu(&ppuEmu, chrrombuf)

	g := &Game{
		NoiseImage: image.NewRGBA(image.Rect(0, 0, ScreenWidth, ScreenHeight)),
		Ppuemu: &ppuEmu,
		Cpuemu: &cpuEmu,
	}

	//go StartPpu(&ppuEmu)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	
}