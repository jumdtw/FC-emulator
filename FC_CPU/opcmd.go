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

	// DEX :
	DEX = 0xca

	// DEY :
	DEY = 0x88

	//EORZERO :
	EORZERO = 0x45

	// INX :
	INX = 0xe8
	// INY :
	INY = 0xc8

	// BCS :
	BCS = 0xb0
	// BEQ :
	BEQ = 0xf0
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
	// ADCABS :
	ADCABS = 0x6d

	// ANDIMM : 
	ANDIMM = 0x29

	// ASLACC :
	ASLACC = 0x0a

	// BITABS :
	BITABS = 0x2c

	// CMPIMM :
	CMPIMM = 0xc9

	// CPX :
	CPXIMM = 0xe0

	// CPY :
	CPYIMM = 0xc0

	// JMP :
	JMP = 0x4c

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

	// SEI :
	SEI = 0x78
)

func ppuwrite(fcEmu *CpuEmu, pos uint16,reg string){
	// 0x2006だったらppuへのアクセス
	if pos == 0x2006 {
		fcEmu.VramAddr = fcEmu.VramAddr << 8
		fcEmu.VramAddr += uint16(fcEmu.Regi[reg])
	}
	if pos == 0x2007 {
		fcEmu.VramWriteFlag = true
		fcEmu.VramWriteValue = fcEmu.Regi[reg]
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
}

func padread(fcEmu *CpuEmu, pos uint16){
	// 0x4016 はパッドの読み込み
	if pos == 0x4016 {
		fcEmu.BotmReadcount++
		switch fcEmu.BotmReadcount {
		// A
		case 1:
			if fcEmu.Abotmflag {
				fcEmu.Abotmflag = false 
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// B
		case 2:
			if fcEmu.Bbotmflag {
				fcEmu.Bbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// Select
		case 3:
			if fcEmu.Selectbotmflag {
				fcEmu.Selectbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// Start
		case 4:
			if fcEmu.Startbotmflag {
				fcEmu.Startbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// up
		case 5:
			if fcEmu.Upbotmflag {
				fcEmu.Upbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// down
		case 6:
			if fcEmu.Downbotmflag {
				fcEmu.Downbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// left
		case 7:
			if fcEmu.Leftbotmflag {
				fcEmu.Leftbotmflag = false
				fcEmu.Memory[pos] = 1
			}else{
				fcEmu.Memory[pos] = 0
			}
		// right
		case 8:
			if fcEmu.Rightbotmflag {
				fcEmu.Rightbotmflag = false
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
	var absposlow uint8 = uint8(fcEmu.Memory[fcEmu.RegPc])
	var xreg uint8 = uint8(fcEmu.Regi["X"])
	var mempos uint16 = uint16(absposlow + xreg)
	var pos uint8 = fcEmu.Memory[mempos]
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) ldaIndY() {
	var mempos uint8 = uint8(fcEmu.Memory[fcEmu.RegPc])
	var pos uint8 = fcEmu.Memory[mempos]
	var yreg uint8 = uint8(fcEmu.Regi["Y"])
	fcEmu.Regi["A"] = fcEmu.Memory[pos] + yreg
	nflagset(fcEmu, fcEmu.Memory[pos] + yreg)
	zflagset(fcEmu, fcEmu.Memory[pos] + yreg)
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
	ppuwrite(fcEmu,pos,"A")
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staAbsX() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	pos = pos + uint16(fcEmu.Regi["X"])
	ppuwrite(fcEmu,pos,"A")
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staAbsY() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	pos = pos + uint16(fcEmu.Regi["Y"])
	ppuwrite(fcEmu,pos,"A")
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	fcEmu.RegPc = fcEmu.RegPc + 2
}
func (fcEmu *CpuEmu) staIndY() {
	var mempos uint8 = uint8(fcEmu.Memory[fcEmu.RegPc])
	var pos uint8 = fcEmu.Memory[mempos]
	var yreg uint8 = uint8(fcEmu.Regi["Y"])
	fcEmu.Memory[pos] = fcEmu.Regi["A"] + yreg
	nflagset(fcEmu, fcEmu.Memory[pos] + yreg)
	zflagset(fcEmu, fcEmu.Memory[pos] + yreg)
	fcEmu.RegPc++
}

// stx
func (fcEmu *CpuEmu) stxZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Memory[pos] = fcEmu.Regi["X"]
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) stxAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	ppuwrite(fcEmu,pos,"X")
	fcEmu.Memory[pos] = fcEmu.Regi["X"]
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
func (fcEmu *CpuEmu) eorZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Regi["A"] ^ fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
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
	if zeroflag != 0b10000000 {
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

// adc
func (fcEmu *CpuEmu) adcImm() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	imm := fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = imm + fcEmu.Regi["A"] + clflag
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	cflagset(fcEmu, uint16(imm) + uint16(fcEmu.Regi["A"]) + uint16(clflag))
	vflagset(fcEmu,imm, fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) adcZero() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	pos := fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Memory[pos] + fcEmu.Regi["A"] + clflag
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	cflagset(fcEmu, uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag))
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) adcAbs() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.Regi["A"] = fcEmu.Memory[pos] + fcEmu.Regi["A"] + clflag
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	cflagset(fcEmu, uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag))
	vflagset(fcEmu, fcEmu.Memory[pos], fcEmu.Regi["A"], clflag,true)
	fcEmu.RegPc+=2
}

// and 
func (fcEmu *CpuEmu) andImm(){
	fcEmu.Regi["A"] = fcEmu.Regi["A"] & fcEmu.Memory[fcEmu.RegPc]
	fcEmu.RegPc++
}

// asl 
func (fcEmu *CpuEmu) aslAcc() {
	fcEmu.Regi["A"] = fcEmu.Regi["A"] >> 1
}

// bit
func (fcEmu *CpuEmu) bitAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var imm uint8 = fcEmu.Memory[pos]
	var regA uint8 = fcEmu.Regi["A"]
	if imm > regA {
		nflagset(fcEmu, 0b10000000)
	}
	if imm == regA {
		zflagset(fcEmu, 0)
	}
	vflagset(fcEmu,imm, regA, 0,false)
	fcEmu.RegPc+=2
}

// cmp レジスタ-メモリ
func (fcEmu *CpuEmu) cmpImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regA uint8 = fcEmu.Regi["A"]
	if imm > regA {
		nflagset(fcEmu, 0b10000000)
		cflagset(fcEmu, 0x0100)
	}
	if imm == regA {
		zflagset(fcEmu, 0)
	}
	fcEmu.RegPc++
}

// cpx レジスタ-メモリ
func (fcEmu *CpuEmu) cpxImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regX uint8 = fcEmu.Regi["X"]
	if imm > regX {
		nflagset(fcEmu, 0b10000000)
		cflagset(fcEmu, 0x0100)
	}
	if imm == regX {
		zflagset(fcEmu, 0)
	}
	fcEmu.RegPc++
}

// cpy レジスタ-メモリ
func (fcEmu *CpuEmu) cpyImm() {
	var imm uint8 = fcEmu.Memory[fcEmu.RegPc]
	var regY uint8 = fcEmu.Regi["Y"]
	if imm > regY {
		nflagset(fcEmu, 0b10000000)
		cflagset(fcEmu, 0x0100)
	}
	if imm == regY {
		zflagset(fcEmu, 0)
	}
	fcEmu.RegPc++
}


// jmp
func (fcEmu *CpuEmu) jmp() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.RegPc = pos
}


// jsr
func (fcEmu *CpuEmu) jsr() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	// スタックに現在のPCを積む
	var sppos uint16 = 0x1000 + uint16(fcEmu.Regi["S"])
	// 0x4050 だったら　50 40 の順でスタックに積む。個々の動作は等で確認が取れていないためこのエミュレータだけの可能性あり。
	lowaddr := fcEmu.RegPc & 0x00ff
	lowaddr = lowaddr - 1
	highaddr := fcEmu.RegPc & 0xff00
	highaddr = highaddr >> 8
	fcEmu.Memory[sppos] = uint8(lowaddr)
	fcEmu.Memory[sppos-1] = uint8(highaddr)
	fcEmu.Regi["S"] = fcEmu.Regi["S"] - 2
	fcEmu.RegPc = pos
}

// rts
func (fcEmu *CpuEmu) rts(){
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont jsr when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x1000 + uint16(fcEmu.Regi["S"])
	highaddr := fcEmu.Memory[sppos]
	var Raddr uint16 = uint16(highaddr)
	Raddr = Raddr << 8
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	sppos = 0x1000 + uint16(fcEmu.Regi["S"])
	lowaddr := fcEmu.Memory[sppos]
	Raddr += uint16(lowaddr)
	fcEmu.RegPc = Raddr
}

// rti ※rstと違いフラッグ処理がたくさん入る
func (fcEmu *CpuEmu) rti(){
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont jsr when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x1000 + uint16(fcEmu.Regi["S"])
	highaddr := fcEmu.Memory[sppos]
	var Raddr uint16 = uint16(highaddr)
	Raddr = Raddr << 8
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	sppos = 0x1000 + uint16(fcEmu.Regi["S"])
	lowaddr := fcEmu.Memory[sppos]
	Raddr += uint16(lowaddr)
	fcEmu.RegPc = Raddr
	fcEmu.InterruptFlag = false
}

// pha
func (fcEmu *CpuEmu) pha() {
	var sppos uint16 = 0x1000 + uint16(fcEmu.Regi["S"])
	fcEmu.Memory[sppos] = fcEmu.Regi["A"]
	fcEmu.Regi["S"] = fcEmu.Regi["S"] - 1
}

// pla
func (fcEmu *CpuEmu) pla() {
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont pla when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x1000 + uint16(fcEmu.Regi["S"])
	fcEmu.Regi["A"] = fcEmu.Memory[sppos]
}

func (fcEmu *CpuEmu) clc() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111110
	cflagset(fcEmu,0)
}

func (fcEmu *CpuEmu) cld() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11110111
}

func (fcEmu *CpuEmu) cli() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111011
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
	// dex
	case DEX:
		fcEmu.deX()
	// dey
	case DEY:
		fcEmu.deY()
	// eor
	case EORZERO:
		fcEmu.eorZero()
	// inx
	case INX:
		fcEmu.inX()
	// iny
	case INY:
		fcEmu.inY()
	// bcs
	case BCS:
		fcEmu.bcs()
	// beq
	case BEQ:
		fcEmu.beq()
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
	case ADCABS:
		fcEmu.adcAbs()
	// and
	case ANDIMM:
		fcEmu.andImm()	
	// asl
	case ASLACC:
		fcEmu.aslAcc()	
	// bit
	case BITABS:
		fcEmu.bitAbs()	
	// cmp
	case CMPIMM:
		fcEmu.cmpImm()	
	// cpx
	case CPXIMM:
		fcEmu.cpxImm()
	// cpy
	case CPYIMM:
		fcEmu.cpyImm()
	// jmp
	case JMP:
		fcEmu.jmp()
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
	// sei 
	case SEI:
		fcEmu.sei()

	default:
		fmt.Printf("not exist function code 0x%x\n", opcd)
		fcEmu.RegPc = fcEmu.RegPc - 1
		panic("error execute")
		//return 1
	}
	return 0

}
