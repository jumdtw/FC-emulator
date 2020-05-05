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

	// BEQ :
	BEQ = 0xf0

	// PHA :
	PHA = 0x48
	// PLA :
	PLA = 0x68
)

// LDA
func (fcEmu *emu) ldaImm() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc++
}
func (fcEmu *emu) ldaZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc++
}
func (fcEmu *emu) ldaZeroX() {
	var immpos uint8 = fcEmu.memory[fcEmu.regPc]
	var pos uint16 = uint16(fcEmu.regi["X"] + immpos)
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc++
}
func (fcEmu *emu) ldaAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldaAbsX() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var xreg uint16 = uint16(fcEmu.regi["X"])
	var pos uint16 = (absposhigh << 8) + absposlow + xreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldaAbsY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var yreg uint16 = uint16(fcEmu.regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}

func (fcEmu *emu) ldaIndX() {
	var absposlow uint8 = uint8(fcEmu.memory[fcEmu.regPc])
	var xreg uint8 = uint8(fcEmu.regi["X"])
	var mempos uint16 = uint16(absposlow + xreg)
	var pos uint8 = fcEmu.memory[mempos]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc++
}

func (fcEmu *emu) ldaIndY() {
	var mempos uint8 = uint8(fcEmu.memory[fcEmu.regPc])
	var pos uint8 = fcEmu.memory[mempos]
	var yreg uint8 = uint8(fcEmu.regi["Y"])
	fcEmu.regi["A"] = fcEmu.memory[pos] + yreg
	fcEmu.regPc++
}

// LDX
func (fcEmu *emu) ldxImm() {
	fcEmu.regi["X"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc++
}
func (fcEmu *emu) ldxZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["X"] = fcEmu.memory[pos]
	fcEmu.regPc++
}
func (fcEmu *emu) ldxZeroY() {
	var immpos uint8 = fcEmu.memory[fcEmu.regPc]
	var pos uint16 = uint16(fcEmu.regi["Y"] + immpos)
	fcEmu.regi["X"] = fcEmu.memory[pos]
	fcEmu.regPc++
}
func (fcEmu *emu) ldxAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regi["X"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu *emu) ldxAbsY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var yreg uint16 = uint16(fcEmu.regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.regi["X"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}

// sta
func (fcEmu *emu) staZero() {
	fcEmu.memory[fcEmu.memory[fcEmu.regPc]] = fcEmu.regi["A"]
	fcEmu.regPc++
}

// beq
func (fcEmu *emu) beq() {
	zeroflag := fcEmu.regi["P"] & 0b00000010
	if zeroflag > 0 {
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

// pha
func (fcEmu *emu) pha() {
	var sppos uint16 = 0x1000 + uint16(fcEmu.regi["S"])
	fcEmu.memory[sppos] = fcEmu.regi["A"]
	fcEmu.regi["S"] = fcEmu.regi["S"] - 1
}

// pla
func (fcEmu *emu) pla() {
	fcEmu.regi["S"] = fcEmu.regi["S"] + 1
	var sppos uint16 = 0x1000 + uint16(fcEmu.regi["S"])
	fcEmu.regi["A"] = fcEmu.memory[sppos]
}

func (fcEmu *emu) nop() {
}

func (fcEmu *emu) Execute() {
	var opcd uint8 = fcEmu.memory[fcEmu.regPc]
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
	// beq
	case BEQ:
		fcEmu.beq()
	// pha
	case PHA:
		fcEmu.pha()
	// pla
	case PLA:
		fcEmu.pla()

	default:
		fmt.Printf("not exist function code 0x%x\n", opcd)
		fcEmu.nop()
		fcEmu.regPc = 0xffff
	}

}
