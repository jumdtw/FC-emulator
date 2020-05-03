package main

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
)

// LDA
func (fcEmu emu) ldaImm() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldaZero() {
	var pos uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldaZeroX() {
	var immpos uint8 = fcEmu.memory[fcEmu.regPc]
	var pos uint8 = fcEmu.regi["X"] + immpos
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldaAbs() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu emu) ldaAbsX() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var xreg uint16 = uint16(fcEmu.regi["X"])
	var pos uint16 = (absposhigh << 8) + absposlow + xreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}
func (fcEmu emu) ldaAbsY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var yreg uint16 = uint16(fcEmu.regi["Y"])
	var pos uint16 = (absposhigh << 8) + absposlow + yreg
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}

func (fcEmu emu) ldaIndX() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var xreg uint16 = uint16(fcEmu.regi["X"])
	var mempos uint16 = (absposhigh << 8) + absposlow + xreg
	var pos uint8 = fcEmu.memory[mempos]
	fcEmu.regi["A"] = fcEmu.memory[pos]
	fcEmu.regPc = fcEmu.regPc + 2
}

func (fcEmu emu) ldaIndY() {
	var absposhigh uint16 = uint16(fcEmu.memory[fcEmu.regPc+1])
	var absposlow uint16 = uint16(fcEmu.memory[fcEmu.regPc])
	var mempos uint16 = (absposhigh << 8) + absposlow
	var pos uint8 = fcEmu.memory[mempos]
	var yreg uint8 = uint8(fcEmu.regi["Y"])
	fcEmu.regi["A"] = fcEmu.memory[pos] + yreg
	fcEmu.regPc = fcEmu.regPc + 2
}

// LDX
func (fcEmu emu) ldxImm() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldxZero() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldxZeroY() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldxAbs() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}
func (fcEmu emu) ldxAbsX() {
	fcEmu.regi["A"] = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1
}

func (fcEmu emu) nop() {
}

func (fcEmu emu) Execute() {
	var opcd uint8 = fcEmu.memory[fcEmu.regPc]
	fcEmu.regPc = fcEmu.regPc + 1

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
		fcEmu.ldxAbsX()

	default:
		fcEmu.nop()
	}

}
