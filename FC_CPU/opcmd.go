package main

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

	// STAZERO :
	STAZERO = 0x85
	// STAABS :
	STAABS = 0x8d

	//EORZERO :
	EORZERO = 0x45

	// BEQ :
	BEQ = 0xf0
	// BNE :
	BNE = 0xd0

	// ADCZERO :
	ADCZERO = 0x65

	// JMP :
	JMP = 0x4c

	// PHA :
	PHA = 0x48
	// PLA :
	PLA = 0x68

	// CLC :
	CLC = 0x18
)

func nflagset(fcEmu *emu,x uint8){
	x = x & 0b10000000
	if x == 0b10000000{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.regi["P"] = 0b01111111 & fcEmu.regi["P"]
		fcEmu.regi["P"] = fcEmu.regi["P"] + 0b10000000
	}else{
		fcEmu.regi["P"] = 0b01111111 & fcEmu.regi["P"]
	}
}

func zflagset(fcEmu *emu,x uint8){
	if x == 0x00{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.regi["P"] = 0b11111101 & fcEmu.regi["P"]
		fcEmu.regi["P"] = fcEmu.regi["P"] + 0b00000010
	}else{
		fcEmu.regi["P"] = 0b11111101 & fcEmu.regi["P"]
	}
}

func cflagset(fcEmu *emu,x uint16){
	if x == 0x0100{
		// いったんゼロにする。加算してフラッグを1にするのですでに1であった場合の対策
		fcEmu.regi["P"] = 0b11111110 & fcEmu.regi["P"]
		fcEmu.regi["P"] = fcEmu.regi["P"] + 0b00000001
	}else{
		fcEmu.regi["P"] = 0b11111110 & fcEmu.regi["P"]
	}
}

// LDA
func (fcEmu *emu) ldaImm() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	nflagset(fcEmu, fcEmu.memory[fcEmu.regPc])
	zflagset(fcEmu, fcEmu.memory[fcEmu.regPc])
	fcEmu.regPc++
}
func (fcEmu *emu) ldaZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc++
}
func (fcEmu *emu) ldaZeroX() {
	var immpos uint8 = fcEmu.memory[fcEmu.regPc]
	var pos uint16 = uint16(fcEmu.regi["X"] + immpos)
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc++
}
func (fcEmu *emu) ldaAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldaAbsX() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var xreg uint16 = uint16(fcEmu.regi["X"])
	var pos uint16 = (absposhigh << 8) + absposlow + xreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldaAbsY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var yreg uint16 = uint16(fcEmu.regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldaIndX() {
	var absposlow uint8 = uint8(fcEmu.memory[fcEmu.regPc])
	var xreg uint8 = uint8(fcEmu.regi["X"])
	var mempos uint16 = uint16(absposlow + xreg)
	var pos uint8 = fcEmu.memory[mempos]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc++
}
func (fcEmu *emu) ldaIndY() {
	var mempos uint8 = uint8(fcEmu.memory[fcEmu.regPc])
	var pos uint8 = fcEmu.memory[mempos]
	var yreg uint8 = uint8(fcEmu.regi["Y"])
	fcEmu.regi["A"] = fcEmu.memory[pos] + yreg
	nflagset(fcEmu, fcEmu.memory[pos] + yreg)
	zflagset(fcEmu, fcEmu.memory[pos] + yreg)
	fcEmu.regPc++
}

// LDX
func (fcEmu *emu) ldxImm() {
	fcEmu.regi["X"] = fcEmu.memory[fcEmu.regPc]
	nflagset(fcEmu, fcEmu.memory[fcEmu.regPc])
	zflagset(fcEmu, fcEmu.memory[fcEmu.regPc])
	fcEmu.regPc++
}
func (fcEmu *emu) ldxZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["X"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc++
}
func (fcEmu *emu) ldxZeroY() {
	var immpos uint8 = fcEmu.memory[fcEmu.regPc]
	var pos uint16 = uint16(fcEmu.regi["Y"] + immpos)
	fcEmu.regi["X"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc++
}
func (fcEmu *emu) ldxAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regi["X"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldxAbsY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var yreg uint16 = uint16(fcEmu.regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.regi["X"] = fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.memory[pos])
	zflagset(fcEmu, fcEmu.memory[pos])
	fcEmu.regPc = fcEmu.regPc + 2
}

// sta
func (fcEmu *emu) staZero() {
	fcEmu.memory[fcEmu.memory[fcEmu.regPc]] = fcEmu.regi["A"]
	fcEmu.regPc++
}
func (fcEmu *emu) staAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.memory[pos] = fcEmu.regi["A"]
	fcEmu.regPc = fcEmu.regPc + 2
}

// eor
func (fcEmu *emu) eorZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["A"] = fcEmu.regi["A"] ^ fcEmu.memory[pos]
	nflagset(fcEmu, fcEmu.regi["A"])
	zflagset(fcEmu, fcEmu.regi["A"])
	fcEmu.regPc++
}

// beq
func (fcEmu *emu) beq() {
	zeroflag := fcEmu.regi["P"] & 0b00000010
	if zeroflag == 0b00000010 {
		bufPC := uint16(fcEmu.regPc)
		minusflag := fcEmu.memory[fcEmu.regPc] & 0b10000000
		var imm uint16
		if minusflag == 0b10000000 {
			var minussign uint16 = 0xff00
			imm = uint16(fcEmu.memory[fcEmu.regPc]) + uint16(minussign)
		} else {
			imm = uint16(fcEmu.memory[fcEmu.regPc])
		}
		bufPC = bufPC + imm
		fcEmu.regPc = uint16(bufPC)
	}
	fcEmu.regPc++
}

// bne
func (fcEmu *emu) bne() {
	zeroflag := fcEmu.regi["P"] & 0b00000010
	if zeroflag == 0x00 {
		bufPC := int16(fcEmu.regPc)
		minusflag := fcEmu.memory[fcEmu.regPc] & 0b10000000
		var imm int16
		if minusflag > 0 {
			var minussign uint16 = 0xff00
			imm = int16(fcEmu.memory[fcEmu.regPc]) + int16(minussign)
		} else {
			imm = int16(fcEmu.memory[fcEmu.regPc])
		}
		bufPC = bufPC + imm
		fcEmu.regPc = uint16(bufPC)
	}
	fcEmu.regPc++
}

// adc
func (fcEmu *emu) adcZero() {
	var clflag uint8 = fcEmu.regi["P"] & 0b00000001
	pos := fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["A"] = fcEmu.memory[pos] + fcEmu.regi["A"] + clflag
	nflagset(fcEmu, fcEmu.regi["A"])
	zflagset(fcEmu, fcEmu.regi["A"])
	cflagset(fcEmu, uint16(fcEmu.memory[pos]) + uint16(fcEmu.regi["A"]) + uint16(clflag))
	fcEmu.regPc++
}

// jmp
func (fcEmu *emu) jmp() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regPc = pos
}

// pha
func (fcEmu *emu) pha() {
	var sppos uint16 = 0x1000 + uint16(fcEmu.regi["S"])
	fcEmu.memory[sppos] = fcEmu.regi["A"]
	fcEmu.regi["S"] = fcEmu.regi["S"] - 1
}

// pla
func (fcEmu *emu) pla() {
	if fcEmu.regi["S"] > 0xff {
		panic("stack error: dont pla when no push.")
	}
	fcEmu.regi["S"] = fcEmu.regi["S"] + 1
	var sppos uint16 = 0x1000 + uint16(fcEmu.regi["S"])
	fcEmu.regi["A"] = fcEmu.memory[sppos]
}

func (fcEmu *emu) clc() {
	fcEmu.regi["P"] = fcEmu.regi["P"] & 0b11111110
	cflagset(fcEmu,0)
}

func (fcEmu *emu) nop() {
}

func (fcEmu *emu) Execute() {
	opcd := fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc++
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
	// sta
	case STAZERO:
		fcEmu.staZero()
	case STAABS:
		fcEmu.staAbs()
	// eor
	case EORZERO:
		fcEmu.eorZero()
	// beq
	case BEQ:
		fcEmu.beq()
	// beq
	case BNE:
		fcEmu.bne()
	// adc
	case ADCZERO:
		fcEmu.adcZero()
	// jmp
	case JMP:
		fcEmu.jmp()
	// pha
	case PHA:
		fcEmu.pha()
	// pla
	case PLA:
		fcEmu.pla()
	// clc
	case CLC:
		fcEmu.clc()

	default:
		fmt.Printf("not exist function code 0x%x\n", opcd)
		fcEmu.nop()
		fcEmu.regPc = 0xffff
	}

}
