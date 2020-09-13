package main

//import "fmt"

const (
	ScreenWidth  = 256
	ScreenHeight = 240
	memCap = 65535	
)

type Sprite struct {
	X int 
	Y int 
	Spritenum int 
	Sflag uint8
}

type Ppu struct {
	// Memory
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
	Memory [memCap]uint8
	Oam [64]Sprite 
}



func Romdatareturn(g *Game,romaddr uint64)(uint64,uint64){
	var highbit, lowbit uint64 = 0, 0
	for i:=0; i<8; i++{
		lowbit = lowbit << 8
		lowbit += uint64(g.Ppuemu.Memory[romaddr+uint64(i)])
	}
	romaddr += 8
	for i:=0; i<8; i++{
		highbit = highbit << 8 
		highbit += uint64(g.Ppuemu.Memory[romaddr+uint64(i)])
	}
	return highbit, lowbit
}

func PatternNumreturn(highbit uint64,lowbit uint64)(uint8,uint64,uint64){
	var patternNum uint8 = 0x00
	//一番最初のビットのみ必要なので取り出す。
	bufhighbit := highbit & 0x8000000000000000
	if bufhighbit == 0x8000000000000000{
		patternNum += 0b00000010
	}
	// 一番最初のビットはいらないので捨てる。
	highbit = highbit << 1
	//一番最初のビットのみ必要なので取り出す。
	buflowbit := lowbit & 0x8000000000000000
	if buflowbit == 0x8000000000000000 {
		patternNum += 0b00000001
	}
	// 一番最初のビットはいらないので捨てる。
	lowbit = lowbit << 1
	return uint8(patternNum),highbit,lowbit
}

// blockが色塗りの最小単位。tile to block と同様に折り返しがあるので注意
func rattributeoffset(numblock int)(int){
	var offset int
	//var highflag bool = false
	// 32で割ることで何行目かがわかる
	line := numblock/32
	// 何列目かを図っている
	var row int = (numblock%16)/2

	offset = line*8 + row

	return offset
}

func Palletnumreturn(g *Game,numblock int,numtile int,Raddr int)(uint8){
	var counter int
	var highflag bool = false 

	// まずnumblockを利用してメモリからattribute情報を取り出す。
	offset := rattributeoffset(numblock)
	var attribute uint8 = g.Ppuemu.Memory[Raddr + offset]

	// numtileの判定。ダルい。
	// 1, まず、16を順に引いていき奇数回なら下側。偶数なら上側(つまり数の少ない方)
	counter = 0
	bufnumblock := numblock
	counter = bufnumblock / 16

	if counter%2==1{
		highflag = true
	}
	if highflag {
		attribute = attribute >> 4
		attribute = attribute & 0b00001111
	}else{
		attribute = attribute & 0b00001111
	}
	// 2, 次にその数自体が偶数なら小さい方。奇数なら大きいほう
	if numblock%2 == 0 {
		attribute = 0b00000011 & attribute
	}else{
		attribute = attribute >> 2
	}

	return attribute
}

func Rgbreturn(g *Game,patternnum uint8,palletnum uint8,bgaddr int)(uint8,uint8,uint8){
	var rc, gc, bc uint8
	palletaddrhead := bgaddr + int(palletnum * 4)
	colornum := g.Ppuemu.Memory[palletaddrhead + int(patternnum)]
	//fmt.Printf("palletaddrhead : 0x%x, palletnum : 0x%x, colornum : 0x%x\n", palletaddrhead, palletnum, colornum)
	
	// ゼロは背景色
	if patternnum == 0 {
		colornum = g.Ppuemu.Memory[0x3f10]
	}
	
	
	switch colornum{
	case 0x00:
		rc = 109
		gc = 109
		bc = 107
	case 0x01:
		rc = 3
		gc = 35
		bc = 144
	case 0x02:
		rc = 1
		gc = 0
		bc = 220
	case 0x3:
		rc = 106 
		gc = 71
		bc = 223
	case 0x04:
		rc = 142
		gc = 1
		bc = 103
	case 0x05:
		rc = 179
		gc = 1
		bc = 113
	case 0x06:
		rc = 186 
		gc = 32
		bc = 4
	case 0x07:
		rc = 148
		gc = 83
		bc = 0
	case 0x08:
		rc = 110
		gc = 73
		bc = 0
	case 0x09:
		rc = 36
		gc = 73
		bc = 0
	case 0x0a:
		rc = 5
		gc = 106
		bc = 36
	case 0x0b:
		rc = 0
		gc = 151
		bc = 0
	case 0x0c:
		rc = 0
		gc = 70
		bc = 76
	case 0x0d:
		rc = 16
		gc = 0
		bc = 17
	case 0x0e:
		rc = 0
		gc = 0
		bc = 0
	case 0x0f:
		rc = 7
		gc = 0
		bc = 0
	case 0x10:
		rc = 182
		gc = 182
		bc = 182
	case 0x11:
		rc = 0
		gc = 108
		bc = 229
	case 0x12:
		rc = 0
		gc = 74
		bc = 251
	case 0x13:
		rc = 144
		gc = 0
		bc = 253
	case 0x14:
		rc = 183
		gc = 0
		bc = 251
	case 0x15:
		rc = 255
		gc = 0
		bc = 141
	case 0x16:
		rc = 239
		gc = 9
		bc = 0
	case 0x17:
		rc = 215
		gc = 109
		bc = 9
	case 0x18:
		rc = 145
		gc = 107
		bc = 6
	case 0x19:
		rc = 42
		gc = 144
		bc = 0
	case 0x1a:
		rc = 2
		gc = 144
		bc = 0
	case 0x1b:
		rc = 0
		gc = 182
		bc = 102
	case 0x1c:
		rc = 0
		gc = 146
		bc = 146
	case 0x1d:
		rc = 0
		gc = 0
		bc = 0
	case 0x1e:
		rc = 0
		gc = 0
		bc = 0
	case 0x1f:
		rc = 0
		gc = 0
		bc = 0
	case 0x20:
		rc = 255
		gc = 255
		bc = 255
	case 0x21:
		rc = 108
		gc = 176
		bc = 251
	case 0x22:
		rc = 153
		gc = 160
		bc = 255
	case 0x23:
		rc = 214
		gc = 110
		bc = 255
	case 0x24:
		rc = 255
		gc = 6
		bc = 236
	case 0x25:
		rc = 253
		gc = 110
		bc = 252
	case 0x26:
		rc = 255
		gc = 144
		bc = 0
	case 0x27:
		rc = 255
		gc = 182
		bc = 0
	case 0x28:
		rc = 217
		gc = 218
		bc = 0
	case 0x29:
		rc = 108
		gc = 219
		bc = 2
	case 0x2a:
		rc = 112
		gc = 225
		bc = 0
	case 0x2b:
		rc = 75
		gc = 254
		bc = 215
	case 0x2c:
		rc = 0
		gc = 251
		bc = 247
	case 0x2d:
		rc = 0
		gc = 0
		bc = 0
	case 0x2e:
		rc = 0
		gc = 0
		bc = 0
	case 0x2f:
		rc = 0
		gc = 0
		bc = 0
	case 0x30:
		rc = 255
		gc = 255
		bc = 255
	case 0x31:
		rc = 190
		gc = 224
		bc = 255
	case 0x32:
		rc = 216
		gc = 184
		bc = 247
	case 0x33:
		rc = 250
		gc = 185
		bc = 253
	case 0x34:
		rc = 255
		gc = 143
		bc = 255
	case 0x35:
		rc = 253
		gc = 183
		bc = 183
	case 0x36:
		rc = 255
		gc = 219
		bc = 143
	case 0x37:
		rc = 255
		gc = 254
		bc = 78
	case 0x38:
		rc = 253
		gc = 255
		bc = 109
	case 0x39:
		rc = 181
		gc = 254
		bc = 77
	case 0x3a:
		rc = 145
		gc = 254
		bc = 109
	case 0x3b:
		rc = 68
		gc = 255
		bc = 217
	case 0x3c:
		rc = 144
		gc = 217
		bc = 255
	case 0x3d:
		rc = 0
		gc = 0
		bc = 0
	case 0x3e:
		rc = 0
		gc = 0
		bc = 0
	case 0x3f:
		rc = 0
		gc = 0
		bc = 0
	default:
		rc = 0xff
		gc = 0
		bc = 0
	}

	return rc, gc, bc
}

func Numblockreturn(i int, k int)(int){
	var numblock int
	i = i/2
	numblock = i*16
	numblock += k/2
	return numblock
}

func Selectbgcartridge(g *Game) uint64 {
	var Raddr uint64 = 0
	bgpatternaddr := g.Cpuemu.Memory[0x2000] & 0b00010000
	if bgpatternaddr == 0b00010000{
		Raddr = 0x1000
	}
	return Raddr
}

func selectspritecartridge(g *Game) uint64 {
	var Raddr uint64 = 0
	bgpatternaddr := g.Cpuemu.Memory[0x2000] & 0b00010000
	if bgpatternaddr == 0b00010000{
		Raddr = 0x1000
	}
	return Raddr
}

func spriteHorVeri(g *Game, highbit uint64, lowbit uint64, vertical uint8, horizontal uint8, palletnum uint8)([]uint8){
	var patternNum uint8
	spritepallet := make([]uint8,64*4)
	if vertical==0x80 && horizontal==0x40 {
		for i :=7; i>=0; i-- {
			for k :=7; k >= 0; k-- {
				patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
				var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f10)
				pp :=(i*8+k)*4
				spritepallet[pp] = rc
				spritepallet[pp+1] = gc
				spritepallet[pp+2] = bc
				spritepallet[pp+3] = 0xff
				if patternNum == 0 {
					spritepallet[pp+3] = 0
				}
			}
		}
	} else if vertical==0x80 {
		for i :=7; i>=0; i-- {
			for k :=0; k < 8; k++ {
				patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
				var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f10)
				pp :=(i*8+k)*4
				spritepallet[pp] = rc
				spritepallet[pp+1] = gc
				spritepallet[pp+2] = bc
				spritepallet[pp+3] = 0xff
				if patternNum == 0 {
					spritepallet[pp+3] = 0
				}
			}
		}
	} else if horizontal==0x40 {
		for i :=0; i<8; i++ {
			for k :=7; k >= 0; k-- {
				patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
				var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f10)
				pp :=(i*8+k)*4
				spritepallet[pp] = rc
				spritepallet[pp+1] = gc
				spritepallet[pp+2] = bc
				spritepallet[pp+3] = 0xff
				if patternNum == 0 {
					spritepallet[pp+3] = 0
				}
			}
		}
	} else {
		for i :=0; i<8; i++ {
			for k :=0; k < 8; k++ {
				patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
				var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f10)
				pp :=(i*8+k)*4
				spritepallet[pp] = rc
				spritepallet[pp+1] = gc
				spritepallet[pp+2] = bc
				spritepallet[pp+3] = 0xff
				if patternNum == 0 {
					spritepallet[pp+3] = 0
				}
			}
		}
	}
	return spritepallet
}

func DrawSprite(g *Game, pp int){
	var palletnum uint8
	oam := g.Ppuemu.Oam[pp]
	palletnum = oam.Sflag & 0b00000011
	var vertical uint8 = oam.Sflag & 0b10000000
	var horizontal uint8 = oam.Sflag & 0b01000000
	spritex := oam.X
	spritey := oam.Y
	// selectspritecartridge
	var raddr uint64 = 0x1000
	bgpatternaddr := g.Cpuemu.Memory[0x2000] & 0b00010000
	if bgpatternaddr == 0b00010000{
		raddr = 0
	}
	highbit, lowbit := Romdatareturn(g,raddr+uint64(oam.Spritenum*16))
	spritepallet := spriteHorVeri(g, highbit, lowbit, vertical, horizontal, palletnum)
	arrayhead := (spritex + spritey*256)*4
	for i :=0; i<8; i++ {
		for k :=0; k < 8; k++ {
			qq :=(i*8+k)*4
			pp := arrayhead + (k+i*256)*4
			if pp < 256*240*4 {
				if spritepallet[qq+3] != 0 {
					g.NoiseImage.Pix[pp] = spritepallet[qq]
					g.NoiseImage.Pix[pp+1] = spritepallet[qq+1]
					g.NoiseImage.Pix[pp+2] = spritepallet[qq+2]
					g.NoiseImage.Pix[pp+3] = spritepallet[qq+3]
				}
			}
		}
	}
}

func Initppuemu(Ppuemu *Ppu,chrrombuf []uint8){
	for i:=0 ;i<0x1000; i++{
		Ppuemu.Memory[0x0000+i] = chrrombuf[i]
	}
	for i:=0 ;i<0x1000; i++{
		Ppuemu.Memory[0x1000+i] = chrrombuf[0x1000+i]
	}
	for i:=0 ;i<64; i++{
		Ppuemu.Oam[i].X = 0x100
		Ppuemu.Oam[i].Y = 0x100
		Ppuemu.Oam[i].Spritenum = 0x100
	}
}
