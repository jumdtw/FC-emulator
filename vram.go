package main

//import "fmt"

func drawTile(g *Game,tileheadaddr int,numblock int,numtile int,nametableaddr int){

	var patternNum uint8
	var palletnum uint8
	
	//var palletnum uint8
	// chrnumにはname tableから、bgまたはspriteの番号を取り出す。
	chrnum := g.Ppuemu.Memory[nametableaddr + numtile]

	// 番号からほしいタイルのメモリ番号の頭アドレス.ここから128bitとりだす。
	Raddr := Selectbgcartridge(g)
	var romaddr uint64 = Raddr + uint64(chrnum)*16
	
	highbit, lowbit := Romdatareturn(g,romaddr)

	var bufvramnum uint8 = 0
	var atriaddr int = 0x23c0

	if nametableaddr==0x2400 {
		bufvramnum = 1
		atriaddr = 0x27c0
	} else if nametableaddr==0x2800 {
		bufvramnum = 2
		atriaddr = 0x2bc0
	} else if nametableaddr==0x2c00 {
		bufvramnum = 3
		atriaddr = 0x2fc0
	}

	// patternNumがどのpalletnumであるか。入りえる値は0~3.
	palletnum = Palletnumreturn(g,numblock,numtile,atriaddr)	



	var pp int
	for i :=0; i<8; i++ {
		for k :=0; k < 8; k++ {
			// pp : 書き込もうとしているピクセルのrcの場所の配列数宇
			pp = tileheadaddr+k*4+i*256*4
			// N tile目のpattern number
			patternNum, highbit, lowbit = PatternNumreturn(highbit,lowbit)
			var rc, gc, bc uint8 = Rgbreturn(g,patternNum,palletnum,0x3f00)
			g.Cpuemu.Bufvram[bufvramnum][pp] = rc
			g.Cpuemu.Bufvram[bufvramnum][pp+1] = gc
			g.Cpuemu.Bufvram[bufvramnum][pp+2] = bc
			g.Cpuemu.Bufvram[bufvramnum][pp+3] = 0xff

		}
	}
}

func Veiwerupdate(g *Game){

	basetable := g.Cpuemu.Memory[0x2000] & 0b00000011
	basetableaddr := 0
	if basetable==1 {
		basetableaddr = 256
	} else if basetable==2 {
		basetableaddr = 512*240
	} else if basetable==3 {
		basetableaddr = 512*240 + 256
	}

	for q :=0; q < 240; q++ {
		if q==int(g.Ppuemu.Oam[0].Y+8) {
			basetableaddr = int(g.Cpuemu.DisplayX) + int(g.Cpuemu.DisplayY) * 512
		}
		for k :=0; k < 256; k++ {
			g.Cpuemu.Veiwbuf[(q*256+k)*4] = g.Cpuemu.FBufvram[(basetableaddr+k+q*512)*4]
			g.Cpuemu.Veiwbuf[(q*256+k)*4+1] = g.Cpuemu.FBufvram[(basetableaddr+k+q*512)*4+1]
			g.Cpuemu.Veiwbuf[(q*256+k)*4+2] = g.Cpuemu.FBufvram[(basetableaddr+k+q*512)*4+2]
			g.Cpuemu.Veiwbuf[(q*256+k)*4+3] = g.Cpuemu.FBufvram[(basetableaddr+k+q*512)*4+3]
		}
	}
}

func Screenupdate(g *Game){
	fbheader := 0
	for i :=0 ; i<4 ; i++{

		if i==1 {
			fbheader = 256
		} else if i==2 {
			fbheader = 512*240
		} else if i==3 {
			fbheader = 512*240 + 256
		}

		for q :=0; q<240; q++ {
			for k :=0; k < 256; k++ {
				g.Cpuemu.FBufvram[(fbheader+k+q*512)*4] = g.Cpuemu.Bufvram[i][(k+q*256)*4]
				g.Cpuemu.FBufvram[(fbheader+k+q*512)*4+1] = g.Cpuemu.Bufvram[i][(k+q*256)*4+1]
				g.Cpuemu.FBufvram[(fbheader+k+q*512)*4+2] = g.Cpuemu.Bufvram[i][(k+q*256)*4+2]
				g.Cpuemu.FBufvram[(fbheader+k+q*512)*4+3] = g.Cpuemu.Bufvram[i][(k+q*256)*4+3]
			}
		}
	}
}