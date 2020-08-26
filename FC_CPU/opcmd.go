package FC_CPU

import (
	"fmt"
	//"os"
)
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
	// LDXABSY :
	LDXABSY = 0xbe

	// LDYIMM:
	LDYIMM = 0xa0
	// LDYZERO:
	LDYZERO = 0xa4
	// LDYZEROX:
	LDYZEROX = 0xb4
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
	// STAINDX:
	STAINDX = 0x81
	// STAINDY:
	STAINDY = 0x91

	// STXZERO :
	STXZERO = 0x86
	// STXZEROY :
	STXZEROY = 0x96
	// STXABS :
	STXABS = 0x8e

	// STYZERO :
	STYZERO = 0x84
	// STYZEROX :
	STYZEROX = 0x94
	// STYABS :
	STYABS = 0x8c

	// DECZERO :
	DECZERO = 0xc6
	// DECZEROX :
	DECZEROX = 0xd6
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
	//EORZEROX :
	EORZEROX = 0x55
	//EORABS :
	EORABS = 0x4d
	//EORABSX :
	EORABSX = 0x5d
	//EORABSY :
	EORABSY = 0x59
	//EORINDX :
	EORINDX = 0x41
	//EORINDY :
	EORINDY = 0x51

	// INCZERO :
	INCZERO = 0xe6
	// INCZEROX :
	INCZEROX = 0xf6
	// INCABS :
	INCABS = 0xee
	// INCABSX :
	INCABSX = 0xfe

	// INX :
	INX = 0xe8
	// INY :
	INY = 0xc8

	// LSRIMM :
	LSRIMM = 0x4a
	// LSRZERO :
	LSRZERO = 0x46
	// LSRZEROX :
	LSRZEROX = 0x56
	// LSRABS :
	LSRABS = 0x4e
	// LSRABSX :
	LSRABSX = 0x5e

	// ORAIMM :
	ORAIMM = 0x09
	// ORAZERO :
	ORAZERO = 0x05
	// ORAZEROX :
	ORAZEROX = 0x15
	// ORAABS :
	ORAABS = 0x0d
	// ORAABSX :
	ORAABSX = 0x1d
	// ORAABSY :
	ORAABSY = 0x19
	// ORAINDX :
	ORAINDX = 0x01
	// ORAINDY :
	ORAINDY = 0x11

	// ROLIMM :
	ROLIMM = 0x2a
	// ROLZERO :
	ROLZERO = 0x26
	// ROLZEROX :
	ROLZEROX = 0x36
	// ROLABS :
	ROLABS = 0x2e
	// ROLABSX :
	ROLABSX = 0x3e


	// RORIMM :
	RORIMM = 0x6a
	// RORZERO :
	RORZERO = 0x66
	// RORZEROX :
	RORZEROX = 0x76
	// RORABS :
	RORABS = 0x6e
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
	// SBCABSX :
	SBCABSX = 0xfd
	// SBCABSY :
	SBCABSY = 0xf9
	// SBCINDX :
	SBCINDX = 0xe1
	// SBCINDY :
	SBCINDY = 0xf1

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
	// BVC :
	BVC = 0x50
	// BVS :
	BVS = 0x70

	// TAX :
	TAX = 0xaa

	// TAY :
	TAY = 0xa8

	// TSX :
	TSX = 0xba

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
	// ADCINDX :
	ADCINDX = 0x61
	// ADCINDY :
	ADCINDY = 0x71

	// ANDIMM : 
	ANDIMM = 0x29
	// ANDZERO : 
	ANDZERO = 0x25
	// ANDZEROX : 
	ANDZEROX = 0x35
	// ANDABS : 
	ANDABS = 0x2d
	// ANDABSX : 
	ANDABSX = 0x3d
	// ANDABSY : 
	ANDABSY = 0x39
	// ANDINDX : 
	ANDINDX = 0x21
	// ANDINDY : 
	ANDINDY = 0x31

	// ASLIMM :
	ASLIMM = 0x0a
	// ASLZERO :
	ASLZERO = 0x06
	// ASLZEROX :
	ASLZEROX = 0x16
	// ASLABS :
	ASLABS = 0x0e
	// ASLABSX :
	ASLABSX = 0x1e

	// BITZERO :
	BITZERO = 0x24
	// BITABS :
	BITABS = 0x2c

	// CMPIMM :
	CMPIMM = 0xc9
	// CMPZERO :
	CMPZERO = 0xc5
	// CMPZEROX :
	CMPZEROX = 0xd5
	// CMPABS :
	CMPABS = 0xcd
	// CMPABSX :
	CMPABSX = 0xdd
	// CMPABSY :
	CMPABSY = 0xd9
	// CMPINDX :
	CMPINDX = 0xc1
	// CMPINDY :
	CMPINDY = 0xd1

	// CPXIMM :
	CPXIMM = 0xe0
	// CPXZERO :
	CPXZERO = 0xe4
	// CPXABS :
	CPXABS = 0xec

	// CPYIMM :
	CPYIMM = 0xc0
	// CPYZERO :
	CPYZERO = 0xc4
	// CPYABS :
	CPYABS = 0xcc

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
	// PHP :
	PHP = 0x08
	// PLA :
	PLA = 0x68
	// PLP :
	PLP = 0x28

	// CLC :
	CLC = 0x18

	// CLD :
	CLD = 0xd8

	// CLI : 
	CLI = 0x58

	// CLV : 
	CLV = 0xb8

	// SEC :
	SEC = 0x38

	// SED :
	SED = 0xf8

	// SEI :
	SEI = 0x78

	// BRK :
	BRK = 0x00

	// NOP :
	NOP = 0xea

	// unofficial code
	// nop
	NOP1A = 0x1a
	NOP3A = 0x3a
	NOP5A = 0x5a
	NOP7A = 0x7a
	NOPDA = 0xda
	NOPFA = 0xfa
	// nopimm
	NOP80 = 0x80
	// nopzero
	NOP4 = 0x4 
	NOP44 = 0x44
	NOP64 = 0x64
	// nopzerox
	NOP14 = 0x14
	NOP34 = 0x34
	NOP54 = 0x54
	NOP74 = 0x74
	NOPD4 = 0xd4
	NOPF4 = 0xf4
	// nop abs
	NOPC = 0x0c
	// nop absx
	NOP1C = 0x1c
	NOP3C = 0x3c
	NOP5C = 0x5c
	NOP7C = 0x7c
	NOPDC = 0xdc
	NOPFC = 0xfc
	// lax 
	LAXZERO = 0xa7
	LAXZEROY = 0xb7
	LAXABS = 0xaf
	LAXABSY = 0xbf
	LAXINDXA3 = 0xa3
	LAXINDY = 0xb3
	// sax
	SAXZERO = 0x87
	SAXZEROY = 0x97
	SAXABS = 0x8f
	SAXINDX = 0x83
	// sbc
	SBCIMMEB = 0xeb
	// dcp
	DCPZERO = 0xc7
	DCPZEROX = 0xd7
	DCPABS = 0xcf
	DCPABSX = 0xdf
	DCPABSY = 0xdb
	DCPINDX = 0xc3
	DCPINDY = 0xd3
	// isc
	ISCZERO = 0xe7
	ISCZEROX = 0xf7
	ISCABS = 0xef
	ISCABSX = 0xff
	ISCABSY = 0xfb
	ISCINDX = 0xe3
	ISCINDY = 0xf3
	// slo
	SLOZERO = 0x07
	SLOZEROX = 0x17
	SLOABS = 0x0f
	SLOABSX = 0x1f
	SLOABSY = 0x1b
	SLOINDX = 0x03
	SLOINDY = 0x13
	// rla
	RLAZERO = 0x27
	RLAZEROX = 0x37
	RLAABS = 0x2f
	RLAABSX = 0x3f
	RLAABSY = 0x3b
	RLAINDX = 0x23
	RLAINDY = 0x33
	// sre
	SREZERO = 0x47
	SREZEROX = 0x57
	SREABS = 0x4f
	SREABSX = 0x5f
	SREABSY = 0x5b
	SREINDX = 0x43
	SREINDY = 0x53
	// rra
	RRAZERO = 0x67
	RRAZEROX = 0x77
	RRAABS = 0x6f
	RRAABSX = 0x7f
	RRAABSY = 0x7b
	RRAINDX = 0x63
	RRAINDY = 0x73

)

func addrImm(fcEmu *CpuEmu)(uint16){
	var pos uint16 = uint16(fcEmu.RegPc)
	fcEmu.RegPc+=1
	return pos
}

func addrZero(fcEmu *CpuEmu)(uint16){
	var pos uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	fcEmu.RegPc+=1
	return pos
}

func addrZeroX(fcEmu *CpuEmu)(uint16){
	var immpos uint8 = fcEmu.Memory[fcEmu.RegPc]
	var pos uint8 = fcEmu.Regi["X"] + immpos
	fcEmu.RegPc+=1
	return uint16(pos)
}

func addrZeroY(fcEmu *CpuEmu)(uint16){
	var immpos uint8 = fcEmu.Memory[fcEmu.RegPc]
	var pos uint8 = fcEmu.Regi["Y"] + immpos
	fcEmu.RegPc+=1
	return uint16(pos)
}

func addrAbs(fcEmu *CpuEmu)(uint16){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.RegPc+=2
	return pos
}

func addrAbsX(fcEmu *CpuEmu)(uint16){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["X"])
	fcEmu.RegPc+=2
	return pos
}

func addrAbsY(fcEmu *CpuEmu)(uint16){
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow + uint16(fcEmu.Regi["Y"])
	fcEmu.RegPc+=2
	return pos
}

func addrIndX(fcEmu *CpuEmu)(uint16){
	var xreg uint8 = fcEmu.Regi["X"]
	var immaddr uint8 = fcEmu.Memory[fcEmu.RegPc] + xreg
	var highpos uint16 = uint16(fcEmu.Memory[immaddr+1])
	var lowpos uint16 = uint16(fcEmu.Memory[immaddr])
	var pos uint16 = (highpos << 8) + lowpos
	fcEmu.RegPc+=1
	return pos
}

func addrIndY(fcEmu *CpuEmu)(uint16){
	var immaddr uint8 = fcEmu.Memory[fcEmu.RegPc]
	var yreg uint16 = uint16(fcEmu.Regi["Y"])
	var highpos uint16 = uint16(fcEmu.Memory[immaddr+1])
	var lowpos uint16 = uint16(fcEmu.Memory[immaddr])
	var pos uint16 = (highpos << 8) + lowpos + yreg
	fcEmu.RegPc+=1
	return pos
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
		//fmt.Printf("PC : 0x%x, regx : 0x%x, regy : 0x%x, vram write : 0x%x\n", fcEmu.RegPc-1 ,fcEmu.Regi["X"] ,fcEmu.Regi["Y"] ,fcEmu.Regi[reg])
	}else {
		fcEmu.VramWriteFlag = false
	}
	// 0x2004 0x2003だったらoamへの
	if pos == 0x2003 {
		fcEmu.Oamnum = int(fcEmu.Regi[reg])
		fcEmu.OamWritecount = 0
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
		fcEmu.Padvalue += uint16(fcEmu.Regi[reg]) & 0x0007
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

// setflag == 0 : calc, setflag == 1 : set, setflag == 2 : clear
func vflagset(fcEmu *CpuEmu, x uint8, y uint8, c uint8, plusflag bool,setflag int){

	// 計算しない命令があるので無条件でフラッグ建てるやつ
	if setflag == 1 {
		fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
		fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
		return
	} else if setflag == 2 {
		fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
		return
	}

	xnflag := x & 0b10000000
	ynflag := y & 0b10000000

	// 足し算の場合
	// オーバーフローは符号が同じじゃないと起きない
	// 正の値
	if plusflag {
		// オバフが起きるのは正＋正か負＋負
		var plus uint16 = uint16(x) + uint16(y) + uint16(c)
		plus = plus & 0x80
		if xnflag == 0x0 && ynflag == 0x0 {
			if plus == 0x80 {
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
			}else{

				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
			}
		} else if xnflag == 0x80 && ynflag == 0x80 {
			if plus == 0x0 {
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
			}else{
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
			}
		} else {
			fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
		}
	}

	//引き算の場合
	if !plusflag {
		// オバフが起きるのは正＋正か負＋負
		y = ^y
		var bufy uint16 = uint16(y) + 1
		var plus uint16 = uint16(x) + bufy + uint16(c)
		plus = plus & 0x80
		if xnflag == 0x0 && ynflag == 0x80 {
			if plus == 0x80 {
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
			}else{
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
			}
		} else if xnflag == 0x80 && ynflag == 0x0 {
			if plus == 0x0 {
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
				fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b01000000
			}else{
				fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
			}
		} else {
			fcEmu.Regi["P"] = 0b10111111 & fcEmu.Regi["P"]
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
func (fcEmu *CpuEmu) lda(pos uint16) {
	padread(fcEmu,pos)
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	if fcEmu.RegPc == 0xc08b {
		//fmt.Printf("0x%x\n",fcEmu.Memory[pos])
	}
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
// LDX
func (fcEmu *CpuEmu) ldx(pos uint16) {
	padread(fcEmu,pos)
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
// LDY
func (fcEmu *CpuEmu) ldy(pos uint16) {
	padread(fcEmu,pos)
	fcEmu.Regi["Y"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
// sta
// 6502 はメモリマップトioなのでstaだけ命令処理以外に色々処理が入る場合が多い
func (fcEmu *CpuEmu) sta(pos uint16) {
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	ppuwrite(fcEmu,pos,"A")
}
// stx
func (fcEmu *CpuEmu) stx(pos uint16) {
	fcEmu.Memory[pos] = fcEmu.Regi["X"]
	ppuwrite(fcEmu,pos,"X")
}
// sty
func (fcEmu *CpuEmu) sty(pos uint16) {
	fcEmu.Memory[pos] = fcEmu.Regi["Y"]
	ppuwrite(fcEmu,pos,"Y")
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
// tsx
func (fcEmu *CpuEmu) tsx() {
	fcEmu.Regi["X"] = fcEmu.Regi["S"]
	nflagset(fcEmu, fcEmu.Regi["X"])
	zflagset(fcEmu, fcEmu.Regi["X"])
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
}
// tya
func (fcEmu *CpuEmu) tya() {
	fcEmu.Regi["A"] = fcEmu.Regi["Y"]
	nflagset(fcEmu, fcEmu.Regi["Y"])
	zflagset(fcEmu, fcEmu.Regi["Y"])
}
// dec
func (fcEmu *CpuEmu) dec(pos uint16) {
	fcEmu.Memory[pos] = fcEmu.Memory[pos] - 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
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
func (fcEmu *CpuEmu) eor(pos uint16) {
	fcEmu.Regi["A"] = fcEmu.Regi["A"] ^ fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
// inc
func (fcEmu *CpuEmu) inc(pos uint16) {
	fcEmu.Memory[pos] = fcEmu.Memory[pos] + 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
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
func (fcEmu *CpuEmu) lsr(pos uint16) {
	lowbit := fcEmu.Memory[pos] & 0x01
	if lowbit == 0x01 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Memory[pos] = fcEmu.Memory[pos] >> 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
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
// ora 
func (fcEmu *CpuEmu) ora(pos uint16){
	fcEmu.Regi["A"] = fcEmu.Regi["A"] | fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
func (fcEmu *CpuEmu) oraImm(){
	fcEmu.Regi["A"] = fcEmu.Regi["A"] | fcEmu.Memory[fcEmu.RegPc]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	fcEmu.RegPc++
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
func (fcEmu *CpuEmu) rol(pos uint16){
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
func (fcEmu *CpuEmu) ror(pos uint16){
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
}
// sbc
func (fcEmu *CpuEmu) sbc(pos uint16) {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var memo uint16 = uint16(fcEmu.Memory[pos])
	var regA uint8 = fcEmu.Regi["A"]
	memo = ^memo
	memo = 0x00ff & memo
	memo = memo + 1
	var ans uint16 = uint16(regA) + uint16(memo)
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], ^clflag,false,0)
	if clflag == 0x00 {
		memo = 0xff
		ans = ans + uint16(memo)	
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
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

// bvc
func (fcEmu *CpuEmu) bvc() {
	zeroflag := fcEmu.Regi["P"] & 0b01000000
	if zeroflag == 0 {
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

// bvs
func (fcEmu *CpuEmu) bvs() {
	zeroflag := fcEmu.Regi["P"] & 0b01000000
	if zeroflag == 0b01000000 {
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
func (fcEmu *CpuEmu) adc(pos uint16) {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	var ans uint16 = uint16(fcEmu.Regi["A"]) + uint16(fcEmu.Memory[pos]) + uint16(clflag)
	vflagset(fcEmu, fcEmu.Regi["A"], fcEmu.Memory[pos], clflag,true,0)
	if ans > 0xff {
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Regi["A"] = uint8(ans)
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
// and 
func (fcEmu *CpuEmu) and(pos uint16){
	fcEmu.Regi["A"] = fcEmu.Memory[pos] & fcEmu.Regi["A"]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}
// asl 
func (fcEmu *CpuEmu) asl(pos uint16) {
	var highbit uint8 = fcEmu.Memory[pos] & 0b10000000
	if highbit == 0x80 {
		cflagset(fcEmu, 0x0100)
	} else {
		cflagset(fcEmu, 0)
	}
	fcEmu.Memory[pos] = fcEmu.Memory[pos] << 1
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
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
// bit
func (fcEmu *CpuEmu) bit(pos uint16) {
	var imm uint8 = fcEmu.Memory[pos]
	nflagset(fcEmu, imm)
	var buf uint8 = imm & 0b01000000
	if buf == 0b01000000 {
		vflagset(fcEmu,0xff, 0xff, 0xff,true,1)
	}else{
		vflagset(fcEmu,0, 0, 0,true,2)
	}
	var regA uint8 = fcEmu.Regi["A"]
	var ans uint8 = regA & imm
	zflagset(fcEmu, ans)
}
// cmp レジスタ-メモリ
// 符号付き
// cmp はcflag を引かないでいい
func (fcEmu *CpuEmu) cmp(pos uint16) {
	var imm uint16 = uint16(fcEmu.Memory[pos])
	var regA uint8 = fcEmu.Regi["A"]
	imm = ^imm
	imm = 0x00ff & imm
	imm = imm + 1
	var ans uint16 = uint16(regA) + imm
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
}
// cpx レジスタ-メモリ
func (fcEmu *CpuEmu) cpx(pos uint16) {
	var imm uint16 = uint16(fcEmu.Memory[pos])
	var regX uint8 = fcEmu.Regi["X"]
	imm = ^imm
	imm = 0x00ff & imm
	imm = imm + 1
	var ans uint16 = uint16(regX) + imm
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
}
// cpy レジスタ-メモリ
func (fcEmu *CpuEmu) cpy(pos uint16) {
	var imm uint16 = uint16(fcEmu.Memory[pos])
	var regY uint8 = fcEmu.Regi["Y"]
	imm = ^imm
	imm = 0x00ff & imm
	imm = imm + 1
	var ans uint16 = uint16(regY) + imm
	if ans >= 0x100{
		cflagset(fcEmu, 0x100)
	} else {
		cflagset(fcEmu, 0)
	}
	nflagset(fcEmu, uint8(ans))
	zflagset(fcEmu, uint8(ans))
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
	var absposlow uint8 = fcEmu.Memory[fcEmu.RegPc]
	// こんなんしるわけないやろ
	// 0x2ffの次は0x200になるっぽい
	var lowindpos uint16 = (absposhigh << 8) + uint16(absposlow)
	var highindpos uint16 = (absposhigh << 8) + uint16(absposlow + 1)
	lowpos := uint16(fcEmu.Memory[lowindpos])
	highpos := uint16(fcEmu.Memory[highindpos])
	pos := (highpos << 8) + lowpos
	fcEmu.RegPc = pos
}


// jsr
func (fcEmu *CpuEmu) jsr() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
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
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11011111
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00100000
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

// php
func (fcEmu *CpuEmu) php() {
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	var buf uint8 = fcEmu.Regi["P"]
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11101111
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00010000
	fcEmu.Memory[sppos] = fcEmu.Regi["P"]
	fcEmu.Regi["P"] = buf
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
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
}

// plp
func (fcEmu *CpuEmu) plp() {
	if fcEmu.Regi["S"] > 0xff {
		panic("stack error: dont pla when no push.")
	}
	fcEmu.Regi["S"] = fcEmu.Regi["S"] + 1
	var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
	fcEmu.Regi["P"] = fcEmu.Memory[sppos] & 0b11101111
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11011111
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00100000
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

func (fcEmu *CpuEmu) clv() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b10111111
}

func (fcEmu *CpuEmu) sec() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111110
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000001
}

func (fcEmu *CpuEmu) sed() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11110111
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00001000
}

func (fcEmu *CpuEmu) sei() {
	fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11111011
	fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00000100
}

func (fcEmu *CpuEmu) brk() {
	var flag uint8 = fcEmu.Regi["P"] & 0b00000100
	if flag == 0{
		fcEmu.InterruptFlag = true
		fcEmu.Regi["P"] = fcEmu.Regi["P"] & 0b11101111
		fcEmu.Regi["P"] = fcEmu.Regi["P"] + 0b00010000
		var sppos uint16 = 0x0100 + uint16(fcEmu.Regi["S"])
		// 0x4050 だったら　50 40 の順でスタックに積む。個々の動作は等で確認が取れていないためこのエミュレータだけの可能性あり。
		lowaddr := fcEmu.RegPc & 0x00ff
		highaddr := fcEmu.RegPc & 0xff00
		highaddr = highaddr >> 8
		fcEmu.Memory[sppos-2] = fcEmu.Regi["P"]
		fcEmu.Memory[sppos-1] = uint8(lowaddr)
		fcEmu.Memory[sppos] = uint8(highaddr)
		fcEmu.Regi["S"] = fcEmu.Regi["S"] - 3
		// 割り込みアドレスの代入
		fcEmu.RegPc = fcEmu.Irqaddr
	}
}

func (fcEmu *CpuEmu) nop() {
}

// unofficial
func (fcEmu *CpuEmu) nopimm(){
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) nopzero(){
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) nopzerox(){
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) nopabs(){
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) nopabsx(){
	fcEmu.RegPc+=2
}
func (fcEmu *CpuEmu) lax(pos uint16){
	fcEmu.Regi["A"] = fcEmu.Memory[pos]
	fcEmu.Regi["X"] = fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Memory[pos])
	zflagset(fcEmu, fcEmu.Memory[pos])
}
func (fcEmu *CpuEmu) sax(pos uint16){
	fcEmu.Memory[pos] = fcEmu.Regi["A"] & fcEmu.Regi["X"]
}
func (fcEmu *CpuEmu) dcp(pos uint16){
	fcEmu.dec(pos)
	fcEmu.cmp(pos)
}
func (fcEmu *CpuEmu) isc(pos uint16){
	fcEmu.inc(pos)
	fcEmu.sbc(pos)
}
func (fcEmu *CpuEmu) slo(pos uint16){
	fcEmu.asl(pos)
	fcEmu.ora(pos)
}
func (fcEmu *CpuEmu) rla(pos uint16){
	fcEmu.rol(pos)
	fcEmu.and(pos)
}
func (fcEmu *CpuEmu) sre(pos uint16){
	fcEmu.lsr(pos)
	fcEmu.eor(pos)
}
func (fcEmu *CpuEmu) rra(pos uint16){
	fcEmu.ror(pos)
	fcEmu.adc(pos)
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
		pos := addrImm(fcEmu)
		fcEmu.lda(pos)
	case LDAZERO:
		pos := addrZero(fcEmu)
		fcEmu.lda(pos)
	case LDAZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.lda(pos)
	case LDAABS:
		pos := addrAbs(fcEmu)
		fcEmu.lda(pos)
	case LDAABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.lda(pos)
	case LDAABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.lda(pos)
	case LDAINDX:
		pos := addrIndX(fcEmu)
		fcEmu.lda(pos)
	case LDAINDY:
		pos := addrIndY(fcEmu)
		fcEmu.lda(pos)
	// ldx
	case LDXIMM:
		pos := addrImm(fcEmu)
		fcEmu.ldx(pos)
	case LDXZERO:
		pos := addrZero(fcEmu)
		fcEmu.ldx(pos)
	case LDXZEROY:
		pos := addrZeroY(fcEmu)
		fcEmu.ldx(pos)
	case LDXABS:
		pos := addrAbs(fcEmu)
		fcEmu.ldx(pos)
	case LDXABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.ldx(pos)
	// ldy
	case LDYIMM:
		pos := addrImm(fcEmu)
		fcEmu.ldy(pos)
	case LDYZERO:
		pos := addrZero(fcEmu)
		fcEmu.ldy(pos)
	case LDYZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.ldy(pos)
	case LDYABS:
		pos := addrAbs(fcEmu)
		fcEmu.ldy(pos)
	case LDYABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.ldy(pos)
	// sta
	case STAZERO:
		pos := addrZero(fcEmu)
		fcEmu.sta(pos)
	case STAZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.sta(pos)
	case STAABS:
		pos := addrAbs(fcEmu)
		fcEmu.sta(pos)
	case STAABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.sta(pos)
	case STAABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.sta(pos)
	case STAINDX:
		pos := addrIndX(fcEmu)
		fcEmu.sta(pos)
	case STAINDY:
		pos := addrIndY(fcEmu)
		fcEmu.sta(pos)
	// stx
	case STXZERO:
		pos := addrZero(fcEmu)
		fcEmu.stx(pos)
	case STXZEROY:
		pos := addrZeroY(fcEmu)
		fcEmu.stx(pos)
	case STXABS:
		pos := addrAbs(fcEmu)
		fcEmu.stx(pos)
	// sty
	case STYZERO:
		pos := addrZero(fcEmu)
		fcEmu.sty(pos)
	case STYZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.sty(pos)
	case STYABS:
		pos := addrAbs(fcEmu)
		fcEmu.sty(pos)
	// dec
	case DECZERO:
		pos := addrZero(fcEmu)
		fcEmu.dec(pos)
	case DECZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.dec(pos)
	case DECABS:
		pos := addrAbs(fcEmu)
		fcEmu.dec(pos)
	case DECABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.dec(pos)
	// dex
	case DEX:
		fcEmu.deX()
	// dey
	case DEY:
		fcEmu.deY()
	// eor
	case EORIMM:
		pos := addrImm(fcEmu)
		fcEmu.eor(pos)
	case EORZERO:
		pos := addrZero(fcEmu)
		fcEmu.eor(pos)
	case EORZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.eor(pos)
	case EORABS:
		pos := addrAbs(fcEmu)
		fcEmu.eor(pos)
	case EORABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.eor(pos)
	case EORABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.eor(pos)
	case EORINDX:
		pos := addrIndX(fcEmu)
		fcEmu.eor(pos)
	case EORINDY:
		pos := addrIndY(fcEmu)
		fcEmu.eor(pos)
	// inc
	case INCZERO:
		pos := addrZero(fcEmu)
		fcEmu.inc(pos)
	case INCZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.inc(pos)
	case INCABS:
		pos := addrAbs(fcEmu)
		fcEmu.inc(pos)
	case INCABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.inc(pos)
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
		pos := addrZero(fcEmu)
		fcEmu.lsr(pos)
	case LSRZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.lsr(pos)
	case LSRABS:
		pos := addrAbs(fcEmu)
		fcEmu.lsr(pos)
	case LSRABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.lsr(pos)
	// ora
	case ORAIMM:
		fcEmu.oraImm()
	case ORAZERO:
		pos := addrZero(fcEmu)
		fcEmu.ora(pos)
	case ORAZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.ora(pos)
	case ORAABS:
		pos := addrAbs(fcEmu)
		fcEmu.ora(pos)
	case ORAABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.ora(pos)
	case ORAABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.ora(pos)
	case ORAINDX:
		pos := addrIndX(fcEmu)
		fcEmu.ora(pos)
	case ORAINDY:
		pos := addrIndY(fcEmu)
		fcEmu.ora(pos)
	// rol
	case ROLIMM:
		fcEmu.rolImm()
	case ROLZERO:
		pos := addrZero(fcEmu)
		fcEmu.rol(pos)
	case ROLZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.rol(pos)
	case ROLABS:
		pos := addrAbs(fcEmu)
		fcEmu.rol(pos)
	case ROLABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.rol(pos)
	// ror
	case RORIMM:
		fcEmu.rorImm()
	case RORZERO:
		pos := addrZero(fcEmu)
		fcEmu.ror(pos)
	case RORZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.ror(pos)
	case RORABS:
		pos := addrAbs(fcEmu)
		fcEmu.ror(pos)
	case RORABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.ror(pos)
	// sbc
	case SBCIMM:
		pos := addrImm(fcEmu)
		fcEmu.sbc(pos)
	case SBCZERO:
		pos := addrZero(fcEmu)
		fcEmu.sbc(pos)
	case SBCZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.sbc(pos)
	case SBCABS:
		pos := addrAbs(fcEmu)
		fcEmu.sbc(pos)
	case SBCABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.sbc(pos)
	case SBCABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.sbc(pos)
	case SBCINDX:
		pos := addrIndX(fcEmu)
		fcEmu.sbc(pos)
	case SBCINDY:
		pos := addrIndY(fcEmu)
		fcEmu.sbc(pos)
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
	// bvc
	case BVC:
		fcEmu.bvc()
	// bvs
	case BVS:
		fcEmu.bvs()
	// tax
	case TAX:
		fcEmu.tax()
	// tay
	case TAY:
		fcEmu.tay()
	// tsx
	case TSX:
		fcEmu.tsx()
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
		pos := addrImm(fcEmu)
		fcEmu.adc(pos)
	case ADCZERO:
		pos := addrZero(fcEmu)
		fcEmu.adc(pos)
	case ADCZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.adc(pos)
	case ADCABS:
		pos := addrAbs(fcEmu)
		fcEmu.adc(pos)
	case ADCABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.adc(pos)
	case ADCABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.adc(pos)
	case ADCINDX:
		pos := addrIndX(fcEmu)
		fcEmu.adc(pos)
	case ADCINDY:
		pos := addrIndY(fcEmu)
		fcEmu.adc(pos)
	// and
	case ANDIMM:
		pos := addrImm(fcEmu)
		fcEmu.and(pos)
	case ANDZERO:
		pos := addrZero(fcEmu)
		fcEmu.and(pos)	
	case ANDZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.and(pos)		
	case ANDABS:
		pos := addrAbs(fcEmu)
		fcEmu.and(pos)		
	case ANDABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.and(pos)
	case ANDABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.and(pos)
	case ANDINDX:
		pos := addrIndX(fcEmu)
		fcEmu.and(pos)
	case ANDINDY:
		pos := addrIndY(fcEmu)
		fcEmu.and(pos)
	// asl
	case ASLIMM:
		fcEmu.aslImm()	
	case ASLZERO:
		pos := addrZero(fcEmu)
		fcEmu.asl(pos)
	case ASLZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.asl(pos)
	case ASLABS:
		pos := addrAbs(fcEmu)
		fcEmu.asl(pos)
	case ASLABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.asl(pos)
	// bit
	case BITZERO:
		pos := addrZero(fcEmu)
		fcEmu.bit(pos)
	case BITABS:
		pos := addrAbs(fcEmu)
		fcEmu.bit(pos)
	// cmp
	case CMPIMM:
		pos := addrImm(fcEmu)
		fcEmu.cmp(pos)
	case CMPZERO:
		pos := addrZero(fcEmu)
		fcEmu.cmp(pos)
	case CMPZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.cmp(pos)
	case CMPABS:
		pos := addrAbs(fcEmu)
		fcEmu.cmp(pos)	
	case CMPABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.cmp(pos)
	case CMPABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.cmp(pos)
	case CMPINDX:
		pos := addrIndX(fcEmu)
		fcEmu.cmp(pos)
	case CMPINDY:
		pos := addrIndY(fcEmu)
		fcEmu.cmp(pos)
	// cpx
	case CPXIMM:
		pos := addrImm(fcEmu)
		fcEmu.cpx(pos)
	case CPXZERO:
		pos := addrZero(fcEmu)
		fcEmu.cpx(pos)
	case CPXABS:
		pos := addrAbs(fcEmu)
		fcEmu.cpx(pos)
	// cpy
	case CPYIMM:
		pos := addrImm(fcEmu)
		fcEmu.cpy(pos)
	case CPYZERO:
		pos := addrZero(fcEmu)
		fcEmu.cpy(pos)
	case CPYABS:
		pos := addrAbs(fcEmu)
		fcEmu.cpy(pos)
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
	// php
	case PHP:
		fcEmu.php()
	// pla
	case PLA:
		fcEmu.pla()
	// plp
	case PLP:
		fcEmu.plp()
	// clc
	case CLC:
		fcEmu.clc()
	// cld
	case CLD:
		fcEmu.cld()
	// cli
	case CLI:
		fcEmu.cli()
	// clv
	case CLV:
		fcEmu.clv()
	// sec 
	case SEC:
		fcEmu.sec()
	// sed 
	case SED:
		fcEmu.sed()
	// sei 
	case SEI:
		fcEmu.sei()
	// brk 
	case BRK:
		fcEmu.brk()
	// nop
	case NOP:
		fcEmu.nop()

	// unofficial code
	// nop 
	case NOP1A,NOP3A,NOP5A,NOP7A,NOPDA,NOPFA:
		fcEmu.nop()
	case NOP80:
		fcEmu.nopimm()
	case NOP4,NOP44,NOP64:
		fcEmu.nopzero()
	case NOP14,NOP34,NOP54,NOP74,NOPD4,NOPF4:
		fcEmu.nopzerox()
	case NOPC:
		fcEmu.nopabs()
	case NOP1C,NOP3C,NOP5C,NOP7C,NOPDC,NOPFC:
		fcEmu.nopabsx()
	// lax
	case LAXZERO:
		pos := addrZero(fcEmu)
		fcEmu.lax(pos)
	case LAXZEROY:
		pos := addrZeroY(fcEmu)
		fcEmu.lax(pos)
	case LAXABS:
		pos := addrAbs(fcEmu)
		fcEmu.lax(pos)
	case LAXABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.lax(pos)
	case LAXINDXA3:
		pos := addrIndX(fcEmu)
		fcEmu.lax(pos)
	case LAXINDY:
		pos := addrIndY(fcEmu)
		fcEmu.lax(pos)
	case SAXZERO:
		pos := addrZero(fcEmu)
		fcEmu.sax(pos)
	case SAXZEROY:
		pos := addrZeroY(fcEmu)
		fcEmu.sax(pos)
	case SAXABS:
		pos := addrAbs(fcEmu)
		fcEmu.sax(pos)
	case SAXINDX:
		pos := addrIndX(fcEmu)
		fcEmu.sax(pos)
	case SBCIMMEB:
		pos := addrImm(fcEmu)
		fcEmu.sbc(pos)
	// dcp
	case DCPZERO:
		pos := addrZero(fcEmu)
		fcEmu.dcp(pos)
	case DCPZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.dcp(pos)
	case DCPABS:
		pos := addrAbs(fcEmu)
		fcEmu.dcp(pos)
	case DCPABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.dcp(pos)
	case DCPABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.dcp(pos)
	case DCPINDX:
		pos := addrIndX(fcEmu)
		fcEmu.dcp(pos)
	case DCPINDY:
		pos := addrIndY(fcEmu)
		fcEmu.dcp(pos)
	// isc
	case ISCZERO:
		pos := addrZero(fcEmu)
		fcEmu.isc(pos)
	case ISCZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.isc(pos)
	case ISCABS:
		pos := addrAbs(fcEmu)
		fcEmu.isc(pos)
	case ISCABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.isc(pos)
	case ISCABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.isc(pos)
	case ISCINDX:
		pos := addrIndX(fcEmu)
		fcEmu.isc(pos)
	case ISCINDY:
		pos := addrIndY(fcEmu)
		fcEmu.isc(pos)
	// slo
	case SLOZERO:
		pos := addrZero(fcEmu)
		fcEmu.slo(pos)
	case SLOZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.slo(pos)
	case SLOABS:
		pos := addrAbs(fcEmu)
		fcEmu.slo(pos)
	case SLOABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.slo(pos)
	case SLOABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.slo(pos)
	case SLOINDX:
		pos := addrIndX(fcEmu)
		fcEmu.slo(pos)
	case SLOINDY:
		pos := addrIndY(fcEmu)
		fcEmu.slo(pos)
	// rla
	case RLAZERO:
		pos := addrZero(fcEmu)
		fcEmu.rla(pos)
	case RLAZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.rla(pos)
	case RLAABS:
		pos := addrAbs(fcEmu)
		fcEmu.rla(pos)
	case RLAABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.rla(pos)
	case RLAABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.rla(pos)
	case RLAINDX:
		pos := addrIndX(fcEmu)
		fcEmu.rla(pos)
	case RLAINDY:
		pos := addrIndY(fcEmu)
		fcEmu.rla(pos)
	// sre
	case SREZERO:
		pos := addrZero(fcEmu)
		fcEmu.sre(pos)
	case SREZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.sre(pos)
	case SREABS:
		pos := addrAbs(fcEmu)
		fcEmu.sre(pos)
	case SREABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.sre(pos)
	case SREABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.sre(pos)
	case SREINDX:
		pos := addrIndX(fcEmu)
		fcEmu.sre(pos)
	case SREINDY:
		pos := addrIndY(fcEmu)
		fcEmu.sre(pos)
	// rra
	case RRAZERO:
		pos := addrZero(fcEmu)
		fcEmu.rra(pos)
	case RRAZEROX:
		pos := addrZeroX(fcEmu)
		fcEmu.rra(pos)
	case RRAABS:
		pos := addrAbs(fcEmu)
		fcEmu.rra(pos)
	case RRAABSX:
		pos := addrAbsX(fcEmu)
		fcEmu.rra(pos)
	case RRAABSY:
		pos := addrAbsY(fcEmu)
		fcEmu.rra(pos)
	case RRAINDX:
		pos := addrIndX(fcEmu)
		fcEmu.rra(pos)
	case RRAINDY:
		pos := addrIndY(fcEmu)
		fcEmu.rra(pos)


	default:
		fmt.Printf("error PC : 0x%x\n", fcEmu.RegPc)
		fmt.Printf("not exist function code 0x%x\n", opcd)
		panic("error execute")
		//return 1
	}
	return 0

}
