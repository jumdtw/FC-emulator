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
	// RIT :
	RTI = 0x40

	// PHA :
	PHA = 0x48
	// PLA :
	PLA = 0x68

	// CLC :
	CLC = 0x18
)

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

// sta
// 6502 はメモリマップトioなのでstaだけ命令処理以外に色々処理が入る場合が多い
func (fcEmu *CpuEmu) staZero() {
	fcEmu.Memory[fcEmu.Memory[fcEmu.RegPc]] = fcEmu.Regi["A"]
	fcEmu.RegPc++
}
func (fcEmu *CpuEmu) staAbs() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	// 0x2006だったらppuへのアクセス
	if pos == 0x2006 {
		fcEmu.VramAddr = fcEmu.VramAddr << 8
		fcEmu.VramAddr += uint16(fcEmu.Regi["A"])
	}
	if pos == 0x2007 {
		fcEmu.VramWriteFlag = true
		fcEmu.VramWriteValue = fcEmu.Regi["A"]
	}else {
		fcEmu.VramWriteFlag = false
	}
	fcEmu.Memory[pos] = fcEmu.Regi["A"]
	fcEmu.RegPc = fcEmu.RegPc + 2
}

// eor
func (fcEmu *CpuEmu) eorZero() {
	var pos uint8 = fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Regi["A"] ^ fcEmu.Memory[pos]
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
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

// adc
func (fcEmu *CpuEmu) adcZero() {
	var clflag uint8 = fcEmu.Regi["P"] & 0b00000001
	pos := fcEmu.Memory[fcEmu.RegPc]
	fcEmu.Regi["A"] = fcEmu.Memory[pos] + fcEmu.Regi["A"] + clflag
	nflagset(fcEmu, fcEmu.Regi["A"])
	zflagset(fcEmu, fcEmu.Regi["A"])
	cflagset(fcEmu, uint16(fcEmu.Memory[pos]) + uint16(fcEmu.Regi["A"]) + uint16(clflag))
	fcEmu.RegPc++
}

// jmp
func (fcEmu *CpuEmu) jmp() {
	var absposhigh uint16 = uint16(fcEmu.Memory[fcEmu.RegPc+1])
	var absposlow uint16 = uint16(fcEmu.Memory[fcEmu.RegPc])
	var pos uint16 = (absposhigh << 8) + absposlow
	fcEmu.RegPc = pos
}

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

func (fcEmu *CpuEmu) nop() {
}

func (fcEmu *CpuEmu) Execute() int {

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

	default:
		fmt.Printf("not exist function code 0x%x\n", opcd)
		fcEmu.nop()
		fcEmu.RegPc = 0xffff
		panic("error execute")
		return 1
	}
	return 0

}
