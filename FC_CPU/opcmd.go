package FC_CPU

import "fmt"

const (
	// LDAIMM : load a reg imm
	LDAIMM = 0xa9
	// LDAZERO :
	LDAZERO = 0xa5
	// LDAZEROX :
	LDAZEROX = 0xb5
	// LDAABS :
	LDAABS = 0xad
	// LDAABSX :
	LDAABSX = 0xbd
	// LDAABSY :
	LDAABSY = 0xb9
	// LDAINDX :
	LDAINDX = 0xa1
	// LDAINDY :
	LDAINDY = 0xb1

	// LDXIMM :
	LDXIMM = 0xa2
	// LDXZERO :
	LDXZERO = 0xa6
	// LDXZEROY :
	LDXZEROY = 0xb6
	// LDXABS :
	LDXABS = 0xae
	// LDXABSX :
	LDXABSX = 0xbe

	// LDYIMM:
	LDYIMM = 0xa0
	// LDYZERO:
	LDYZERO = 0xa4
	// LDYABS:
	LDYABS = 0xac
	// LDYABSX:
	LDYABSX = 0xbc

	// STAZERO :
	STAZERO = 0x85
	// STAZEROX :
	STAZEROX = 0x95
	// STAABS :
	STAABS = 0x8d
	// STAABSX :
	STAABSX = 0x9d
	// STAABSY:
	STAABSY = 0x99
	// STAINDY:
	STAINDY = 0x91

	// STXZERO :
	STXZERO = 0x86
	// STXABS :
	STXABS = 0x8e

	// STYZERO :
	STYZERO = 0x84
	// STYABS :
	STYABS = 0x8c

	// DECZERO :
	DECZERO = 0xc6
	// DECABS :
	DECABS = 0xce
	// DECABSX :
	DECABSX = 0xde

	// DEX :
	DEX = 0xca

	// DEY :
	DEY = 0x88

	//EORIMM :
	EORIMM = 0x49
	//EORZERO :
	EORZERO = 0x45

	// INCZERO :
	INCZERO = 0xe6
	// INCABS :
	INCABS = 0xee

	// INX :
	INX = 0xe8
	// INY :
	INY = 0xc8

	// LSRIMM :
	LSRIMM = 0x4a
	// LSRZERO :
	LSRZERO = 0x46
	// LSRABS :
	LSRABS = 0x4e

	// ORAIMM :
	ORAIMM = 0x09
	// ORAZERO :
	ORAZERO = 0x05
	// ORAABS :
	ORAABS = 0x0d

	// ROLIMM :
	ROLIMM = 0x2a
	// ROLZERO :
	ROLZERO = 0x26
	// ROLABS :
	ROLABS = 0x2e

	// RORIMM :
	RORIMM = 0x6a
	// RORABSX :
	RORABSX = 0x7e

	// SBCIMM :
	SBCIMM = 0xe9
	// SBCZERO :
	SBCZERO = 0xe5
	// SBCZEROX :
	SBCZEROX = 0xf5
	// SBCABS :
	SBCABS = 0xed
	// SBCABSY :
	SBCABSY = 0xf9

	// BCC :
	BCC = 0x90
	// BCS :
	BCS = 0xb0
	// BEQ :
	BEQ = 0xf0
	// BMI :
	BMI = 0x30
	// BNE :
	BNE = 0xd0
	// BPL :
	BPL = 0x10

	// TAX :
	TAX = 0xaa

	// TAY :
	TAY = 0xa8

	// TXA :
	TXA = 0x8a

	// TXS :
	TXS = 0x9a

	// TYA :
	TYA = 0x98

	// ADCIMM : 
	ADCIMM = 0x69
	// ADCZERO :
	ADCZERO = 0x65
	// ADCZEROX :
	ADCZEROX = 0x75
	// ADCABS :
	ADCABS = 0x6d
	// ADCABSX :
	ADCABSX = 0x7d
	// ADCABSY :
	ADCABSY = 0x79

	// ANDIMM : 
	ANDIMM = 0x29
	// ANDABS : 
	ANDABS = 0x2d
	// ANDABSX : 
	ANDABSX = 0x3d

	// ASLIMM :
	ASLIMM = 0x0a
	// ASLABS :
	ASLABS = 0x0e

	// BITZERO :
	BITZERO = 0x24
	// BITABS :
	BITABS = 0x2c

	// CMPIMM :
	CMPIMM = 0xc9
	// CMPZERO :
	CMPZERO = 0xc5
	// CMPABS :
	CMPABS = 0xcd
	// CMPABSX :
	CMPABSX = 0xdd
	// CMPABSY :
	CMPABSY = 0xd9

	// CPX :
	CPXIMM = 0xe0

	// CPY :
	CPYIMM = 0xc0

	// JMPABS :
	JMPABS = 0x4c
	// JMPIND :
	JMPIND = 0x6c

	// JSR :
	JSR = 0x20

	// RIS :
	RTS = 0x60

	// RIT :
	RTI = 0x40

	// PHA :
	PHA = 0x48
	// PLA :
	PLA = 0x68

	// CLC :
	CLC = 0x18

	// CLD :
	CLD = 0xd8

	// CLI : 
	CLI = 0x58

	// SEC :
	SEC = 0x38

	// SEI :
	SEI = 0x78
)

func read_ppustatic(fcEmu *CpuEmu, pos uint16){

}

func ppuwrite(fcEmu *CpuEmu, pos uint16,reg string){

	//fmt.Printf("now sta pos : 0x%x, memoryvalue : 0x%x\n",pos,fcEmu.Memory[pos])

	// 0x2006だったらppuへのアクセス
	if pos == 0x2006 {
		fcEmu.VramAddr = fcEmu.VramAddr << 8
		fcEmu.VramAddr += uint16(fcEmu.Regi[reg])
		//fmt.Printf("change vram pos : 0x%x !!!!!!!!!!!!!!!!!!\n",fcEmu.Regi[reg])
	}
	if pos == 0x2007 {
		fcEmu.VramWriteFlag = true
		fcEmu.VramWriteValue = fcEmu.Regi[reg]
		//fmt.Printf("vram write : 0x%x ????????????????\n", fcEmu.Regi[reg])
	}else {
		fcEmu.VramWriteFlag = false
	}
	// 0x2004 0x2003だったらoamへの
	if pos == 0x2003 {
		fcEmu.Oamnum = int(fcEmu.Regi[reg])
	}
	if pos == 0x2004 {
		fcEmu.OamWriteFlag = true
		fcEmu.OamWriteValue = int(fcEmu.Regi[reg])
		if fcEmu.OamWritecount == 4{
			fcEmu.Oamnum++
			if fcEmu.Oamnum >= 255{
				fcEmu.Oamnum = 0
			}
			fcEmu.OamWritecount = 0
		}
		fcEmu.OamWritecount++
	}

	if pos == 0x2002{
		fcEmu.VramAddr = 0
		//fmt.Println("reset scroll")
	}

	if pos == 0x2005 {
		//fmt.Printf("scroll x or y : 0x%x\n",fcEmu.Regi[reg])
	}

	if pos == 0x4014 {
		//fmt.Println("send sprite DMA!!!!")
		fcEmu.DAMflag = true
		fcEmu.DAMvalue = fcEmu.Regi[reg]
	}

	// pad
	if pos == 0x4016 {
		fcEmu.Padvalue = fcEmu.Padvalue << 8
		fcEmu.Padvalue += uint16(fcEmu.Regi[reg])
		if fcEmu.Padvalue == 0x100 {
			fcEmu.Abotmflag = false
			fcEmu.Bbotmflag = false
			fcEmu.Selectbotmflag = false
			fcEmu.Startbotmflag = false
			fcEmu.Leftbotmflag = false
			fcEmu.Upbotmflag = false
			fcEmu.Downbotmflag = false
			fcEmu.Rightbotmflag = false
		}
	}
}

func padread(fcEmu *CpuEmu, pos uint16){
	//fmt.Printf("pad read : 0x%x\n",pos)
	// 0x4016 はパッドの読み込み
	if pos == 0x4016 {
		fcEmu.BotmReadcount++
		switch fcEmu.BotmReadcount {
		// A
		case 1:
			if fcEmu.Abotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// B
		case 2:
			if fcEmu.Bbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// Select
		case 3:
			if fcEmu.Selectbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// Start
		case 4:
			if fcEmu.Startbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// up
		case 5:
			if fcEmu.Upbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// down
		case 6:
			if fcEmu.Downbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// left
		case 7:
			if fcEmu.Leftbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// right
		case 8:
			if fcEmu.Rightbotmflag {
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
			fcEmu.BotmReadcount = 0
		}
	}else{
		fcEmu.BotmReadcount = 0
	}
}

func nflagset(fcEmu *CpuEmu,x uint8){
	x = x & 0b10000000
	if x == 0b10000000{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.Regi["P"] = 0b01111111 & fcEmu.Regi["P"]
		fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b10000000
	}else{
		fcEmu.Regi["P"] = 0b01111111 & fcEmu.Regi["P"]
	}
}

func vflagset(fcEmu *CpuEmu, x uint8, y uint8, c uint8, plusflag bool){

	xnflag := x & 0b10000000
	ynflag := y & 0b10000000

	// 足し算の場合
	// オーバーフローは符号が同じじゃないと起きない
	// 正の値
	if plusflag {
		var plus uint16 = uint16(x) + uint16(y) + uint16(c)
		if plus > 0b0000000001111111{
			fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
			fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
		}else{
			fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
		}
	}

	//引き算の場合
	if !plusflag {
		if xnflag != 0b10000000 && ynflag == 0b10000000 {

			if plusflag {
				var plus uint16 = uint16(x) - uint16(y) - uint16(c)
				if plus > 0b0000000001111111{
					fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
					fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
				}else{
					fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				}
			}
	
		}else if xnflag == 0b10000000 && ynflag != 0b10000000 {
			if plusflag {
				var plus uint16 = uint16(y) - uint16(x) - uint16(c)
				if plus > 0b0000000001111111{
					fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
					fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
				}else{
					fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				}
			}
		}
	}
}

func zflagset(fcEmu *CpuEmu,x uint8){
	if x == 0x00{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.Regi["P"] = 0b11111101 & fcEmu.Regi["P"]
		fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000010
	}else{
		fcEmu.Regi["P"] = 0b11111101 & fcEmu.Regi["P"]
	}
}

func cflagset(fcEmu *CpuEmu,x uint16){
	if x == 0x0100{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.Regi["P"] = 0b11111110 & fcEmu.Regi["P"]
		fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000001
	}else{
		fcEmu.Regi["P"] = 0b11111110 & fcEmu.Regi["P"]
	}
}




// LDA
func (fcEmu *CpuEmu) ldaImm() {
	fcEmu.Regi["A"] = fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	zflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldaZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldaZeroX() {
	var immpos uint8 = fcEmu.Memory[fcEmu.RegPc]
	var pos uint16 = uint16(fcEmu.Regi["X"] + immpos)
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldaAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	padread(fcEmu,pos)
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) ldaAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var xreg uint16 = uint16(fcEmu.Regi["X"])
	var pos uint16 = (absposhigh << 8) + absposlow + xreg
	padread(fcEmu,pos)
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) ldaAbsY() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var yreg uint16 = uint16(fcEmu.Regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	padread(fcEmu,pos)
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) ldaIndX() {
	var xreg uint16 = uint16(fcEmu.Regi["X"])
	var immaddr uint16 = uint16(fcEmu.Memory[fcEmu.RegPc]) + xreg
	var highpos uint16 = uint16(fcEmu.Memory[immaddr+1])
	var lowpos uint16 = uint16(fcEmu.Memory[immaddr])
	var pos uint16 = (highpos << 8) + lowpos
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	padread(fcEmu,pos)
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldaIndY() {
	var immaddr uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var yreg uint16 = uint16(fcEmu.Regi["Y"])
	var highpos uint16 = uint16(fcEmu.Memory[immaddr+1])
	var lowpos uint16 = uint16(fcEmu.Memory[immaddr])
	var pos uint16 = (highpos << 8) + lowpos + yreg
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	padread(fcEmu,pos)
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}

// LDX
func (fcEmu *CpuEmu) ldxImm() {
	fcEmu.Regi["X"] = fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	zflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldxZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldxZeroY() {
	var immpos uint8 = fcEmu.Memory[fcEmu.RegPc]
	var pos uint16 = uint16(fcEmu.Regi["Y"] + immpos)
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldxAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	padread(fcEmu,pos)
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) ldxAbsY() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var yreg uint16 = uint16(fcEmu.Regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	padread(fcEmu,pos)
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// LDY
func (fcEmu *CpuEmu) ldyImm() {
	fcEmu.Regi["Y"] = fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	zflagset(fcEmu, fcEmu.Memory[fcEmu.RegPc])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldyZero() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	padread(fcEmu,pos)
	fcEmu.Regi["Y"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 1
}
func (fcEmu *CpuEmu) ldyAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	padread(fcEmu,pos)
	fcEmu.Regi["Y"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) ldyAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	padread(fcEmu,pos)
	fcEmu.Regi["Y"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// sta
// 6502 はメモリマップトioなのでstaだけ命令処理以外に色々処理が入る場合が多い
func (fcEmu *CpuEmu) staZero() {
	fcEmu.Memory[fcEmu.Memory[fcEmu.RegPc]] = fcEmu.Regi["A"]
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) staZeroX() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(fcEmu.Regi["X"])
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) staAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	ppuwrite(fcEmu,pos,"A")
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	pos = pos + uint16(fcEmu.Regi["X"])
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	ppuwrite(fcEmu,pos,"A")
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staAbsY() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	pos = pos + uint16(fcEmu.Regi["Y"])
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	ppuwrite(fcEmu,pos,"A")
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staIndY() {
	var immaddr uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var yreg uint16 = uint16(fcEmu.Regi["Y"])
	var highpos uint16 = uint16(fcEmu.Memory[immaddr+1])
	var lowpos uint16 = uint16(fcEmu.Memory[immaddr])
	var pos uint16 = (highpos << 8) + lowpos + yreg
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	ppuwrite(fcEmu,pos,"A")
	fcEmu.RegPc++
}

// stx
func (fcEmu *CpuEmu) stxZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Memory[pos] = fcEmu.Regi["X"]
	ppuwrite(fcEmu,uint16(pos),"X")
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) stxAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Memory[pos] = fcEmu.Regi["X"]
	ppuwrite(fcEmu,pos,"X")
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// sty
func (fcEmu *CpuEmu) styZero() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	fcEmu.Memory[pos] = fcEmu.Regi["Y"]
	ppuwrite(fcEmu,pos,"Y")
	fcEmu.RegPc = fcEmu.RegPc + 1
}
func (fcEmu *CpuEmu) styAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Memory[pos] = fcEmu.Regi["Y"]
	ppuwrite(fcEmu,pos,"Y")
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// tax
func (fcEmu *CpuEmu) tax() {
	fcEmu.Regi["X"] = fcEmu.Regi["A"]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}

// tay
func (fcEmu *CpuEmu) tay() {
	fcEmu.Regi["Y"] = fcEmu.Regi["A"]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}

// txa
func (fcEmu *CpuEmu) txa() {
	fcEmu.Regi["A"] = fcEmu.Regi["X"]
	nflagset(fcEmu, fcEmu.Regi["X"])
	zflagset(fcEmu, fcEmu.Regi["X"])
}

// txs
func (fcEmu *CpuEmu) txs() {
	fcEmu.Regi["S"] = fcEmu.Regi["X"]
	nflagset(fcEmu, fcEmu.Regi["X"])
	zflagset(fcEmu, fcEmu.Regi["X"])
}

// tya
func (fcEmu *CpuEmu) tya() {
	fcEmu.Regi["A"] = fcEmu.Regi["Y"]
	nflagset(fcEmu, fcEmu.Regi["Y"])
	zflagset(fcEmu, fcEmu.Regi["Y"])
}

// dec
func (fcEmu *CpuEmu) decZero() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	fcEmu.Memory[pos] = fcEmu.Memory[pos] - 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) decAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Memory[pos] = fcEmu.Memory[pos] - 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) decAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	fcEmu.Memory[pos] = fcEmu.Memory[pos] - 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// dex
func (fcEmu *CpuEmu) deX() {
	fcEmu.Regi["X"] = fcEmu.Regi["X"] - 1
	nflagset(fcEmu, fcEmu.Regi["X"])
	zflagset(fcEmu, fcEmu.Regi["X"])
}

// dey
func (fcEmu *CpuEmu) deY() {
	fcEmu.Regi["Y"] = fcEmu.Regi["Y"] - 1
	nflagset(fcEmu, fcEmu.Regi["Y"])
	zflagset(fcEmu, fcEmu.Regi["Y"])
}

// eor
func (fcEmu *CpuEmu) eorImm() {
	fcEmu.Regi["A"] = fcEmu.Regi["A"] ^ fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) eorZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Regi["A"] ^ fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
}

// inc
func (fcEmu *CpuEmu) incZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) incAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// inx
func (fcEmu *CpuEmu) inX() {
	fcEmu.Regi["X"] = fcEmu.Regi["X"] + 1
	nflagset(fcEmu, fcEmu.Regi["X"])
	zflagset(fcEmu, fcEmu.Regi["X"])
}

// iny
func (fcEmu *CpuEmu) inY() {
	fcEmu.Regi["Y"] = fcEmu.Regi["Y"] + 1
	nflagset(fcEmu, fcEmu.Regi["Y"])
	zflagset(fcEmu, fcEmu.Regi["Y"])
}

// lsr
func (fcEmu *CpuEmu) lsrImm() {
	lowbit := fcEmu.Regi["A"] & 0x01
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = fcEmu.Regi["A"] >> 1
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
func (fcEmu *CpuEmu) lsrZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	lowbit := fcEmu.Memory[pos] & 0x01
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Memory[pos] = fcEmu.Memory[pos] >> 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) lsrAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	lowbit := fcEmu.Memory[pos] & 0x01
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Memory[pos] = fcEmu.Memory[pos] >> 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc+=2
}

// ora 
func (fcEmu *CpuEmu) oraImm(){
	fcEmu.Regi["A"] = fcEmu.Regi["A"] | fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) oraZero(){
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Regi["A"] | fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) oraAbs(){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Regi["A"] = fcEmu.Regi["A"] | fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc+=2
}

// rol
func (fcEmu *CpuEmu) rolImm(){
	var highbit uint8 = fcEmu.Regi["A"] & 0b10000000
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	fcEmu.Regi["A"] = fcEmu.Regi["A"] << 1
	fcEmu.Regi["A"] = fcEmu.Regi["A"] + clflag
	if highbit == 0b10000000 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
func (fcEmu *CpuEmu) rolZero(){
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	var highbit uint8 = fcEmu.Memory[pos] & 0b10000000
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	fcEmu.Memory[pos] = fcEmu.Memory[pos] << 1
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + clflag
	if highbit == 0b10000000 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) rolAbs(){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var highbit uint8 = fcEmu.Memory[pos] & 0b10000000
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	fcEmu.Memory[pos] = fcEmu.Memory[pos] << 1
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + clflag
	if highbit == 0b10000000 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc+=2
}

// ror
func (fcEmu *CpuEmu) rorImm(){
	var lowbit uint8 = fcEmu.Regi["A"] & 0b00000001
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	clflag = clflag << 7
	fcEmu.Regi["A"] = fcEmu.Regi["A"] >> 1
	fcEmu.Regi["A"] = fcEmu.Regi["A"] + clflag
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
func (fcEmu *CpuEmu) rorAbsX(){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	var lowbit uint8 = fcEmu.Memory[pos] & 0b00000001
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	clflag = clflag << 7
	fcEmu.Memory[pos] = fcEmu.Memory[pos] >> 1
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + clflag
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc+=2
}

// sbc
func (fcEmu *CpuEmu) sbcImm() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var memo uint8 = fcEmu.Memory[fcEmu.RegPc]
	memo = ^memo + 1
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(memo)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)	
	}
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[fcEmu.RegPc], clflag,false)
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) sbcZero() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var memo uint8 = fcEmu.Memory[pos]
	memo = ^memo + 1
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(memo)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)
	}
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], clflag,false)
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) sbcZeroX() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(fcEmu.Regi["X"])
	var memo uint8 = fcEmu.Memory[pos]
	memo = ^memo + 1
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(memo)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)	
	}
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], clflag,false)
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) sbcAbs() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var memo uint8 = fcEmu.Memory[pos]
	memo = ^memo + 1
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(memo)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)	
	}
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	cflagset(fcEmu, uint16(fcEmu.Regi["A"]) - uint16(fcEmu.Memory[pos]) - uint16(clflag))
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], clflag,false)
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) sbcAbsY() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["Y"])
	var memo uint8 = fcEmu.Memory[pos]
	memo = ^memo + 1
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(memo)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)	
	}
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], clflag,false)
	fcEmu.RegPc+=2
}

// bcc
func (fcEmu *CpuEmu) bcc() {
	zeroflag := fcEmu.Regi["P"] & 0b00000001
	if zeroflag == 0b00000000 {
		bufPC := uint16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm uint16
		if minusflag == 0b10000000 {
			var minussign uint16 = 0xff00
			imm = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(minussign)
		} else {
			imm = uint16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// bcs
func (fcEmu *CpuEmu) bcs() {
	zeroflag := fcEmu.Regi["P"] & 0b00000001
	if zeroflag == 0b00000001 {
		bufPC := uint16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm uint16
		if minusflag == 0b10000000 {
			var minussign uint16 = 0xff00
			imm = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(minussign)
		} else {
			imm = uint16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// beq
func (fcEmu *CpuEmu) beq() {
	zeroflag := fcEmu.Regi["P"] & 0b00000010
	if zeroflag == 0b00000010 {
		bufPC := uint16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm uint16
		if minusflag == 0b10000000 {
			var minussign uint16 = 0xff00
			imm = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(minussign)
		} else {
			imm = uint16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// bmi
func (fcEmu *CpuEmu) bmi() {
	nflag := fcEmu.Regi["P"] & 0b10000000
	if nflag == 0x80 {
		bufPC := int16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm int16
		if minusflag > 0 {
			var minussign uint16 = 0xff00
			imm = int16(fcEmu.Memory[fcEmu.RegPc]) + int16(minussign)
		} else {
			imm = int16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// bne
func (fcEmu *CpuEmu) bne() {
	zeroflag := fcEmu.Regi["P"] & 0b00000010
	if zeroflag == 0x00 {
		bufPC := int16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm int16
		if minusflag > 0 {
			var minussign uint16 = 0xff00
			imm = int16(fcEmu.Memory[fcEmu.RegPc]) + int16(minussign)
		} else {
			imm = int16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// bpl
func (fcEmu *CpuEmu) bpl() {
	zeroflag := fcEmu.Regi["P"] & 0b10000000
	if zeroflag == 0b00000000 {
		bufPC := int16(fcEmu.RegPc)
		minusflag := fcEmu.Memory[fcEmu.RegPc] & 0b10000000
		var imm int16
		if minusflag == 0b10000000 {
			var minussign uint16 = 0xff00
			imm = int16(fcEmu.Memory[fcEmu.RegPc]) + int16(minussign)
		} else {
			imm = int16(fcEmu.Memory[fcEmu.RegPc])
		}
		bufPC = bufPC + imm
		fcEmu.RegPc = uint16(bufPC)
	}
	fcEmu.RegPc++
}

// adc
func (fcEmu *CpuEmu) adcImm() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var ans uint16 = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu,uint8(ans), fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) adcZero() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	pos := fcEmu.Memory[fcEmu.RegPc]
	var ans uint16 = uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) adcZeroX() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc]) + uint16(fcEmu.Regi["X"])
	var ans uint16 = uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) adcAbs() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var ans uint16 = uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) adcAbsX() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	var ans uint16 = uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) adcAbsY() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["Y"])
	var ans uint16 = uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc+=2
}

// and 
func (fcEmu *CpuEmu) andImm(){
	fcEmu.Regi["A"] = fcEmu.Regi["A"] & fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) andAbs(){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Regi["A"] = fcEmu.Memory[pos] & fcEmu.Regi["A"]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) andAbsX(){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var xreg uint16 = uint16(fcEmu.Regi["X"])
	var pos uint16 = (absposhigh << 8) + absposlow + xreg
	fcEmu.Regi["A"] = fcEmu.Memory[pos] & fcEmu.Regi["A"]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// asl 
func (fcEmu *CpuEmu) aslImm() {
	var highbit uint8 = fcEmu.Regi["A"] & 0b10000000
	if highbit == 0x80 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = fcEmu.Regi["A"] << 1
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
func (fcEmu *CpuEmu) aslAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var highbit uint8 = fcEmu.Memory[pos] & 0b10000000
	if highbit == 0x80 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Memory[pos] = fcEmu.Memory[pos] << 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// bit
/*
// cmp は符号ありでこっちはなしで減算する
func (fcEmu *CpuEmu) bitZero() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var imm uint8 = fcEmu.Memory[pos]
	imm = ^imm + 1
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint16 = uint16(regA) + uint16(imm)
	if clflag == 0x00 {
		imm = 0xff
		ans = ans + uint16(imm)
	}
	nflagset(fcEmu, 0x80)
	zflagset(fcEmu, uint8(ans))
	vflagset(fcEmu,uint8(imm), regA, 0,false)
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) bitAbs() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var imm uint8 = fcEmu.Memory[pos]
	imm = ^imm + 1
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint16 = uint16(regA) + uint16(imm)
	if clflag == 0x00 {
		imm = 0xff
		ans = ans + uint16(imm)
	}
	nflagset(fcEmu, 0x80)
	zflagset(fcEmu, uint8(ans))
	vflagset(fcEmu,uint8(imm), regA, 0,false)
	fcEmu.RegPc+=2
}
*/
func (fcEmu *CpuEmu) bitZero() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint8 = regA & imm
	nflagset(fcEmu, ans)
	zflagset(fcEmu, ans)
	vflagset(fcEmu,0, 0, 0,true)
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) bitAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint8 = regA & imm
	nflagset(fcEmu, ans)
	zflagset(fcEmu, ans)
	vflagset(fcEmu,0, 0, 0,true)
	fcEmu.RegPc+=2
}

// cmp レジスタ-メモリ
// cmp はcflag を引かないでいい
func (fcEmu *CpuEmu) cmpImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regA uint8 = fcEmu.Regi["A"]
	imm = ^imm + 1
	var ans uint16 = uint16(regA) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) cmpZero() {
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	imm = ^imm + 1
	var ans uint16 = uint16(regA) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc+=1
}
func (fcEmu *CpuEmu) cmpAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	imm = ^imm + 1
	var ans uint16 = uint16(regA) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) cmpAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	var imm uint8 = fcEmu.Memory[pos]
	imm = ^imm + 1
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint16 = uint16(regA) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) cmpAbsY() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["Y"])
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	imm = ^imm + 1
	var ans uint16 = uint16(regA) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc+=2
}

// cpx レジスタ-メモリ
func (fcEmu *CpuEmu) cpxImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regX uint8 = fcEmu.Regi["X"]
	imm = ^imm + 1
	var ans uint16 = uint16(regX) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.RegPc++
}

// cpy レジスタ-メモリ
func (fcEmu *CpuEmu) cpyImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regY uint8 = fcEmu.Regi["Y"]
	imm = ^imm + 1
	var ans uint16 = uint16(regY) + uint16(imm)
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	
	fcEmu.RegPc++
}


// jmp
func (fcEmu *CpuEmu) jmpAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.RegPc = pos
}
func (fcEmu *CpuEmu) jmpInd() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var indpos uint16 = (absposhigh << 8) + absposlow
	lowpos := uint16(fcEmu.Memory[indpos])
	highpos := uint16(fcEmu.Memory[indpos+1])
	pos := (highpos << 8) + lowpos
	fcEmu.RegPc = pos
}


// jsr
func (fcEmu *CpuEmu) jsr() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	// スタックに現在のPC＋２を積む
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	// 0x4050 だったら　50 40 の順でスタックに積む。個々の動作は等で確認が取れていないためこのエミュレータだけの可能性あり。
	lowaddr := (fcEmu.RegPc+1) & 0x00ff
	highaddr := (fcEmu.RegPc+1) & 0xff00
	highaddr = highaddr >> 8
	fcEmu.Memory[sppos-1] = uint8(lowaddr)
	fcEmu.Memory[sppos] = uint8(highaddr)
	fcEmu.Regi["S"] = fcEmu.Regi["S"] - 2
	fcEmu.RegPc = pos
}

// rts
func (fcEmu *CpuEmu) rts(){
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont jsr when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	lowaddr := fcEmu.Memory[sppos]
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	sppos = 0x0100 + uint16(fcEmu.Regi["S"])
	highaddr := fcEmu.Memory[sppos]
	var Raddr uint16 = uint16(highaddr)
	Raddr = Raddr << 8
	Raddr += uint16(lowaddr)
	fcEmu.RegPc = Raddr + 1
}

// rti ※rstと違いフラッグ処理がたくさん入る
func (fcEmu *CpuEmu) rti(){
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont jsr when no push.")
	}
	// p reg 
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	fcEmu.Regi["P"] = fcEmu.Memory[sppos]
	// return addr
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	sppos = 0x0100 + uint16(fcEmu.Regi["S"])
	lowaddr := fcEmu.Memory[sppos]
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	sppos = 0x0100 + uint16(fcEmu.Regi["S"])
	highaddr := fcEmu.Memory[sppos]
	var Raddr uint16 = uint16(highaddr)
	Raddr = Raddr << 8
	Raddr += uint16(lowaddr)
	fcEmu.RegPc = Raddr
	fcEmu.InterruptFlag = false
}

// pha
func (fcEmu *CpuEmu) pha() {
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	fcEmu.Memory[sppos] = fcEmu.Regi["A"]
	fcEmu.Regi["S"] = fcEmu.Regi["S"] - 1
}

// pla
func (fcEmu *CpuEmu) pla() {
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont pla when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	fcEmu.Regi["A"] = fcEmu.Memory[sppos]
}

func (fcEmu *CpuEmu) clc() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111110
}

func (fcEmu *CpuEmu) cld() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11110111
}

func (fcEmu *CpuEmu) cli() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111011
}

func (fcEmu *CpuEmu) sec() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111110
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000001
}

func (fcEmu *CpuEmu) sei() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111011
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000100
}

func (fcEmu *CpuEmu) nop() {
}

func (fcEmu *CpuEmu) Execute() int {
	if fcEmu.RegPc == 0xffff{
		return 0
	}
	opcd := fcEmu.Memory[fcEmu.RegPc]
	var exeflag bool = true
	for _, e := range fcEmu.Exeopcdlist{
		if e == opcd {
			exeflag = false
		}
	}	
	if exeflag {
		fcEmu.Exeopcdlist = append(fcEmu.Exeopcdlist,opcd)
	}

	fcEmu.RegPc++
	switch opcd {
	// lda
	case LDAIMM:
		fcEmu.ldaImm()
	case LDAZERO:
		fcEmu.ldaZero()
	case LDAZEROX:
		fcEmu.ldaZeroX()
	case LDAABS:
		fcEmu.ldaAbs()
	case LDAABSX:
		fcEmu.ldaAbsX()
	case LDAABSY:
		fcEmu.ldaAbsY()
	case LDAINDX:
		fcEmu.ldaIndX()
	case LDAINDY:
		fcEmu.ldaIndY()
	// ldx
	case LDXIMM:
		fcEmu.ldxImm()
	case LDXZERO:
		fcEmu.ldxZero()
	case LDXZEROY:
		fcEmu.ldxZeroY()
	case LDXABS:
		fcEmu.ldxAbs()
	case LDXABSX:
		fcEmu.ldxAbsY()
	// ldy
	case LDYIMM:
		fcEmu.ldyImm()
	case LDYZERO:
		fcEmu.ldyZero()
	case LDYABS:
		fcEmu.ldyAbs()
	case LDYABSX:
		fcEmu.ldyAbsX()
	// sta
	case STAZERO:
		fcEmu.staZero()
	case STAZEROX:
		fcEmu.staZeroX()
	case STAABS:
		fcEmu.staAbs()
	case STAABSX:
		fcEmu.staAbsX()
	case STAABSY:
		fcEmu.staAbsY()
	case STAINDY:
		fcEmu.staIndY()
	// stx
	case STXZERO:
		fcEmu.stxZero()
	case STXABS:
		fcEmu.stxAbs()
	// sty
	case STYZERO:
		fcEmu.styZero()
	case STYABS:
		fcEmu.styAbs()
	// dec
	case DECZERO:
		fcEmu.decZero()
	case DECABS:
		fcEmu.decAbs()
	case DECABSX:
		fcEmu.decAbsX()
	// dex
	case DEX:
		fcEmu.deX()
	// dey
	case DEY:
		fcEmu.deY()
	// eor
	case EORIMM:
		fcEmu.eorImm()
	case EORZERO:
		fcEmu.eorZero()
	// inc
	case INCZERO:
		fcEmu.incZero()
	case INCABS:
		fcEmu.incAbs()
	// inx
	case INX:
		fcEmu.inX()
	// iny
	case INY:
		fcEmu.inY()
	// lsr
	case LSRIMM:
		fcEmu.lsrImm()
	case LSRZERO:
		fcEmu.lsrZero()
	case LSRABS:
		fcEmu.lsrAbs()
	// ora
	case ORAIMM:
		fcEmu.oraImm()
	case ORAZERO:
		fcEmu.oraZero()
	case ORAABS:
		fcEmu.oraAbs()
	// rol
	case ROLIMM:
		fcEmu.rolImm()
	case ROLZERO:
		fcEmu.rolZero()
	case ROLABS:
		fcEmu.rolAbs()
	// ror
	case RORIMM:
		fcEmu.rorImm()
	case RORABSX:
		fcEmu.rorAbsX()
	// sbc
	case SBCIMM:
		fcEmu.sbcImm()
	case SBCZERO:
		fcEmu.sbcZero()
	case SBCZEROX:
		fcEmu.sbcZeroX()
	case SBCABS:
		fcEmu.sbcAbs()
	case SBCABSY:
		fcEmu.sbcAbsY()
	// bcc
	case BCC:
		fcEmu.bcc()
	// bcs
	case BCS:
		fcEmu.bcs()
	// beq
	case BEQ:
		fcEmu.beq()
	// bne
	case BMI:
		fcEmu.bmi()
	// bne
	case BNE:
		fcEmu.bne()
	// bpl
	case BPL:
		fcEmu.bpl()
	// tax
	case TAX:
		fcEmu.tax()
	// tay
	case TAY:
		fcEmu.tay()
	// txa
	case TXA:
		fcEmu.txa()
	// txs
	case TXS:
		fcEmu.txs()
	// tya
	case TYA:
		fcEmu.tya()
	// adc
	case ADCIMM:
		fcEmu.adcImm()
	case ADCZERO:
		fcEmu.adcZero()
	case ADCZEROX:
		fcEmu.adcZeroX()
	case ADCABS:
		fcEmu.adcAbs()
	case ADCABSX:
		fcEmu.adcAbsX()
	case ADCABSY:
		fcEmu.adcAbsY()
	// and
	case ANDIMM:
		fcEmu.andImm()
	case ANDABS:
		fcEmu.andAbs()	
	case ANDABSX:
		fcEmu.andAbsX()	
	// asl
	case ASLIMM:
		fcEmu.aslImm()	
	case ASLABS:
		fcEmu.aslAbs()
	// bit
	case BITZERO:
		fcEmu.bitZero()
	case BITABS:
		fcEmu.bitAbs()	
	// cmp
	case CMPIMM:
		fcEmu.cmpImm()
	case CMPZERO:
		fcEmu.cmpZero()
	case CMPABS:
		fcEmu.cmpAbs()	
	case CMPABSX:
		fcEmu.cmpAbsX()
	case CMPABSY:
		fcEmu.cmpAbsY()
	// cpx
	case CPXIMM:
		fcEmu.cpxImm()
	// cpy
	case CPYIMM:
		fcEmu.cpyImm()
	// jmp
	case JMPABS:
		fcEmu.jmpAbs()
	case JMPIND:
		fcEmu.jmpInd()
	// jsr
	case JSR:
		fcEmu.jsr()
	// rts
	case RTS:
		fcEmu.rts()
	// rti
	case RTI:
		fcEmu.rti()
	// pha
	case PHA:
		fcEmu.pha()
	// pla
	case PLA:
		fcEmu.pla()
	// clc
	case CLC:
		fcEmu.clc()
	// cld
	case CLD:
		fcEmu.cld()
	// cli
	case CLI:
		fcEmu.cli()
	// sec 
	case SEC:
		fcEmu.sec()
	// sei 
	case SEI:
		fcEmu.sei()

	default:
		fmt.Printf("not exist function code 0x%x\n", opcd)
		panic("error execute")
		//return 1
	}
	return 0

}
